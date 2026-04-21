package interactive_test

import (
	"testing"

	"github.com/Lightfld/lightfield-cli/internal/interactive"
	cmdpkg "github.com/Lightfld/lightfield-cli/pkg/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildCatalogFromGeneratedCommandTree(t *testing.T) {
	catalog := interactive.BuildCatalog(cmdpkg.Command)

	require.NotEmpty(t, catalog.Resources)
	require.NotEmpty(t, catalog.GlobalFields)

	resourceIdx := indexResource(catalog, "account")
	require.NotEqual(t, -1, resourceIdx)

	actionIdx := indexAction(catalog.Resources[resourceIdx], "create")
	require.NotEqual(t, -1, actionIdx)

	action := catalog.Resources[resourceIdx].Actions[actionIdx]
	assert.Equal(t, "account", action.ResourceName)
	assert.Equal(t, "create", action.Name)
	assert.NotEmpty(t, action.Fields)

	apiKeyIdx := indexField(catalog.GlobalFields, "api-key")
	require.NotEqual(t, -1, apiKeyIdx)
	assert.Equal(t, interactive.FieldKindText, catalog.GlobalFields[apiKeyIdx].Kind)
	assert.Equal(t, -1, indexField(catalog.GlobalFields, "help"))
	assert.Equal(t, -1, indexField(catalog.GlobalFields, "version"))

	fieldsIdx := indexField(action.Fields, "fields")
	require.NotEqual(t, -1, fieldsIdx)
	assert.Equal(t, interactive.FieldKindMultiline, action.Fields[fieldsIdx].Kind)
	assert.NotEmpty(t, action.Fields[fieldsIdx].Usage)
	assert.True(t, action.Fields[fieldsIdx].Required)
	assert.Equal(t, -1, indexAction(catalog.Resources[resourceIdx], "help"))
	assert.NotContains(t, catalog.Resources[resourceIdx].Actions[indexAction(catalog.Resources[resourceIdx], "list")].Usage, "<u>")
	assert.NotContains(t, catalog.Resources[resourceIdx].Actions[indexAction(catalog.Resources[resourceIdx], "list")].Usage, "](")
}

func TestBuildArgsUsesGlobalAndLocalFlagsInOrder(t *testing.T) {
	catalog := interactive.Catalog{
		GlobalFields: []interactive.FieldSpec{
			{Name: "api-key", Kind: interactive.FieldKindText},
		},
		Resources: []interactive.ResourceSpec{
			{
				Name: "account",
				Actions: []interactive.ActionSpec{
					{
						Name: "list",
						Fields: []interactive.FieldSpec{
							{Name: "limit", Kind: interactive.FieldKindNumber},
						},
					},
				},
			},
		},
	}

	state := interactive.SessionState{
		ResourceIdx: 0,
		ActionIdx:   0,
		Values: map[string]string{
			"api-key": "sk test",
			"limit":   "5",
		},
	}

	args, display, err := interactive.BuildArgs(catalog, state)
	require.NoError(t, err)
	assert.Equal(t, []string{"--api-key", "sk test", "account", "list", "--limit", "5"}, args)
	assert.Equal(t, "lightfield --api-key 'sk test' account list --limit 5", display)
}

func TestRootCommandHasInteractiveAction(t *testing.T) {
	require.NotNil(t, cmdpkg.Command.Action)
}

func indexResource(catalog interactive.Catalog, name string) int {
	for i, resource := range catalog.Resources {
		if resource.Name == name {
			return i
		}
	}
	return -1
}

func indexAction(resource interactive.ResourceSpec, name string) int {
	for i, action := range resource.Actions {
		if action.Name == name {
			return i
		}
	}
	return -1
}

func indexField(fields []interactive.FieldSpec, name string) int {
	for i, field := range fields {
		if field.Name == name {
			return i
		}
	}
	return -1
}
