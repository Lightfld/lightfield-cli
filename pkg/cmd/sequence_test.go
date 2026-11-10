// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestSequenceCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "create",
			"--definition", "{description: x, steps: [{deliveryMode: AUTO, name: x, prompt: x, type: EMAIL, mode: AGENT}], messaging: {callToAction: P, persona: P, trigger: P, valueProp: P}, pacing: {EMAIL: {dailyLimit: 1}, LINKEDIN_INVITE: {dailyLimit: 1}, LINKEDIN_MESSAGE: {dailyLimit: 1}}, priority: HIGH, sendWindow: {daysOfWeek: [0], endHour: 1, startHour: 0, timezone: timezone}}",
			"--name", "x",
			"--reviewer", "x",
			"--sender", "x",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sequenceCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "create",
			"--definition.description", "x",
			"--definition.steps", "[{deliveryMode: AUTO, name: x, prompt: x, type: EMAIL, mode: AGENT}]",
			"--definition.messaging", "{callToAction: P, persona: P, trigger: P, valueProp: P}",
			"--definition.pacing", "{EMAIL: {dailyLimit: 1}, LINKEDIN_INVITE: {dailyLimit: 1}, LINKEDIN_MESSAGE: {dailyLimit: 1}}",
			"--definition.priority", "HIGH",
			"--definition.send-window", "{daysOfWeek: [0], endHour: 1, startHour: 0, timezone: timezone}",
			"--name", "x",
			"--reviewer", "x",
			"--sender", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"definition:\n" +
			"  description: x\n" +
			"  steps:\n" +
			"    - deliveryMode: AUTO\n" +
			"      name: x\n" +
			"      prompt: x\n" +
			"      type: EMAIL\n" +
			"      mode: AGENT\n" +
			"  messaging:\n" +
			"    callToAction: P\n" +
			"    persona: P\n" +
			"    trigger: P\n" +
			"    valueProp: P\n" +
			"  pacing:\n" +
			"    EMAIL:\n" +
			"      dailyLimit: 1\n" +
			"    LINKEDIN_INVITE:\n" +
			"      dailyLimit: 1\n" +
			"    LINKEDIN_MESSAGE:\n" +
			"      dailyLimit: 1\n" +
			"  priority: HIGH\n" +
			"  sendWindow:\n" +
			"    daysOfWeek:\n" +
			"      - 0\n" +
			"    endHour: 1\n" +
			"    startHour: 0\n" +
			"    timezone: timezone\n" +
			"name: x\n" +
			"reviewers:\n" +
			"  - x\n" +
			"senders:\n" +
			"  - x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sequence", "create",
		)
	})
}

func TestSequenceRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "retrieve",
			"--id", "id",
		)
	})
}

func TestSequenceList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestSequenceActivate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "activate",
			"--id", "id",
			"--body", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sequenceActivate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "activate",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sequence", "activate",
			"--id", "id",
		)
	})
}

func TestSequenceDeactivate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "deactivate",
			"--id", "id",
			"--body", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sequenceDeactivate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "deactivate",
			"--id", "id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("{}")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sequence", "deactivate",
			"--id", "id",
		)
	})
}

func TestSequenceDiscardDraft(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "discard-draft",
			"--id", "id",
			"--expected-version", "expectedVersion",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("expectedVersion: expectedVersion")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sequence", "discard-draft",
			"--id", "id",
		)
	})
}

func TestSequenceListEnrollments(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "list-enrollments",
			"--id", "id",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestSequenceListMessages(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "list-messages",
			"--id", "id",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestSequenceListTemplateVariables(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "list-template-variables",
		)
	})
}

func TestSequencePublishDraft(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "publish-draft",
			"--id", "id",
			"--expected-version", "expectedVersion",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("expectedVersion: expectedVersion")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sequence", "publish-draft",
			"--id", "id",
		)
	})
}

func TestSequenceRetrieveDraft(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "retrieve-draft",
			"--id", "id",
		)
	})
}

func TestSequenceRetrieveEnrollment(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "retrieve-enrollment",
			"--id", "id",
			"--enrollment-id", "enrollmentId",
		)
	})
}

