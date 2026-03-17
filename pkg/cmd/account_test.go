// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/lightfield-cli/internal/mocktest"
	"github.com/stainless-sdks/lightfield-cli/internal/requestflag"
)

func TestAccountCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "account", "create",
			"--api-key", "string",
			"--fields", "{system_name: system_name, system_facebook: system_facebook, system_headcount: system_headcount, system_industry: [string], system_instagram: system_instagram, system_lastFundingType: system_lastFundingType, system_linkedIn: system_linkedIn, system_primaryAddress: {foo: string}, system_twitter: system_twitter, system_website: [string]}",
			"--relationships", "{system_contact: string, system_owner: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "account", "create",
			"--api-key", "string",
			"--fields.system-name", "system_name",
			"--fields.system-facebook", "system_facebook",
			"--fields.system-headcount", "system_headcount",
			"--fields.system-industry", "[string]",
			"--fields.system-instagram", "system_instagram",
			"--fields.system-last-funding-type", "system_lastFundingType",
			"--fields.system-linked-in", "system_linkedIn",
			"--fields.system-primary-address", "{foo: string}",
			"--fields.system-twitter", "system_twitter",
			"--fields.system-website", "[string]",
			"--relationships.system-contact", "string",
			"--relationships.system-owner", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  system_name: system_name\n" +
			"  system_facebook: system_facebook\n" +
			"  system_headcount: system_headcount\n" +
			"  system_industry:\n" +
			"    - string\n" +
			"  system_instagram: system_instagram\n" +
			"  system_lastFundingType: system_lastFundingType\n" +
			"  system_linkedIn: system_linkedIn\n" +
			"  system_primaryAddress:\n" +
			"    foo: string\n" +
			"  system_twitter: system_twitter\n" +
			"  system_website:\n" +
			"    - string\n" +
			"relationships:\n" +
			"  system_contact: string\n" +
			"  system_owner: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "account", "create",
			"--api-key", "string",
		)
	})
}

func TestAccountRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "account", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAccountUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "account", "update",
			"--api-key", "string",
			"--id", "id",
			"--fields", "{system_facebook: system_facebook, system_headcount: system_headcount, system_industry: [string], system_instagram: system_instagram, system_lastFundingType: system_lastFundingType, system_linkedIn: system_linkedIn, system_name: system_name, system_primaryAddress: {foo: string}, system_twitter: system_twitter, system_website: [string]}",
			"--relationships", "{system_contact: {add: string, remove: string, replace: string}, system_owner: {add: string, remove: string, replace: string}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "account", "update",
			"--api-key", "string",
			"--id", "id",
			"--fields.system-facebook", "system_facebook",
			"--fields.system-headcount", "system_headcount",
			"--fields.system-industry", "[string]",
			"--fields.system-instagram", "system_instagram",
			"--fields.system-last-funding-type", "system_lastFundingType",
			"--fields.system-linked-in", "system_linkedIn",
			"--fields.system-name", "system_name",
			"--fields.system-primary-address", "{foo: string}",
			"--fields.system-twitter", "system_twitter",
			"--fields.system-website", "[string]",
			"--relationships.system-contact", "{add: string, remove: string, replace: string}",
			"--relationships.system-owner", "{add: string, remove: string, replace: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  system_facebook: system_facebook\n" +
			"  system_headcount: system_headcount\n" +
			"  system_industry:\n" +
			"    - string\n" +
			"  system_instagram: system_instagram\n" +
			"  system_lastFundingType: system_lastFundingType\n" +
			"  system_linkedIn: system_linkedIn\n" +
			"  system_name: system_name\n" +
			"  system_primaryAddress:\n" +
			"    foo: string\n" +
			"  system_twitter: system_twitter\n" +
			"  system_website:\n" +
			"    - string\n" +
			"relationships:\n" +
			"  system_contact:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  system_owner:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "account", "update",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestAccountList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "account", "list",
			"--api-key", "string",
			"--limit", "1",
			"--offset", "0",
		)
	})
}
