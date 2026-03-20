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

var accountCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new account record. The `$name` field is required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values for the new account. System fields use a `$` prefix (e.g. `$name`, `$website`); custom attributes use their bare slug (e.g. `tier`, `renewalDate`). Required: `$name` (string). Fields of type `SINGLE_SELECT` or `MULTI_SELECT` accept either an option ID or label from the field's `typeConfiguration.options` — call the <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> to discover available fields and options. See <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for value type details.",
			Required: true,
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationships to set on the new account. System relationships use a `$` prefix (e.g. `$owner`, `$contact`); custom relationships use their bare slug. Each value is a single entity ID or an array of IDs. Call the <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> to list available relationship keys.",
			BodyPath: "relationships",
		},
	},
	Action:          handleAccountCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[string]{
			Name:       "fields.name",
			Usage:      "Display name of the account.",
			InnerField: "$name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.facebook",
			Usage:      "Facebook handle or profile identifier (`SOCIAL_HANDLE`).",
			InnerField: "$facebook",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.headcount",
			Usage:      "Employee count range (`SINGLE_SELECT`). Pass the option ID or label from the field definition.",
			InnerField: "$headcount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.industry",
			Usage:      "Industries the account operates in (`MULTI_SELECT`). Pass option IDs or labels from the field definition.",
			InnerField: "$industry",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.instagram",
			Usage:      "Instagram handle or profile identifier (`SOCIAL_HANDLE`).",
			InnerField: "$instagram",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.last-funding-type",
			Usage:      "Most recent funding round type (`SINGLE_SELECT`). Pass the option ID or label from the field definition.",
			InnerField: "$lastFundingType",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.linked-in",
			Usage:      "LinkedIn handle or profile identifier (`SOCIAL_HANDLE`).",
			InnerField: "$linkedIn",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.primary-address",
			Usage:      "Primary address (`ADDRESS`).",
			InnerField: "$primaryAddress",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.twitter",
			Usage:      "Twitter/X handle (`SOCIAL_HANDLE`).",
			InnerField: "$twitter",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.website",
			Usage:      "Website URLs associated with the account (`URL`, multi-value).",
			InnerField: "$website",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[any]{
			Name:       "relationships.contact",
			Usage:      "ID(s) of contacts to associate with this account.",
			InnerField: "$contact",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.owner",
			Usage:      "ID of the user who owns this account.",
			InnerField: "$owner",
		},
	},
})

var accountRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single account by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique identifier of the account to retrieve.",
			Required: true,
		},
	},
	Action:          handleAccountRetrieve,
	HideHelpCommand: true,
}

var accountUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing account by ID. Only included fields and relationships are\nmodified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique identifier of the account to update.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values to update — only provided fields are modified; omitted fields are left unchanged. System fields use a `$` prefix (e.g. `$name`); custom attributes use their bare slug. `SINGLE_SELECT` and `MULTI_SELECT` fields accept an option ID or label — call the <u>[definitions endpoint](/api/resources/account/methods/definitions)</u> for available options. See <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for value type details.",
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationship operations to apply. System relationships use a `$` prefix (e.g. `$owner`, `$contact`). Each value is an operation object with `add`, `remove`, or `replace`.",
			BodyPath: "relationships",
		},
	},
	Action:          handleAccountUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[any]{
			Name:       "fields.facebook",
			Usage:      "Facebook handle or profile identifier (`SOCIAL_HANDLE`).",
			InnerField: "$facebook",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.headcount",
			Usage:      "Employee count range (`SINGLE_SELECT`). Pass the option ID or label from the field definition.",
			InnerField: "$headcount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.industry",
			Usage:      "Industries the account operates in (`MULTI_SELECT`). Pass option IDs or labels from the field definition.",
			InnerField: "$industry",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.instagram",
			Usage:      "Instagram handle or profile identifier (`SOCIAL_HANDLE`).",
			InnerField: "$instagram",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.last-funding-type",
			Usage:      "Most recent funding round type (`SINGLE_SELECT`). Pass the option ID or label from the field definition.",
			InnerField: "$lastFundingType",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.linked-in",
			Usage:      "LinkedIn handle or profile identifier (`SOCIAL_HANDLE`).",
			InnerField: "$linkedIn",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.name",
			Usage:      "Display name of the account.",
			InnerField: "$name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.primary-address",
			Usage:      "Primary address (`ADDRESS`).",
			InnerField: "$primaryAddress",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.twitter",
			Usage:      "Twitter/X handle (`SOCIAL_HANDLE`).",
			InnerField: "$twitter",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.website",
			Usage:      "Website URLs associated with the account (`URL`, multi-value).",
			InnerField: "$website",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.contact",
			Usage:      "Operation to modify associated contacts.",
			InnerField: "$contact",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.owner",
			Usage:      "Operation to modify the account owner.",
			InnerField: "$owner",
		},
	},
})

var accountList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of accounts. Use `offset` and `limit` to paginate\nthrough results. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for\nmore information about pagination.",
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
	Action:          handleAccountList,
	HideHelpCommand: true,
}

var accountDefinitions = cli.Command{
	Name:            "definitions",
	Usage:           "Returns the schema for all field and relationship definitions available on\naccounts, including both system-defined and custom fields. Useful for\nunderstanding the shape of account data before creating or updating records. See\n<u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for\nmore details.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAccountDefinitions,
	HideHelpCommand: true,
}

func handleAccountCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomlightfldlightfieldgo.AccountNewParams{}

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
	_, err = client.Account.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "account create", obj, format, transform)
}

func handleAccountRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Account.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "account retrieve", obj, format, transform)
}

func handleAccountUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomlightfldlightfieldgo.AccountUpdateParams{}

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
	_, err = client.Account.Update(
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
	return ShowJSON(os.Stdout, "account update", obj, format, transform)
}

func handleAccountList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomlightfldlightfieldgo.AccountListParams{}

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
	_, err = client.Account.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "account list", obj, format, transform)
}

func handleAccountDefinitions(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Account.Definitions(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "account definitions", obj, format, transform)
}
