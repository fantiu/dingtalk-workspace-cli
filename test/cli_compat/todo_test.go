package cli_compat_test

import (
	"testing"
)

// ── task create ─────────────────────────────────────────────

func TestTodoTaskCreate_should_call_create_todo_task(t *testing.T) {
	cap := setupTestDeps(t, "todo")
	root := buildRoot()
	err := execCmd(t, root, []string{"todo", "task", "create"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "create_todo_task")
}
