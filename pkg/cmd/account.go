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
			Required: true,
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			BodyPath: "relationships",
		},
	},
	Action:          handleAccountCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[string]{
			Name:       "fields.name",
			InnerField: "$name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.facebook",
			InnerField: "$facebook",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.headcount",
			InnerField: "$headcount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.industry",
			InnerField: "$industry",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.instagram",
			InnerField: "$instagram",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.last-funding-type",
			InnerField: "$lastFundingType",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.linked-in",
			InnerField: "$linkedIn",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.primary-address",
			InnerField: "$primaryAddress",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.twitter",
			InnerField: "$twitter",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.website",
			InnerField: "$website",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[any]{
			Name:       "relationships.contacts",
			InnerField: "$contacts",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.owner",
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
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			BodyPath: "relationships",
		},
	},
	Action:          handleAccountUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[any]{
			Name:       "fields.facebook",
			InnerField: "$facebook",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.headcount",
			InnerField: "$headcount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.industry",
			InnerField: "$industry",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.instagram",
			InnerField: "$instagram",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.last-funding-type",
			InnerField: "$lastFundingType",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.linked-in",
			InnerField: "$linkedIn",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.name",
			InnerField: "$name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.primary-address",
			InnerField: "$primaryAddress",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.twitter",
			InnerField: "$twitter",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.website",
			InnerField: "$website",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.contacts",
			InnerField: "$contacts",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.owner",
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
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
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
