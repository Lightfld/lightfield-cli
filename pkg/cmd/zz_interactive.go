package cmd

import "github.com/Lightfld/lightfield-cli/internal/interactive"

func init() {
	if Command != nil {
		Command.Action = interactive.Run
	}
}
