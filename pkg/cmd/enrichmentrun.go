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

var enrichmentRunEnrich = cli.Command{
	Name:    "enrich",
	Usage:   "Looks up missing information for a record from Lightfield's enrichment providers\nand writes it back to it. Accepts `contacts` and `accounts`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "entity-slug",
			Usage:     "The type of record to enrich.",
			Required:  true,
			PathParam: "entitySlug",
		},
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "The ID of the record to enrich.",
			Required:  true,
			PathParam: "entityId",
		},
		&requestflag.Flag[[]string]{
			Name:     "field",
			Usage:    "Fields to enrich, e.g. `[\"email\", \"title\"]`. Named fields are refreshed even when they already hold a value, so this is how a stale value is replaced. Omit to fill only the fields the record is missing; if any of those the workspace enriches only on request, no run starts and the response names them (`manual_fields_unconfirmed`). Naming a field the workspace's settings enrich only on request enriches it, and its answer is raised as a suggestion rather than written over a value the record already holds. Naming a field the settings turn off is rejected with `400 enrichment_field_disabled`. A named field is not used to look the record up — the run derives it fresh from the record's other identifiers — so naming every identifier (a contact's name and email together) leaves nothing to search by and is rejected with `400 enrichment_fields_not_enrichable`, creating no run; refresh them one at a time. `phone` is turned off unless the workspace opts into it, so naming it is rejected with `400 enrichment_field_disabled`.",
			BodyPath: "fields",
		},
	},
	Action:          handleEnrichmentRunEnrich,
	HideHelpCommand: true,
}

var enrichmentRunGetEnrichmentRun = cli.Command{
	Name:    "get-enrichment-run",
	Usage:   "Returns the status of an enrichment run by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "run-id",
			Usage:     "The enrichment run ID.",
			Required:  true,
			PathParam: "runId",
		},
	},
	Action:          handleEnrichmentRunGetEnrichmentRun,
	HideHelpCommand: true,
}

func handleEnrichmentRunEnrich(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.EnrichmentRunEnrichParams{
		EntitySlug: githubcomlightfldlightfieldgo.EnrichmentRunEnrichParamsEntitySlug(cmd.Value("entity-slug").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EnrichmentRun.Enrich(
		ctx,
		cmd.Value("entity-id").(string),
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
		Title:          "enrichment-run enrich",
		Transform:      transform,
	})
}

func handleEnrichmentRunGetEnrichmentRun(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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
	_, err = client.EnrichmentRun.GetEnrichmentRun(ctx, cmd.Value("run-id").(string), options...)
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
		Title:          "enrichment-run get-enrichment-run",
		Transform:      transform,
	})
}
