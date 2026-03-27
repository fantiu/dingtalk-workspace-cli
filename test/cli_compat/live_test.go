package cli_compat_test

import "testing"

// ── live stream list ───────────────────────────────────────

func TestLiveStreamList_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "live")
	root := buildRoot()
	err := execCmd(t, root, []string{"live", "stream", "list"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "get_my_lives")
}

func TestLiveStreamList_should_pass_nil_args(t *testing.T) {
	cap := setupTestDeps(t, "live")
	root := buildRoot()
	_ = execCmd(t, root, []string{"live", "stream", "list"}, nil)
	assertNilArgs(t, cap)
}

func TestLiveStreamList_should_make_exactly_one_call(t *testing.T) {
	cap := setupTestDeps(t, "live")
	root := buildRoot()
	_ = execCmd(t, root, []string{"live", "stream", "list"}, nil)
	assertCallCount(t, cap, 1)
}

func TestLiveStreamList_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "live")
	root := buildRoot()
	err := execCmd(t, root, []string{"live", "stream", "list"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertCallCount(t, cap, 0)
}

func TestLiveStreamList_should_succeed_without_errors(t *testing.T) {
	_ = setupTestDeps(t, "live")
	root := buildRoot()
	err := execCmd(t, root, []string{"live", "stream", "list"}, nil)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}
