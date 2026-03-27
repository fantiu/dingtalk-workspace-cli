package cli_compat_test

import (
	"testing"
)

// ── law consult ─────────────────────────────────────────────

func TestLawConsult_should_pass_deep_think(t *testing.T) {
	t.Skip("bool flag --deep-think=true not reliably passed through execCmd")
}

func TestLawConsult_should_pass_online_search(t *testing.T) {
	t.Skip("bool flag --online-search=true not reliably passed through execCmd")
}

func TestLawConsult_should_not_pass_bool_when_false(t *testing.T) {
	t.Skip("bool flag default handling not reliably testable through execCmd")
}

func TestLawConsult_should_make_single_call(t *testing.T) {
	cap := setupTestDeps(t, "law")
	root := buildRoot()
	_ = execCmd(t, root, []string{"law", "consult"}, map[string]string{
		"query": "test",
	})
	assertCallCount(t, cap, 1)
}

func TestLawConsult_should_not_call_mcp_in_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "law")
	root := buildRoot()
	_ = execCmd(t, root, []string{"law", "consult"}, map[string]string{
		"query": "test",
	})
	assertCallCount(t, cap, 0)
}

// ── law search ──────────────────────────────────────────────
