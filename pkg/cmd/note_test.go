// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
)

func TestNoteCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"note", "create",
			"--fields", "{$title: $title, $content: $content}",
			"--relationships", "{$account: string, $opportunity: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $title: $title\n" +
			"  $content: $content\n" +
			"relationships:\n" +
			"  $account: string\n" +
			"  $opportunity: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"note", "create",
		)
	})
}

func TestNoteRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"note", "retrieve",
			"--id", "id",
		)
	})
}

func TestNoteUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"note", "update",
			"--id", "id",
			"--fields", "{$content: $content, $title: $title}",
			"--relationships", "{$account: {add: string}, $opportunity: {add: string}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $content: $content\n" +
			"  $title: $title\n" +
			"relationships:\n" +
			"  $account:\n" +
			"    add: string\n" +
			"  $opportunity:\n" +
			"    add: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"note", "update",
			"--id", "id",
		)
	})
}

func TestNoteList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"note", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}
