// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
)

func TestMeetingCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting", "create",
			"--fields", "{$endDate: $endDate, $startDate: $startDate, $title: $title, $attendeeEmails: [string], $description: $description, $meetingUrl: $meetingUrl, $organizerEmail: $organizerEmail, $privacySetting: FULL}",
			"--auto-create-records=true",
			"--relationships", "{$transcript: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $endDate: $endDate\n" +
			"  $startDate: $startDate\n" +
			"  $title: $title\n" +
			"  $attendeeEmails:\n" +
			"    - string\n" +
			"  $description: $description\n" +
			"  $meetingUrl: $meetingUrl\n" +
			"  $organizerEmail: $organizerEmail\n" +
			"  $privacySetting: FULL\n" +
			"autoCreateRecords: true\n" +
			"relationships:\n" +
			"  $transcript: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"meeting", "create",
		)
	})
}

func TestMeetingRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting", "retrieve",
			"--id", "id",
		)
	})
}

func TestMeetingUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting", "update",
			"--id", "id",
			"--fields", "{$privacySetting: FULL}",
			"--relationships", "{$transcript: {replace: replace}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $privacySetting: FULL\n" +
			"relationships:\n" +
			"  $transcript:\n" +
			"    replace: replace\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"meeting", "update",
			"--id", "id",
		)
	})
}

func TestMeetingList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"meeting", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}
