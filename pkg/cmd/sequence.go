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

var sequenceCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new sequence and its first draft, and returns the draft. The sequence\nsends nothing until the draft is published and the sequence is activated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "definition",
			Usage:    "Configuration of the first draft. `description` and `steps` are required.",
			Required: true,
			BodyPath: "definition",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the sequence.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "reviewer",
			Usage:    "Extra members who may approve this sequence’s `FOR_REVIEW` messages before they send. The sending member and the sequence’s creator always review, whether or not they are listed. Member IDs of active members. Omit for none.",
			BodyPath: "reviewers",
		},
		&requestflag.Flag[[]string]{
			Name:     "sender",
			Usage:    "The senders this sequence sends from; new contacts rotate across them. Sender IDs of unarchived senders. Omit for none; activating needs a sender for each channel the steps use.",
			BodyPath: "senders",
		},
	},
	Action:          handleSequenceCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"definition": {
		&requestflag.InnerFlag[string]{
			Name:       "definition.description",
			Usage:      "What this version of the sequence is for.",
			InnerField: "description",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "definition.steps",
			Usage:      "The ordered step list.",
			InnerField: "steps",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "definition.messaging",
			Usage:      "Messaging inputs that frame every step. Omitted or null fields are left unset.",
			InnerField: "messaging",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "definition.pacing",
			Usage:      "How the whole sequence paces each step type it sends, keyed by sequence step type. Omitted step types have no daily cap.",
			InnerField: "pacing",
		},
		&requestflag.InnerFlag[string]{
			Name:       "definition.priority",
			Usage:      "Send priority relative to other sequences. Defaults to `NORMAL`.",
			InnerField: "priority",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "definition.send-window",
			Usage:      "When the sequence may send. Omitted fields take the default: 9:00–17:00, Monday to Friday, UTC.",
			InnerField: "sendWindow",
		},
	},
})

var sequenceRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single sequence by its ID. Edits that have not been published are\nnot included; those can be read with\n[Retrieve a sequence draft](/api/resources/sequence/methods/retrieveDraft).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence to retrieve.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSequenceRetrieve,
	HideHelpCommand: true,
}

var sequenceList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of sequences, including ones that have never been\npublished. Results are sorted by when the sequence was created (`createdAt`),\nnewest first. Use `offset` and `limit` to paginate through results, and\nbracketed filters to narrow the list. See\n<u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information\nabout pagination.",
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
	Action:          handleSequenceList,
	HideHelpCommand: true,
}

var sequenceActivate = cli.Command{
	Name:    "activate",
	Usage:   "Starts sending a published sequence to its enrolled contacts and returns it.\nRequires a published version and an available sender for each channel the steps\nuse; senders are added through the draft. An active sequence is returned\nunchanged. The sequence’s status changes immediately; its enrollments’ statuses\nfollow asynchronously. Returns 400 `sequence_not_published` before the first\npublish, 400 `no_selectable_sender` when a channel has no available sender, and\n409 when the sequence changed while the request ran: re-read it to see the\nchange, then retry.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleSequenceActivate,
	HideHelpCommand: true,
}

var sequenceDeactivate = cli.Command{
	Name:    "deactivate",
	Usage:   "Pauses sending for a published sequence and returns it. Enrollments stay in\nplace and resume when the sequence is activated again; a paused sequence is\nreturned unchanged. The sequence’s status changes immediately; its enrollments’\nstatuses follow asynchronously. Returns 400 `sequence_not_published` before the\nfirst publish, and 409 when the sequence changed while the request ran: re-read\nit to see the change, then retry.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handleSequenceDeactivate,
	HideHelpCommand: true,
}

var sequenceDiscardDraft = cli.Command{
	Name:    "discard-draft",
	Usage:   "Discards the open draft at `expectedVersion`. A published sequence keeps its\npublished version. A sequence that has never been published is deleted with its\ndraft, and `sequenceDeleted` is true. Returns 404 `draft_not_found` when no\ndraft is open, and 409 when `expectedVersion` is not the open draft’s `version`;\nretrieve the draft and retry.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "expected-version",
			Usage:    "The open draft’s `version` as last read.",
			Required: true,
			BodyPath: "expectedVersion",
		},
	},
	Action:          handleSequenceDiscardDraft,
	HideHelpCommand: true,
}

