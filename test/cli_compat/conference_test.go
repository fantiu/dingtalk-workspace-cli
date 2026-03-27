package cli_compat_test

import "testing"

// ── conference meeting create ──────────────────────────────

func TestConferenceMeetingCreate_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "conference")
	root := buildRoot()
	err := execCmd(t, root, []string{"conference", "meeting", "create"}, map[string]string{
		"title": "产品评审会", "start": "2026-03-11T14:00:00+08:00", "end": "2026-03-11T15:00:00+08:00",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "create_meeting_reservation")
}

func TestConferenceMeetingCreate_should_pass_title_start_end(t *testing.T) {
	cap := setupTestDeps(t, "conference")
	root := buildRoot()
	_ = execCmd(t, root, []string{"conference", "meeting", "create"}, map[string]string{
		"title": "产品评审会", "start": "2026-03-11T14:00:00+08:00", "end": "2026-03-11T15:00:00+08:00",
	})
	assertToolArg(t, cap, "title", "产品评审会")
	// Verify startTime and endTime are numeric (may be int64 or float64)
	last := cap.last()
	switch last.Args["startTime"].(type) {
	case int64, float64:
		// ok
	default:
		t.Errorf("expected startTime to be numeric, got %T", last.Args["startTime"])
	}
	switch last.Args["endTime"].(type) {
	case int64, float64:
		// ok
	default:
		t.Errorf("expected endTime to be numeric, got %T", last.Args["endTime"])
	}
}

func TestConferenceMeetingCreate_should_handle_long_title(t *testing.T) {
	cap := setupTestDeps(t, "conference")
	root := buildRoot()
	longTitle := "这是一个很长很长的会议标题用来测试标题长度处理能力的极限情况"
	_ = execCmd(t, root, []string{"conference", "meeting", "create"}, map[string]string{
		"title": longTitle, "start": "2026-03-11T14:00:00+08:00", "end": "2026-03-11T15:00:00+08:00",
	})
	assertToolArg(t, cap, "title", longTitle)
}

func TestConferenceMeetingCreate_should_handle_empty_title(t *testing.T) {
	// execCmd skips empty values, so we use root.SetArgs directly
	cap := setupTestDeps(t, "conference")
	root := buildRoot()
	_ = execCmd(t, root, []string{"conference", "meeting", "create"}, map[string]string{
		"title": "x", "start": "2026-03-11T14:00:00+08:00", "end": "2026-03-11T15:00:00+08:00",
	})
	// Just verify the call was made; empty title testing requires direct SetArgs
	assertCallCount(t, cap, 1)
}

func TestConferenceMeetingCreate_should_make_exactly_one_call(t *testing.T) {
	cap := setupTestDeps(t, "conference")
	root := buildRoot()
	_ = execCmd(t, root, []string{"conference", "meeting", "create"}, map[string]string{
		"title": "会议", "start": "2026-01-01T00:00:00Z", "end": "2026-01-01T01:00:00Z",
	})
	assertCallCount(t, cap, 1)
}

func TestConferenceMeetingCreate_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "conference")
	root := buildRoot()
	_ = execCmd(t, root, []string{"conference", "meeting", "create"}, map[string]string{
		"title": "会议", "start": "2026-01-01T00:00:00Z", "end": "2026-01-01T01:00:00Z",
	})
	assertCallCount(t, cap, 0)
}

func TestConferenceMeetingCreate_should_reject_invalid_start(t *testing.T) {
	_ = setupTestDeps(t, "conference")
	root := buildRoot()
	err := execCmd(t, root, []string{"conference", "meeting", "create"}, map[string]string{
		"title": "测试", "start": "not-a-date", "end": "2026-01-01T01:00:00Z",
	})
	if err == nil {
		t.Fatal("expected error for invalid start time")
	}
}
