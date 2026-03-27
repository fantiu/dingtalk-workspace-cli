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

package helpers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cobracmd"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/errors"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/executor"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/i18n"
	"github.com/spf13/cobra"
)

type dingtalkVendorHandler struct {
	subtree string
}

func init() {
	for _, subtree := range []string{"discovery", "oa-plus", "ai-sincere-hire"} {
		subtree := subtree
		RegisterHiddenDingTalk(func() Handler {
			return dingtalkVendorHandler{subtree: subtree}
		})
	}
}

func (d dingtalkVendorHandler) Name() string {
	return d.subtree
}

func (d dingtalkVendorHandler) Command(runner executor.Runner) *cobra.Command {
	switch d.subtree {
	case "discovery":
		return newDiscoveryCommand(runner)
	case "oa-plus":
		return newOAPlusCommand(runner)
	case "ai-sincere-hire":
		return newAISincereHireCommand(runner)
	default:
		return newHiddenGroup("dingtalk", "Hidden DingTalk vendor extensions")
	}
}

func registerDingTalkFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("json", "", "NewsFeedPushRequest JSON payload")
	cmd.PersistentFlags().String("source", "", "Crawl source ID")
	cmd.PersistentFlags().String("filenames", "", "Comma-separated file names")
	cmd.PersistentFlags().String("keyword", "", "Subscription keyword")
	cmd.PersistentFlags().String("instance-id", "", "Approval instance ID")
	cmd.PersistentFlags().String("process-code", "", "Approval process code")
	cmd.PersistentFlags().String("size", "20", "Result size")
	cmd.PersistentFlags().String("cursor", "", "Cursor")
}

func newHiddenGroup(use, short string) *cobra.Command {
	return cobracmd.NewHiddenGroupCommand(use, short)
}

func newVisibleGroup(use, short string) *cobra.Command {
	return cobracmd.NewGroupCommand(use, short)
}

func newDiscoveryCommand(runner executor.Runner) *cobra.Command {
	cmd := newVisibleGroup("discovery", "Content discovery")
	registerDingTalkFlags(cmd)
	media := newVisibleGroup("media", "Media content")
	media.AddCommand(&cobra.Command{
		Use:               "save",
		Short:             "Save media content",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDiscoveryMediaSave(cmd, runner)
		},
	})
	media.AddCommand(&cobra.Command{
		Use:               "subscribe",
		Short:             "Subscribe to media source",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDiscoveryMediaSubscribe(cmd, runner)
		},
	})

	oss := newVisibleGroup("oss", "Upload credentials")
	oss.AddCommand(&cobra.Command{
		Use:               "get-upload-url",
		Short:             "Get OSS upload URL",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDiscoveryOSSGetUploadURL(cmd, runner)
		},
	})

	subscribe := newVisibleGroup("subscribe", "Keyword subscription")
	subscribe.AddCommand(&cobra.Command{
		Use:               "save",
		Short:             i18n.T("保存Keyword subscription规则"),
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDiscoverySubscribeSave(cmd, runner)
		},
	})

	cmd.AddCommand(media, oss, subscribe)
	return cmd
}