var sequenceListEnrollments = cli.Command{
	Name:    "list-enrollments",
	Usage:   "Returns a paginated list of a sequence’s enrollments, including completed and\ncancelled ones. Results are sorted by when the contact was enrolled\n(`enrolledAt`), newest first. Test-run enrollments are not included. Use\n`offset` and `limit` to paginate through results, and bracketed filters to\nnarrow the list. See <u>[List endpoints](/using-the-api/list-endpoints/)</u> for\nmore information about pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
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
	Action:          handleSequenceListEnrollments,
	HideHelpCommand: true,
}

var sequenceListMessages = cli.Command{
	Name:    "list-messages",
	Usage:   "Returns a paginated list of a sequence’s messages: every outbound email and\nLinkedIn message, and each enrollment’s first reply. Results are sorted by\n`createdAt`, newest first. Superseded outbound attempts are not included, and\nlist items omit content; retrieve a message to read it. Use `offset` and `limit`\nto paginate through results, and bracketed filters to narrow the list. See\n<u>[List endpoints](/using-the-api/list-endpoints/)</u> for more information\nabout pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
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
	Action:          handleSequenceListMessages,
	HideHelpCommand: true,
}

var sequenceListTemplateVariables = cli.Command{
	Name:            "list-template-variables",
	Usage:           "Lists the variables a sequence template may use: the first and last name, the\ncontact’s and its first account’s fields of a type a message can print, and the\nsender’s name and email. Use an entry’s `ref` in a send step’s subject or body.\nA variable needs a fallback for contacts missing the value: `{{ {attr:ad_…} }}`\nis written `{{ coalesce({attr:ad_…}, \"fallback\") }}`, and one without a fallback\nis refused when the draft is saved. Anything else inside `{{ }}` is refused too.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleSequenceListTemplateVariables,
	HideHelpCommand: true,
}

var sequencePublishDraft = cli.Command{
	Name:    "publish-draft",
	Usage:   "Publishes the open draft at `expectedVersion` and returns the sequence. A first\npublish leaves the sequence `DEACTIVATED`; activate it to start sending.\nContacts already enrolled continue on the new version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "expected-version",
			Usage:    "The open draft’s `version` as last read.",
			Required: true,
			BodyPath: "expectedVersion",
		},
	},
	Action:          handleSequencePublishDraft,
	HideHelpCommand: true,
}

var sequenceRetrieveDraft = cli.Command{
	Name:    "retrieve-draft",
	Usage:   "Retrieves the sequence’s open draft, the same one editors see in Lightfield.\nReturns 404 `draft_not_found` when no draft is open.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSequenceRetrieveDraft,
	HideHelpCommand: true,
}

var sequenceRetrieveEnrollment = cli.Command{
	Name:    "retrieve-enrollment",
	Usage:   "Retrieves a single enrollment by its ID: the contact’s position in the cadence\nand its latest activity. Test-run enrollments are not returned.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "enrollment-id",
			Usage:     "Unique identifier of the enrollment to retrieve.",
			Required:  true,
			PathParam: "enrollmentId",
		},
	},
	Action:          handleSequenceRetrieveEnrollment,
	HideHelpCommand: true,
}

var sequenceRetrieveMessage = cli.Command{
	Name:    "retrieve-message",
	Usage:   "Retrieves a single sequence message by its ID: an outbound email or LinkedIn\nmessage, or an enrollment’s first reply, with its plain-text content. Content is\nreturned with `sequences:read` regardless of the email’s privacy; see\n<u>[Sequence message access](/using-the-api/scopes/#sequence-message-access)</u>.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "message-id",
			Usage:     "Unique identifier of the message to retrieve.",
			Required:  true,
			PathParam: "messageId",
		},
	},
	Action:          handleSequenceRetrieveMessage,
	HideHelpCommand: true,
}

var sequenceRetrievePerformance = cli.Command{
	Name:    "retrieve-performance",
	Usage:   "Retrieves a sequence’s lifetime engagement, totalled by step type and broken\ndown per step, with the live enrollments currently on each step. Test runs are\nnot counted. A sequence that has never been published reports zero counts and no\nsteps.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleSequenceRetrievePerformance,
	HideHelpCommand: true,
}

var sequenceRetrieveTest = cli.Command{
	Name:    "retrieve-test",
	Usage:   "Retrieves the result of a test started with\n[Start a sequence test](/api/resources/sequence/methods/startTest): its status,\nand for every step of the tested version whether it composed a message, was\nskipped or failed. Poll it until `status` is `COMPLETED` or `FAILED`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "test-id",
			Usage:     "The `id` returned when the test was started.",
			Required:  true,
			PathParam: "testId",
		},
	},
	Action:          handleSequenceRetrieveTest,
	HideHelpCommand: true,
}

