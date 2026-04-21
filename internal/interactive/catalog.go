package interactive

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/urfave/cli/v3"
)

type requestFlagDetails interface {
	GetQueryPath() string
	GetHeaderPath() string
	GetBodyPath() string
	IsBodyRoot() bool
}

type requiredFlagOrStdin interface {
	IsRequiredAsFlagOrStdin() bool
}

var (
	htmlTagPattern      = regexp.MustCompile(`</?u>`)
	markdownLinkPattern = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
)

func BuildCatalog(root *cli.Command) Catalog {
	catalog := Catalog{
		GlobalFields: extractFields(root.Flags, true, nil),
		Resources:    make([]ResourceSpec, 0, len(root.Commands)),
	}

	for _, resource := range root.Commands {
		if resource == nil || resource.Hidden || strings.HasPrefix(resource.Name, "@") || strings.HasPrefix(resource.Name, "__") {
			continue
		}

		res := ResourceSpec{
			Name:     resource.Name,
			Category: resource.Category,
			Usage:    resource.Usage,
			Actions:  make([]ActionSpec, 0, len(resource.Commands)),
		}

		for _, action := range resource.Commands {
			if action == nil || action.Hidden || action.Name == "help" {
				continue
			}
			res.Actions = append(res.Actions, ActionSpec{
				ResourceName: resource.Name,
				Name:         action.Name,
				Usage:        sanitizeUsage(action.Usage),
				Command:      action,
				Fields:       extractFields(action.Flags, false, action),
			})
		}

		if len(res.Actions) > 0 {
			catalog.Resources = append(catalog.Resources, res)
		}
	}

	return catalog
}

func extractFields(flags []cli.Flag, global bool, cmd *cli.Command) []FieldSpec {
	fields := make([]FieldSpec, 0, len(flags))
	for _, flag := range flags {
		if flag == nil {
			continue
		}
		if visible, ok := flag.(cli.VisibleFlag); ok && !visible.IsVisible() {
			continue
		}
		names := flag.Names()
		if len(names) == 0 {
			continue
		}
		if hiddenFromInteractive(names[0]) {
			continue
		}

		field := FieldSpec{
			Name:        names[0],
			Aliases:     names[1:],
			Global:      global,
			Usage:       usageForFlag(flag, names[0]),
			DefaultText: defaultTextForFlag(flag),
			TypeName:    typeNameForFlag(flag),
			MultiValue:  isMultiValue(flag),
			Required:    isRequired(flag),
			Kind:        classifyField(flag),
			RequestHint: requestHint(flag),
		}
		if cmd != nil && cmd.IsSet(field.Name) {
			field.CurrentValue = stringifyValue(flag.Get(), field.Kind)
		}
		fields = append(fields, field)
	}
	return fields
}

func usageForFlag(flag cli.Flag, name string) string {
	if doc, ok := flag.(cli.DocGenerationFlag); ok {
		if usage := strings.TrimSpace(doc.GetUsage()); usage != "" {
			return sanitizeUsage(usage)
		}
	}

	switch name {
	case "api-key":
		return "API key used to authenticate requests."
	default:
		return ""
	}
}

func sanitizeUsage(value string) string {
	value = htmlTagPattern.ReplaceAllString(value, "")
	value = markdownLinkPattern.ReplaceAllString(value, "$1")
	value = strings.TrimSpace(value)
	return value
}

func hiddenFromInteractive(name string) bool {
	switch name {
	case "help", "version", "generate-shell-completion":
		return true
	default:
		return false
	}
}

func defaultTextForFlag(flag cli.Flag) string {
	if doc, ok := flag.(cli.DocGenerationFlag); ok {
		return doc.GetDefaultText()
	}
	return ""
}

func typeNameForFlag(flag cli.Flag) string {
	if doc, ok := flag.(cli.DocGenerationFlag); ok {
		if typeName := strings.TrimSpace(doc.TypeName()); typeName != "" {
			return typeName
		}
	}

	t := reflect.TypeOf(flag)
	if t == nil {
		return ""
	}
	return t.String()
}

func isRequired(flag cli.Flag) bool {
	if required, ok := flag.(requiredFlagOrStdin); ok {
		return required.IsRequiredAsFlagOrStdin()
	}
	if required, ok := flag.(cli.RequiredFlag); ok {
		return required.IsRequired()
	}
	return false
}

func isMultiValue(flag cli.Flag) bool {
	if multi, ok := flag.(cli.DocGenerationMultiValueFlag); ok {
		return multi.IsMultiValueFlag()
	}
	return false
}

func classifyField(flag cli.Flag) FieldKind {
	typeName := strings.ToLower(typeNameForFlag(flag))

	switch {
	case strings.Contains(typeName, "boolean") || strings.Contains(typeName, "bool"):
		return FieldKindBool
	case strings.Contains(typeName, "float"), strings.Contains(typeName, "int"), strings.Contains(typeName, "number"):
		if isMultiValue(flag) {
			return FieldKindList
		}
		return FieldKindNumber
	case strings.Contains(typeName, "map["), strings.Contains(typeName, "="), strings.Contains(typeName, "any"):
		return FieldKindMultiline
	case isMultiValue(flag):
		return FieldKindList
	default:
		return FieldKindText
	}
}

func requestHint(flag cli.Flag) string {
	requestFlag, ok := flag.(requestFlagDetails)
	if !ok {
		return ""
	}

	switch {
	case requestFlag.IsBodyRoot():
		return "body"
	case requestFlag.GetBodyPath() != "":
		return "body"
	case requestFlag.GetQueryPath() != "":
		return "query"
	case requestFlag.GetHeaderPath() != "":
		return "header"
	default:
		return ""
	}
}

func CollectFields(catalog Catalog, state SessionState) []FieldSpec {
	if len(catalog.Resources) == 0 {
		return nil
	}
	if state.ResourceIdx < 0 || state.ResourceIdx >= len(catalog.Resources) {
		return nil
	}
	resource := catalog.Resources[state.ResourceIdx]
	if state.ActionIdx < 0 || state.ActionIdx >= len(resource.Actions) {
		return nil
	}
	action := resource.Actions[state.ActionIdx]

	fields := make([]FieldSpec, 0, len(catalog.GlobalFields)+len(action.Fields))
	fields = append(fields, catalog.GlobalFields...)
	fields = append(fields, action.Fields...)

	for i := range fields {
		if value, ok := state.Values[fields[i].Name]; ok {
			fields[i].CurrentValue = value
		}
	}

	slices.SortStableFunc(fields, func(a, b FieldSpec) int {
		switch {
		case a.Required && !b.Required:
			return -1
		case !a.Required && b.Required:
			return 1
		case !a.Global && b.Global:
			return -1
		case a.Global && !b.Global:
			return 1
		default:
			return strings.Compare(a.Name, b.Name)
		}
	})

	return fields
}

func MissingRequiredFields(fields []FieldSpec, values map[string]string) []FieldSpec {
	missing := make([]FieldSpec, 0)
	for _, field := range fields {
		if !field.Required {
			continue
		}
		if strings.TrimSpace(values[field.Name]) == "" {
			missing = append(missing, field)
		}
	}
	return missing
}

func stringifyValue(value any, kind FieldKind) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return v
	case bool:
		if kind == FieldKindBool {
			if v {
				return "true"
			}
			return ""
		}
		if v {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprint(v)
	}
}
