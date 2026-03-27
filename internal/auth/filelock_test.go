package auth

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestAcquireTokenLock_Basic(t *testing.T) {
	t.Parallel()

	configDir := t.TempDir()

	lock, err := acquireTokenLock(configDir)
	if err != nil {
		t.Fatalf("acquireTokenLock() error = %v", err)
	}

	// Lock file should exist while held.
	lockPath := filepath.Join(configDir, lockFileName)
	if _, err := os.Stat(lockPath); err != nil {
		t.Fatalf("lock file should exist while held, stat error = %v", err)
	}

	lock.release()

	// Lock file may still exist on disk after release (flock does not remove
	// the file), but we should be able to re-acquire it.
	lock2, err := acquireTokenLock(configDir)
	if err != nil {
		t.Fatalf("re-acquire after release error = %v", err)
	}
	lock2.release()
}

func TestAcquireTokenLock_CreatesDirectory(t *testing.T) {
	t.Parallel()

	base := t.TempDir()
	configDir := filepath.Join(base, "a", "b", "c")

	lock, err := acquireTokenLock(configDir)
	if err != nil {
		t.Fatalf("acquireTokenLock() error = %v", err)
	}
	defer lock.release()

	info, err := os.Stat(configDir)
	if err != nil {
		t.Fatalf("Stat(configDir) error = %v", err)
	}
	if !info.IsDir() {
		t.Fatal("configDir should be a directory")
	}
}

func TestAcquireTokenLock_DoubleRelease(t *testing.T) {
	t.Parallel()

	configDir := t.TempDir()

	lock, err := acquireTokenLock(configDir)
	if err != nil {
		t.Fatalf("acquireTokenLock() error = %v", err)
	}

	// First release should work fine.
	lock.release()

	// Second release should not panic.
	lock.release()
}

func TestAcquireTokenLock_Contention(t *testing.T) {
	t.Parallel()

	configDir := t.TempDir()

	// Goroutine 1 acquires the lock first.
	lock1, err := acquireTokenLock(configDir)
	if err != nil {
		t.Fatalf("acquireTokenLock() g1 error = %v", err)
	}

	acquired := make(chan struct{})
	var g2Err error
	var wg sync.WaitGroup
	wg.Add(1)

	// Goroutine 2 tries to acquire — should block until g1 releases.
	go func() {
		defer wg.Done()
		lock2, err := acquireTokenLock(configDir)
		if err != nil {
			g2Err = err
			close(acquired)
			return
		}
		close(acquired)
		lock2.release()
	}()

	// Give goroutine 2 a moment to start blocking.
	time.Sleep(100 * time.Millisecond)

	// Verify goroutine 2 has not acquired yet.
	select {
	case <-acquired:
		t.Fatal("goroutine 2 should not have acquired the lock while goroutine 1 holds it")
	default:
		// Expected: goroutine 2 is still waiting.
	}

	// Release lock1 so goroutine 2 can proceed.
	lock1.release()

	// Wait for goroutine 2 to finish.
	wg.Wait()

	if g2Err != nil {
		t.Fatalf("acquireTokenLock() g2 error = %v", g2Err)
	}
}

func TestAcquireTokenLock_LockFilePermissions(t *testing.T) {
	t.Parallel()

	configDir := t.TempDir()

	lock, err := acquireTokenLock(configDir)
	if err != nil {
		t.Fatalf("acquireTokenLock() error = %v", err)
	}
	defer lock.release()

	lockPath := filepath.Join(configDir, lockFileName)
	info, err := os.Stat(lockPath)
	if err != nil {
		t.Fatalf("Stat(lock file) error = %v", err)
	}
	perm := info.Mode().Perm()
	if perm != 0o600 {
		t.Fatalf("lock file permissions = %o, want 0600", perm)
	}
}
