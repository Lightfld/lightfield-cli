// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/lightfield-cli/internal/mocktest"
	"github.com/stainless-sdks/lightfield-cli/internal/requestflag"
)

func TestContactCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "create",
			"--api-key", "string",
			"--fields", "{system_email: [string], system_name: {firstName: firstName, lastName: lastName}, system_profilePhotoUrl: system_profilePhotoUrl}",
			"--relationships", "{system_account: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(contactCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "create",
			"--api-key", "string",
			"--fields.system-email", "[string]",
			"--fields.system-name", "{firstName: firstName, lastName: lastName}",
			"--fields.system-profile-photo-url", "system_profilePhotoUrl",
			"--relationships.system-account", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  system_email:\n" +
			"    - string\n" +
			"  system_name:\n" +
			"    firstName: firstName\n" +
			"    lastName: lastName\n" +
			"  system_profilePhotoUrl: system_profilePhotoUrl\n" +
			"relationships:\n" +
			"  system_account: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "contact", "create",
			"--api-key", "string",
		)
	})
}

func TestContactRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestContactUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "update",
			"--api-key", "string",
			"--id", "id",
			"--fields", "{system_email: [string], system_name: {firstName: firstName, lastName: lastName}, system_profilePhotoUrl: system_profilePhotoUrl}",
			"--relationships", "{system_account: {add: string, remove: string, replace: string}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(contactUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "update",
			"--api-key", "string",
			"--id", "id",
			"--fields.system-email", "[string]",
			"--fields.system-name", "{firstName: firstName, lastName: lastName}",
			"--fields.system-profile-photo-url", "system_profilePhotoUrl",
			"--relationships.system-account", "{add: string, remove: string, replace: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  system_email:\n" +
			"    - string\n" +
			"  system_name:\n" +
			"    firstName: firstName\n" +
			"    lastName: lastName\n" +
			"  system_profilePhotoUrl: system_profilePhotoUrl\n" +
			"relationships:\n" +
			"  system_account:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "contact", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestContactList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "list",
			"--api-key", "string",
			"--limit", "1",
			"--offset", "0",
		)
	})
}