func runDiscoveryMediaSave(cmd *cobra.Command, runner executor.Runner) error {
	payload, err := cmd.Flags().GetString("json")
	if err != nil {
		return errors.NewInternal("failed to read --json")
	}
	if strings.TrimSpace(payload) == "" {
		return errors.NewValidation("--json is required")
	}

	var request any
	if err := json.Unmarshal([]byte(payload), &request); err != nil {
		return errors.NewValidation(fmt.Sprintf("--json must be valid JSON: %v", err))
	}

	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk discovery media save",
		"dingtalk-discovery",
		"save_video_and_image",
		map[string]any{
			"NewsFeedPushRequest": request,
		},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func runDiscoveryMediaSubscribe(cmd *cobra.Command, runner executor.Runner) error {
	source, err := cmd.Flags().GetString("source")
	if err != nil {
		return errors.NewInternal("failed to read --source")
	}
	if strings.TrimSpace(source) == "" {
		return errors.NewValidation("--source is required")
	}

	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk discovery media subscribe",
		"dingtalk-discovery",
		"save_media_subscribe_rule",
		map[string]any{
			"skillCrawlId": source,
		},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func runDiscoveryOSSGetUploadURL(cmd *cobra.Command, runner executor.Runner) error {
	filenames, err := cmd.Flags().GetString("filenames")
	if err != nil {
		return errors.NewInternal("failed to read --filenames")
	}
	if strings.TrimSpace(filenames) == "" {
		return errors.NewValidation("--filenames is required")
	}

	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk discovery oss get-upload-url",
		"dingtalk-discovery",
		"batch_get_oss_temp_upload_url",
		map[string]any{
			"filenames": filenames,
		},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func runDiscoverySubscribeSave(cmd *cobra.Command, runner executor.Runner) error {
	keyword, err := cmd.Flags().GetString("keyword")
	if err != nil {
		return errors.NewInternal("failed to read --keyword")
	}
	if strings.TrimSpace(keyword) == "" {
		return errors.NewValidation("--keyword is required")
	}

	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk discovery subscribe save",
		"dingtalk-discovery",
		"save_keyword_subscribe_rule",
		map[string]any{
			"keyword": keyword,
		},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func newOAPlusCommand(runner executor.Runner) *cobra.Command {
	cmd := newVisibleGroup("oa-plus", "OA approval enhanced")
	registerDingTalkFlags(cmd)
	approval := newVisibleGroup("approval", i18n.T("审批实例管理"))
	approval.AddCommand(&cobra.Command{
		Use:               "get",
		Short:             i18n.T("获取审批实例详情"),
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runOAPlusApprovalGet(cmd, runner)
		},
	})
	approval.AddCommand(&cobra.Command{
		Use:               "list",
		Short:             i18n.T("分页查询审批实例"),
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runOAPlusApprovalList(cmd, runner)
		},
	})
	cmd.AddCommand(approval)
	return cmd
}

func runOAPlusApprovalGet(cmd *cobra.Command, runner executor.Runner) error {
	instanceID, err := cmd.Flags().GetString("instance-id")
	if err != nil {
		return errors.NewInternal("failed to read --instance-id")
	}
	if strings.TrimSpace(instanceID) == "" {
		return errors.NewValidation("--instance-id is required")
	}

	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk oa-plus approval get",
		"dingtalk-oa-plus",
		"get_approval_instance",
		map[string]any{
			"instanceId": instanceID,
		},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func runOAPlusApprovalList(cmd *cobra.Command, runner executor.Runner) error {
	processCode, err := cmd.Flags().GetString("process-code")
	if err != nil {
		return errors.NewInternal("failed to read --process-code")
	}
	if strings.TrimSpace(processCode) == "" {
		return errors.NewValidation("--process-code is required")
	}

	size, err := cmd.Flags().GetString("size")
	if err != nil {
		return errors.NewInternal("failed to read --size")
	}
	cursor, err := cmd.Flags().GetString("cursor")
	if err != nil {
		return errors.NewInternal("failed to read --cursor")
	}

	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk oa-plus approval list",
		"dingtalk-oa-plus",
		"list_approval_instances",
		map[string]any{
			"processCode": processCode,
			"size":        size,
			"cursor":      cursor,
		},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func newAISincereHireCommand(runner executor.Runner) *cobra.Command {
	cmd := newVisibleGroup("ai-sincere-hire", i18n.T("AI诚聘"))
	registerDingTalkFlags(cmd)
	cmd.AddCommand(&cobra.Command{
		Use:               "guide",
		Short:             i18n.T("获取使用指引"),
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAISincereHireGuide(cmd, runner)
		},
	})
	job := newVisibleGroup("job", i18n.T("岗位查询"))
	job.AddCommand(&cobra.Command{
		Use:               "list",
		Short:             i18n.T("查询在招岗位"),
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAISincereHireJobList(cmd, runner)
		},
	})
	talent := newVisibleGroup("talent", i18n.T("人才查询"))
	talent.AddCommand(&cobra.Command{
		Use:               "list",
		Short:             i18n.T("查询入职人才"),
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAISincereHireTalentList(cmd, runner)
		},
	})
	cmd.AddCommand(job, talent)
	return cmd
}

func runAISincereHireGuide(cmd *cobra.Command, runner executor.Runner) error {
	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk ai-sincere-hire guide",
		"dingtalk-ai-sincere-hire",
		"query_guide_url",
		map[string]any{},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func runAISincereHireJobList(cmd *cobra.Command, runner executor.Runner) error {
	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk ai-sincere-hire job list",
		"dingtalk-ai-sincere-hire",
		"query_opening_job_list",
		map[string]any{},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}

func runAISincereHireTalentList(cmd *cobra.Command, runner executor.Runner) error {
	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(
		"dingtalk ai-sincere-hire talent list",
		"dingtalk-ai-sincere-hire",
		"query_success_talent_list",
		map[string]any{},
	))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}
