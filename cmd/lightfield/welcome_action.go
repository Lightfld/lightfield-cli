package main

import (
	"context"
	"os"

	"github.com/Lightfld/lightfield-cli/internal/welcome"
	"github.com/Lightfld/lightfield-cli/pkg/cmd"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"github.com/urfave/cli/v3"
)

func init() {
	cmd.Command.Action = showWelcome
}

func showWelcome(ctx context.Context, c *cli.Command) error {
	if !shouldShowWelcome(os.Stdin, os.Stdout) {
		return cli.ShowAppHelp(c)
	}

	// WithAltScreen isolates the TUI in an alternate buffer so no escape
	// sequences (ANSI colors, OSC 8 hyperlinks) leak into the user's main
	// scrollback on exit.
	//
	// Do NOT add tea.WithMouseCellMotion or tea.WithMouseAllMotion here.
	// If the process is killed while mouse tracking is active, the terminal
	// is left capturing mouse events as garbage escape codes until the user
	// manually runs "reset".
	program := tea.NewProgram(
		welcome.NewModel(c.Version),
		tea.WithAltScreen(),
		tea.WithContext(ctx),
		tea.WithInput(os.Stdin),
		tea.WithOutput(os.Stdout),
	)

	_, err := program.Run()
	return err
}

func shouldShowWelcome(stdin, stdout *os.File) bool {
	if os.Getenv("CI") != "" || os.Getenv("LIGHTFIELD_NO_WELCOME") != "" {
		return false
	}
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	return term.IsTerminal(stdin.Fd()) && term.IsTerminal(stdout.Fd())
}
