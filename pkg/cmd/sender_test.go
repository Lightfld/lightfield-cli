// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestSenderRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sender", "retrieve",
			"--id", "id",
		)
	})
}

func TestSenderUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sender", "update",
			"--id", "id",
			"--pacing", "{EMAIL: {dailyLimit: -9007199254740991, maxGapSeconds: -9007199254740991, minGapSeconds: -9007199254740991}, LINKEDIN_INVITE: {dailyLimit: -9007199254740991, maxGapSeconds: -9007199254740991, minGapSeconds: -9007199254740991}, LINKEDIN_MESSAGE: {dailyLimit: -9007199254740991, maxGapSeconds: -9007199254740991, minGapSeconds: -9007199254740991}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(senderUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sender", "update",
			"--id", "id",
			"--pacing.email", "{dailyLimit: -9007199254740991, maxGapSeconds: -9007199254740991, minGapSeconds: -9007199254740991}",
			"--pacing.linkedin-invite", "{dailyLimit: -9007199254740991, maxGapSeconds: -9007199254740991, minGapSeconds: -9007199254740991}",
			"--pacing.linkedin-message", "{dailyLimit: -9007199254740991, maxGapSeconds: -9007199254740991, minGapSeconds: -9007199254740991}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"pacing:\n" +
			"  EMAIL:\n" +
			"    dailyLimit: -9007199254740991\n" +
			"    maxGapSeconds: -9007199254740991\n" +
			"    minGapSeconds: -9007199254740991\n" +
			"  LINKEDIN_INVITE:\n" +
			"    dailyLimit: -9007199254740991\n" +
			"    maxGapSeconds: -9007199254740991\n" +
			"    minGapSeconds: -9007199254740991\n" +
			"  LINKEDIN_MESSAGE:\n" +
			"    dailyLimit: -9007199254740991\n" +
			"    maxGapSeconds: -9007199254740991\n" +
			"    minGapSeconds: -9007199254740991\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sender", "update",
			"--id", "id",
		)
	})
}

func TestSenderList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sender", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestSenderMarkHealthy(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sender", "mark-healthy",
			"--id", "id",
			"--body", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(senderMarkHealthy)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sender", "mark-healthy",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sender", "mark-healthy",
			"--id", "id",
		)
	})
}
