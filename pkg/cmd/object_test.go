// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/lightfield-cli/internal/mocktest"
)

func TestObjectCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "object", "create",
			"--api-key", "string",
			"--entity-type", "accounts",
			"--fields", "{foo: string}",
			"--relationships", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  foo: string\n" +
			"relationships:\n" +
			"  foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "object", "create",
			"--api-key", "string",
			"--entity-type", "accounts",
		)
	})
}

func TestObjectRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "object", "retrieve",
			"--api-key", "string",
			"--entity-type", "accounts",
			"--id", "id",
		)
	})
}

func TestObjectUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "object", "update",
			"--api-key", "string",
			"--entity-type", "accounts",
			"--id", "id",
			"--fields", "{foo: string}",
			"--relationships", "{foo: {add: string, remove: string}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  foo: string\n" +
			"relationships:\n" +
			"  foo:\n" +
			"    add: string\n" +
			"    remove: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "object", "update",
			"--api-key", "string",
			"--entity-type", "accounts",
			"--id", "id",
		)
	})
}

func TestObjectList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "object", "list",
			"--api-key", "string",
			"--entity-type", "accounts",
			"--limit", "1",
			"--offset", "0",
		)
	})
}
