package cli_compat_test

import (
	"testing"
)

// ── create ──────────────────────────────────────────────────

func TestAiappCreate_should_call_create_ai_app(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	err := execCmd(t, root, []string{"aiapp", "create"}, map[string]string{
		"prompt": "创建一个天气查询应用",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "create_ai_app")
	assertToolArg(t, cap, "prompt", "创建一个天气查询应用")
}

func TestAiappCreate_should_include_skills_when_provided(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "create"}, map[string]string{
		"prompt": "翻译应用", "skills": "s1,s2",
	})
	assertToolArg(t, cap, "officialSkillUids", []string{"s1", "s2"})
}

func TestAiappCreate_should_not_include_skills_when_empty(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "create"}, map[string]string{
		"prompt": "P", "skills": "",
	})
	assertArgNotPresent(t, cap, "officialSkillUids")
}

func TestAiappCreate_should_handle_long_prompt(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	longPrompt := "创建一个能够处理多语言翻译、支持语音输入、并且具备历史记录功能的AI翻译助手应用"
	_ = execCmd(t, root, []string{"aiapp", "create"}, map[string]string{"prompt": longPrompt})
	assertToolArg(t, cap, "prompt", longPrompt)
}

func TestAiappCreate_should_make_single_call(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "create"}, map[string]string{"prompt": "P"})
	assertCallCount(t, cap, 1)
}

func TestAiappCreate_should_handle_single_skill(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "create"}, map[string]string{
		"prompt": "P", "skills": "only_skill",
	})
	assertToolArg(t, cap, "officialSkillUids", []string{"only_skill"})
}

func TestAiappCreate_should_not_call_mcp_in_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "create"}, map[string]string{"prompt": "P"})
	assertCallCount(t, cap, 0)
}

// ── query ───────────────────────────────────────────────────

func TestAiappQuery_should_call_query_ai_app(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	err := execCmd(t, root, []string{"aiapp", "query"}, map[string]string{
		"task-id": "TASK001",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "query_ai_app")
	assertToolArg(t, cap, "taskId", "TASK001")
}

func TestAiappQuery_should_handle_different_task_ids(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "query"}, map[string]string{"task-id": "T_XYZ"})
	assertToolArg(t, cap, "taskId", "T_XYZ")
}

func TestAiappQuery_should_make_single_call(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "query"}, map[string]string{"task-id": "T1"})
	assertCallCount(t, cap, 1)
}

func TestAiappQuery_should_return_no_error(t *testing.T) {
	_ = setupTestDeps(t, "aiapp")
	root := buildRoot()
	err := execCmd(t, root, []string{"aiapp", "query"}, map[string]string{"task-id": "T1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAiappQuery_should_not_call_mcp_in_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "query"}, map[string]string{"task-id": "T1"})
	assertCallCount(t, cap, 0)
}

// ── modify ──────────────────────────────────────────────────

func TestAiappModify_should_call_modify_ai_app(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	err := execCmd(t, root, []string{"aiapp", "modify"}, map[string]string{
		"prompt":    "改为翻译应用",
		"thread-id": "THREAD001",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "modify_ai_app")
	assertToolArg(t, cap, "prompt", "改为翻译应用")
	assertToolArg(t, cap, "threadId", "THREAD001")
}

func TestAiappModify_should_include_skills_when_provided(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "modify"}, map[string]string{
		"prompt": "P", "thread-id": "TH1", "skills": "s1,s2,s3",
	})
	assertToolArg(t, cap, "officialSkillUids", []string{"s1", "s2", "s3"})
}

func TestAiappModify_should_not_include_skills_when_empty(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "modify"}, map[string]string{
		"prompt": "P", "thread-id": "TH1", "skills": "",
	})
	assertArgNotPresent(t, cap, "officialSkillUids")
}

func TestAiappModify_should_pass_thread_id(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "modify"}, map[string]string{
		"prompt": "P", "thread-id": "TH_XYZ",
	})
	assertToolArg(t, cap, "threadId", "TH_XYZ")
}

func TestAiappModify_should_make_single_call(t *testing.T) {
	cap := setupTestDeps(t, "aiapp")
	root := buildRoot()
	_ = execCmd(t, root, []string{"aiapp", "modify"}, map[string]string{
		"prompt": "P", "thread-id": "TH1",
	})
	assertCallCount(t, cap, 1)
}
