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
			"--fields", "{$name: $name, $objectType: account}",
			"--relationships", "{$accounts: string}",
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
			"--fields.object-type", "account",
			"--relationships", "{$accounts: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $name: $name\n" +
			"  $objectType: account\n" +
			"relationships:\n" +
			"  $accounts: string\n")
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
			"--fields", "{$name: $name}",
			"--relationships", "{$accounts: {add: string, remove: string}}",
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
			"--fields.name", "$name",
			"--relationships", "{$accounts: {add: string, remove: string}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $name: $name\n" +
			"relationships:\n" +
			"  $accounts:\n" +
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
