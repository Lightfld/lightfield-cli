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
			"--relationships", "{$account: {add: string, remove: string, replace: string}, $assignedTo: {add: string, remove: string, replace: string}, $opportunity: {add: string, remove: string, replace: string}}",
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
			"--relationships", "{$account: {add: string, remove: string, replace: string}, $assignedTo: {add: string, remove: string, replace: string}, $opportunity: {add: string, remove: string, replace: string}}",
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
			"    replace: string\n" +
			"  $assignedTo:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  $opportunity:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n")
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
