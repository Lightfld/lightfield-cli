// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
)

func TestAccountCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account", "create",
			"--fields", "{$name: $name, $facebook: $facebook, $headcount: $headcount, $industry: [string], $instagram: $instagram, $lastFundingType: $lastFundingType, $linkedIn: $linkedIn, $primaryAddress: {foo: string}, $twitter: $twitter, $website: [string]}",
			"--relationships", "{$contacts: string, $owner: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account", "create",
			"--fields.name", "$name",
			"--fields.facebook", "$facebook",
			"--fields.headcount", "$headcount",
			"--fields.industry", "[string]",
			"--fields.instagram", "$instagram",
			"--fields.last-funding-type", "$lastFundingType",
			"--fields.linked-in", "$linkedIn",
			"--fields.primary-address", "{foo: string}",
			"--fields.twitter", "$twitter",
			"--fields.website", "[string]",
			"--relationships.contacts", "string",
			"--relationships.owner", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $name: $name\n" +
			"  $facebook: $facebook\n" +
			"  $headcount: $headcount\n" +
			"  $industry:\n" +
			"    - string\n" +
			"  $instagram: $instagram\n" +
			"  $lastFundingType: $lastFundingType\n" +
			"  $linkedIn: $linkedIn\n" +
			"  $primaryAddress:\n" +
			"    foo: string\n" +
			"  $twitter: $twitter\n" +
			"  $website:\n" +
			"    - string\n" +
			"relationships:\n" +
			"  $contacts: string\n" +
			"  $owner: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"account", "create",
		)
	})
}

func TestAccountRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account", "retrieve",
			"--id", "id",
		)
	})
}

func TestAccountUpdate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account", "update",
			"--id", "id",
			"--fields", "{$facebook: $facebook, $headcount: $headcount, $industry: [string], $instagram: $instagram, $lastFundingType: $lastFundingType, $linkedIn: $linkedIn, $name: $name, $primaryAddress: {foo: string}, $twitter: $twitter, $website: [string]}",
			"--relationships", "{$contacts: {add: string, remove: string, replace: string}, $owner: {add: string, remove: string, replace: string}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(accountUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account", "update",
			"--id", "id",
			"--fields.facebook", "$facebook",
			"--fields.headcount", "$headcount",
			"--fields.industry", "[string]",
			"--fields.instagram", "$instagram",
			"--fields.last-funding-type", "$lastFundingType",
			"--fields.linked-in", "$linkedIn",
			"--fields.name", "$name",
			"--fields.primary-address", "{foo: string}",
			"--fields.twitter", "$twitter",
			"--fields.website", "[string]",
			"--relationships.contacts", "{add: string, remove: string, replace: string}",
			"--relationships.owner", "{add: string, remove: string, replace: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"fields:\n" +
			"  $facebook: $facebook\n" +
			"  $headcount: $headcount\n" +
			"  $industry:\n" +
			"    - string\n" +
			"  $instagram: $instagram\n" +
			"  $lastFundingType: $lastFundingType\n" +
			"  $linkedIn: $linkedIn\n" +
			"  $name: $name\n" +
			"  $primaryAddress:\n" +
			"    foo: string\n" +
			"  $twitter: $twitter\n" +
			"  $website:\n" +
			"    - string\n" +
			"relationships:\n" +
			"  $contacts:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n" +
			"  $owner:\n" +
			"    add: string\n" +
			"    remove: string\n" +
			"    replace: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"account", "update",
			"--id", "id",
		)
	})
}

func TestAccountList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account", "list",
			"--limit", "1",
			"--offset", "0",
		)
	})
}

func TestAccountDefinitions(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account", "definitions",
		)
	})
}
