// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestMergeGetMerge(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "get-merge",
			"--id", "id",
		)
	})
}

func TestMergeMergeAccounts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "merge-accounts",
			"--duplicate-id", "duplicateId",
			"--primary-id", "primaryId",
			"--field-resolutions", "{foo: primary}",
			"--options", "{multiSelectUnion: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mergeMergeAccounts)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "merge-accounts",
			"--duplicate-id", "duplicateId",
			"--primary-id", "primaryId",
			"--field-resolutions", "{foo: primary}",
			"--options.multi-select-union=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"duplicateId: duplicateId\n" +
			"primaryId: primaryId\n" +
			"fieldResolutions:\n" +
			"  foo: primary\n" +
			"options:\n" +
			"  multiSelectUnion: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"merge", "merge-accounts",
		)
	})
}

func TestMergeMergeContacts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "merge-contacts",
			"--duplicate-id", "duplicateId",
			"--primary-id", "primaryId",
			"--field-resolutions", "{foo: primary}",
			"--options", "{multiSelectUnion: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mergeMergeContacts)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "merge-contacts",
			"--duplicate-id", "duplicateId",
			"--primary-id", "primaryId",
			"--field-resolutions", "{foo: primary}",
			"--options.multi-select-union=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"duplicateId: duplicateId\n" +
			"primaryId: primaryId\n" +
			"fieldResolutions:\n" +
			"  foo: primary\n" +
			"options:\n" +
			"  multiSelectUnion: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"merge", "merge-contacts",
		)
	})
}

func TestMergeMergeObjectValues(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "merge-object-values",
			"--entity-slug", "entitySlug",
			"--duplicate-id", "duplicateId",
			"--primary-id", "primaryId",
			"--field-resolutions", "{foo: primary}",
			"--options", "{multiSelectUnion: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mergeMergeObjectValues)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "merge-object-values",
			"--entity-slug", "entitySlug",
			"--duplicate-id", "duplicateId",
			"--primary-id", "primaryId",
			"--field-resolutions", "{foo: primary}",
			"--options.multi-select-union=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"duplicateId: duplicateId\n" +
			"primaryId: primaryId\n" +
			"fieldResolutions:\n" +
			"  foo: primary\n" +
			"options:\n" +
			"  multiSelectUnion: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"merge", "merge-object-values",
			"--entity-slug", "entitySlug",
		)
	})
}

func TestMergeMergeOpportunities(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "merge-opportunities",
			"--duplicate-id", "duplicateId",
			"--primary-id", "primaryId",
			"--field-resolutions", "{foo: primary}",
			"--options", "{multiSelectUnion: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mergeMergeOpportunities)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"merge", "merge-opportunities",
			"--duplicate-id", "duplicateId",
			"--primary-id", "primaryId",
			"--field-resolutions", "{foo: primary}",
			"--options.multi-select-union=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"duplicateId: duplicateId\n" +
			"primaryId: primaryId\n" +
			"fieldResolutions:\n" +
			"  foo: primary\n" +
			"options:\n" +
			"  multiSelectUnion: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"merge", "merge-opportunities",
		)
	})
}
