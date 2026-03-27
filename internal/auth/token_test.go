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
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func setupTestMAC(t *testing.T) {
	t.Helper()
}

func TestTokenSaveLoadAndDelete(t *testing.T) {
	setupTestMAC(t)

	configDir := t.TempDir()
	now := time.Now().UTC()
	original := &TokenData{
		AccessToken:    "at_test_123",
		RefreshToken:   "rt_test_456",
		PersistentCode: "pc_test_789",
		ExpiresAt:      now.Add(2 * time.Hour),
		RefreshExpAt:   now.Add(30 * 24 * time.Hour),
		CorpID:         "ding123",
		UserID:         "user001",
		UserName:       "张三",
		CorpName:       "测试科技",
	}

	if err := SaveTokenData(configDir, original); err != nil {
		t.Fatalf("SaveTokenData() error = %v", err)
	}

	// Verify .data file was created with correct permissions.
	dataPath := filepath.Join(configDir, secureDataFile)
	info, err := os.Stat(dataPath)
	if err != nil {
		t.Fatalf("Stat(.data) error = %v", err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf(".data perms = %o, want 600", info.Mode().Perm())
	}
	if _, err := os.Stat(dataPath + ".tmp"); !os.IsNotExist(err) {
		t.Fatalf(".data.tmp should not remain, stat err = %v", err)
	}

	// Verify .data file is NOT valid plaintext JSON (it's encrypted).
	raw, err := os.ReadFile(dataPath)
	if err != nil {
		t.Fatalf("ReadFile(.data) error = %v", err)
	}
	var probe map[string]any
	if json.Unmarshal(raw, &probe) == nil {
		t.Fatal("saved .data should be encrypted, not plain JSON")
	}

	loaded, err := LoadTokenData(configDir)
	if err != nil {
		t.Fatalf("LoadTokenData() error = %v", err)
	}
	if loaded.AccessToken != original.AccessToken || loaded.PersistentCode != original.PersistentCode {
		t.Fatalf("loaded token = %#v, want access/persistent code preserved", loaded)
	}
	if loaded.UserID != original.UserID {
		t.Fatalf("loaded user id = %q, want %q", loaded.UserID, original.UserID)
	}
	if loaded.CorpID != original.CorpID {
		t.Fatalf("loaded corp_id = %q, want %q", loaded.CorpID, original.CorpID)
	}

	if err := DeleteTokenData(configDir); err != nil {
		t.Fatalf("DeleteTokenData() error = %v", err)
	}
	if _, err := LoadTokenData(configDir); err == nil {
		t.Fatal("LoadTokenData() error = nil after delete, want failure")
	}
}

func TestTokenDecryptionFailsWithCorruptedData(t *testing.T) {
	configDir := t.TempDir()
	data := &TokenData{
		AccessToken:  "at_test",
		RefreshToken: "rt_test",
		ExpiresAt:    time.Now().Add(time.Hour),
		RefreshExpAt: time.Now().Add(24 * time.Hour),
	}
	if err := SaveTokenData(configDir, data); err != nil {
		t.Fatalf("SaveTokenData() error = %v", err)
	}

	dataPath := filepath.Join(configDir, secureDataFile)
	raw, err := os.ReadFile(dataPath)
	if err != nil {
		t.Fatalf("ReadFile(.data) error = %v", err)
	}
	raw[len(raw)-1] ^= 0xFF
	if err := os.WriteFile(dataPath, raw, 0o600); err != nil {
		t.Fatalf("WriteFile(.data) error = %v", err)
	}

	if _, err := LoadTokenData(configDir); err == nil {
		t.Fatal("LoadTokenData with corrupted ciphertext should fail")
	}
}

func TestSecureDataExists(t *testing.T) {
	configDir := t.TempDir()
	if SecureDataExists(configDir) {
		t.Fatal("SecureDataExists() should be false before save")
	}

	data := &TokenData{
		AccessToken: "at_test",
		ExpiresAt:   time.Now().Add(time.Hour),
	}
	if err := SaveTokenData(configDir, data); err != nil {
		t.Fatalf("SaveTokenData() error = %v", err)
	}
	if !SecureDataExists(configDir) {
		t.Fatal("SecureDataExists() should be true after save")
	}
}

func TestTokenValidityChecks(t *testing.T) {
	t.Parallel()

	valid := &TokenData{
		AccessToken:  "at_valid",
		RefreshToken: "rt_valid",
		ExpiresAt:    time.Now().Add(2 * time.Hour),
		RefreshExpAt: time.Now().Add(24 * time.Hour),
	}
	if !valid.IsAccessTokenValid() {
		t.Fatal("access token expiring in 2h should be valid")
	}
	if !valid.IsRefreshTokenValid() {
		t.Fatal("refresh token expiring in 24h should be valid")
	}

	expiringSoon := &TokenData{
		AccessToken: "at_soon",
		ExpiresAt:   time.Now().Add(3 * time.Minute),
	}
	if expiringSoon.IsAccessTokenValid() {
		t.Fatal("access token expiring inside 5m buffer should be invalid")
	}

	expiredRefresh := &TokenData{
		RefreshToken: "rt_expired",
		RefreshExpAt: time.Now().Add(-1 * time.Hour),
	}
	if expiredRefresh.IsRefreshTokenValid() {
		t.Fatal("expired refresh token should be invalid")
	}
}
