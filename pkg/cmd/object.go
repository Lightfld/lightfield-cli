// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/Lightfld/lightfield-cli/internal/apiquery"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
	"github.com/Lightfld/lightfield-go"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var objectCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new record for the specified custom object type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The slug of the custom object type.",
			Required:  true,
			PathParam: "entitySlug",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field names to values for the new record.",
			Required: true,
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationship names to entity ID(s) to associate.",
			BodyPath: "relationships",
		},
	},
	Action:          handleObjectCreate,
	HideHelpCommand: true,
}

var objectRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single record by ID for the specified custom object type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The slug of the custom object type.",
			Required:  true,
			PathParam: "entitySlug",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The ID of the record to retrieve.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleObjectRetrieve,
	HideHelpCommand: true,
}

var objectUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates an existing record by ID for the specified custom object type. Only\nincluded fields and relationships are modified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The slug of the custom object type.",
			Required:  true,
			PathParam: "entitySlug",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The ID of the record to update.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field names to values. Only provided fields are modified.",
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationship names to operations (`add`, `remove`, or `replace`).",
			BodyPath: "relationships",
		},
	},
	Action:          handleObjectUpdate,
	HideHelpCommand: true,
}

var objectList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of records for the specified custom object type. Use\n`offset` and `limit` to paginate through results, and `$field` query parameters\nto filter. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more\ninformation about <u>[pagination](/using-the-api/list-endpoints/#pagination)</u>\nand <u>[filtering](/using-the-api/list-endpoints/#filtering)</u>.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The slug of the custom object type.",
			Required:  true,
			PathParam: "entitySlug",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of records to return. Defaults to 25, maximum 25.",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "Number of records to skip for pagination. Defaults to 0.",
			QueryPath: "offset",
		},
	},
	Action:          handleObjectList,
	HideHelpCommand: true,
}

var objectDelete = cli.Command{
	Name:    "delete",
	Usage:   "Moves a custom object record to the trash. The record is soft-deleted and may be\nrestored from the Lightfield UI.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The slug of the custom object type.",
			Required:  true,
			PathParam: "entitySlug",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The ID of the record to delete.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleObjectDelete,
	HideHelpCommand: true,
}

var objectDefinitions = cli.Command{
	Name:    "definitions",
	Usage:   "Returns field and relationship definitions for the specified custom object type.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The slug of the custom object type.",
			Required:  true,
			PathParam: "entitySlug",
		},
	},
	Action:          handleObjectDefinitions,
	HideHelpCommand: true,
}

var objectFieldHistory = cli.Command{
	Name:    "field-history",
	Usage:   "Returns the value-change history for a single field on a record, newest first.\nConsecutive identical values are collapsed. History is cursor-paginated: pass\n`nextCursor` from the previous response as `after` to page through older values.\nOnly attribute-backed fields (custom attributes and attribute-backed system\nfields) have history — column-backed system fields return an error.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The slug of the custom object type.",
			Required:  true,
			PathParam: "entitySlug",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the record.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "field-key",
			Usage:     "Field key whose value history to return. System fields use a `$` prefix (e.g. `$status`); custom attributes use their bare slug.",
			Required:  true,
			PathParam: "fieldKey",
		},
		&requestflag.Flag[string]{
			Name:      "after",
			Usage:     "Cursor from a previous response’s `nextCursor` to fetch the next page.",
			QueryPath: "after",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of history entries to return. Defaults to 20, maximum 100.",
			QueryPath: "limit",
		},
	},
	Action:          handleObjectFieldHistory,
	HideHelpCommand: true,
}

var objectListDefinitions = cli.Command{
	Name:            "list-definitions",
	Usage:           "Returns all custom object types available to the caller.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleObjectListDefinitions,
	HideHelpCommand: true,
}

func handleObjectCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-slug") && len(unusedArgs) > 0 {
		cmd.Set("entity-slug", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomlightfldlightfieldgo.ObjectNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Object.New(
		ctx,
		cmd.Value("entity-slug").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "object create",
		Transform:      transform,
	})
}

func handleObjectRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomlightfldlightfieldgo.ObjectGetParams{
		EntitySlug: cmd.Value("entity-slug").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Object.Get(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "object retrieve",
		Transform:      transform,
	})
}

func handleObjectUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomlightfldlightfieldgo.ObjectUpdateParams{
		EntitySlug: cmd.Value("entity-slug").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Object.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "object update",
		Transform:      transform,
	})
}

func handleObjectList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-slug") && len(unusedArgs) > 0 {
		cmd.Set("entity-slug", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomlightfldlightfieldgo.ObjectListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Object.List(
		ctx,
		cmd.Value("entity-slug").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "object list",
		Transform:      transform,
	})
}

func handleObjectDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomlightfldlightfieldgo.ObjectDeleteParams{
		EntitySlug: cmd.Value("entity-slug").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Object.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "object delete",
		Transform:      transform,
	})
}

func handleObjectDefinitions(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-slug") && len(unusedArgs) > 0 {
		cmd.Set("entity-slug", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Object.Definitions(ctx, cmd.Value("entity-slug").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "object definitions",
		Transform:      transform,
	})
}

func handleObjectFieldHistory(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("field-key") && len(unusedArgs) > 0 {
		cmd.Set("field-key", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomlightfldlightfieldgo.ObjectFieldHistoryParams{
		EntitySlug: cmd.Value("entity-slug").(string),
		ID:         cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Object.FieldHistory(
		ctx,
		cmd.Value("field-key").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "object field-history",
		Transform:      transform,
	})
}

func handleObjectListDefinitions(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Object.ListDefinitions(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "object list-definitions",
		Transform:      transform,
	})
}
