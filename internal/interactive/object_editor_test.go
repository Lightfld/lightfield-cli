package interactive

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseObjectRowsRoundTrip(t *testing.T) {
	rows := parseObjectRows(`{
  "$name": "Acme",
  "tier": "enterprise",
  "active": true
}`)

	require.Len(t, rows, 3)
	assert.Equal(t, "$name", rows[0].Key)
	assert.Equal(t, "Acme", rows[0].Value)
	assert.Equal(t, "active", rows[1].Key)
	assert.Equal(t, "true", rows[1].Value)
	assert.Equal(t, "tier", rows[2].Key)
	assert.Equal(t, "enterprise", rows[2].Value)
}

func TestSerializeObjectRowsCoercesSimpleValues(t *testing.T) {
	value, err := serializeObjectRows([]objectRow{
		{Key: "$name", Value: "Acme"},
		{Key: "employeeCount", Value: "42"},
		{Key: "active", Value: "true"},
		{Key: "segments", Value: `["enterprise","mid-market"]`},
	})

	require.NoError(t, err)
	assert.JSONEq(t, `{
  "$name": "Acme",
  "employeeCount": 42,
  "active": true,
  "segments": ["enterprise", "mid-market"]
}`, value)
}
