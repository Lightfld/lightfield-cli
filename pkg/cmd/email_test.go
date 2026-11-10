// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestEmailRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "retrieve",
			"--id", "id",
		)
	})
}

func TestEmailList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestEmailDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "delete",
			"--id", "id",
			"--body", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailDelete)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "delete",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email", "delete",
			"--id", "id",
		)
	})
}

func TestEmailDraft(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "draft",
			"--from", "sales@acme.com",
			"--attachment", "file_01abc2def3ghi4jkl5mno6pqr",
			"--bcc", "lead@example.com",
			"--cc", "lead@example.com",
			"--message-body", "{content: '<p>Hi there,</p><p>Following up on our chat earlier this week.</p>', contentType: HTML}",
			"--subject", "subject",
			"--to", "lead@example.com",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailDraft)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "draft",
			"--from", "sales@acme.com",
			"--attachment", "file_01abc2def3ghi4jkl5mno6pqr",
			"--bcc", "lead@example.com",
			"--cc", "lead@example.com",
			"--message-body.content", "<p>Hi there,</p><p>Following up on our chat earlier this week.</p>",
			"--message-body.content-type", "HTML",
			"--subject", "subject",
			"--to", "lead@example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"from: sales@acme.com\n" +
			"attachments:\n" +
			"  - file_01abc2def3ghi4jkl5mno6pqr\n" +
			"bcc:\n" +
			"  - lead@example.com\n" +
			"cc:\n" +
			"  - lead@example.com\n" +
			"messageBody:\n" +
			"  content: <p>Hi there,</p><p>Following up on our chat earlier this week.</p>\n" +
			"  contentType: HTML\n" +
			"subject: subject\n" +
			"to:\n" +
			"  - lead@example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email", "draft",
		)
	})
}

func TestEmailLog(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "log",
			"--fields", "{$from: $from, $sentAt: $sentAt, $bcc: [string], $body: {html: html, text: text}, $cc: [string], $privacySetting: FULL, $subject: $subject, $threadId: x, $to: [string]}",
			"--auto-create-records=true",
			"--headers", "{foo: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailLog)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "log",
			"--fields.from", "$from",
			"--fields.sent-at", "$sentAt",
			"--fields.bcc", "[string]",
			"--fields.body", "{html: html, text: text}",
			"--fields.cc", "[string]",
			"--fields.privacy-setting", "FULL",
			"--fields.subject", "$subject",
			"--fields.thread-id", "x",
			"--fields.to", "[string]",
			"--auto-create-records=true",
			"--headers", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $from: $from\n" +
			"  $sentAt: $sentAt\n" +
			"  $bcc:\n" +
			"    - string\n" +
			"  $body:\n" +
			"    html: html\n" +
			"    text: text\n" +
			"  $cc:\n" +
			"    - string\n" +
			"  $privacySetting: FULL\n" +
			"  $subject: $subject\n" +
			"  $threadId: x\n" +
			"  $to:\n" +
			"    - string\n" +
			"autoCreateRecords: true\n" +
			"headers:\n" +
			"  foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email", "log",
		)
	})
}

func TestEmailRestore(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "restore",
			"--id", "id",
			"--body", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailRestore)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "restore",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email", "restore",
			"--id", "id",
		)
	})
}

func TestEmailSend(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "send",
			"--from", "sales@acme.com",
			"--message-body", "{content: '<p>Hi there,</p><p>Following up on our chat earlier this week.</p>', contentType: HTML}",
			"--subject", "Following up on our chat",
			"--to", "lead@example.com",
			"--attachment", "file_01abc2def3ghi4jkl5mno6pqr",
			"--bcc", "lead@example.com",
			"--cc", "lead@example.com",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(emailSend)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "send",
			"--from", "sales@acme.com",
			"--message-body.content", "<p>Hi there,</p><p>Following up on our chat earlier this week.</p>",
			"--message-body.content-type", "HTML",
			"--subject", "Following up on our chat",
			"--to", "lead@example.com",
			"--attachment", "file_01abc2def3ghi4jkl5mno6pqr",
			"--bcc", "lead@example.com",
			"--cc", "lead@example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"from: sales@acme.com\n" +
			"messageBody:\n" +
			"  content: <p>Hi there,</p><p>Following up on our chat earlier this week.</p>\n" +
			"  contentType: HTML\n" +
			"subject: Following up on our chat\n" +
			"to:\n" +
			"  - lead@example.com\n" +
			"attachments:\n" +
			"  - file_01abc2def3ghi4jkl5mno6pqr\n" +
			"bcc:\n" +
			"  - lead@example.com\n" +
			"cc:\n" +
			"  - lead@example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email", "send",
		)
	})
}
