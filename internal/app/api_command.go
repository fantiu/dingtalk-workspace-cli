// Copyright 2026 Alibaba Group
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/apiclient"
	authpkg "github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/auth"
	apperrors "github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/errors"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/output"
	"github.com/spf13/cobra"
)

// apiFlags holds the flags specific to the `dws api` command.
type apiFlags struct {
	params    string
	data      string
	pageAll   bool
	pageLimit int
	pageDelay int
	baseURL   string
}

// newAPICommand creates the `dws api` subcommand for raw DingTalk OpenAPI calls.
func newAPICommand(flags *GlobalFlags) *cobra.Command {
	af := &apiFlags{}

	cmd := &cobra.Command{
		Use:   "api <METHOD> <PATH> [flags]",
		Short: "调用钉钉 OpenAPI (Raw HTTP)",
		Long: `直接调用钉钉 OpenAPI，支持新版和旧版两种 API 形态。

新版 API (api.dingtalk.com):
  Token 通过 HTTP Header (x-acs-dingtalk-access-token) 传递。
  路径格式: /v1.0/xxx 或 /v2.0/xxx

旧版 API (oapi.dingtalk.com):
  Token 通过 URL 查询参数 (access_token) 传递。
  路径格式: /topapi/v2/xxx 或完整 URL https://oapi.dingtalk.com/topapi/...

仅限使用自有应用凭证（--client-id/--client-secret）登录后使用。
通过 MCP 默认凭证登录获取的加密 token 不支持 raw API 调用。

示例:
  # === 新版 API (api.dingtalk.com) ===

  # 获取当前用户信息
  dws api GET /v1.0/contact/users/me

  # 搜索用户 (POST + JSON body)
  dws api POST /v1.0/contact/users/search \
    --data '{"queryWord":"张三","offset":0,"size":10}'

  # 创建日历事件
  dws api POST /v1.0/calendar/users/me/calendars/primary/events \
    --data '{"summary":"Team Meeting","start":{"dateTime":"2026-01-01T10:00:00+08:00"}}'

  # === 旧版 API (oapi.dingtalk.com) ===

  # 获取用户详情 (旧版，使用 --base-url)
  dws api POST /topapi/v2/user/get \
    --base-url https://oapi.dingtalk.com \
    --data '{"userid":"manager123"}'

  # 也可以直接使用完整 URL
  dws api POST https://oapi.dingtalk.com/topapi/v2/user/get \
    --data '{"userid":"manager123"}'

  # === 通用功能 ===

  # 分页获取所有结果
  dws api GET /v1.0/attendance/groups --page-all --page-limit 5

  # Dry-run 预览请求
  dws api GET /v1.0/contact/users/me --dry-run

  # 使用 jq 过滤输出
  dws api GET /v1.0/contact/users/me --jq '.nick'`,
		Args:              cobra.ExactArgs(2),
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAPI(cmd, args, flags, af)
		},
	}

	cmd.Flags().StringVar(&af.params, "params", "", "查询参数 JSON (支持 - 从 stdin 读取)")
	cmd.Flags().StringVar(&af.data, "data", "", "请求体 JSON (支持 - 从 stdin 读取)")
	cmd.Flags().BoolVar(&af.pageAll, "page-all", false, "自动遍历所有分页")
	cmd.Flags().IntVar(&af.pageLimit, "page-limit", apiclient.DefaultPageLimit, "最大分页数 (0 = 无限)")
	cmd.Flags().IntVar(&af.pageDelay, "page-delay", apiclient.DefaultPageDelay, "分页间隔毫秒")
	cmd.Flags().StringVar(&af.baseURL, "base-url", "", "覆盖 API 基础 URL (默认 https://api.dingtalk.com)")

	return cmd
}

