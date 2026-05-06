// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
)

func TestObjectCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"object", "create",
			"--entity-slug", "entitySlug",
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
			t, pipeData,
			"--api-key", "string",
			"object", "create",
			"--entity-slug", "entitySlug",
		)
	})
}

func TestObjectRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"object", "retrieve",
			"--entity-slug", "entitySlug",
			"--id", "id",
		)
	})
}

func TestObjectUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"object", "update",
			"--entity-slug", "entitySlug",
			"--id", "id",
			"--fields", "{foo: string}",
			"--relationships", "{foo: {add: string, remove: string, replace: string}}",
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
			"    remove: string\n" +
			"    replace: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"object", "update",
			"--entity-slug", "entitySlug",
			"--id", "id",
		)
	})
}

func TestObjectList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"object", "list",
			"--entity-slug", "entitySlug",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestObjectDefinitions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"object", "definitions",
			"--entity-slug", "entitySlug",
		)
	})
}

func TestObjectListDefinitions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"object", "list-definitions",
		)
	})
}
