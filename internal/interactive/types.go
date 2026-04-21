package interactive

import "github.com/urfave/cli/v3"

type FieldKind string

const (
	FieldKindBool      FieldKind = "bool"
	FieldKindText      FieldKind = "text"
	FieldKindNumber    FieldKind = "number"
	FieldKindList      FieldKind = "list"
	FieldKindMultiline FieldKind = "multiline"
)

type FieldSpec struct {
	Name         string
	Aliases      []string
	Usage        string
	DefaultText  string
	TypeName     string
	Kind         FieldKind
	Required     bool
	MultiValue   bool
	Global       bool
	RequestHint  string
	CurrentValue string
}

type ActionSpec struct {
	ResourceName string
	Name         string
	Usage        string
	Command      *cli.Command
	Fields       []FieldSpec
}

type ResourceSpec struct {
	Name     string
	Category string
	Usage    string
	Actions  []ActionSpec
}

type Catalog struct {
	GlobalFields []FieldSpec
	Resources    []ResourceSpec
}

type SessionState struct {
	ResourceIdx int
	ActionIdx   int
	Values      map[string]string
}

func (s SessionState) Clone() SessionState {
	out := SessionState{
		ResourceIdx: s.ResourceIdx,
		ActionIdx:   s.ActionIdx,
		Values:      make(map[string]string, len(s.Values)),
	}
	for k, v := range s.Values {
		out.Values[k] = v
	}
	return out
}
