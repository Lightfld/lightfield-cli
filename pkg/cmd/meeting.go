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

var meetingCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new meeting record. This endpoint only supports creation of meetings\nin the past. The `$title`, `$startDate`, and `$endDate` fields are required.\nOnly the `$transcript` relationship is writable on create; all other meeting\nrelationships are system-managed. The `$account` and `$opportunity`\nrelationships are read-only and are derived from the meeting's `$contact`\nattendees and the accounts/opportunities those contacts belong to. The response\nis privacy-aware and includes a read-only `accessLevel`. See\n<u>[Uploading meeting transcripts](/using-the-api/uploading-meeting-transcripts/)</u>\nfor the full file upload and transcript attachment flow.",
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
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[string]{
			Name:       "fields.end-date",
			Usage:      "The end time of the meeting in ISO 8601 format. Must be in the past.",
			InnerField: "$endDate",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fields.start-date",
			Usage:      "The start time of the meeting in ISO 8601 format. Must be in the past.",
			InnerField: "$startDate",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fields.title",
			Usage:      "The title of the meeting.",
			InnerField: "$title",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "fields.attendee-emails",
			Usage:      "A list of attendee email addresses.",
			InnerField: "$attendeeEmails",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.description",
			Usage:      "A description of the meeting.",
			InnerField: "$description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.meeting-url",
			Usage:      "The URL for the meeting.",
			InnerField: "$meetingUrl",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.organizer-email",
			Usage:      "The email address of the meeting organizer. This field accepts a single email address.",
			InnerField: "$organizerEmail",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.privacy-setting",
			Usage:      "The privacy setting for the meeting (`FULL` or `METADATA`).",
			InnerField: "$privacySetting",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[any]{
			Name:       "relationships.transcript",
			Usage:      "The file ID to attach and set as the meeting canonical transcript.",
			InnerField: "$transcript",
		},
	},
})

var meetingRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single meeting by its ID. `$transcript` is the file ID of the\ncanonical transcript. `$transcripts` lists the file ID of every transcript on\nthe meeting. `$note` lists the note IDs attached to the meeting. Meeting fields\nand transcript visibility are redacted based on the caller-specific privacy\nresolution, and the response includes a read-only `accessLevel`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the meeting to retrieve.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMeetingRetrieve,
	HideHelpCommand: true,
}

var meetingUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing meeting by ID. Only included fields and relationships are\nmodified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the meeting to update.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values to update. Only `$privacySetting`, `$attendeeEmails`, `$meetingPrep` and `$postMeetingSummary` are writable, and omitted fields are left unchanged.",
			BodyPath: "fields",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "relationships",
			Usage:    "Relationship operations to apply. Only `$transcript`, with exactly one of `replace` or `remove`, is supported; clearing `$transcript` is not supported.",
			BodyPath: "relationships",
		},
	},
	Action:          handleMeetingUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[[]string]{
			Name:       "fields.attendee-emails",
			Usage:      "The full list of attendee email addresses, up to 500. Only meetings that have ended can be edited.",
			InnerField: "$attendeeEmails",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.meeting-prep",
			Usage:      "The meeting prep, as Markdown. Lightfield regenerates it as new activity arrives, which replaces an edit; null or an empty string clears it until then.",
			InnerField: "$meetingPrep",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.post-meeting-summary",
			Usage:      "The meeting summary, as Markdown. Lightfield regenerates it as new activity arrives, which replaces an edit; null or an empty string clears it until then.",
			InnerField: "$postMeetingSummary",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.privacy-setting",
			Usage:      "The privacy setting for the meeting.",
			InnerField: "$privacySetting",
		},
	},
	"relationships": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "relationships.transcript",
			InnerField: "$transcript",
		},
	},
})

