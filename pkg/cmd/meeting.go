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

var meetingCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new meeting record. This endpoint only supports creation of meetings\nin the past. The `$title`, `$startDate`, and `$endDate` fields are required.\nOnly the `$transcript` relationship is writable on create; all other meeting\nrelationships are system-managed. The response is privacy-aware and includes a\nread-only `accessLevel`. See\n<u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>\nfor the full file upload and transcript attachment flow.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values for the new MANUAL meeting. System fields use a `$` prefix (for example `$title`, `$startDate`, `$endDate`). Required: `$title`, `$startDate`, and `$endDate`. `$organizerEmail` accepts a single email address when provided; `$attendeeEmails` accepts an array of email addresses. See <u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for value type details.",
			Required: true,
			BodyPath: "fields",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-create-records",
			Usage:    "When true, the initial post-create meeting sync may auto-create account and contact records for external attendees.",
			BodyPath: "autoCreateRecords",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationships to set on the new meeting. Only `$transcript` is writable on create; all other meeting relationships are system-managed.",
			BodyPath: "relationships",
		},
	},
	Action:          handleMeetingCreate,
	HideHelpCommand: true,
}

var meetingRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single meeting by its ID. Meeting fields and transcript visibility\nare redacted based on the caller-specific privacy resolution, and the response\nincludes a read-only `accessLevel`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique identifier of the meeting to retrieve.",
			Required: true,
		},
	},
	Action:          handleMeetingRetrieve,
	HideHelpCommand: true,
}

var meetingUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates an existing meeting by ID. Only included fields and relationships are\nmodified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Unique identifier of the meeting to update.",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values to update. Only `$privacySetting` is writable, and omitted fields are left unchanged.",
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationship operations to apply. Only `$transcript.replace` is supported; removing or clearing `$transcript` is not supported.",
			BodyPath: "relationships",
		},
	},
	Action:          handleMeetingUpdate,
	HideHelpCommand: true,
}

var meetingList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of meetings. Use `offset` and `limit` to paginate\nthrough results. Each meeting is privacy-filtered per caller, includes a\nread-only `accessLevel`, and may redact transcript or content fields based on\nthe caller-specific privacy resolution. See\n<u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information\nabout pagination.",
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
	Action:          handleMeetingList,
	HideHelpCommand: true,
}

func handleMeetingCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomlightfldlightfieldgo.MeetingNewParams{}

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
	_, err = client.Meeting.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "meeting create", obj, format, explicitFormat, transform)
}

func handleMeetingRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Meeting.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "meeting retrieve", obj, format, explicitFormat, transform)
}

func handleMeetingUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomlightfldlightfieldgo.MeetingUpdateParams{}

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
	_, err = client.Meeting.Update(
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
	return ShowJSON(os.Stdout, os.Stderr, "meeting update", obj, format, explicitFormat, transform)
}

func handleMeetingList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomlightfldlightfieldgo.MeetingListParams{}

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
	_, err = client.Meeting.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "meeting list", obj, format, explicitFormat, transform)
}
