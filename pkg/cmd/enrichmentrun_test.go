// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
)

func TestEnrichmentRunEnrich(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enrichment-run", "enrich",
			"--entity-slug", "contacts",
			"--entity-id", "entityId",
			"--field", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"enrichment-run", "enrich",
			"--entity-slug", "contacts",
			"--entity-id", "entityId",
		)
	})
}

func TestEnrichmentRunGetEnrichmentRun(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"enrichment-run", "get-enrichment-run",
			"--run-id", "runId",
		)
	})
}
