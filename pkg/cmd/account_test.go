// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/mocktest"
)

func TestAccountCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"account", "create",
			"--fields", "{$name: $name, $facebook: $facebook, $headcount: $headcount, $industry: [string], $instagram: $instagram, $lastFundingType: $lastFundingType, $linkedIn: $linkedIn, $primaryAddress: {city: city, country: country, latitude: 0, longitude: 0, postalCode: postalCode, state: state, street: street, street2: street2}, $twitter: $twitter, $website: [string]}",
			"--relationships", "{$contact: string, $owner: string}",
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
			"    city: city\n" +
			"    country: country\n" +
			"    latitude: 0\n" +
			"    longitude: 0\n" +
			"    postalCode: postalCode\n" +
			"    state: state\n" +
			"    street: street\n" +
			"    street2: street2\n" +
			"  $twitter: $twitter\n" +
			"  $website:\n" +
			"    - string\n" +
			"relationships:\n" +
			"  $contact: string\n" +
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
			"--fields", "{$facebook: $facebook, $headcount: $headcount, $industry: [string], $instagram: $instagram, $lastFundingType: $lastFundingType, $linkedIn: $linkedIn, $name: $name, $primaryAddress: {city: city, country: country, latitude: 0, longitude: 0, postalCode: postalCode, state: state, street: street, street2: street2}, $twitter: $twitter, $website: [string]}",
			"--relationships", "{$contact: {add: string, remove: string, replace: string}, $owner: {add: string, remove: string, replace: string}}",
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
			"    city: city\n" +
			"    country: country\n" +
			"    latitude: 0\n" +
			"    longitude: 0\n" +
			"    postalCode: postalCode\n" +
			"    state: state\n" +
			"    street: street\n" +
			"    street2: street2\n" +
			"  $twitter: $twitter\n" +
			"  $website:\n" +
			"    - string\n" +
			"relationships:\n" +
			"  $contact:\n" +
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
