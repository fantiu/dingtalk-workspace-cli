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

package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/executor"
	"github.com/DingTalk-Real-AI/dingtalk-workspace-cli/internal/ir"
)

func TestBuildFlagSpecsGeneratesOnlySupportedTopLevelFlags(t *testing.T) {
	t.Parallel()

	specs := BuildFlagSpecs(map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title": map[string]any{
				"type":        "string",
				"description": "Document title",
			},
			"notify": map[string]any{
				"type": "boolean",
			},
			"metadata": map[string]any{
				"type": "object",
			},
			"tags": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "string",
				},
			},
		},
	}, map[string]ir.CLIFlagHint{
		"title": {
			Shorthand: "t",
			Alias:     "name",
		},
	})

	if len(specs) != 4 {
		t.Fatalf("BuildFlagSpecs() len = %d, want 4", len(specs))
	}
	if specs[0].PropertyName != "metadata" || specs[1].PropertyName != "notify" || specs[2].PropertyName != "tags" || specs[3].PropertyName != "title" {
		t.Fatalf("BuildFlagSpecs() unexpected order = %#v", specs)
	}
	if specs[0].Kind != "json" {
		t.Fatalf("BuildFlagSpecs() metadata kind = %q, want json", specs[0].Kind)
	}
	if specs[3].Alias != "name" || specs[3].Shorthand != "t" {
		t.Fatalf("BuildFlagSpecs() title hints = %#v, want alias=name shorthand=t", specs[3])
	}
}

func TestFixtureLoaderLoadsCatalog(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	fixturePath := filepath.Join(dir, "catalog.json")
	data := []byte(`{"products":[{"id":"doc","display_name":"文档","server_key":"doc-key","endpoint":"https://example.com/server/doc","tools":[{"rpc_name":"create_document","canonical_path":"doc.create_document"}]}]}`)
	if err := os.WriteFile(fixturePath, data, 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	catalog, err := FixtureLoader{Path: fixturePath}.Load(context.Background())
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if len(catalog.Products) != 1 || catalog.Products[0].ID != "doc" {
		t.Fatalf("Load() catalog = %#v, want doc product", catalog)
	}
}

func TestSchemaPayloadFindsTool(t *testing.T) {
	t.Parallel()

	payload, err := schemaPayload(ir.Catalog{
		Products: []ir.CanonicalProduct{
			{
				ID: "doc",
				Tools: []ir.ToolDescriptor{
					{
						RPCName:       "create_document",
						CanonicalPath: "doc.create_document",
						InputSchema: map[string]any{
							"type": "object",
							"required": []any{
								"title",
							},
						},
					},
				},
			},
		},
	}, []string{"doc.create_document"})
	if err != nil {
		t.Fatalf("schemaPayload() error = %v", err)
	}
	if payload["kind"] != "schema" {
		t.Fatalf("schemaPayload() kind = %#v, want schema", payload["kind"])
	}
}

func TestNewMCPCommandReturnsLoaderErrorForInvocations(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("fixture missing")
	cmd := NewMCPCommand(context.Background(), errorLoader{err: wantErr}, executor.EchoRunner{})
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"doc", "create_document"})

	err := cmd.Execute()
	if !errors.Is(err, wantErr) {
		t.Fatalf("Execute() error = %v, want %v", err, wantErr)
	}
}

func TestNewMCPCommandSkipsProductsMarkedSkip(t *testing.T) {
	t.Parallel()

	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "doc",
					CLI: &ir.ProductCLIMetadata{
						Skip: true,
					},
				},
				{
					ID: "drive",
				},
			},
		},
	}, executor.EchoRunner{})

	if got := cmd.Commands(); len(got) != 1 || got[0].Name() != "drive" {
		t.Fatalf("mcp commands = %#v, want only drive", got)
	}
}

