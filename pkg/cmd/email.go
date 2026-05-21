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

var emailSend = requestflag.WithInnerFlags(cli.Command{
	Name:    "send",
	Usage:   "Sends an email via the connected email account that owns the `from` address.\nCurrently supports new sends only; replies and forwards are not yet supported.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			Required: true,
			BodyPath: "body",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "Bare email address (no display name). Must match a connected email account owned by the API key user. Compared case-insensitively. Used as the From header when sending.",
			Required: true,
			BodyPath: "from",
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
			Usage:    "Optional list of file IDs (uploaded via the Files API) to attach to the email. Maximum 5 attachments per email. Each attachment must be ≤ 3MB and the total across all attachments must be ≤ 25MB.",
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
	"body": {
		&requestflag.InnerFlag[string]{
			Name:       "body.content",
			Usage:      "Email body content.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.content-type",
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
