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
	"path/filepath"
	"strings"
	"time"

	authpkg "github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/auth"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/config"
	apperrors "github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/errors"
	"github.com/spf13/cobra"
)

type authLoginConfig struct {
	Token  string
	Force  bool
	Device bool
}

func buildAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "auth",
		Short:             "认证管理",
		Long:              "管理钉钉 CLI 的认证凭证。支持 OAuth 扫码登录和 Device Flow。",
		Args:              cobra.NoArgs,
		TraverseChildren:  true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(newAuthLoginCommand())
	cmd.AddCommand(
		newAuthLogoutCommand(),
		newAuthStatusCommand(),
		newAuthImportCommand(),
		newAuthExchangeCommand(),
		newAuthResetCommand(),
	)
	return cmd
}

func newAuthLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "login",
		Short:             "登录钉钉（自动刷新 token，必要时扫码）",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := resolveAuthLoginConfig(cmd)
			if err != nil {
				return err
			}
			configDir := defaultConfigDir()
			var tokenData *authpkg.TokenData

			switch {
			case strings.TrimSpace(cfg.Token) != "":
				tokenData = &authpkg.TokenData{
					AccessToken: cfg.Token,
					ExpiresAt:   time.Now().Add(config.ManualTokenExpiry),
				}
				if err := authpkg.SaveTokenData(configDir, tokenData); err != nil {
					return apperrors.NewInternal(fmt.Sprintf("failed to persist auth token: %v", err))
				}
			case cfg.Device:
				loginCtx, cancel := context.WithTimeout(cmd.Context(), config.DeviceFlowTimeout)
				defer cancel()

				provider := authpkg.NewDeviceFlowProvider(configDir, nil)
				provider.Output = cmd.ErrOrStderr()
				tokenData, err = provider.Login(loginCtx)
				if err != nil {
					return apperrors.NewAuth(fmt.Sprintf("device authorization failed: %v", err))
				}
			default:
				loginCtx, cancel := context.WithTimeout(cmd.Context(), config.OAuthFlowTimeout)
				defer cancel()

				provider := authpkg.NewOAuthProvider(configDir, nil)
				provider.Output = cmd.ErrOrStderr()
				configureOAuthProviderCompatibility(provider, configDir)
				tokenData, err = provider.Login(loginCtx, cfg.Force)
				if err != nil {
					return apperrors.NewAuth(fmt.Sprintf("dingtalk login failed: %v", err))
				}
			}

			clearCompatCache()

			w := cmd.OutOrStdout()
			fmt.Fprintln(w)
			if !cfg.Device && tokenData != nil && tokenData.IsAccessTokenValid() && !cfg.Force {
				fmt.Fprintf(w, "[OK] Token 有效，无需重新登录\n")
			} else {
				fmt.Fprintf(w, "[OK] 登录成功！\n")
			}
			if tokenData != nil {
				if tokenData.CorpName != "" {
					fmt.Fprintf(w, "%-16s%s\n", "企业:", tokenData.CorpName)
				}
				if tokenData.CorpID != "" {
					fmt.Fprintf(w, "%-16s%s\n", "企业 ID:", tokenData.CorpID)
				}
				if tokenData.UserName != "" {
					fmt.Fprintf(w, "%-16s%s\n", "用户:", tokenData.UserName)
				}
				if expiry := authLoginDisplayExpiry(tokenData); expiry != "" {
					fmt.Fprintf(w, "%-16s%s\n", "有效期:", expiry)
				}
			}
			fmt.Fprintf(w, "Token 将自动刷新，无需重复登录\n")
			return nil
		},
	}
	cmd.Flags().String("token", "", "Access token")
	cmd.Flags().Bool("device", false, "Use device authorization flow (compatibility flag)")
	cmd.Flags().Bool("force", false, "Force interactive login flow (compatibility flag)")
	cmd.Flags().String("redirect-url", "", "Loopback redirect URL compatibility flag")
	cmd.Flags().String("scopes", "", "Space-separated DingTalk OAuth scopes")
	cmd.Flags().String("authorize-url", "", "Override DingTalk authorization URL")
	cmd.Flags().String("token-url", "", "Override DingTalk token exchange URL")
	cmd.Flags().String("refresh-url", "", "Override DingTalk refresh token URL")
	cmd.Flags().Int("login-timeout", 0, "Compatibility flag for login timeout seconds")
	cmd.Flags().Bool("no-browser", false, "Compatibility flag for browser launch suppression")
	return cmd
}

func newAuthLogoutCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "logout",
		Short:             "清除认证信息",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			configDir := defaultConfigDir()
			revokeCtx, cancel := context.WithTimeout(cmd.Context(), 15*time.Second)
			defer cancel()
			_ = authpkg.RevokeTokenRemote(revokeCtx)

			if err := authpkg.DeleteTokenData(configDir); err != nil {
				return apperrors.NewInternal(fmt.Sprintf("failed to clear token data: %v", err))
			}
			_ = os.Remove(filepath.Join(configDir, "mcp_url"))
			_ = os.Remove(filepath.Join(configDir, "token"))
			clearCompatCache()
			w := cmd.OutOrStdout()
			fmt.Fprintln(w, "[OK] 已清除所有认证信息")
			fmt.Fprintln(w, "请运行 dws auth login 重新登录")
			return nil
		},
	}
}

func newAuthStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "status",
		Short:             "查看认证状态",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			configDir := defaultConfigDir()

			authenticated := false
			updatedAt := ""
			refreshed := false
			provider := authpkg.NewOAuthProvider(configDir, nil)
			configureOAuthProviderCompatibility(provider, configDir)
			if data, err := provider.Status(); err == nil {
				if !data.IsAccessTokenValid() && data.IsRefreshTokenValid() {
					refreshCtx, cancel := context.WithTimeout(cmd.Context(), 15*time.Second)
					_, refreshErr := provider.GetAccessToken(refreshCtx)
					cancel()
					if refreshErr == nil {
						if updatedData, statusErr := provider.Status(); statusErr == nil {
							data = updatedData
							refreshed = true
						}
					}
				}
				if authStatusAuthenticated(data) {
					authenticated = true
					updatedAt = authStatusUpdatedAt(data)
				}
			}

			w := cmd.OutOrStdout()
			if authenticated {
				if refreshed {
					fmt.Fprintf(w, "%-16s%s\n", "状态:", "已登录 ✅")
					fmt.Fprintln(w, "Token 已自动刷新")
				} else {
					fmt.Fprintf(w, "%-16s%s\n", "状态:", "已登录 ✅")
				}
				if updatedAt != "" {
					fmt.Fprintf(w, "%-16s%s\n", "有效期:", updatedAt)
				}
			} else {
				fmt.Fprintf(w, "%-16s%s\n", "状态:", "未登录")
				fmt.Fprintln(w, "运行 dws auth login 进行登录")
			}
			return nil
		},
	}
}

func newAuthImportCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "import <file>",
		Short:             "导入认证信息",
		Hidden:            true,
		Args:              cobra.ExactArgs(1),
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			configDir := defaultConfigDir()
			if err := validateOptionalPath("credentials file", args[0]); err != nil {
				return err
			}
			if _, err := authpkg.LoadExportedCredentials(cmd.Context(), args[0], configDir); err != nil {
				return apperrors.NewValidation(fmt.Sprintf("failed to import credentials: %v", err))
			}

			provider := authpkg.NewOAuthProvider(configDir, nil)
			refreshCtx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()
			token, refreshErr := provider.GetAccessToken(refreshCtx)
			tokenData, statusErr := provider.Status()
			if statusErr != nil {
				return apperrors.NewInternal(fmt.Sprintf("failed to load imported token data: %v", statusErr))
			}
			if refreshErr == nil {
				tokenData.AccessToken = token
			}
			clearCompatCache()

			w := cmd.OutOrStdout()
			fmt.Fprintln(w, "[OK] 认证信息导入成功")
			if refreshErr != nil {
				fmt.Fprintf(w, "[WARN] 凭证暂时无法刷新: %v\n", refreshErr)
			}
			fmt.Fprintln(w, "Token 将自动刷新，无需重复登录")
			return nil
		},
	}
}

func newAuthExchangeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "exchange",
		Short:             "Exchange an authorization code for credentials",
		Hidden:            true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			code, err := cmd.Flags().GetString("code")
			if err != nil {
				return apperrors.NewInternal("failed to read --code")
			}
			code = strings.TrimSpace(code)
			if code == "" {
				return apperrors.NewValidation("--code is required")
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return apperrors.NewInternal("failed to read --uid")
			}

			configDir := defaultConfigDir()
			provider := authpkg.NewOAuthProvider(configDir, nil)
			configureOAuthProviderCompatibility(provider, configDir)
			exchangeCtx, cancel := context.WithTimeout(cmd.Context(), time.Minute)
			defer cancel()
			tokenData, err := provider.ExchangeAuthCode(exchangeCtx, code, strings.TrimSpace(uid))
			if err != nil {
				return apperrors.NewAuth(fmt.Sprintf("failed to exchange authorization code: %v", err))
			}
			clearCompatCache()

			w := cmd.OutOrStdout()
			fmt.Fprintln(w, "[OK] 授权码兑换成功！")
			if strings.TrimSpace(uid) != "" {
				fmt.Fprintf(w, "%-16s%s\n", "用户:", strings.TrimSpace(uid))
			}
			if strings.TrimSpace(tokenData.CorpID) != "" {
				fmt.Fprintf(w, "%-16s%s\n", "企业 ID:", tokenData.CorpID)
			}
			if !tokenData.ExpiresAt.IsZero() {
				fmt.Fprintf(w, "%-16s%s\n", "有效期:", authLoginFormatExpiry(tokenData.ExpiresAt))
			}
			return nil
		},
	}
	cmd.Flags().String("code", "", "Authorization code")
	cmd.Flags().String("uid", "", "Optional user identifier for compatibility")
	cmd.Flags().String("client-id", "", "Compatibility flag")
	cmd.Flags().String("authorize-url", "", "Compatibility flag")
	cmd.Flags().String("token-url", "", "Compatibility flag")
	cmd.Flags().String("refresh-url", "", "Compatibility flag")
	cmd.Flags().String("redirect-url", "", "Compatibility flag")
	cmd.Flags().String("scopes", "", "Compatibility flag")
	return cmd
}

func newAuthResetCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "reset",
		Short:             "重置认证信息（清除本地 Token，触发重新授权）",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			configDir := defaultConfigDir()
			if err := authpkg.DeleteTokenData(configDir); err != nil {
				return apperrors.NewInternal(fmt.Sprintf("failed to reset token data: %v", err))
			}
			_ = os.Remove(filepath.Join(configDir, "mcp_url"))
			_ = os.Remove(filepath.Join(configDir, "token"))
			clearCompatCache()
			w := cmd.OutOrStdout()
			fmt.Fprintln(w, "[OK] 认证信息已重置")
			fmt.Fprintln(w, "请运行 dws auth login 重新登录")
			return nil
		},
	}
}

func timeOrEmpty(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format(time.RFC3339)
}

func authLoginFormatExpiry(t time.Time) string {
	remaining := time.Until(t)
	if remaining <= 0 {
		return "已过期"
	}
	if remaining > 24*time.Hour {
		return fmt.Sprintf("%.0f 天后", remaining.Hours()/24)
	}
	return fmt.Sprintf("%.0f 小时后", remaining.Hours())
}

// authLoginDisplayExpiry 返回用于显示的有效期（优先显示 refresh token 有效期）
func authLoginDisplayExpiry(data *authpkg.TokenData) string {
	if data == nil {
		return ""
	}
	// 优先使用 refresh token 有效期（更长，对用户更有意义）
	if data.IsRefreshTokenValid() {
		return authLoginFormatExpiry(data.RefreshExpAt)
	}
	// 回退到 access token 有效期
	if !data.ExpiresAt.IsZero() {
		return authLoginFormatExpiry(data.ExpiresAt)
	}
	return ""
}

func clearCompatCache() {
	store := cacheStoreFromEnv()
	if store != nil {
		_ = os.RemoveAll(store.Root)
	}
}

func resolveAuthLoginConfig(cmd *cobra.Command) (authLoginConfig, error) {
	token, err := cmd.Flags().GetString("token")
	if err != nil {
		return authLoginConfig{}, apperrors.NewInternal("failed to read --token")
	}
	device, err := cmd.Flags().GetBool("device")
	if err != nil {
		return authLoginConfig{}, apperrors.NewInternal("failed to read --device")
	}
	force, err := cmd.Flags().GetBool("force")
	if err != nil {
		return authLoginConfig{}, apperrors.NewInternal("failed to read --force")
	}
	return authLoginConfig{
		Token:  strings.TrimSpace(token),
		Force:  force,
		Device: device,
	}, nil
}

func authStatusAuthenticated(data *authpkg.TokenData) bool {
	if data == nil {
		return false
	}
	return data.IsAccessTokenValid() || data.IsRefreshTokenValid()
}

func authStatusUpdatedAt(data *authpkg.TokenData) string {
	if data == nil {
		return ""
	}
	if data.IsAccessTokenValid() {
		return timeOrEmpty(data.ExpiresAt)
	}
	if data.IsRefreshTokenValid() {
		return timeOrEmpty(data.RefreshExpAt)
	}
	return ""
}
