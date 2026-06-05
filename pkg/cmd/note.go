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

var noteCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new note record.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values for the new note. `$title` is required; `$content` is optional. See **[Fields and relationships](/using-the-api/fields-and-relationships/)** for value type details.",
			Required: true,
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationships to set on the new note. System relationships use a `$` prefix (e.g. `$account`, `$opportunity`, `$contact`); custom relationships use their bare slug. Each value is a single entity ID or an array of IDs. The note author is automatically set to the API key owner.",
			BodyPath: "relationships",
		},
	},
	Action:          handleNoteCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[string]{
			Name:       "fields.title",
			Usage:      "Title of the note.",
			InnerField: "$title",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.content",
			Usage:      "Content of the note as markdown formatted text.",
			InnerField: "$content",
		},
	},
})

var noteRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single note by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the note to retrieve.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleNoteRetrieve,
	HideHelpCommand: true,
}

var noteUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing note by ID. Only included fields and relationships are\nmodified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the note to update.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values to update — only provided fields are modified; omitted fields are left unchanged. See **[Fields and relationships](/using-the-api/fields-and-relationships/)** for value type details.",
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationship operations to apply. System relationships use a `$` prefix (e.g. `$account`, `$opportunity`, `$contact`); custom relationships use their bare slug. Each value is an operation object with `add` or `remove`.",
			BodyPath: "relationships",
		},
	},
	Action:          handleNoteUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[*string]{
			Name:       "fields.content",
			Usage:      "Content of the note as markdown formatted text.",
			InnerField: "$content",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.title",
			Usage:      "Title of the note.",
			InnerField: "$title",
		},
	},
})

var noteList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of notes. Use `offset` and `limit` to paginate through\nresults. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for more\ninformation about pagination.",
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
	Action:          handleNoteList,
	HideHelpCommand: true,
}

var noteDelete = cli.Command{
	Name:    "delete",
	Usage:   "Moves a note to the trash. The note is soft-deleted and may be restored from the\nLightfield UI.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the note to delete.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleNoteDelete,
	HideHelpCommand: true,
}

var noteDefinitions = cli.Command{
	Name:            "definitions",
	Usage:           "Returns the schema for the field and relationship definitions available on\nnotes. Useful for understanding the shape of note data before creating or\nupdating records. See\n<u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for\nmore details.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleNoteDefinitions,
	HideHelpCommand: true,
}

func handleNoteCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.NoteNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Note.New(ctx, params, options...)
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
		Title:          "note create",
		Transform:      transform,
	})
}

func handleNoteRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Note.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "note retrieve",
		Transform:      transform,
	})
}

func handleNoteUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.NoteUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Note.Update(
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
		Title:          "note update",
		Transform:      transform,
	})
}

func handleNoteList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.NoteListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Note.List(ctx, params, options...)
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
		Title:          "note list",
		Transform:      transform,
	})
}

func handleNoteDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.NoteDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Note.Delete(
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
		Title:          "note delete",
		Transform:      transform,
	})
}

func handleNoteDefinitions(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Note.Definitions(ctx, options...)
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
		Title:          "note definitions",
		Transform:      transform,
	})
}
