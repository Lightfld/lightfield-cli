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

var mergeGetMerge = cli.Command{
	Name:    "get-merge",
	Usage:   "Returns the status and details of a merge operation by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The merge operation ID.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMergeGetMerge,
	HideHelpCommand: true,
}

var mergeMergeAccounts = requestflag.WithInnerFlags(cli.Command{
	Name:    "merge-accounts",
	Usage:   "Merges two accounts into one. The primary account retains its ID; the duplicate\nis soft-deleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "duplicate-id",
			Usage:    "ID of the duplicate record to merge into the primary and then discard.",
			Required: true,
			BodyPath: "duplicateId",
		},
		&requestflag.Flag[string]{
			Name:     "primary-id",
			Usage:    "ID of the record to keep.",
			Required: true,
			BodyPath: "primaryId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "field-resolutions",
			Usage:    "Per-field resolution overrides keyed by attribute slug.",
			BodyPath: "fieldResolutions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
	},
	Action:          handleMergeMergeAccounts,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[bool]{
			Name:       "options.multi-select-union",
			Usage:      "When true, multi-select fields are merged by union rather than primary-takes-all.",
			InnerField: "multiSelectUnion",
		},
	},
})

var mergeMergeContacts = requestflag.WithInnerFlags(cli.Command{
	Name:    "merge-contacts",
	Usage:   "Merges two contacts into one. The primary contact retains its ID; the duplicate\nis soft-deleted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "duplicate-id",
			Usage:    "ID of the duplicate record to merge into the primary and then discard.",
			Required: true,
			BodyPath: "duplicateId",
		},
		&requestflag.Flag[string]{
			Name:     "primary-id",
			Usage:    "ID of the record to keep.",
			Required: true,
			BodyPath: "primaryId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "field-resolutions",
			Usage:    "Per-field resolution overrides keyed by attribute slug.",
			BodyPath: "fieldResolutions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
	},
	Action:          handleMergeMergeContacts,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[bool]{
			Name:       "options.multi-select-union",
			Usage:      "When true, multi-select fields are merged by union rather than primary-takes-all.",
			InnerField: "multiSelectUnion",
		},
	},
})

var mergeMergeObjectValues = requestflag.WithInnerFlags(cli.Command{
	Name:    "merge-object-values",
	Usage:   "Merges two records of the specified custom object type into one. The primary\nrecord retains its ID; the duplicate is soft-deleted. Both records must belong\nto the custom object type named in the path.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The slug of the custom object type.",
			Required:  true,
			PathParam: "entitySlug",
		},
		&requestflag.Flag[string]{
			Name:     "duplicate-id",
			Usage:    "ID of the duplicate record to merge into the primary and then discard.",
			Required: true,
			BodyPath: "duplicateId",
		},
		&requestflag.Flag[string]{
			Name:     "primary-id",
			Usage:    "ID of the record to keep.",
			Required: true,
			BodyPath: "primaryId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "field-resolutions",
			Usage:    "Per-field resolution overrides keyed by attribute slug.",
			BodyPath: "fieldResolutions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
	},
	Action:          handleMergeMergeObjectValues,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[bool]{
			Name:       "options.multi-select-union",
			Usage:      "When true, multi-select fields are merged by union rather than primary-takes-all.",
			InnerField: "multiSelectUnion",
		},
	},
})

var mergeMergeOpportunities = requestflag.WithInnerFlags(cli.Command{
	Name:    "merge-opportunities",
	Usage:   "Merges two opportunities into one. The primary opportunity retains its ID; the\nduplicate is soft-deleted. Both opportunities must belong to the same account.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "duplicate-id",
			Usage:    "ID of the duplicate record to merge into the primary and then discard.",
			Required: true,
			BodyPath: "duplicateId",
		},
		&requestflag.Flag[string]{
			Name:     "primary-id",
			Usage:    "ID of the record to keep.",
			Required: true,
			BodyPath: "primaryId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "field-resolutions",
			Usage:    "Per-field resolution overrides keyed by attribute slug.",
			BodyPath: "fieldResolutions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "options",
			BodyPath: "options",
		},
	},
	Action:          handleMergeMergeOpportunities,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"options": {
		&requestflag.InnerFlag[bool]{
			Name:       "options.multi-select-union",
			Usage:      "When true, multi-select fields are merged by union rather than primary-takes-all.",
			InnerField: "multiSelectUnion",
		},
	},
})

func handleMergeGetMerge(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Merge.GetMerge(ctx, cmd.Value("id").(string), options...)
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
		Title:          "merge get-merge",
		Transform:      transform,
	})
}

func handleMergeMergeAccounts(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MergeMergeAccountsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Merge.MergeAccounts(ctx, params, options...)
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
		Title:          "merge merge-accounts",
		Transform:      transform,
	})
}

func handleMergeMergeContacts(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MergeMergeContactsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Merge.MergeContacts(ctx, params, options...)
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
		Title:          "merge merge-contacts",
		Transform:      transform,
	})
}

func handleMergeMergeObjectValues(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MergeMergeObjectValuesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Merge.MergeObjectValues(
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
		Title:          "merge merge-object-values",
		Transform:      transform,
	})
}

func handleMergeMergeOpportunities(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MergeMergeOpportunitiesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Merge.MergeOpportunities(ctx, params, options...)
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
		Title:          "merge merge-opportunities",
		Transform:      transform,
	})
}