func TestProductCommandUsesCLICommandAlias(t *testing.T) {
	t.Parallel()

	runner := &captureRunner{}
	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "doc",
					CLI: &ir.ProductCLIMetadata{
						Command: "documents",
					},
					Tools: []ir.ToolDescriptor{
						{RPCName: "create_document"},
					},
				},
			},
		},
	}, runner)

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"documents", "create_document"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if runner.last.CanonicalProduct != "doc" {
		t.Fatalf("runner.last.CanonicalProduct = %q, want doc", runner.last.CanonicalProduct)
	}
}

func TestNewMCPCommandAddsGroupedRoutesFromCLIMetadata(t *testing.T) {
	t.Parallel()

	runner := &captureRunner{}
	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "doc",
					CLI: &ir.ProductCLIMetadata{
						Command: "documents",
						Group:   "office/collab",
					},
					Tools: []ir.ToolDescriptor{
						{RPCName: "create_document"},
					},
				},
			},
		},
	}, runner)

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"office", "collab", "documents", "create_document"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if runner.last.CanonicalProduct != "doc" {
		t.Fatalf("runner.last.CanonicalProduct = %q, want doc", runner.last.CanonicalProduct)
	}
}

func TestToolCommandUsesCLINameAndFlagHints(t *testing.T) {
	t.Parallel()

	runner := &captureRunner{}
	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "doc",
					Tools: []ir.ToolDescriptor{
						{
							RPCName: "create_document",
							CLIName: "create",
							InputSchema: map[string]any{
								"type": "object",
								"properties": map[string]any{
									"title": map[string]any{"type": "string"},
								},
							},
							FlagHints: map[string]ir.CLIFlagHint{
								"title": {
									Alias:     "name",
									Shorthand: "t",
								},
							},
						},
					},
				},
			},
		},
	}, runner)
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"doc", "create", "--name", "hello"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if runner.last.Tool != "create_document" {
		t.Fatalf("runner.last.Tool = %q, want create_document", runner.last.Tool)
	}
	if runner.last.Params["title"] != "hello" {
		t.Fatalf("runner.last.Params[title] = %#v, want hello", runner.last.Params["title"])
	}
}

func TestToolCommandValidatesInputSchemaBeforeRun(t *testing.T) {
	t.Parallel()

	runner := &captureRunner{}
	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "doc",
					Tools: []ir.ToolDescriptor{
						{
							RPCName: "create_document",
							InputSchema: map[string]any{
								"type": "object",
								"required": []any{
									"title",
								},
								"properties": map[string]any{
									"title": map[string]any{"type": "string"},
								},
							},
						},
					},
				},
			},
		},
	}, runner)

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"doc", "create_document", "--params", `{"title":"ok","unknown":"x"}`})

	err := cmd.Execute()
	if err == nil {
		t.Fatal("Execute() error = nil, want schema validation error")
	}
	if !strings.Contains(err.Error(), "$.unknown is not allowed") {
		t.Fatalf("Execute() error = %v, want unknown-property validation", err)
	}
	if runner.called != 0 {
		t.Fatalf("runner called = %d, want 0", runner.called)
	}
}

