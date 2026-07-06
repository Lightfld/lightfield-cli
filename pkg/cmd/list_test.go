// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestListCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "create",
			"--fields", "{$name: $name, $objectType: $objectType, $description: $description, $kind: target}",
			"--relationships", "{$accounts: string, $contacts: string, $opportunities: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "create",
			"--fields.name", "$name",
			"--fields.object-type", "$objectType",
			"--fields.description", "$description",
			"--fields.kind", "target",
			"--relationships.accounts", "string",
			"--relationships.contacts", "string",
			"--relationships.opportunities", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $name: $name\n" +
			"  $objectType: $objectType\n" +
			"  $description: $description\n" +
			"  $kind: target\n" +
			"relationships:\n" +
			"  $accounts: string\n" +
			"  $contacts: string\n" +
			"  $opportunities: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"list", "create",
		)
	})
}

func TestListRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "retrieve",
			"--id", "id",
		)
	})
}

func TestListUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "update",
			"--id", "id",
			"--fields", "{$description: $description, $kind: target, $name: $name}",
			"--relationships", "{$accounts: {add: string, remove: string}, $contacts: {add: string, remove: string}, $opportunities: {add: string, remove: string}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "update",
			"--id", "id",
			"--fields.description", "$description",
			"--fields.kind", "target",
			"--fields.name", "$name",
			"--relationships.accounts", "{add: string, remove: string}",
			"--relationships.contacts", "{add: string, remove: string}",
			"--relationships.opportunities", "{add: string, remove: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $description: $description\n" +
			"  $kind: target\n" +
			"  $name: $name\n" +
			"relationships:\n" +
			"  $accounts:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"  $contacts:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"  $opportunities:\n" +
			"    add: string\n" +
			"    remove: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"list", "update",
			"--id", "id",
		)
	})
}

func TestListList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestListDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "delete",
			"--id", "id",
			"--body", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listDelete)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "delete",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"list", "delete",
			"--id", "id",
		)
	})
}

func TestListListAccounts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "list-accounts",
			"--list-id", "listId",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestListListContacts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "list-contacts",
			"--list-id", "listId",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestListListOpportunities(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"list", "list-opportunities",
			"--list-id", "listId",
			"--limit", "1",
			"--offset", "0",
		)
	})
}
