package cli_compat_test

// 暂未上线，保留测试源码但不执行
// 对应 register.go 中注释掉的 tbCmd

/*

import "testing"

// ── tb project list ────────────────────────────────────────

func TestTbProjectList_should_call_list_projects(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list"}, map[string]string{"name": ""})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "list_projects")
}

func TestTbProjectList_should_include_name_when_provided(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list"}, map[string]string{"name": "Q1"})
	assertToolArg(t, cap, "name", "Q1")
}

func TestTbProjectList_should_omit_name_when_empty(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list"}, map[string]string{"name": ""})
	last := cap.last()
	if _, ok := last.Args["name"]; ok {
		t.Error("expected name to be omitted when empty")
	}
}

func TestTbProjectList_should_handle_chinese_name(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list"}, map[string]string{"name": "产品迭代"})
	assertToolArg(t, cap, "name", "产品迭代")
}

func TestTbProjectList_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list"}, map[string]string{"name": ""})
	assertCallCount(t, cap, 0)
}

// ── tb project list-mine ───────────────────────────────────

func TestTbProjectListMine_should_call_get_user_projects(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list-mine"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "get_user_projects")
}

func TestTbProjectListMine_should_pass_nil_args(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-mine"}, nil)
	assertNilArgs(t, cap)
}

func TestTbProjectListMine_should_make_exactly_one_call(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-mine"}, nil)
	assertCallCount(t, cap, 1)
}

func TestTbProjectListMine_should_succeed(t *testing.T) {
	_ = setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list-mine"}, nil)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestTbProjectListMine_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-mine"}, nil)
	assertCallCount(t, cap, 0)
}

// ── tb project create ──────────────────────────────────────

func TestTbProjectCreate_should_call_create_project(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "create"}, map[string]string{"name": "新项目"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "create_project")
}

func TestTbProjectCreate_should_pass_name(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "create"}, map[string]string{"name": "Q1 产品迭代"})
	assertToolArg(t, cap, "name", "Q1 产品迭代")
}

func TestTbProjectCreate_should_handle_chinese_name(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "create"}, map[string]string{"name": "钉钉开发项目"})
	assertToolArg(t, cap, "name", "钉钉开发项目")
}

func TestTbProjectCreate_should_handle_empty_name(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "create"}, map[string]string{"name": ""})
	assertToolArg(t, cap, "name", "")
}

func TestTbProjectCreate_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "create"}, map[string]string{"name": "test"})
	assertCallCount(t, cap, 0)
}

// ── tb project update ──────────────────────────────────────

func TestTbProjectUpdate_should_call_update_project_info(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "update"}, map[string]string{"id": "PID1", "name": "新名称"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "update_project_info")
}

func TestTbProjectUpdate_should_pass_project_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "update"}, map[string]string{"id": "PID123", "name": ""})
	assertToolArg(t, cap, "projectId", "PID123")
}

func TestTbProjectUpdate_should_include_name_when_provided(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "update"}, map[string]string{"id": "PID1", "name": "Updated"})
	assertToolArg(t, cap, "name", "Updated")
}

func TestTbProjectUpdate_should_omit_name_when_empty(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "update"}, map[string]string{"id": "PID1", "name": ""})
	last := cap.last()
	if _, ok := last.Args["name"]; ok {
		t.Error("expected name to be omitted when empty")
	}
}

func TestTbProjectUpdate_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "update"}, map[string]string{"id": "PID1", "name": ""})
	assertCallCount(t, cap, 0)
}

// ── tb project list-members ────────────────────────────────

func TestTbProjectListMembers_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list-members"}, map[string]string{"id": "PID1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "get_project_member_list")
}

func TestTbProjectListMembers_should_pass_project_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-members"}, map[string]string{"id": "PID999"})
	assertToolArg(t, cap, "projectId", "PID999")
}

func TestTbProjectListMembers_should_make_one_call(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-members"}, map[string]string{"id": "PID1"})
	assertCallCount(t, cap, 1)
}

func TestTbProjectListMembers_should_handle_empty_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-members"}, map[string]string{"id": ""})
	assertToolArg(t, cap, "projectId", "")
}

func TestTbProjectListMembers_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-members"}, map[string]string{"id": "PID1"})
	assertCallCount(t, cap, 0)
}

// ── tb project add-member ──────────────────────────────────

func TestTbProjectAddMember_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "add-member"}, map[string]string{"id": "PID1", "users": "userId1,userId2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "add_project_members")
}

func TestTbProjectAddMember_should_pass_project_id_and_users(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "add-member"}, map[string]string{"id": "PID1", "users": "userId1,userId2"})
	assertToolArg(t, cap, "projectId", "PID1")
	assertToolArg(t, cap, "userIds", "userId1,userId2")
}

func TestTbProjectAddMember_should_handle_single_user(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "add-member"}, map[string]string{"id": "PID1", "users": "userId1"})
	assertToolArg(t, cap, "userIds", "userId1")
}

func TestTbProjectAddMember_should_handle_empty_users(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "add-member"}, map[string]string{"id": "PID1", "users": ""})
	assertToolArg(t, cap, "userIds", "")
}

func TestTbProjectAddMember_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "add-member"}, map[string]string{"id": "PID1", "users": "userId1"})
	assertCallCount(t, cap, 0)
}

// ── tb project list-task-types ─────────────────────────────

func TestTbProjectListTaskTypes_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list-task-types"}, map[string]string{"id": "PID1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "get_project_task_types")
	assertToolArg(t, cap, "projectId", "PID1")
}

func TestTbProjectListTaskTypes_should_handle_empty_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-task-types"}, map[string]string{"id": ""})
	assertToolArg(t, cap, "projectId", "")
}

func TestTbProjectListTaskTypes_should_make_one_call(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-task-types"}, map[string]string{"id": "PID1"})
	assertCallCount(t, cap, 1)
}

func TestTbProjectListTaskTypes_should_succeed(t *testing.T) {
	_ = setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list-task-types"}, map[string]string{"id": "PID1"})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
}

func TestTbProjectListTaskTypes_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-task-types"}, map[string]string{"id": "PID1"})
	assertCallCount(t, cap, 0)
}

// ── tb project list-workflow ───────────────────────────────

func TestTbProjectListWorkflow_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list-workflow"}, map[string]string{"id": "PID1", "keyword": ""})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "search_project_workflow_status")
}

func TestTbProjectListWorkflow_should_include_keyword_when_provided(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-workflow"}, map[string]string{"id": "PID1", "keyword": "进行中"})
	assertToolArg(t, cap, "keyword", "进行中")
}

func TestTbProjectListWorkflow_should_omit_keyword_when_empty(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-workflow"}, map[string]string{"id": "PID1", "keyword": ""})
	last := cap.last()
	if _, ok := last.Args["keyword"]; ok {
		t.Error("expected keyword to be omitted when empty")
	}
}

func TestTbProjectListWorkflow_should_pass_project_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-workflow"}, map[string]string{"id": "PID1", "keyword": ""})
	assertToolArg(t, cap, "projectId", "PID1")
}

func TestTbProjectListWorkflow_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-workflow"}, map[string]string{"id": "PID1", "keyword": ""})
	assertCallCount(t, cap, 0)
}

// ── tb project list-priorities ─────────────────────────────

func TestTbProjectListPriorities_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list-priorities"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "list_org_task_priorities")
}

func TestTbProjectListPriorities_should_pass_nil_args(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-priorities"}, nil)
	assertNilArgs(t, cap)
}

func TestTbProjectListPriorities_should_make_one_call(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-priorities"}, nil)
	assertCallCount(t, cap, 1)
}

func TestTbProjectListPriorities_should_succeed(t *testing.T) {
	_ = setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "project", "list-priorities"}, nil)
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
}

func TestTbProjectListPriorities_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "project", "list-priorities"}, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task create ─────────────────────────────────────────

func TestTbTaskCreate_should_call_create_task(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "task", "create"}, map[string]string{
		"project": "PID1", "title": "开发登录模块", "content": "实现 OAuth 登录", "executor": "",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "create_task")
}

func TestTbTaskCreate_should_return_error_when_content_empty(t *testing.T) {
	_ = setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "task", "create"}, map[string]string{
		"project": "PID1", "title": "test", "content": "", "executor": "",
	})
	if err == nil {
		t.Fatal("expected error for empty content")
	}
}

func TestTbTaskCreate_should_include_executor_when_provided(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "task", "create"}, map[string]string{
		"project": "PID1", "title": "test", "content": "desc", "executor": "exec001",
	})
	assertToolArg(t, cap, "executorId", "exec001")
}

func TestTbTaskCreate_should_omit_executor_when_empty(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "task", "create"}, map[string]string{
		"project": "PID1", "title": "test", "content": "desc", "executor": "",
	})
	last := cap.last()
	if _, ok := last.Args["executorId"]; ok {
		t.Error("expected executorId to be omitted when empty")
	}
}

func TestTbTaskCreate_should_pass_all_required_fields(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "task", "create"}, map[string]string{
		"project": "PID1", "title": "任务标题", "content": "任务描述", "executor": "",
	})
	assertToolArg(t, cap, "projectId", "PID1")
	assertToolArg(t, cap, "title", "任务标题")
	assertToolArg(t, cap, "content", "任务描述")
}

// ── tb task get ────────────────────────────────────────────

func TestTbTaskGet_should_call_get_task_detail(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "get_task_detail")
	assertToolArg(t, cap, "taskId", "TID1")
}

func TestTbTaskGet_should_handle_different_task_ids(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get")
	setParentFlags(cmd, map[string]string{"id": "TID999"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "TID999")
}

func TestTbTaskGet_should_make_one_call(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 1)
}

func TestTbTaskGet_should_handle_empty_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get")
	setParentFlags(cmd, map[string]string{"id": ""})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "")
}

func TestTbTaskGet_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task search ─────────────────────────────────────────

func TestTbTaskSearch_should_call_search_task_ids_by_tql(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "task", "search"}, map[string]string{
		"tql": "isDone = false ORDER BY priority DESC",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "search_task_ids_by_tql")
}

func TestTbTaskSearch_should_pass_tql_query(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "task", "search"}, map[string]string{"tql": "isDone = false"})
	assertToolArg(t, cap, "tql", "isDone = false")
}

func TestTbTaskSearch_should_handle_complex_tql(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	tql := `isDone = false AND priority = "high" ORDER BY created DESC LIMIT 50`
	_ = execCmd(t, root, []string{"tb", "task", "search"}, map[string]string{"tql": tql})
	assertToolArg(t, cap, "tql", tql)
}

func TestTbTaskSearch_should_handle_tql_with_chinese_chars(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	tql := `statusName = "进行中"`
	_ = execCmd(t, root, []string{"tb", "task", "search"}, map[string]string{"tql": tql})
	assertToolArg(t, cap, "tql", tql)
}

func TestTbTaskSearch_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "task", "search"}, map[string]string{"tql": "isDone = false"})
	assertCallCount(t, cap, 0)
}

// ── tb task update-title ───────────────────────────────────

func TestTbTaskUpdateTitle_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-title")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "新标题"})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "update_task_title")
}

func TestTbTaskUpdateTitle_should_return_error_when_title_empty(t *testing.T) {
	_ = setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-title")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": ""})
	err := cmd.RunE(cmd, nil)
	if err == nil {
		t.Fatal("expected error for empty title")
	}
}

func TestTbTaskUpdateTitle_should_pass_task_id_and_content(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-title")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "Updated"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "TID1")
	assertToolArg(t, cap, "content", "Updated")
}

func TestTbTaskUpdateTitle_should_handle_chinese_title(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-title")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "完成开发"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "content", "完成开发")
}

func TestTbTaskUpdateTitle_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-title")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "test"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task update-status ──────────────────────────────────

func TestTbTaskUpdateStatus_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-status")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"status": "已完成"})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "update_task_status")
}

func TestTbTaskUpdateStatus_should_pass_task_id_and_status(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-status")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"status": "进行中"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "TID1")
	assertToolArg(t, cap, "statusName", "进行中")
}

func TestTbTaskUpdateStatus_should_handle_english_status(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-status")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"status": "Done"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "statusName", "Done")
}

func TestTbTaskUpdateStatus_should_handle_empty_status(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-status")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"status": ""})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "statusName", "")
}

func TestTbTaskUpdateStatus_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-status")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"status": "Done"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task update-priority ────────────────────────────────

func TestTbTaskUpdatePriority_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-priority")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"priority": "高"})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "update_task_priority")
}

func TestTbTaskUpdatePriority_should_pass_priority_name(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-priority")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"priority": "紧急"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "priorityName", "紧急")
}

func TestTbTaskUpdatePriority_should_handle_english_priority(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-priority")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"priority": "High"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "priorityName", "High")
}

func TestTbTaskUpdatePriority_should_handle_empty_priority(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-priority")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"priority": ""})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "priorityName", "")
}

func TestTbTaskUpdatePriority_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-priority")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"priority": "High"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task assign ─────────────────────────────────────────

func TestTbTaskAssign_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "assign")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"executor": "exec001"})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "assign_task_assignees")
}

func TestTbTaskAssign_should_pass_task_id_and_executor(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "assign")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"executor": "exec001"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "TID1")
	assertToolArg(t, cap, "executorId", "exec001")
}

func TestTbTaskAssign_should_handle_different_executor(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "assign")
	setParentFlags(cmd, map[string]string{"id": "TID2"})
	setFlags(cmd, map[string]string{"executor": "exec999"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "executorId", "exec999")
}

func TestTbTaskAssign_should_handle_empty_executor(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "assign")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"executor": ""})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "executorId", "")
}

func TestTbTaskAssign_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "assign")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"executor": "exec001"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task update-due ─────────────────────────────────────

func TestTbTaskUpdateDue_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-due")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"date": "2026-03-31T00:00:00Z"})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "set_task_due_date")
}

func TestTbTaskUpdateDue_should_pass_date(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-due")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"date": "2026-12-31"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "dueDate", "2026-12-31")
}

func TestTbTaskUpdateDue_should_pass_iso_date(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-due")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"date": "2026-03-31T23:59:59+08:00"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "dueDate", "2026-03-31T23:59:59+08:00")
}

func TestTbTaskUpdateDue_should_handle_empty_date(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-due")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"date": ""})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "dueDate", "")
}

func TestTbTaskUpdateDue_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "update-due")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"date": "2026-03-31"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task comment ────────────────────────────────────────

func TestTbTaskComment_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "comment")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"content": "做得好"})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "add_task_comment")
}

func TestTbTaskComment_should_pass_task_id_and_content(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "comment")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"content": "评论内容"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "TID1")
	assertToolArg(t, cap, "content", "评论内容")
}

func TestTbTaskComment_should_handle_long_content(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "comment")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	long := "这是一段很长的评论内容用于测试长文本处理能力的边界情况请注意观察"
	setFlags(cmd, map[string]string{"content": long})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "content", long)
}

func TestTbTaskComment_should_handle_empty_content(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "comment")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"content": ""})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "content", "")
}

func TestTbTaskComment_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "comment")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"content": "test"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task add-progress ───────────────────────────────────

func TestTbTaskAddProgress_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "add-progress")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "进展1", "content": "完成50%", "status": ""})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "create_task_progress")
}

func TestTbTaskAddProgress_should_include_status_when_provided(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "add-progress")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "进展1", "content": "有风险", "status": "2"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "status", "2")
}

func TestTbTaskAddProgress_should_omit_status_when_empty(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "add-progress")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "进展1", "content": "desc", "status": ""})
	_ = cmd.RunE(cmd, nil)
	last := cap.last()
	if _, ok := last.Args["status"]; ok {
		t.Error("expected status to be omitted when empty")
	}
}

func TestTbTaskAddProgress_should_pass_all_fields(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "add-progress")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "周报", "content": "完成核心功能", "status": "1"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "TID1")
	assertToolArg(t, cap, "title", "周报")
	assertToolArg(t, cap, "content", "完成核心功能")
	assertToolArg(t, cap, "status", "1")
}

func TestTbTaskAddProgress_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "add-progress")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	setFlags(cmd, map[string]string{"title": "t", "content": "c", "status": ""})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb task get-progress ───────────────────────────────────

func TestTbTaskGetProgress_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get-progress")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	err := cmd.RunE(cmd, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "get_task_progress")
	assertToolArg(t, cap, "taskId", "TID1")
}

func TestTbTaskGetProgress_should_handle_different_task_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get-progress")
	setParentFlags(cmd, map[string]string{"id": "TID999"})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "TID999")
}

func TestTbTaskGetProgress_should_make_one_call(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get-progress")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 1)
}

func TestTbTaskGetProgress_should_handle_empty_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get-progress")
	setParentFlags(cmd, map[string]string{"id": ""})
	_ = cmd.RunE(cmd, nil)
	assertToolArg(t, cap, "taskId", "")
}

func TestTbTaskGetProgress_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	cmd := findCmd(root, "tb", "task", "get-progress")
	setParentFlags(cmd, map[string]string{"id": "TID1"})
	_ = cmd.RunE(cmd, nil)
	assertCallCount(t, cap, 0)
}

// ── tb worktime list ───────────────────────────────────────

func TestTbWorktimeList_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "worktime", "list"}, map[string]string{"task": "TID1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "get_actual_work_hours_by_task_id")
}

func TestTbWorktimeList_should_pass_task_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "list"}, map[string]string{"task": "TID123"})
	assertToolArg(t, cap, "taskId", "TID123")
}

func TestTbWorktimeList_should_handle_empty_task_id(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "list"}, map[string]string{"task": ""})
	assertToolArg(t, cap, "taskId", "")
}

func TestTbWorktimeList_should_make_one_call(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "list"}, map[string]string{"task": "TID1"})
	assertCallCount(t, cap, 1)
}

func TestTbWorktimeList_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "list"}, map[string]string{"task": "TID1"})
	assertCallCount(t, cap, 0)
}

// ── tb worktime create ─────────────────────────────────────

func TestTbWorktimeCreate_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "worktime", "create"}, map[string]string{
		"task": "TID1", "executor": "exec001", "start": "2026-03-01", "end": "2026-03-02", "hours": "3600000",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "create_actual_work_hour_record")
}

func TestTbWorktimeCreate_should_pass_all_required_fields(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "create"}, map[string]string{
		"task": "TID1", "executor": "exec001", "start": "2026-03-01", "end": "2026-03-02", "hours": "3600000",
	})
	assertToolArg(t, cap, "taskId", "TID1")
	assertToolArg(t, cap, "executorId", "exec001")
	assertToolArg(t, cap, "startDate", "2026-03-01")
	assertToolArg(t, cap, "endDate", "2026-03-02")
	assertToolArg(t, cap, "actualHour", "3600000")
}

func TestTbWorktimeCreate_should_handle_large_hours(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "create"}, map[string]string{
		"task": "TID1", "executor": "exec001", "start": "2026-03-01", "end": "2026-03-02", "hours": "999999999",
	})
	assertToolArg(t, cap, "actualHour", "999999999")
}

func TestTbWorktimeCreate_should_handle_empty_fields(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "create"}, map[string]string{
		"task": "", "executor": "", "start": "", "end": "", "hours": "",
	})
	assertToolArg(t, cap, "taskId", "")
}

func TestTbWorktimeCreate_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "create"}, map[string]string{
		"task": "TID1", "executor": "exec001", "start": "2026-03-01", "end": "2026-03-02", "hours": "100",
	})
	assertCallCount(t, cap, 0)
}

// ── tb worktime update ─────────────────────────────────────

func TestTbWorktimeUpdate_should_call_correct_tool(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	err := execCmd(t, root, []string{"tb", "worktime", "update"}, map[string]string{
		"id": "WH001", "executor": "exec001", "date": "2026-03-01",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertToolName(t, cap, "update_actual_work_hour_record")
}

func TestTbWorktimeUpdate_should_pass_all_fields(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "update"}, map[string]string{
		"id": "WH001", "executor": "exec001", "date": "2026-03-15",
	})
	assertToolArg(t, cap, "workHourId", "WH001")
	assertToolArg(t, cap, "executorId", "exec001")
	assertToolArg(t, cap, "date", "2026-03-15")
}

func TestTbWorktimeUpdate_should_handle_different_ids(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "update"}, map[string]string{
		"id": "WH999", "executor": "exec999", "date": "2026-12-31",
	})
	assertToolArg(t, cap, "workHourId", "WH999")
	assertToolArg(t, cap, "executorId", "exec999")
}

func TestTbWorktimeUpdate_should_handle_empty_fields(t *testing.T) {
	cap := setupTestDeps(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "update"}, map[string]string{
		"id": "", "executor": "", "date": "",
	})
	assertToolArg(t, cap, "workHourId", "")
}

func TestTbWorktimeUpdate_should_not_call_when_dry_run(t *testing.T) {
	cap := setupTestDepsWithDryRun(t, "tb")
	root := buildRoot()
	skipIfNotRegistered(t, root, "tb")
	_ = execCmd(t, root, []string{"tb", "worktime", "update"}, map[string]string{
		"id": "WH001", "executor": "exec001", "date": "2026-03-01",
	})
	assertCallCount(t, cap, 0)
}

*/
