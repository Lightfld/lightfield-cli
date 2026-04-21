package interactive

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type RunResult struct {
	Command string
	Output  string
	Err     error
}

type runFinishedMsg struct {
	Result RunResult
}

func BuildArgs(catalog Catalog, state SessionState) ([]string, string, error) {
	resource, action, ok := selectedAction(catalog, state)
	if !ok {
		return nil, "", fmt.Errorf("no action selected")
	}

	args := make([]string, 0)
	appendFieldArgs := func(fields []FieldSpec) {
		for _, field := range fields {
			value := strings.TrimSpace(state.Values[field.Name])
			if value == "" {
				continue
			}
			switch field.Kind {
			case FieldKindBool:
				if value == "true" {
					args = append(args, "--"+field.Name)
				}
			default:
				args = append(args, "--"+field.Name, value)
			}
		}
	}

	appendFieldArgs(catalog.GlobalFields)
	args = append(args, resource.Name, action.Name)
	appendFieldArgs(action.Fields)

	return args, "lightfield " + quoteArgs(args), nil
}

func RunCommandCmd(ctx context.Context, args []string, display string) tea.Cmd {
	return func() tea.Msg {
		executable, err := os.Executable()
		if err != nil {
			return runFinishedMsg{Result: RunResult{Command: display, Err: err}}
		}

		cmd := exec.CommandContext(ctx, executable, args...)
		cmd.Env = os.Environ()
		output, err := cmd.CombinedOutput()
		return runFinishedMsg{
			Result: RunResult{
				Command: display,
				Output:  string(output),
				Err:     err,
			},
		}
	}
}

func quoteArgs(args []string) string {
	quoted := make([]string, 0, len(args))
	for _, arg := range args {
		if arg == "" {
			quoted = append(quoted, "''")
			continue
		}
		if strings.ContainsAny(arg, " \t\n'\"\\$`|&;<>*?[]{}()!") {
			quoted = append(quoted, "'"+strings.ReplaceAll(arg, "'", `'"'"'`)+"'")
			continue
		}
		quoted = append(quoted, arg)
	}
	return strings.Join(quoted, " ")
}

func selectedAction(catalog Catalog, state SessionState) (ResourceSpec, ActionSpec, bool) {
	if state.ResourceIdx < 0 || state.ResourceIdx >= len(catalog.Resources) {
		return ResourceSpec{}, ActionSpec{}, false
	}
	resource := catalog.Resources[state.ResourceIdx]
	if state.ActionIdx < 0 || state.ActionIdx >= len(resource.Actions) {
		return ResourceSpec{}, ActionSpec{}, false
	}
	return resource, resource.Actions[state.ActionIdx], true
}
