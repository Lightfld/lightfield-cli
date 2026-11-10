// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestTaskCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "create",
			"--fields", "{$status: $status, $title: $title, $description: $description, $dueAt: $dueAt}",
			"--relationships", "{$assignedTo: string, $account: string, $createdBy: string, $opportunity: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(taskCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "create",
			"--fields.status", "$status",
			"--fields.title", "$title",
			"--fields.description", "$description",
			"--fields.due-at", "$dueAt",
			"--relationships", "{$assignedTo: string, $account: string, $createdBy: string, $opportunity: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $status: $status\n" +
			"  $title: $title\n" +
			"  $description: $description\n" +
			"  $dueAt: $dueAt\n" +
			"relationships:\n" +
			"  $assignedTo: string\n" +
			"  $account: string\n" +
			"  $createdBy: string\n" +
			"  $opportunity: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"task", "create",
		)
	})
}

func TestTaskRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "retrieve",
			"--id", "id",
		)
	})
}

func TestTaskUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "update",
			"--id", "id",
			"--fields", "{$description: $description, $dueAt: $dueAt, $status: $status, $title: $title}",
			"--relationships", "{$account: {add: string, remove: string, replace: replace}, $assignedTo: {add: string, remove: string, replace: replace}, $opportunity: {add: string, remove: string, replace: replace}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(taskUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "update",
			"--id", "id",
			"--fields.description", "$description",
			"--fields.due-at", "$dueAt",
			"--fields.status", "$status",
			"--fields.title", "$title",
			"--relationships", "{$account: {add: string, remove: string, replace: replace}, $assignedTo: {add: string, remove: string, replace: replace}, $opportunity: {add: string, remove: string, replace: replace}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $description: $description\n" +
			"  $dueAt: $dueAt\n" +
			"  $status: $status\n" +
			"  $title: $title\n" +
			"relationships:\n" +
			"  $account:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: replace\n" +
			"  $assignedTo:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: replace\n" +
			"  $opportunity:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: replace\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"task", "update",
			"--id", "id",
		)
	})
}

func TestTaskList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "list",
			"--limit", "1",
			"--offset", "0",
			"--status", "active",
		)
	})
}

func TestTaskDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "delete",
			"--id", "id",
			"--body", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(taskDelete)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "delete",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"task", "delete",
			"--id", "id",
		)
	})
}

func TestTaskDefinitions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "definitions",
		)
	})
}

func TestTaskRestore(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "restore",
			"--id", "id",
			"--body", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(taskRestore)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"task", "restore",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"task", "restore",
			"--id", "id",
		)
	})
}
