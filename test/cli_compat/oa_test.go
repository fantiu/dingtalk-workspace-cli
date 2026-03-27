package cli_compat_test

import "testing"

// ── oa approval list-pending ───────────────────────────────

func TestOaApprovalListPending_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "oa")
	root := buildRoot()
	err := execCmd(t, root, []string{"oa", "approval", "list-pending"}, map[string]string{
		"start": "2026-03-10T00:00:00+08:00", "end": "2026-03-10T23:59:59+08:00",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "list_pending_approvals")
}

func TestOaApprovalListPending_should_pass_timestamps(t *testing.T) {
	cap := setupTestDeps(t, "oa")
	root := buildRoot()
	_ = execCmd(t, root, []string{"oa", "approval", "list-pending"}, map[string]string{
		"start": "2026-03-10T00:00:00+08:00", "end": "2026-03-10T23:59:59+08:00",
	})
	// Source converts ISO-8601 to float64 millis; MCP uses starTime (not startTime)
	last := cap.last()
	if last == nil {
		t.Fatal("no MCP call captured")
	}
	if _, ok := last.Args["starTime"].(float64); !ok {
		t.Errorf("expected starTime as float64, got %T", last.Args["starTime"])
	}
	if _, ok := last.Args["endTime"].(float64); !ok {
		t.Errorf("expected endTime as float64, got %T", last.Args["endTime"])
	}
}

func TestOaApprovalListPending_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "oa")
	root := buildRoot()
	_ = execCmd(t, root, []string{"oa", "approval", "list-pending"}, map[string]string{
		"start": "2026-03-10T00:00:00+08:00", "end": "2026-03-10T23:59:59+08:00",
	})
	assertCallCount(t, cap, 0)
}

func TestOaApprovalListPending_should_reject_invalid_iso_start(t *testing.T) {
	_ = setupTestDeps(t, "oa")
	root := buildRoot()
	err := execCmd(t, root, []string{"oa", "approval", "list-pending"}, map[string]string{
		"start": "not-a-date", "end": "2026-03-10T23:59:59+08:00",
	})
	if err == nil {
		t.Fatal("expected error for invalid start time")
	}
}
