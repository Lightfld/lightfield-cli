package cmd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runAgainstServer(t *testing.T, args ...string) (*http.Request, error) {
	t.Helper()
	var captured *http.Request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"object":"list","data":[],"totalCount":0}`))
	}))
	t.Cleanup(server.Close)

	argv := append([]string{"lightfield", "--base-url", server.URL, "--api-key", "key", "--format", "json"}, args...)
	err := Command.Run(context.Background(), argv)
	return captured, err
}

func TestListFilterSendsEachFilterAsQueryParam(t *testing.T) {
	req, err := runAgainstServer(t,
		"opportunity", "list",
		"--limit", "25",
		"--filter", "$stage[equal]=Qualified",
		"--filter", "$name[equal]=Acme, Inc.",
		"--filter", "$updatedAt[greaterThanOrEqual]=2026-11-11T00:00:00Z",
	)
	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Equal(t, "/v1/opportunities", req.URL.Path)
	assert.Equal(t, url.Values{
		"limit":                          {"25"},
		"$stage[equal]":                  {"Qualified"},
		"$name[equal]":                   {"Acme, Inc."},
		"$updatedAt[greaterThanOrEqual]": {"2026-11-11T00:00:00Z"},
	}, req.URL.Query())
}

func TestListFilterKeepsCustomFieldKeysAndEmptyValues(t *testing.T) {
	req, err := runAgainstServer(t,
		"object", "list",
		"--entity-slug", "deal-product",
		"--filter", "budget-status=approved",
		"--filter", "$owner[-equal]=",
	)
	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Equal(t, "/v1/objects/deal-product", req.URL.Path)
	assert.Equal(t, url.Values{
		"budget-status":  {"approved"},
		"$owner[-equal]": {""},
	}, req.URL.Query())
}

func TestListFilterRejectsMalformedFilter(t *testing.T) {
	req, err := runAgainstServer(t, "contact", "list", "--filter", "$name")
	require.ErrorContains(t, err, "expected FIELD[OPERATOR]=VALUE")
	assert.Nil(t, req)
}

func TestListFilterRejectsKeysOwnedByOtherFlags(t *testing.T) {
	req, err := runAgainstServer(t, "account", "list", "--filter", "limit=5")
	require.ErrorContains(t, err, "use --limit instead")
	assert.Nil(t, req)
}