// runAPI is the main execution logic for `dws api`.
func runAPI(cmd *cobra.Command, args []string, gf *GlobalFlags, af *apiFlags) error {
	ctx := cmd.Context()
	method := args[0]
	path := args[1]

	// 1. Validate HTTP method.
	method, err := apiclient.ValidateMethod(method)
	if err != nil {
		return apperrors.NewValidation(err.Error())
	}

	// 2. Validate API path.
	if err := apiclient.ValidatePath(path); err != nil {
		return apperrors.NewValidation(err.Error())
	}

	// 3. Validate input safety for params and data.
	if err := apiclient.ValidateUserInput(af.params, "--params"); err != nil {
		return apperrors.NewValidation(err.Error())
	}
	if err := apiclient.ValidateUserInput(af.data, "--data"); err != nil {
		return apperrors.NewValidation(err.Error())
	}

	// 4. Validate mutual exclusion.
	if err := apiclient.ValidateStdinExclusion(af.params, af.data); err != nil {
		return apperrors.NewValidation(err.Error())
	}
	if err := apiclient.ValidateFlagExclusion(gf.Output, af.pageAll); err != nil {
		return apperrors.NewValidation(err.Error())
	}

	// 5. Parse --params.
	params, err := apiclient.ParseJSONMap(af.params, "--params", os.Stdin)
	if err != nil {
		return apperrors.NewValidation(err.Error())
	}

	// 6. Parse --data.
	body, err := apiclient.ParseOptionalBody(method, af.data, os.Stdin)
	if err != nil {
		return apperrors.NewValidation(err.Error())
	}

	// 7. Determine API style (new vs legacy) for token resolution.
	fullURL := apiclient.NormalisePath(path, af.baseURL)
	isLegacy := apiclient.IsLegacyAPI(fullURL)

	// 7b. Security: validate target host is a trusted DingTalk domain.
	if err := apiclient.ValidateTargetHost(fullURL); err != nil {
		return apperrors.NewValidation(err.Error())
	}

	// 8. Resolve app-level token based on API style (with timeout).
	tokenCtx, tokenCancel := context.WithTimeout(ctx, 15*time.Second)
	defer tokenCancel()
	token, err := resolveRawAPIToken(tokenCtx, gf.Token, isLegacy)
	if err != nil {
		return err
	}

	// 9. Build request.
	req := apiclient.RawAPIRequest{
		Method: method,
		Path:   path,
		Params: params,
		Data:   body,
	}

	baseURL := af.baseURL

	// 10. Dry-run mode.
	if gf.DryRun {
		return apiclient.PrintDryRun(cmd.OutOrStdout(), req, baseURL, token)
	}

	// 11. Create client with timeout.
	client := apiclient.NewClient(token, baseURL)
	if gf.Timeout > 0 {
		client.HTTPClient.Timeout = time.Duration(gf.Timeout) * time.Second
	}

	// 12. Execute request (with or without pagination).
	format := output.Format(gf.Format)
	respOpts := apiclient.ResponseOptions{
		OutputPath: gf.Output,
		Format:     format,
		JqExpr:     gf.JQ,
		Fields:     gf.Fields,
		Out:        cmd.OutOrStdout(),
		ErrOut:     cmd.ErrOrStderr(),
	}

	if af.pageAll {
		return runPaginated(ctx, client, req, af, respOpts)
	}

	resp, err := client.Do(ctx, req)
	if err != nil {
		return apperrors.NewAPI(fmt.Sprintf("API 请求失败: %v", err))
	}
	return apiclient.HandleResponse(resp, respOpts)
}

// runPaginated executes a paginated API request and outputs all results.
func runPaginated(ctx context.Context, client *apiclient.APIClient, req apiclient.RawAPIRequest, af *apiFlags, opts apiclient.ResponseOptions) error {
	pages, err := client.PaginateAll(ctx, req, apiclient.PaginationOptions{
		PageLimit: af.pageLimit,
		PageDelay: af.pageDelay,
	})
	if err != nil && len(pages) == 0 {
		return apperrors.NewAPI(fmt.Sprintf("分页请求失败: %v", err))
	}

	// Output all pages as a JSON array.
	return output.WriteFiltered(opts.Out, opts.Format, pages, opts.Fields, opts.JqExpr)
}

// resolveRawAPIToken resolves an app-level access token for raw API calls.
// It uses AppTokenProvider to obtain the correct token type based on API style:
//   - isLegacy=false: fetches from POST api.dingtalk.com/v1.0/oauth2/accessToken
//   - isLegacy=true:  fetches from GET oapi.dingtalk.com/gettoken
//
// Tokens are cached in keychain and auto-refreshed when expired.
func resolveRawAPIToken(ctx context.Context, explicitToken string, isLegacy bool) (string, error) {
	// Explicit --token flag takes priority (user knows what they're doing).
	if t := strings.TrimSpace(explicitToken); t != "" {
		return t, nil
	}

	// Resolve app credentials (clientID/clientSecret).
	appKey := authpkg.ClientID()
	appSecret := authpkg.ClientSecret()

	if appKey == "" || appSecret == "" || strings.HasPrefix(appKey, "<") || strings.HasPrefix(appSecret, "<") {
		return "", apperrors.NewAuth(
			"缺少应用凭证。dws api 需要使用自有应用的 AppKey/AppSecret 获取 accessToken。\n\n" +
				"解决方法:\n" +
				"  1. 使用自有应用凭证登录:\n" +
				"     dws auth login --client-id <APP_KEY> --client-secret <APP_SECRET>\n\n" +
				"  2. 或通过环境变量设置:\n" +
				"     export DWS_CLIENT_ID=<APP_KEY>\n" +
				"     export DWS_CLIENT_SECRET=<APP_SECRET>\n" +
				"     dws auth login\n\n" +
				"说明: 通过 MCP 默认凭证登录的加密 token 无法用于 raw API 调用。",
		)
	}

	// Use AppTokenProvider for automatic caching and refresh.
	configDir := defaultConfigDir()
	provider := &authpkg.AppTokenProvider{
		ConfigDir: configDir,
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	token, err := provider.GetToken(ctx, isLegacy)
	if err != nil {
		return "", apperrors.NewAuth(fmt.Sprintf("获取应用级访问令牌失败: %v", err))
	}
	if strings.TrimSpace(token) == "" {
		return "", apperrors.NewAuth("应用级访问令牌为空，请检查应用凭证是否正确")
	}

	return strings.TrimSpace(token), nil
}
