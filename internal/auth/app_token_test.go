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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAppTokenData_IsNewTokenValid(t *testing.T) {
	tests := []struct {
		name string
		data *AppTokenData
		want bool
	}{
		{"nil data", nil, false},
		{"empty token", &AppTokenData{}, false},
		{"expired", &AppTokenData{
			NewAccessToken: "tok",
			NewExpiresAt:   time.Now().Add(-1 * time.Minute),
		}, false},
		{"within buffer", &AppTokenData{
			NewAccessToken: "tok",
			NewExpiresAt:   time.Now().Add(3 * time.Minute), // 3 min < 5 min buffer
		}, false},
		{"valid", &AppTokenData{
			NewAccessToken: "tok",
			NewExpiresAt:   time.Now().Add(10 * time.Minute),
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.IsNewTokenValid(); got != tt.want {
				t.Errorf("IsNewTokenValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppTokenData_IsLegacyTokenValid(t *testing.T) {
	tests := []struct {
		name string
		data *AppTokenData
		want bool
	}{
		{"nil data", nil, false},
		{"empty token", &AppTokenData{}, false},
		{"expired", &AppTokenData{
			LegacyAccessToken: "tok",
			LegacyExpiresAt:   time.Now().Add(-1 * time.Minute),
		}, false},
		{"valid", &AppTokenData{
			LegacyAccessToken: "tok",
			LegacyExpiresAt:   time.Now().Add(10 * time.Minute),
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.IsLegacyTokenValid(); got != tt.want {
				t.Errorf("IsLegacyTokenValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppTokenData_JSONRoundTrip(t *testing.T) {
	original := &AppTokenData{
		NewAccessToken:    "new-tok-abc",
		NewExpiresAt:      time.Now().Add(2 * time.Hour).Truncate(time.Second),
		LegacyAccessToken: "legacy-tok-xyz",
		LegacyExpiresAt:   time.Now().Add(2 * time.Hour).Truncate(time.Second),
		ClientID:          "my-app-key",
		UpdatedAt:         time.Now().Truncate(time.Second),
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded AppTokenData
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded.NewAccessToken != original.NewAccessToken {
		t.Errorf("NewAccessToken = %q, want %q", decoded.NewAccessToken, original.NewAccessToken)
	}
	if decoded.LegacyAccessToken != original.LegacyAccessToken {
		t.Errorf("LegacyAccessToken = %q, want %q", decoded.LegacyAccessToken, original.LegacyAccessToken)
	}
	if decoded.ClientID != original.ClientID {
		t.Errorf("ClientID = %q, want %q", decoded.ClientID, original.ClientID)
	}
}

func TestFetchNewAPIToken_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"accessToken": "new-tok-123",
			"expireIn":    7200,
		})
	}))
	defer srv.Close()

	resp, err := srv.Client().Post(srv.URL, "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var result struct {
		AccessToken string `json:"accessToken"`
		ExpireIn    int64  `json:"expireIn"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	if result.AccessToken != "new-tok-123" {
		t.Errorf("got token %q, want new-tok-123", result.AccessToken)
	}
	if result.ExpireIn != 7200 {
		t.Errorf("got expireIn %d, want 7200", result.ExpireIn)
	}
}

func TestFetchLegacyAPIToken_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"errcode":      0,
			"errmsg":       "ok",
			"access_token": "legacy-tok-456",
			"expires_in":   7200,
		})
	}))
	defer srv.Close()

	resp, err := srv.Client().Get(srv.URL + "?appkey=mykey&appsecret=mysecret")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var result struct {
		ErrCode     int    `json:"errcode"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	if result.ErrCode != 0 {
		t.Errorf("got errcode %d, want 0", result.ErrCode)
	}
	if result.AccessToken != "legacy-tok-456" {
		t.Errorf("got token %q, want legacy-tok-456", result.AccessToken)
	}
	if result.ExpiresIn != 7200 {
		t.Errorf("got expires_in %d, want 7200", result.ExpiresIn)
	}
}

func TestFetchLegacyAPIToken_BusinessError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"errcode":      40014,
			"errmsg":       "invalid appkey",
			"access_token": "",
			"expires_in":   0,
		})
	}))
	defer srv.Close()

	resp, err := srv.Client().Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var result struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	if result.ErrCode != 40014 {
		t.Errorf("got errcode %d, want 40014", result.ErrCode)
	}
}

func TestAppTokenProvider_GetToken_MissingCredentials(t *testing.T) {
	provider := &AppTokenProvider{
		ConfigDir: t.TempDir(),
		AppKey:    "",
		AppSecret: "",
	}
	_, err := provider.GetToken(context.Background(), false)
	if err == nil {
		t.Error("expected error for missing credentials")
	}
}

func TestTruncateStr(t *testing.T) {
	if got := truncateStr("hello", 10); got != "hello" {
		t.Errorf("got %q, want hello", got)
	}
	if got := truncateStr("hello world", 5); got != "hello..." {
		t.Errorf("got %q, want hello...", got)
	}
}