var sequenceStartTest = cli.Command{
	Name:    "start-test",
	Usage:   "Tries the sequence on one contact without sending anything: the contact is\nenrolled in a test, its steps run back to back with waits skipped, and each\nmessage is composed from the contact’s real data and kept as a draft. Use it to\ncheck what a step, a template variable or its fallback produces before\npublishing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "contact-id",
			Usage:    "ID of the contact to test the sequence on. The test composes each message from this contact’s real data and sends nothing.",
			Required: true,
			BodyPath: "contactId",
		},
		&requestflag.Flag[string]{
			Name:     "expected-version",
			Usage:    "The open draft’s `version` as last read. When sent, the test runs exactly that draft, and is refused with 409 if the draft has changed or none is open. Omit it to test whatever is open now.",
			BodyPath: "expectedVersion",
		},
	},
	Action:          handleSequenceStartTest,
	HideHelpCommand: true,
}

var sequenceUpdateDraft = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-draft",
	Usage:   "Saves changes to the sequence’s open draft, or opens one from the published\nversion when none is open. Only the fields sent are changed; `definition.steps`\nreplaces the whole step list. Changes take effect when the draft is published.\nSee\n<u>[Using sequences with the API](/sequences/using-the-api#editing-a-draft)</u>\nfor the draft and version model.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "definition",
			Usage:    "Fields to change. Omitted fields are unchanged.",
			BodyPath: "definition",
		},
		&requestflag.Flag[string]{
			Name:     "expected-version",
			Usage:    "The open draft’s `version` as last read. Required when a draft is open; omit it only when none is, to start one from the published version.",
			BodyPath: "expectedVersion",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the sequence.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "reviewers",
			Usage:    "Extra members who may approve this sequence’s `FOR_REVIEW` messages before they send. The sending member and the sequence’s creator always review, whether or not they are listed. Changes the list by adding and removing members; omit to keep it.",
			BodyPath: "reviewers",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "senders",
			Usage:    "The senders this sequence sends from; new contacts rotate across them. Changes the list by adding and removing senders; omit to keep it.",
			BodyPath: "senders",
		},
	},
	Action:          handleSequenceUpdateDraft,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"definition": {
		&requestflag.InnerFlag[string]{
			Name:       "definition.description",
			Usage:      "What this version of the sequence is for.",
			InnerField: "description",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "definition.messaging",
			Usage:      "Messaging inputs that frame every step. Omitted fields are unchanged; null clears a field.",
			InnerField: "messaging",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "definition.pacing",
			Usage:      "How the whole sequence paces each step type it sends, keyed by sequence step type. Omitted step types are unchanged.",
			InnerField: "pacing",
		},
		&requestflag.InnerFlag[string]{
			Name:       "definition.priority",
			Usage:      "Send priority relative to other sequences.",
			InnerField: "priority",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "definition.send-window",
			Usage:      "When the sequence may send. Omitted fields are unchanged.",
			InnerField: "sendWindow",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "definition.steps",
			Usage:      "The complete ordered step list, replacing the current one. Keep a step by sending its `id`; omit `id` for a new step.",
			InnerField: "steps",
		},
	},
	"reviewers": {
		&requestflag.InnerFlag[[]string]{
			Name:       "reviewers.add",
			Usage:      "Member IDs of active members to add. A member already listed is unchanged.",
			InnerField: "add",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "reviewers.remove",
			Usage:      "Member IDs to remove, including members who have left the workspace. A member not listed is unchanged.",
			InnerField: "remove",
		},
	},
	"senders": {
		&requestflag.InnerFlag[[]string]{
			Name:       "senders.add",
			Usage:      "Sender IDs of unarchived senders to add. A sender already listed is unchanged.",
			InnerField: "add",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "senders.remove",
			Usage:      "Sender IDs to remove, including archived senders. A sender not listed is unchanged.",
			InnerField: "remove",
		},
	},
})

