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
	"fmt"
	"strings"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/cobracmd"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/executor"
	"github.com/spf13/cobra"
)

func init() {
	RegisterPublic(func() Handler {
		return creditHelper{}
	})
}

type creditHelper struct{}

func (creditHelper) Name() string {
	return "credit"
}

func (creditHelper) Command(runner executor.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "credit",
		Short:             "Enterprise credit search and risk helpers",
		Args:              cobra.NoArgs,
		TraverseChildren:  true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(
		newCreditSearchCommand(runner),
		newCreditRiskCommand(runner),
		newCreditEquityCommand(runner),
	)

	return cmd
}

func newCreditSearchCommand(runner executor.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "search",
		Short:             "Enterprise name search",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			name = strings.TrimSpace(name)
			if name == "" {
				return fmt.Errorf("--name is required")
			}

			params := map[string]any{
				"company_name": name,
			}
			if cmd.Flags().Changed("page") {
				page, err := cmd.Flags().GetInt("page")
				if err != nil {
					return err
				}
				params["page"] = page
			}
			if cmd.Flags().Changed("size") {
				size, err := cmd.Flags().GetInt("size")
				if err != nil {
					return err
				}
				params["size"] = size
			}

			return runHelper(cmd, runner, "credit-ep", "ep_info_search_query", params)
		},
	}
	cmd.Flags().String("name", "", "Enterprise name keyword")
	cmd.Flags().Int("page", 0, "Page offset")
	cmd.Flags().Int("size", 0, "Page size")
	return cmd
}

func newCreditRiskCommand(runner executor.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "risk",
		Short:             "Enterprise risk information",
		Args:              cobra.NoArgs,
		TraverseChildren:  true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	court := &cobra.Command{
		Use:               "court",
		Short:             "Court notice",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cert, err := cmd.Flags().GetString("cert")
			if err != nil {
				return err
			}
			cert = strings.TrimSpace(cert)
			if cert == "" {
				return fmt.Errorf("--cert is required")
			}

			params := map[string]any{
				"ep_cert_no": cert,
			}
			if cmd.Flags().Changed("page") {
				page, err := cmd.Flags().GetInt("page")
				if err != nil {
					return err
				}
				params["page"] = page
			}
			if cmd.Flags().Changed("size") {
				size, err := cmd.Flags().GetInt("size")
				if err != nil {
					return err
				}
				params["size"] = size
			}

			return runHelper(cmd, runner, "credit-risk", "ep_dossier_courtnotice_query", params)
		},
	}
	court.Flags().String("cert", "", "Enterprise registration number or credit code")
	court.Flags().Int("page", 0, "Page offset")
	court.Flags().Int("size", 0, "Page size")
	cmd.AddCommand(court)
	return cmd
}

func newCreditEquityCommand(runner executor.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "equity",
		Short:             "Enterprise equity information",
		Args:              cobra.NoArgs,
		TraverseChildren:  true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	shareholder := &cobra.Command{
		Use:               "shareholder",
		Short:             "Shareholder information",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cert, err := cmd.Flags().GetString("cert")
			if err != nil {
				return err
			}
			cert = strings.TrimSpace(cert)
			if cert == "" {
				return fmt.Errorf("--cert is required")
			}

			params := map[string]any{
				"ep_cert_no": cert,
			}
			if cmd.Flags().Changed("page") {
				page, err := cmd.Flags().GetInt("page")
				if err != nil {
					return err
				}
				params["page"] = page
			}
			if cmd.Flags().Changed("size") {
				size, err := cmd.Flags().GetInt("size")
				if err != nil {
					return err
				}
				params["size"] = size
			}

			return runHelper(cmd, runner, "credit-equity", "ep_dossier_shareholder_query", params)
		},
	}
	shareholder.Flags().String("cert", "", "Enterprise registration number or credit code")
	shareholder.Flags().Int("page", 0, "Page offset")
	shareholder.Flags().Int("size", 0, "Page size")
	cmd.AddCommand(shareholder)
	return cmd
}

func runHelper(cmd *cobra.Command, runner executor.Runner, canonicalProduct, tool string, params map[string]any) error {
	result, err := runner.Run(cmd.Context(), executor.NewHelperInvocation(cobracmd.LegacyCommandPath(cmd), canonicalProduct, tool, params))
	if err != nil {
		return err
	}
	return writeCommandPayload(cmd, result)
}
