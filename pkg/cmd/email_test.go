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

func TestEmailSend(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"email", "send",
			"--body", "{content: x, contentType: HTML}",
			"--from", `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
			"--subject", "x",
			"--to", `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
			"--attachment", "string",
			"--bcc", `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
			"--cc", `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
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
			"--body.content", "x",
			"--body.content-type", "HTML",
			"--from", `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
			"--subject", "x",
			"--to", `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
			"--attachment", "string",
			"--bcc", `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
			"--cc", `"S?oC"g*W"5"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw`,
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"body:\n" +
			"  content: x\n" +
			"  contentType: HTML\n" +
			"from: >-\n" +
			"  \"S?oC\"g*W\"5\"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw\n" +
			"subject: x\n" +
			"to:\n" +
			"  - >-\n" +
			"    \"S?oC\"g*W\"5\"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw\n" +
			"attachments:\n" +
			"  - string\n" +
			"bcc:\n" +
			"  - >-\n" +
			"    \"S?oC\"g*W\"5\"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw\n" +
			"cc:\n" +
			"  - >-\n" +
			"    \"S?oC\"g*W\"5\"@m-0-9.V9.w4-o2-l7--.TJdq6.1k.H8n-SjA-.1.U3k7.-F86WnfNI.-R6O-N68g-4-.AmqyytAVIw\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"email", "send",
		)
	})
}