func TestSequenceRetrieveMessage(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "retrieve-message",
			"--id", "id",
			"--message-id", "messageId",
		)
	})
}

func TestSequenceRetrievePerformance(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "retrieve-performance",
			"--id", "id",
		)
	})
}

func TestSequenceRetrieveTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "retrieve-test",
			"--id", "id",
			"--test-id", "testId",
		)
	})
}

func TestSequenceStartTest(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "start-test",
			"--id", "id",
			"--contact-id", "x",
			"--expected-version", "expectedVersion",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contactId: x\n" +
			"expectedVersion: expectedVersion\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sequence", "start-test",
			"--id", "id",
		)
	})
}

func TestSequenceUpdateDraft(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "update-draft",
			"--id", "id",
			"--definition", "{description: x, messaging: {callToAction: P, persona: P, trigger: P, valueProp: P}, pacing: {EMAIL: {dailyLimit: 1}, LINKEDIN_INVITE: {dailyLimit: 1}, LINKEDIN_MESSAGE: {dailyLimit: 1}}, priority: HIGH, sendWindow: {daysOfWeek: [0], endHour: 1, startHour: 0, timezone: timezone}, steps: [{deliveryMode: AUTO, name: x, prompt: x, type: EMAIL, id: x, mode: AGENT}]}",
			"--expected-version", "expectedVersion",
			"--name", "x",
			"--reviewers", "{add: [x], remove: [x]}",
			"--senders", "{add: [x], remove: [x]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sequenceUpdateDraft)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "update-draft",
			"--id", "id",
			"--definition.description", "x",
			"--definition.messaging", "{callToAction: P, persona: P, trigger: P, valueProp: P}",
			"--definition.pacing", "{EMAIL: {dailyLimit: 1}, LINKEDIN_INVITE: {dailyLimit: 1}, LINKEDIN_MESSAGE: {dailyLimit: 1}}",
			"--definition.priority", "HIGH",
			"--definition.send-window", "{daysOfWeek: [0], endHour: 1, startHour: 0, timezone: timezone}",
			"--definition.steps", "[{deliveryMode: AUTO, name: x, prompt: x, type: EMAIL, id: x, mode: AGENT}]",
			"--expected-version", "expectedVersion",
			"--name", "x",
			"--reviewers.add", "[x]",
			"--reviewers.remove", "[x]",
			"--senders.add", "[x]",
			"--senders.remove", "[x]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"definition:\n" +
			"  description: x\n" +
			"  messaging:\n" +
			"    callToAction: P\n" +
			"    persona: P\n" +
			"    trigger: P\n" +
			"    valueProp: P\n" +
			"  pacing:\n" +
			"    EMAIL:\n" +
			"      dailyLimit: 1\n" +
			"    LINKEDIN_INVITE:\n" +
			"      dailyLimit: 1\n" +
			"    LINKEDIN_MESSAGE:\n" +
			"      dailyLimit: 1\n" +
			"  priority: HIGH\n" +
			"  sendWindow:\n" +
			"    daysOfWeek:\n" +
			"      - 0\n" +
			"    endHour: 1\n" +
			"    startHour: 0\n" +
			"    timezone: timezone\n" +
			"  steps:\n" +
			"    - deliveryMode: AUTO\n" +
			"      name: x\n" +
			"      prompt: x\n" +
			"      type: EMAIL\n" +
			"      id: x\n" +
			"      mode: AGENT\n" +
			"expectedVersion: expectedVersion\n" +
			"name: x\n" +
			"reviewers:\n" +
			"  add:\n" +
			"    - x\n" +
			"  remove:\n" +
			"    - x\n" +
			"senders:\n" +
			"  add:\n" +
			"    - x\n" +
			"  remove:\n" +
			"    - x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sequence", "update-draft",
			"--id", "id",
		)
	})
}

func TestSequenceUpdateEnrollments(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sequence", "update-enrollments",
			"--id", "id",
			"--add", "x",
			"--remove", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"add:\n" +
			"  - x\n" +
			"remove:\n" +
			"  - x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sequence", "update-enrollments",
			"--id", "id",
		)
	})
}
