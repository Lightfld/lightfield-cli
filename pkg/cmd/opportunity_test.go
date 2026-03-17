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
			"--fields", "{system_name: system_name, system_stage: system_stage}",
			"--relationships", "{system_account: string, system_champion: string, system_createdBy: string, system_evaluator: string, system_owner: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(opportunityCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "opportunity", "create",
			"--api-key", "string",
			"--fields.system-name", "system_name",
			"--fields.system-stage", "system_stage",
			"--relationships.system-account", "string",
			"--relationships.system-champion", "string",
			"--relationships.system-created-by", "string",
			"--relationships.system-evaluator", "string",
			"--relationships.system-owner", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  system_name: system_name\n" +
			"  system_stage: system_stage\n" +
			"relationships:\n" +
			"  system_account: string\n" +
			"  system_champion: string\n" +
			"  system_createdBy: string\n" +
			"  system_evaluator: string\n" +
			"  system_owner: string\n")
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
			"--fields", "{system_name: system_name, system_stage: system_stage}",
			"--relationships", "{system_account: {add: string, remove: string, replace: string}, system_champion: {add: string, remove: string, replace: string}, system_createdBy: {add: string, remove: string, replace: string}, system_evaluator: {add: string, remove: string, replace: string}, system_owner: {add: string, remove: string, replace: string}}",
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
			"--fields.system-name", "system_name",
			"--fields.system-stage", "system_stage",
			"--relationships.system-account", "{add: string, remove: string, replace: string}",
			"--relationships.system-champion", "{add: string, remove: string, replace: string}",
			"--relationships.system-created-by", "{add: string, remove: string, replace: string}",
			"--relationships.system-evaluator", "{add: string, remove: string, replace: string}",
			"--relationships.system-owner", "{add: string, remove: string, replace: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  system_name: system_name\n" +
			"  system_stage: system_stage\n" +
			"relationships:\n" +
			"  system_account:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  system_champion:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  system_createdBy:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  system_evaluator:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  system_owner:\n" +
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
