// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestContactCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "create",
			"--api-key", "string",
			"--fields", "{$email: [string], $name: {firstName: firstName, lastName: lastName}, $profilePhotoUrl: $profilePhotoUrl}",
			"--relationships", "{$accounts: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(contactCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "create",
			"--api-key", "string",
			"--fields.email", "[string]",
			"--fields.name", "{firstName: firstName, lastName: lastName}",
			"--fields.profile-photo-url", "$profilePhotoUrl",
			"--relationships.accounts", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $email:\n" +
			"    - string\n" +
			"  $name:\n" +
			"    firstName: firstName\n" +
			"    lastName: lastName\n" +
			"  $profilePhotoUrl: $profilePhotoUrl\n" +
			"relationships:\n" +
			"  $accounts: string\n")
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
			"--fields", "{$email: [string], $name: {firstName: firstName, lastName: lastName}, $profilePhotoUrl: $profilePhotoUrl}",
			"--relationships", "{$accounts: {add: string, remove: string, replace: string}}",
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
			"--fields.email", "[string]",
			"--fields.name", "{firstName: firstName, lastName: lastName}",
			"--fields.profile-photo-url", "$profilePhotoUrl",
			"--relationships.accounts", "{add: string, remove: string, replace: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $email:\n" +
			"    - string\n" +
			"  $name:\n" +
			"    firstName: firstName\n" +
			"    lastName: lastName\n" +
			"  $profilePhotoUrl: $profilePhotoUrl\n" +
			"relationships:\n" +
			"  $accounts:\n" +
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

func TestContactDefinitions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "contact", "definitions",
			"--api-key", "string",
		)
	})
}
