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

import (
	stderrors "errors"
	"path/filepath"
	"strings"
	"unicode"
)

var (
	ErrInvalidResourceName = stderrors.New("invalid resource name")
	ErrUnsafePath          = stderrors.New("unsafe path detected")
)

func ResourceName(name string) error {
	if name == "" {
		return stderrors.New("resource name cannot be empty")
	}
	if len([]rune(name)) > 128 {
		return stderrors.New("resource name too long (max 128 characters)")
	}
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' {
			continue
		}
		return ErrInvalidResourceName
	}
	if unicode.IsDigit([]rune(name)[0]) {
		return ErrInvalidResourceName
	}
	return nil
}

func SafePath(path string) error {
	if path == "" {
		return stderrors.New("path cannot be empty")
	}

	cleanPath := filepath.Clean(path)
	if strings.Contains(cleanPath, "..") {
		return ErrUnsafePath
	}
	if strings.ContainsRune(path, '\x00') {
		return ErrUnsafePath
	}

	lowerPath := strings.ToLower(path)
	for _, pattern := range []string{
		"..",
		"~",
		"$(",
		"`",
		"|",
		";",
		"&",
		"<",
		">",
		"\n",
		"\r",
	} {
		if strings.Contains(lowerPath, pattern) {
			return ErrUnsafePath
		}
	}
	return nil
}