var sequenceUpdateEnrollments = cli.Command{
	Name:    "update-enrollments",
	Usage:   "Enrolls the contacts in `add` and cancels the open enrollments of the contacts\nin `remove`, in one request: if any addition is invalid, nothing changes. Send\nadd, remove or both, with at most 25 contact IDs in total; duplicate IDs are\nignored, and an ID cannot be in both lists. The response has one result per\ncontact: `enrolled`, `already_enrolled`, `unenrolled`, or `no_active_enrollment`\nfor a removal with nothing open.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Unique identifier of the sequence.",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "add",
			Usage:    "Contact IDs to enroll. Open enrollments are unchanged; contacts whose prior enrollment ended are enrolled again.",
			BodyPath: "add",
		},
		&requestflag.Flag[[]string]{
			Name:     "remove",
			Usage:    "Contact IDs to unenroll. Cancels open enrollments and preserves their history.",
			BodyPath: "remove",
		},
	},
	Action:          handleSequenceUpdateEnrollments,
	HideHelpCommand: true,
}

func handleSequenceCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.New(ctx, params, options...)
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
		Title:          "sequence create",
		Transform:      transform,
	})
}

func handleSequenceRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Sequence.Get(ctx, cmd.Value("id").(string), options...)
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
		Title:          "sequence retrieve",
		Transform:      transform,
	})
}

func handleSequenceList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.List(ctx, params, options...)
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
		Title:          "sequence list",
		Transform:      transform,
	})
}

func handleSequenceActivate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceActivateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.Activate(
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
		Title:          "sequence activate",
		Transform:      transform,
	})
}

func handleSequenceDeactivate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceDeactivateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.Deactivate(
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
		Title:          "sequence deactivate",
		Transform:      transform,
	})
}

func handleSequenceDiscardDraft(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceDiscardDraftParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.DiscardDraft(
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
		Title:          "sequence discard-draft",
		Transform:      transform,
	})
}

func handleSequenceListEnrollments(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceListEnrollmentsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.ListEnrollments(
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
		Title:          "sequence list-enrollments",
		Transform:      transform,
	})
}

func handleSequenceListMessages(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceListMessagesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.ListMessages(
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
		Title:          "sequence list-messages",
		Transform:      transform,
	})
}

func handleSequenceListTemplateVariables(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Sequence.ListTemplateVariables(ctx, options...)
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
		Title:          "sequence list-template-variables",
		Transform:      transform,
	})
}

func handleSequencePublishDraft(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequencePublishDraftParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.PublishDraft(
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
		Title:          "sequence publish-draft",
		Transform:      transform,
	})
}

func handleSequenceRetrieveDraft(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Sequence.GetDraft(ctx, cmd.Value("id").(string), options...)
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
		Title:          "sequence retrieve-draft",
		Transform:      transform,
	})
}

func handleSequenceRetrieveEnrollment(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("enrollment-id") && len(unusedArgs) > 0 {
		cmd.Set("enrollment-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.SequenceGetEnrollmentParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.GetEnrollment(
		ctx,
		cmd.Value("enrollment-id").(string),
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
		Title:          "sequence retrieve-enrollment",
		Transform:      transform,
	})
}

func handleSequenceRetrieveMessage(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.SequenceGetMessageParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.GetMessage(
		ctx,
		cmd.Value("message-id").(string),
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
		Title:          "sequence retrieve-message",
		Transform:      transform,
	})
}

func handleSequenceRetrievePerformance(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Sequence.GetPerformance(ctx, cmd.Value("id").(string), options...)
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
		Title:          "sequence retrieve-performance",
		Transform:      transform,
	})
}

func handleSequenceRetrieveTest(ctx context.Context, cmd *cli.Command) error {
	client := githubcomlightfldlightfieldgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("test-id") && len(unusedArgs) > 0 {
		cmd.Set("test-id", unusedArgs[0])
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

	params := githubcomlightfldlightfieldgo.SequenceGetTestParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.GetTest(
		ctx,
		cmd.Value("test-id").(string),
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
		Title:          "sequence retrieve-test",
		Transform:      transform,
	})
}

func handleSequenceStartTest(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceStartTestParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.StartTest(
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
		Title:          "sequence start-test",
		Transform:      transform,
	})
}

func handleSequenceUpdateDraft(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceUpdateDraftParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.UpdateDraft(
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
		Title:          "sequence update-draft",
		Transform:      transform,
	})
}

func handleSequenceUpdateEnrollments(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomlightfldlightfieldgo.SequenceUpdateEnrollmentsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sequence.UpdateEnrollments(
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
		Title:          "sequence update-enrollments",
		Transform:      transform,
	})
}