func TestToolCommandSupportsDryRunWithoutSensitiveConfirmation(t *testing.T) {
	t.Parallel()

	runner := &captureRunner{}
	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "doc",
					Tools: []ir.ToolDescriptor{
						{
							RPCName:   "create_document",
							Sensitive: true,
							InputSchema: map[string]any{
								"type": "object",
							},
						},
					},
				},
			},
		},
	}, runner)

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.PersistentFlags().Bool("dry-run", false, "Preview the operation without executing it")
	cmd.SetArgs([]string{"doc", "create_document", "--dry-run"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !runner.last.DryRun {
		t.Fatalf("runner.last.DryRun = %t, want true", runner.last.DryRun)
	}
	if runner.called != 1 {
		t.Fatalf("runner called = %d, want 1", runner.called)
	}
}

func TestDeprecatedLifecycleAddsWarningToResult(t *testing.T) {
	t.Parallel()

	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "legacy-doc",
					Lifecycle: &ir.LifecycleInfo{
						DeprecatedBy:    9527,
						DeprecationDate: "2026-04-01T00:00:00Z",
						MigrationURL:    "https://example.com/migration",
					},
					Tools: []ir.ToolDescriptor{
						{
							RPCName: "search_documents",
						},
					},
				},
			},
		},
	}, executor.EchoRunner{})

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errOut)
	cmd.SetArgs([]string{"legacy-doc", "search_documents"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	var payload struct {
		Response map[string]any `json:"response"`
	}
	if err := json.Unmarshal(out.Bytes(), &payload); err != nil {
		t.Fatalf("json.Unmarshal() error = %v\noutput:\n%s", err, out.String())
	}
	if payload.Response["warning"] == "" {
		t.Fatalf("warning is empty, payload=%#v", payload.Response)
	}
	warning, _ := payload.Response["warning"].(string)
	if !strings.Contains(warning, "deprecated_by_mcpId=9527") {
		t.Fatalf("warning = %q, want deprecated_by_mcpId=9527", warning)
	}
}

func TestDeprecatedLifecyclePrintsWarningToStderr(t *testing.T) {
	t.Parallel()

	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "legacy-doc",
					Lifecycle: &ir.LifecycleInfo{
						DeprecatedBy: 9527,
						MigrationURL: "https://example.com/migration",
					},
					Tools: []ir.ToolDescriptor{
						{RPCName: "search_documents"},
					},
				},
			},
		},
	}, executor.EchoRunner{})

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errOut)
	cmd.SetArgs([]string{"legacy-doc", "search_documents"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	stderr := errOut.String()
	if !strings.Contains(stderr, "warning: product legacy-doc is deprecated") {
		t.Fatalf("stderr = %q, want deprecation warning", stderr)
	}
	if !strings.Contains(stderr, "migration=https://example.com/migration") {
		t.Fatalf("stderr = %q, want migration hint", stderr)
	}
}

func TestSensitiveToolConfirmationWorksWithoutYesFlag(t *testing.T) {
	t.Parallel()

	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "doc",
					Tools: []ir.ToolDescriptor{
						{
							RPCName:   "create_document",
							CLIName:   "create-document",
							Sensitive: true,
						},
					},
				},
			},
		},
	}, executor.EchoRunner{})

	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetIn(strings.NewReader("yes\n"))
	cmd.SetArgs([]string{"doc", "create-document"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
}

func TestLegacyCandidateLifecycleAddsWarningToResult(t *testing.T) {
	t.Parallel()

	cmd := NewMCPCommand(context.Background(), StaticLoader{
		Catalog: ir.Catalog{
			Products: []ir.CanonicalProduct{
				{
					ID: "legacy-candidate",
					Lifecycle: &ir.LifecycleInfo{
						DeprecatedCandidate: true,
					},
					Tools: []ir.ToolDescriptor{
						{RPCName: "search_documents"},
					},
				},
			},
		},
	}, executor.EchoRunner{})

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errOut)
	cmd.SetArgs([]string{"legacy-candidate", "search_documents"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	var payload struct {
		Response map[string]any `json:"response"`
	}
	if err := json.Unmarshal(out.Bytes(), &payload); err != nil {
		t.Fatalf("json.Unmarshal() error = %v\noutput:\n%s", err, out.String())
	}
	warning, _ := payload.Response["warning"].(string)
	if !strings.Contains(warning, "legacy candidate") {
		t.Fatalf("warning = %q, want legacy candidate marker", warning)
	}
}

type errorLoader struct {
	err error
}

func (l errorLoader) Load(context.Context) (ir.Catalog, error) {
	return ir.Catalog{}, l.err
}

type captureRunner struct {
	last   executor.Invocation
	called int
}

func (r *captureRunner) Run(_ context.Context, invocation executor.Invocation) (executor.Result, error) {
	r.last = invocation
	r.called++
	return executor.Result{Invocation: invocation}, nil
}
