// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestOpportunityCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "opportunity", "create",
			"--api-key", "string",
			"--fields", "{$name: $name, $stage: $stage}",
			"--relationships", "{$account: string, $champion: string, $createdBy: string, $evaluator: string, $owner: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(opportunityCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "opportunity", "create",
			"--api-key", "string",
			"--fields.name", "$name",
			"--fields.stage", "$stage",
			"--relationships.account", "string",
			"--relationships.champion", "string",
			"--relationships.created-by", "string",
			"--relationships.evaluator", "string",
			"--relationships.owner", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $name: $name\n" +
			"  $stage: $stage\n" +
			"relationships:\n" +
			"  $account: string\n" +
			"  $champion: string\n" +
			"  $createdBy: string\n" +
			"  $evaluator: string\n" +
			"  $owner: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "opportunity", "create",
			"--api-key", "string",
		)
	})
}

func TestOpportunityRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "opportunity", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestOpportunityUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "opportunity", "update",
			"--api-key", "string",
			"--id", "id",
			"--fields", "{$name: $name, $stage: $stage}",
			"--relationships", "{$champion: {add: string, remove: string, replace: string}, $evaluator: {add: string, remove: string, replace: string}, $owner: {add: string, remove: string, replace: string}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(opportunityUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "opportunity", "update",
			"--api-key", "string",
			"--id", "id",
			"--fields.name", "$name",
			"--fields.stage", "$stage",
			"--relationships.champion", "{add: string, remove: string, replace: string}",
			"--relationships.evaluator", "{add: string, remove: string, replace: string}",
			"--relationships.owner", "{add: string, remove: string, replace: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $name: $name\n" +
			"  $stage: $stage\n" +
			"relationships:\n" +
			"  $champion:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  $evaluator:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  $owner:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "opportunity", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestOpportunityList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "opportunity", "list",
			"--api-key", "string",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestOpportunityDefinitions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "opportunity", "definitions",
			"--api-key", "string",
		)
	})
}
