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

var emailRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single email by its ID. Email fields are redacted based on the\ncaller-specific privacy resolution, and the response includes a read-only\n`accessLevel`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the email to retrieve.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEmailRetrieve,
	HideHelpCommand: true,
}

var emailList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of emails. Results are sorted by when the record was\ncreated in Lightfield (`createdAt`), newest first. This is not the time the\nemail was sent or received. Use `offset` and `limit` to paginate through\nresults. Each email is privacy-filtered per caller, includes a read-only\n`accessLevel`, and may redact the subject at metadata-only access.\n`fields.$body` is not included on list items; use GET `/v1/emails/{id}` for\nmessage body HTML. As in the Lightfield app, an email is listed only when it is\nlinked to an account that is not trashed; this applies to every caller,\nincluding admins and workspace API keys. See\n<u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information\nabout pagination.",
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
	Action:          handleEmailList,
	HideHelpCommand: true,
}

var emailDelete = cli.Command{
	Name:    "delete",
	Usage:   "Moves an email logged through `POST /v1/emails/log` to the trash. Synced emails\ncannot be deleted here. Only participants on the email and workspace owners may\ndelete it.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the email to delete.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleEmailDelete,
	HideHelpCommand: true,
}

var emailDraft = requestflag.WithInnerFlags(cli.Command{
	Name:    "draft",
	Usage:   "Creates a draft in the connected email account that owns the `from` address (its\nprimary address or one of its verified aliases). Mirrors native email-client\nbehavior: only `from` is required — `to`, `cc`, `bcc`, `subject`, `messageBody`,\nand `attachments` are all optional. At least one of those optional fields must\nbe populated; sending only `from` returns a 400.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Bare email address (no display name). The primary address or a verified alias of a connected email account owned by the API key user. Compared case-insensitively. Mailbox where the draft is created.",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[[]string]{
			Name:     "attachment",
			Usage:    "Optional list of file IDs (uploaded via the Files API) to attach to the draft. Maximum 5 attachments per draft. Gmail: up to 20 MB combined raw attachment size. Outlook: up to 3 MB per attachment. Uploads with purpose email_attachment accept up to 20 MB per file; sender-specific limits apply when sending or creating a draft.",
			BodyPath: "attachments",
		},
		&requestflag.Flag[[]string]{
			Name:     "bcc",
			Usage:    "Bcc recipients (same shape as `to`).",
			BodyPath: "bcc",
		},
		&requestflag.Flag[[]string]{
			Name:     "cc",
			Usage:    "Cc recipients (same shape as `to`).",
			BodyPath: "cc",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "message-body",
			Usage:    "Email message body (HTML or plain text).",
			BodyPath: "messageBody",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Email subject.",
			BodyPath: "subject",
		},
		&requestflag.Flag[[]string]{
			Name:     "to",
			Usage:    "Recipient email addresses (bare, no display names). Up to 500.",
			BodyPath: "to",
		},
	},
	Action:          handleEmailDraft,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message-body": {
		&requestflag.InnerFlag[string]{
			Name:       "message-body.content",
			Usage:      "Email body content.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message-body.content-type",
			Usage:      "Defaults to `HTML`.",
			InnerField: "contentType",
		},
	},
})

var emailLog = requestflag.WithInnerFlags(cli.Command{
	Name:    "log",
	Usage:   "Logs an email that no connected account will sync, such as history from a\nprevious CRM or a mailbox export. The email appears on account, contact, and\nopportunity timelines at `$sentAt` and threads with related emails. `$sentAt`\nmust be in the past, and at least one participant must be a workspace member and\none must be external.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "fields",
			Usage:    "Field values for the logged email. System fields use a `$` prefix. Required: `$sentAt` and `$from`, plus at least one recipient across `$to`, `$cc`, and `$bcc`.",
			Required: true,
			BodyPath: "fields",
		},
		&requestflag.Flag[bool]{
			Name:     "auto-create-records",
			Usage:    "When true, the post-create sync may create account and contact records for unknown external participants.",
			BodyPath: "autoCreateRecords",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			Usage:    "Raw RFC 5322 headers, passed through as read (any key case). Only `Message-ID`, `In-Reply-To`, and `References` are used. At most 64 KB.",
			BodyPath: "headers",
		},
	},
	Action:          handleEmailLog,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"fields": {
		&requestflag.InnerFlag[string]{
			Name:       "fields.from",
			Usage:      "The sender address. This field accepts a single email address.",
			InnerField: "$from",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fields.sent-at",
			Usage:      "When the email was sent, in ISO 8601 format. Must be in the past.",
			InnerField: "$sentAt",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "fields.bcc",
			Usage:      "Bcc-recipient addresses.",
			InnerField: "$bcc",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "fields.body",
			Usage:      "The email body. Pass whichever parts you have.",
			InnerField: "$body",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "fields.cc",
			Usage:      "Cc-recipient addresses.",
			InnerField: "$cc",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "fields.privacy-setting",
			Usage:      "The privacy setting for the email (`FULL` or `METADATA`). Defaults to the most permissive setting among the participants’ connected accounts, else `METADATA`.",
			InnerField: "$privacySetting",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fields.subject",
			Usage:      "The email subject line.",
			InnerField: "$subject",
		},
		&requestflag.InnerFlag[string]{
			Name:       "fields.thread-id",
			Usage:      "Any stable grouping key. Emails with the same value share a thread. Overrides threading derived from `headers`.",
			InnerField: "$threadId",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "fields.to",
			Usage:      "To-recipient addresses.",
			InnerField: "$to",
		},
	},
})

