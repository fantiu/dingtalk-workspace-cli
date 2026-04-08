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

// Package runtimetoken resolves API bearer tokens for features that bypass
// the MCP runner (e.g. A2A gateway) but should behave like tool calls.
package runtimetoken

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strings"

	authpkg "github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/auth"
)

// ResolveAccessToken returns a non-empty bearer token using the same sources
// as MCP: explicitToken when non-empty, else OAuth access_token, else legacy token file.
func ResolveAccessToken(ctx context.Context, configDir string, explicitToken string) (string, error) {
	if t := strings.TrimSpace(explicitToken); t != "" {
		return t, nil
	}
	if strings.TrimSpace(configDir) == "" {
		return "", fmt.Errorf("config directory is empty")
	}
	disc := slog.New(slog.NewTextHandler(io.Discard, nil))
	provider := authpkg.NewOAuthProvider(configDir, disc)
	token, err := provider.GetAccessToken(ctx)
	if err == nil && strings.TrimSpace(token) != "" {
		return strings.TrimSpace(token), nil
	}
	if err != nil && errors.Is(err, authpkg.ErrTokenDecryption) {
		return "", err
	}
	manager := authpkg.NewManager(configDir, nil)
	if legacy, _, merr := manager.GetToken(); merr == nil && strings.TrimSpace(legacy) != "" {
		return strings.TrimSpace(legacy), nil
	}
	if err != nil {
		return "", err
	}
	return "", fmt.Errorf("no credentials found, run: dws auth login")
}
