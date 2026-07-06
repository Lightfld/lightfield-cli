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

var listCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new list. The `$name` and `$objectType` fields are required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values for the new list. Required: `$name` (string) and `$objectType`.",
			Required: true,
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationships to set on the new list.",
			BodyPath: "relationships",
		},
	},
	Action:          handleListCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[string]{
			Name:       "fields.name",
			Usage:      "Display name of the list.",
			InnerField: "$name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fields.object-type",
			Usage:      "The type of entities this list contains. Use `$account`, `$contact`, or `$opportunity` (the `$` prefix identifies system types). Bare slugs without the prefix (e.g. `account`) are accepted for backward compatibility but are deprecated.",
			InnerField: "$objectType",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.description",
			Usage:      "Optional description of what this list is for.",
			InnerField: "$description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.kind",
			Usage:      "Product-level list purpose. Normally omit to create a regular list; reserve `target` for the single account-prioritization (APG) demand list per org, and only for account/contact lists.",
			InnerField: "$kind",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[any]{
			Name:       "relationships.accounts",
			Usage:      "Account ID(s) to add as initial members. List `$objectType` must be `account`.",
			InnerField: "$accounts",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.contacts",
			Usage:      "Contact ID(s) to add as initial members. List `$objectType` must be `contact`.",
			InnerField: "$contacts",
		},
		&requestflag.InnerFlag[any]{
			Name:       "relationships.opportunities",
			Usage:      "Opportunity ID(s) to add as initial members. List `$objectType` must be `opportunity`.",
			InnerField: "$opportunities",
		},
	},
})

var listRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single list by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the list to retrieve.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleListRetrieve,
	HideHelpCommand: true,
}

var listUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing list by ID. Only included fields are modified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the list to update.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values to update — only provided fields are modified; omitted fields are left unchanged.",
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationship operations. Use the key matching the list's `$objectType` (e.g. `$accounts` for an account list).",
			BodyPath: "relationships",
		},
	},
	Action:          handleListUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[*string]{
			Name:       "fields.description",
			Usage:      "Optional description of what this list is for. Pass null to clear.",
			InnerField: "$description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.kind",
			Usage:      "Product-level list purpose. Normally leave unset; reserve `target` for the single account-prioritization (APG) demand list per org (account/contact lists only). Pass null to clear.",
			InnerField: "$kind",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fields.name",
			Usage:      "Display name of the list.",
			InnerField: "$name",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.accounts",
			Usage:      "Add/remove accounts. List `$objectType` must be `account`.",
			InnerField: "$accounts",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.contacts",
			Usage:      "Add/remove contacts. List `$objectType` must be `contact`.",
			InnerField: "$contacts",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.opportunities",
			Usage:      "Add/remove opportunities. List `$objectType` must be `opportunity`.",
			InnerField: "$opportunities",
		},
	},
})

var listList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of lists. Use `offset` and `limit` to paginate through\nresults. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more\ninformation about pagination.",
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
	Action:          handleListList,
	HideHelpCommand: true,
}

var listDelete = cli.Command{
	Name:    "delete",
	Usage:   "Moves a list to the trash. The list is soft-deleted and may be restored from the\nLightfield UI. Member entities (accounts, contacts, or opportunities) are not\naffected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the list to delete.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleListDelete,
	HideHelpCommand: true,
}

var listListAccounts = cli.Command{
	Name:    "list-accounts",
	Usage:   "Returns a paginated list of accounts that belong to the specified list.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "list-id",
			Usage:     "Unique identifier of the list.",
			Required:  true,
			PathParam: "listId",
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
	Action:          handleListListAccounts,
	HideHelpCommand: true,
}

var listListContacts = cli.Command{
	Name:    "list-contacts",
	Usage:   "Returns a paginated list of contacts that belong to the specified list.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "list-id",
			Usage:     "Unique identifier of the list.",
			Required:  true,
			PathParam: "listId",
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
	Action:          handleListListContacts,
	HideHelpCommand: true,
}

var listListOpportunities = cli.Command{
	Name:    "list-opportunities",
	Usage:   "Returns a paginated list of opportunities that belong to the specified list.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "list-id",
			Usage:     "Unique identifier of the list.",
			Required:  true,
			PathParam: "listId",
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
	Action:          handleListListOpportunities,
	HideHelpCommand: true,
}

func handleListCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.ListNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.List.New(ctx, params, options...)
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
		Title:          "list create",
		Transform:      transform,
	})
}

func handleListRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.List.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "list retrieve",
		Transform:      transform,
	})
}

func handleListUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.ListUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.List.Update(
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
		Title:          "list update",
		Transform:      transform,
	})
}

func handleListList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.ListListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.List.List(ctx, params, options...)
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
		Title:          "list list",
		Transform:      transform,
	})
}

func handleListDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.ListDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.List.Delete(
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
		Title:          "list delete",
		Transform:      transform,
	})
}

func handleListListAccounts(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.ListListAccountsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.List.ListAccounts(
		ctx,
		cmd.Value("list-id").(string),
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
		Title:          "list list-accounts",
		Transform:      transform,
	})
}

func handleListListContacts(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.ListListContactsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.List.ListContacts(
		ctx,
		cmd.Value("list-id").(string),
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
		Title:          "list list-contacts",
		Transform:      transform,
	})
}

func handleListListOpportunities(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.ListListOpportunitiesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.List.ListOpportunities(
		ctx,
		cmd.Value("list-id").(string),
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
		Title:          "list list-opportunities",
		Transform:      transform,
	})
}
