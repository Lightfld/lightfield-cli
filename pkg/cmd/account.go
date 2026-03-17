// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/lightfield-cli/internal/apiquery"
	"github.com/stainless-sdks/lightfield-cli/internal/requestflag"
	"github.com/stainless-sdks/lightfield-go"
	"github.com/stainless-sdks/lightfield-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var accountCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Perform create operation",
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
			Name:       "fields.system-name",
			InnerField: "system_name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-facebook",
			InnerField: "system_facebook",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-headcount",
			InnerField: "system_headcount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-industry",
			InnerField: "system_industry",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-instagram",
			InnerField: "system_instagram",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-last-funding-type",
			InnerField: "system_lastFundingType",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-linked-in",
			InnerField: "system_linkedIn",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-primary-address",
			InnerField: "system_primaryAddress",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-twitter",
			InnerField: "system_twitter",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-website",
			InnerField: "system_website",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[any]{
			Name:       "relationships.system-contact",
			InnerField: "system_contact",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.system-owner",
			InnerField: "system_owner",
		},
	},
})

var accountRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Perform retrieve operation",
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
	Usage:   "Perform update operation",
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
			Name:       "fields.system-facebook",
			InnerField: "system_facebook",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-headcount",
			InnerField: "system_headcount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-industry",
			InnerField: "system_industry",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-instagram",
			InnerField: "system_instagram",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-last-funding-type",
			InnerField: "system_lastFundingType",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-linked-in",
			InnerField: "system_linkedIn",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-name",
			InnerField: "system_name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-primary-address",
			InnerField: "system_primaryAddress",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-twitter",
			InnerField: "system_twitter",
		},
		&requestflag.InnerFlag[any]{
			Name:       "fields.system-website",
			InnerField: "system_website",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.system-contact",
			InnerField: "system_contact",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.system-owner",
			InnerField: "system_owner",
		},
	},
})

var accountList = cli.Command{
	Name:    "list",
	Usage:   "Perform list operation",
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

func handleAccountCreate(ctx context.Context, cmd *cli.Command) error {
	client := lightfield.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lightfield.AccountNewParams{}

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
	client := lightfield.NewClient(getDefaultRequestOptions(cmd)...)
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
	client := lightfield.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lightfield.AccountUpdateParams{}

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
	client := lightfield.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := lightfield.AccountListParams{}

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
