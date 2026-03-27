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
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/config"
)

const (
	lockFileName   = ".data.lock"
	lockRetryDelay = 50 * time.Millisecond
)

// tokenFileLock provides cross-process file locking for token operations.
// It prevents concurrent refresh from multiple CLI processes,
// which can corrupt token data when two processes refresh simultaneously.
//
// Platform support:
//   - Unix/macOS: flock(2) system call
//   - Windows: LockFileEx / UnlockFileEx from kernel32.dll
type tokenFileLock struct {
	path string
	file *os.File
}

// acquireTokenLock acquires an exclusive file lock for token operations.
// It blocks (with timeout) if another process holds the lock.
// The caller MUST call release() when done.
func acquireTokenLock(configDir string) (*tokenFileLock, error) {
	if err := os.MkdirAll(configDir, config.DirPerm); err != nil {
		return nil, fmt.Errorf("creating config dir for lock: %w", err)
	}

	lockPath := filepath.Join(configDir, lockFileName)
	f, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, config.FilePerm)
	if err != nil {
		return nil, fmt.Errorf("opening lock file: %w", err)
	}

	deadline := time.Now().Add(config.LockTimeout)
	for {
		if err := lockFile(f); err == nil {
			return &tokenFileLock{path: lockPath, file: f}, nil
		}

		if time.Now().After(deadline) {
			_ = f.Close()
			return nil, fmt.Errorf("timeout acquiring token lock after %v (another dws process may be running)", config.LockTimeout)
		}

		time.Sleep(lockRetryDelay)
	}
}

// release releases the file lock.
func (l *tokenFileLock) release() {
	if l.file != nil {
		unlockFile(l.file)
		_ = l.file.Close()
		l.file = nil
	}
}
