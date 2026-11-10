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

var opportunityCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new opportunity record. The `$name` field and the `$account`\nrelationship are required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values for the new opportunity. System fields use a `$` prefix (e.g. `$name`, `$stage`); custom attributes use their bare slug. Required: `$name` (string). `$stage` accepts an option ID or label when provided. Fields of type `SINGLE_SELECT` or `MULTI_SELECT` accept either an option ID or label from the field's `typeConfiguration.options` — call the <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> to discover available fields and options. See <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for value type details.\n\nRead definitions for the object type’s configured `$pipeline` and `$stage` fields and options. See <u>[Pipelines and stages](/using-the-api/fields-and-relationships/#pipelines-and-stages)</u> for write rules.",
			Required: true,
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationships to set on the new opportunity. System relationships use a `$` prefix (e.g. `$account`, `$owner`); custom relationships use their bare slug. `$account` is required. Each value is a single entity ID or an array of IDs. Call the <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> to list available relationship keys.",
			Required: true,
			BodyPath: "relationships",
		},
	},
	Action:          handleOpportunityCreate,
	HideHelpCommand: true,
}

var opportunityRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single opportunity by its ID. A record you may see by name only is\nreturned as `{ id, restricted: { name, reason } }`, with no `fields` or\n`relationships`; check for `restricted` before reading them.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the opportunity to retrieve.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleOpportunityRetrieve,
	HideHelpCommand: true,
}

var opportunityUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates an existing opportunity by ID. Only included fields and relationships\nare modified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the opportunity to update.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values to update — only provided fields are modified; omitted fields are left unchanged. System fields use a `$` prefix (e.g. `$name`, `$stage`); custom attributes use their bare slug. `SINGLE_SELECT` and `MULTI_SELECT` fields accept an option ID or label — call the <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> for available options. See <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for value type details.\n\nRead definitions for the object type’s configured `$pipeline` and `$stage` fields and options. See <u>[Pipelines and stages](/using-the-api/fields-and-relationships/#pipelines-and-stages)</u> for write rules.",
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationship operations to apply. System relationships use a `$` prefix (e.g. `$owner`, `$champion`). Each value is an operation object with `add`, `remove`, or `replace`.",
			BodyPath: "relationships",
		},
	},
	Action:          handleOpportunityUpdate,
	HideHelpCommand: true,
}

var opportunityList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of opportunities. Results are sorted by when the record\nwas created in Lightfield (`createdAt`), newest first. Each page holds only the\nrecords you can currently read, so a page can contain fewer records than\n`limit`, or none, while more remain, and `totalCount` can include records you\ncan no longer read. To read every record, keep advancing `offset` by `limit`\nuntil it reaches `totalCount`; do not stop at a short or empty page. Use\n`offset` and `limit` to paginate through results, and `$field` query parameters\nto filter. To list the opportunities on a given account, filter by the\n`$account` relationship — e.g. `?$account=acc_123` (`?$account[-equal]=acc_123`\nexcludes them instead).",
	Suggest: true,
	Flags: []cli.Flag{
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
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Which records to return by trash state. Defaults to \"active\". Use \"trashed\" to list records currently in the trash; this additionally requires this resource's `restore` scope.",
			QueryPath: "status",
		},
	},
	Action:          handleOpportunityList,
	HideHelpCommand: true,
}

var opportunityDelete = cli.Command{
	Name:    "delete",
	Usage:   "Moves an opportunity to the trash. The opportunity is soft-deleted and may be\nrestored from the Lightfield UI. Associated tasks, notes, and meetings are not\ndeleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the opportunity to delete.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleOpportunityDelete,
	HideHelpCommand: true,
}

var opportunityDefinitions = cli.Command{
	Name:            "definitions",
	Usage:           "Returns the schema for all field and relationship definitions available on\nopportunities, including both system-defined and custom fields. Useful for\nunderstanding the shape of opportunity data before creating or updating records.\nSee <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u>\nfor more details.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleOpportunityDefinitions,
	HideHelpCommand: true,
}

var opportunityFieldHistory = cli.Command{
	Name:    "field-history",
	Usage:   "Returns the value-change history for a single field on a record, newest first.\nConsecutive identical values are collapsed. History is cursor-paginated: pass\n`nextCursor` from the previous response as `after` to page through older values.\nOnly attribute-backed fields (custom attributes and attribute-backed system\nfields) have history — column-backed system fields return an error. A record you\nmay see by name only has no visible history and returns 404, like a record that\ndoes not exist.",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleOpportunityFieldHistory,
	HideHelpCommand: true,
}

var opportunityRestore = cli.Command{
	Name:    "restore",
	Usage:   "Restores a trashed opportunity, bringing it back to active. Find trashed\nopportunities via `GET /v1/opportunities?status=trashed`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the opportunity to restore.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleOpportunityRestore,
	HideHelpCommand: true,
}

func handleOpportunityCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := githubcomlightfldlightfieldgo.OpportunityNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Opportunity.New(ctx, params, options...)
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
		Title:          "opportunity create",
		Transform:      transform,
	})
}

func handleOpportunityRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Opportunity.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "opportunity retrieve",
		Transform:      transform,
	})
}

func handleOpportunityUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.OpportunityUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Opportunity.Update(
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
		Title:          "opportunity update",
		Transform:      transform,
	})
}

func handleOpportunityList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.OpportunityListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Opportunity.List(ctx, params, options...)
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
		Title:          "opportunity list",
		Transform:      transform,
	})
}

func handleOpportunityDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.OpportunityDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Opportunity.Delete(
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
		Title:          "opportunity delete",
		Transform:      transform,
	})
}

func handleOpportunityDefinitions(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Opportunity.Definitions(ctx, options...)
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
		Title:          "opportunity definitions",
		Transform:      transform,
	})
}

func handleOpportunityFieldHistory(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.OpportunityFieldHistoryParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Opportunity.FieldHistory(
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
		Title:          "opportunity field-history",
		Transform:      transform,
	})
}

func handleOpportunityRestore(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.OpportunityRestoreParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Opportunity.Restore(
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
		Title:          "opportunity restore",
		Transform:      transform,
	})
}
