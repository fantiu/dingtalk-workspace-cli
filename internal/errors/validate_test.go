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

package errors

import "testing"

func TestResourceName(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "valid", input: "search_open_platform_docs"},
		{name: "valid-cjk", input: "审批查询"},
		{name: "leading-digit", input: "1tool", wantErr: true},
		{name: "shell-char", input: "tool;rm", wantErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := ResourceName(tc.input)
			if tc.wantErr && err == nil {
				t.Fatalf("ResourceName(%q) error = nil, want failure", tc.input)
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("ResourceName(%q) error = %v, want nil", tc.input, err)
			}
		})
	}
}

func TestSafePath(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "relative", input: "skills/generated"},
		{name: "absolute", input: "/tmp/dws/export.json"},
		{name: "traversal", input: "../secret", wantErr: true},
		{name: "shell", input: "out;rm -rf /", wantErr: true},
		{name: "null-byte", input: "bad\x00path", wantErr: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := SafePath(tc.input)
			if tc.wantErr && err == nil {
				t.Fatalf("SafePath(%q) error = nil, want failure", tc.input)
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("SafePath(%q) error = %v, want nil", tc.input, err)
			}
		})
	}
}