var meetingList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of meetings. Results are sorted by when the record was\ncreated in Lightfield (`createdAt`), newest first. This is not the meeting start\ntime. Use `offset` and `limit` to paginate through results. Each meeting is\nprivacy-filtered per caller, includes a read-only `accessLevel`, and may redact\ntranscript or content fields based on the caller-specific privacy resolution. As\nin the Lightfield app, a meeting is listed only when it is internal-only or\nlinked to an account or an active contact; this applies to every caller,\nincluding admins and workspace API keys. Pass `status=trashed` to list meetings\ncurrently in the trash instead of live records. Listing trashed records\nadditionally requires the `meetings:restore` scope. See\n<u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information\nabout pagination.",
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
	Action:          handleMeetingList,
	HideHelpCommand: true,
}

var meetingDelete = cli.Command{
	Name:    "delete",
	Usage:   "Moves a meeting to the trash. The meeting is soft-deleted and may be restored\nfrom the Lightfield UI. Only callers who can edit the meeting (meeting\nparticipants or workspace admins) may delete it. Calling delete on a meeting\nthat is already trashed returns a 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the meeting to delete.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleMeetingDelete,
	HideHelpCommand: true,
}

var meetingCreateRecording = cli.Command{
	Name:    "create-recording",
	Usage:   "A recording is one session of the meeting: its transcript and, when there is\none, its audio or video. Recordings come from the Lightfield notetaker, Gong,\nGranola, or an upload. Only uploaded recordings can be created or given media\nthrough the API. Creates an uploaded recording from files already uploaded\nthrough `POST /v1/files`: a transcript (purpose `meeting_transcript`, required)\nand optionally its audio or video (purpose `meeting_recording`). Both files\nattach in one write, so a failed request attaches neither. Media can also be\nadded later by updating the recording. Sending a transcript file that is already\non the meeting returns that recording. Only meeting participants and workspace\nowners may upload.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the meeting.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "transcript-file-id",
			Usage:    "File ID of a completed upload with purpose `meeting_transcript`.",
			Required: true,
			BodyPath: "transcriptFileId",
		},
		&requestflag.Flag[bool]{
			Name:     "is-canonical",
			Usage:    "Pin this as the meeting's primary recording. Defaults to `true`. `false` attaches it without pinning; on a meeting with no pinned recording, the primary is chosen by source (notetaker, Gong, Granola, then uploads) and then recency, so a newer upload can still become primary.",
			Default:  true,
			BodyPath: "isCanonical",
		},
		&requestflag.Flag[string]{
			Name:     "media-file-id",
			Usage:    "File ID of a completed upload with purpose `meeting_recording`. Can also be added later by updating the recording.",
			BodyPath: "mediaFileId",
		},
	},
	Action:          handleMeetingCreateRecording,
	HideHelpCommand: true,
}

var meetingDefinitions = cli.Command{
	Name:            "definitions",
	Usage:           "Returns the schema for the field and relationship definitions available on\nmeetings. Useful for understanding the shape of meeting data before creating or\nupdating records. See\n<u>[Fields and relationships](/using-the-api/fields-and-relationships/)</u> for\nmore details.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMeetingDefinitions,
	HideHelpCommand: true,
}

var meetingListRecordings = cli.Command{
	Name:    "list-recordings",
	Usage:   "A recording is one session of the meeting: its transcript and, when there is\none, its audio or video. Recordings come from the Lightfield notetaker, Gong,\nGranola, or an upload. Only uploaded recordings can be created or given media\nthrough the API. Returns every recording on the meeting without media URLs;\nretrieve a single recording for its media. Requires full access to the meeting;\na caller who sees only its metadata gets a 403.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the meeting.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMeetingListRecordings,
	HideHelpCommand: true,
}

var meetingRestore = cli.Command{
	Name:    "restore",
	Usage:   "Restores a trashed meeting, bringing it back to active. Find trashed meetings\nvia `GET /v1/meetings?status=trashed`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the meeting to restore.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleMeetingRestore,
	HideHelpCommand: true,
}

