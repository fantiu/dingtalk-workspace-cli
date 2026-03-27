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

package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// ExportedCredentials represents the JSON structure of an exported credentials file,
// used by auth import to restore credentials on another machine.
type ExportedCredentials struct {
	RefreshToken   string `json:"refresh_token,omitempty"`
	PersistentCode string `json:"persistent_code,omitempty"`
	CorpID         string `json:"corp_id"`
	UserID         string `json:"user_id,omitempty"`
	UserName       string `json:"user_name,omitempty"`
	CorpName       string `json:"corp_name,omitempty"`
	ExportedAt     string `json:"exported_at"`
}

func LoadExportedCredentials(ctx context.Context, path, configDir string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading credentials file: %w", err)
	}

	var creds ExportedCredentials
	if err := json.Unmarshal(b, &creds); err != nil {
		return "", fmt.Errorf("parsing credentials file: %w", err)
	}
	// Accept either persistent_code or refresh_token as a valid credential.
	if creds.PersistentCode == "" && creds.RefreshToken == "" {
		return "", fmt.Errorf("credentials file has no usable credential (need persistent_code or refresh_token)")
	}

	data := &TokenData{
		PersistentCode: creds.PersistentCode,
		RefreshToken:   creds.RefreshToken,
		RefreshExpAt:   time.Now().Add(30 * 24 * time.Hour),
		CorpID:         creds.CorpID,
		UserID:         creds.UserID,
		UserName:       creds.UserName,
		CorpName:       creds.CorpName,
	}
	if err := SaveTokenData(configDir, data); err != nil {
		return "", fmt.Errorf("saving imported credentials: %w", err)
	}
	return "", nil
}
