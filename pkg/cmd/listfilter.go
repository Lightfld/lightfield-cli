package cmd

import (
	"fmt"
	"strings"

	"github.com/Lightfld/lightfield-go/option"
	"github.com/urfave/cli/v3"
)

const filterFlagName = "filter"

// filterValues keeps each --filter occurrence whole. A cli.StringSliceFlag
// would split on commas, breaking values such as `$name[equal]=Acme, Inc.`.
type filterValues []string

func (f *filterValues) Set(s string) error {
	*f = append(*f, s)
	return nil
}

func (f *filterValues) String() string { return strings.Join(*f, " ") }

func (f *filterValues) Get() any { return []string(*f) }

// Query keys a list command already sets through its own flags.
var reservedFilterKeys = map[string]string{
	"limit":  "--limit",
	"offset": "--offset",
	"status": "--status",
}

func init() {
	for _, c := range []*cli.Command{
		&accountList,
		&channelList,
		&contactList,
		&emailList,
		&meetingList,
		&memberList,
		&messageList,
		&noteList,
		&objectList,
		&opportunityList,
		&senderList,
		&sequenceList,
		&sequenceListEnrollments,
		&sequenceListMessages,
		&taskList,
	} {
		c.Flags = append(c.Flags, &cli.GenericFlag{
			Name:  filterFlagName,
			Usage: "Filter by a field, as `FIELD[OPERATOR]=VALUE`, e.g. '$stage[equal]=Qualified'; the operator defaults to equal. Repeat to require every filter. Single-quote each filter so the shell doesn't expand $.",
			Value: &filterValues{},
		})
	}
}

// listFilterOptions turns each --filter into a raw query parameter, which the
// API reads as a `$field[operator]=value` filter.
func listFilterOptions(cmd *cli.Command) ([]option.RequestOption, error) {
	filters, ok := cmd.Value(filterFlagName).([]string)
	if !ok {
		return nil, nil
	}
	var options []option.RequestOption
	for _, filter := range filters {
		key, value, found := strings.Cut(filter, "=")
		if !found || key == "" {
			return nil, fmt.Errorf("Invalid --filter %q: expected FIELD[OPERATOR]=VALUE, e.g. '$stage[equal]=Qualified'", filter)
		}
		if flag, reserved := reservedFilterKeys[key]; reserved {
			return nil, fmt.Errorf("Invalid --filter %q: use %s instead", filter, flag)
		}
		options = append(options, option.WithQueryAdd(key, value))
	}
	return options, nil
}