var emailRestore = cli.Command{
	Name:    "restore",
	Usage:   "Restores a trashed email that was logged through `POST /v1/emails/log`, bringing\nit back to active. Synced emails cannot be restored here. Calling restore on an\nemail that is not trashed is a no-op and returns the existing record.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the email to restore.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleEmailRestore,
	HideHelpCommand: true,
}

var emailSend = requestflag.WithInnerFlags(cli.Command{
	Name:    "send",
	Usage:   "Sends an email via the connected email account that owns the `from` address (its\nprimary address or one of its verified aliases). Currently supports new sends\nonly; replies and forwards are not yet supported.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Bare email address (no display name). The primary address or a verified alias of a connected email account owned by the API key user. Compared case-insensitively. Used as the From header when sending.",
			Required: true,
			BodyPath: "from",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "message-body",
			Usage:    "Email message body (HTML or plain text).",
			Required: true,
			BodyPath: "messageBody",
		},
		&requestflag.Flag[string]{
			Name:     "subject",
			Usage:    "Email subject. Cannot be empty.",
			Required: true,
			BodyPath: "subject",
		},
		&requestflag.Flag[[]string]{
			Name:     "to",
			Usage:    "Recipient email addresses (bare, no display names). At least 1, at most 500.",
			Required: true,
			BodyPath: "to",
		},
		&requestflag.Flag[[]string]{
			Name:     "attachment",
			Usage:    "Optional list of file IDs (uploaded via the Files API) to attach to the email. Maximum 5 attachments per email. Gmail: up to 20 MB combined raw attachment size. Outlook: up to 3 MB per attachment. Uploads with purpose email_attachment accept up to 20 MB per file; sender-specific limits apply when sending or creating a draft.",
			BodyPath: "attachments",
		},
		&requestflag.Flag[[]string]{
			Name:     "bcc",
			Usage:    "Bcc recipients (same shape as `to`).",
			BodyPath: "bcc",
		},
		&requestflag.Flag[[]string]{
			Name:     "cc",
			Usage:    "Cc recipients (same shape as `to`).",
			BodyPath: "cc",
		},
	},
	Action:          handleEmailSend,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message-body": {
		&requestflag.InnerFlag[string]{
			Name:       "message-body.content",
			Usage:      "Email body content.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message-body.content-type",
			Usage:      "Defaults to `HTML`.",
			InnerField: "contentType",
		},
	},
})

func handleEmailRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Email.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "email retrieve",
		Transform:      transform,
	})
}

func handleEmailList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.EmailListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.List(ctx, params, options...)
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
		Title:          "email list",
		Transform:      transform,
	})
}

func handleEmailDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.EmailDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.Delete(
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
		Title:          "email delete",
		Transform:      transform,
	})
}

func handleEmailDraft(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.EmailDraftParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.Draft(ctx, params, options...)
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
		Title:          "email draft",
		Transform:      transform,
	})
}

func handleEmailLog(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.EmailLogParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.Log(ctx, params, options...)
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
		Title:          "email log",
		Transform:      transform,
	})
}

func handleEmailRestore(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.EmailRestoreParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.Restore(
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
		Title:          "email restore",
		Transform:      transform,
	})
}

func handleEmailSend(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.EmailSendParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Email.Send(ctx, params, options...)
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
		Title:          "email send",
		Transform:      transform,
	})
}