var meetingRetrieveRecording = cli.Command{
	Name:    "retrieve-recording",
	Usage:   "A recording is one session of the meeting: its transcript and, when there is\none, its audio or video. Recordings come from the Lightfield notetaker, Gong,\nGranola, or an upload. Only uploaded recordings can be created or given media\nthrough the API. Returns the recording with a short-lived `media.url` for its\naudio or video. For notetaker and Gong recordings the URL is fetched from the\nprovider on each call. Requires full access to the meeting; a caller who sees\nonly its metadata gets a 403.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the meeting.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "recording-id",
			Usage:     "Unique identifier of the recording.",
			Required:  true,
			PathParam: "recordingId",
		},
	},
	Action:          handleMeetingRetrieveRecording,
	HideHelpCommand: true,
}

var meetingUpdateRecording = cli.Command{
	Name:    "update-recording",
	Usage:   "A recording is one session of the meeting: its transcript and, when there is\none, its audio or video. Recordings come from the Lightfield notetaker, Gong,\nGranola, or an upload. Only uploaded recordings can be created or given media\nthrough the API. Changes only what is sent. `isCanonical: true` makes the\nrecording the meeting primary, for any provider. `mediaFileId` adds or replaces\nthe audio or video of an uploaded recording; a replaced file is deleted.\nNotetaker, Gong and Granola recordings cannot be given media. Only meeting\nparticipants and workspace owners may update.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the meeting.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "recording-id",
			Usage:     "Unique identifier of the recording.",
			Required:  true,
			PathParam: "recordingId",
		},
		&requestflag.Flag[bool]{
			Name:     "is-canonical",
			Usage:    "`true` pins this as the meeting's primary recording. Works for every provider, but the recording needs a transcript file (`transcript` is not null). `false` is rejected: to move the primary, set it on another recording.",
			BodyPath: "isCanonical",
		},
		&requestflag.Flag[string]{
			Name:     "media-file-id",
			Usage:    "File ID of a completed upload with purpose `meeting_recording`. Adds or replaces the media; a replaced file is deleted. Uploaded recordings only.",
			BodyPath: "mediaFileId",
		},
	},
	Action:          handleMeetingUpdateRecording,
	HideHelpCommand: true,
}

func handleMeetingCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MeetingNewParams{}

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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "meeting create",
		Transform:      transform,
	})
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "meeting retrieve",
		Transform:      transform,
	})
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

	params := githubcomlightfldlightfieldgo.MeetingUpdateParams{}

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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "meeting update",
		Transform:      transform,
	})
}

func handleMeetingList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MeetingListParams{}

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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "meeting list",
		Transform:      transform,
	})
}

func handleMeetingDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MeetingDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Meeting.Delete(
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
		Title:          "meeting delete",
		Transform:      transform,
	})
}

func handleMeetingCreateRecording(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MeetingNewRecordingParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Meeting.NewRecording(
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
		Title:          "meeting create-recording",
		Transform:      transform,
	})
}

func handleMeetingDefinitions(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Meeting.Definitions(ctx, options...)
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
		Title:          "meeting definitions",
		Transform:      transform,
	})
}

func handleMeetingListRecordings(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Meeting.ListRecordings(ctx, cmd.Value("id").(string), options...)
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
		Title:          "meeting list-recordings",
		Transform:      transform,
	})
}

func handleMeetingRestore(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.MeetingRestoreParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Meeting.Restore(
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
		Title:          "meeting restore",
		Transform:      transform,
	})
}

func handleMeetingRetrieveRecording(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-id") && len(unusedArgs) > 0 {
		cmd.Set("recording-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.MeetingGetRecordingParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Meeting.GetRecording(
		ctx,
		cmd.Value("recording-id").(string),
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
		Title:          "meeting retrieve-recording",
		Transform:      transform,
	})
}

func handleMeetingUpdateRecording(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("recording-id") && len(unusedArgs) > 0 {
		cmd.Set("recording-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.MeetingUpdateRecordingParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Meeting.UpdateRecording(
		ctx,
		cmd.Value("recording-id").(string),
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
		Title:          "meeting update-recording",
		Transform:      transform,
	})
}
