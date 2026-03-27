package helpers

import (
	"bytes"
	"context"
	"testing"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/executor"
)

type captureRunner struct {
	last executor.Invocation
}

func (r *captureRunner) Run(_ context.Context, invocation executor.Invocation) (executor.Result, error) {
	r.last = invocation
	return executor.Result{Invocation: invocation}, nil
}

func TestChatMessageSendIgnoresLegacyRealBuildModeEnv(t *testing.T) {
	t.Setenv("DWS_"+"BUILD_MODE", "real")

	runner := &captureRunner{}
	cmd := newChatMessageSendCommand(runner)

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"--user", "user-001", "hello"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v\noutput:\n%s", err, out.String())
	}

	if got := runner.last.Params["clawType"]; got != "default" {
		t.Fatalf("clawType = %#v, want default", got)
	}
}
