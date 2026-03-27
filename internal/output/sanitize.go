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

package output

// SanitizeForTerminal strips control characters and dangerous Unicode markers
// before untrusted text is printed to a terminal or log sink.
func SanitizeForTerminal(text string) string {
	out := make([]rune, 0, len(text))
	for _, r := range text {
		if r == '\n' || r == '\t' {
			out = append(out, r)
			continue
		}
		if r < 0x20 || r == 0x7f {
			continue
		}
		if isDangerousUnicode(r) {
			continue
		}
		out = append(out, r)
	}
	return string(out)
}

func isDangerousUnicode(r rune) bool {
	switch {
	case r >= '\u200B' && r <= '\u200D':
		return true
	case r == '\uFEFF':
		return true
	case r >= '\u202A' && r <= '\u202E':
		return true
	case r >= '\u2028' && r <= '\u2029':
		return true
	case r >= '\u2066' && r <= '\u2069':
		return true
	default:
		return false
	}
}
