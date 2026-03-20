// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Lightfld/lightfield-cli/internal/apiquery"
	"github.com/Lightfld/lightfield-cli/internal/requestflag"
	"github.com/Lightfld/lightfield-go"
	"github.com/Lightfld/lightfield-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var opportunityCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new opportunity record. The `$name` and `$stage` fields and the\n`$account` relationship are required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values for the new opportunity. System fields use a `$` prefix (e.g. `$name`, `$stage`); custom attributes use their bare slug. Required: `$name` (string) and `$stage` (option ID or label). Fields of type `SINGLE_SELECT` or `MULTI_SELECT` accept either an option ID or label from the field's `typeConfiguration.options` — call the <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> to discover available fields and options. See <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for value type details.",
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
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[string]{
			Name:       "fields.name",
			Usage:      "Display name of the opportunity.",
			InnerField: "$name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fields.stage",
			Usage:      "Pipeline stage (`SINGLE_SELECT`). Pass the option ID or label from the field definition.",
			InnerField: "$stage",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[any]{
			Name:       "relationships.account",
			Usage:      "ID of the account this opportunity belongs to.",
			InnerField: "$account",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.champion",
			Usage:      "ID of the contact who is the internal champion.",
			InnerField: "$champion",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.created-by",
			Usage:      "ID of the user who created this opportunity.",
			InnerField: "$createdBy",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.evaluator",
			Usage:      "ID of the contact who is the evaluator.",
			InnerField: "$evaluator",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.owner",
			Usage:      "ID of the user who owns this opportunity.",
			InnerField: "$owner",
		},
	},
})

var opportunityRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single opportunity by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique identifier of the opportunity to retrieve.",
			Required: true,
		},
	},
	Action:          handleOpportunityRetrieve,
	HideHelpCommand: true,
}

var opportunityUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing opportunity by ID. Only included fields and relationships\nare modified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique identifier of the opportunity to update.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values to update — only provided fields are modified; omitted fields are left unchanged. System fields use a `$` prefix (e.g. `$name`, `$stage`); custom attributes use their bare slug. `SINGLE_SELECT` and `MULTI_SELECT` fields accept an option ID or label — call the <u>[definitions endpoint](/api/resources/opportunity/methods/definitions)</u> for available options. See <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for value type details.",
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
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[any]{
			Name:       "fields.name",
			Usage:      "Display name of the opportunity.",
			InnerField: "$name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.stage",
			Usage:      "Pipeline stage (`SINGLE_SELECT`). Pass the option ID or label from the field definition.",
			InnerField: "$stage",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.champion",
			Usage:      "Operation to modify the internal champion.",
			InnerField: "$champion",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.evaluator",
			Usage:      "Operation to modify the evaluator.",
			InnerField: "$evaluator",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.owner",
			Usage:      "Operation to modify the opportunity owner.",
			InnerField: "$owner",
		},
	},
})

var opportunityList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of opportunities. Use `offset` and `limit` to paginate\nthrough results. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for\nmore information about pagination.",
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
	},
	Action:          handleOpportunityList,
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

func handleOpportunityCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomlightfldlightfieldgo.OpportunityNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Opportunity.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "opportunity create", obj, format, transform)
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "opportunity retrieve", obj, format, transform)
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

	params := githubcomlightfldlightfieldgo.OpportunityUpdateParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "opportunity update", obj, format, transform)
}

func handleOpportunityList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomlightfldlightfieldgo.OpportunityListParams{}

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
	_, err = client.Opportunity.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "opportunity list", obj, format, transform)
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "opportunity definitions", obj, format, transform)
}
