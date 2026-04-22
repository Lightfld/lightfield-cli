package welcome

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func BenchmarkRenderCard(b *testing.B) {
	s := newCardStyles()
	heroLines := []heroContentLine{
		{text: "Lightfield CLI v0.4.1", style: s.eyebrow},
		{text: "The agent-native CRM platform.", style: s.hero},
	}
	m := Model{
		width:   100,
		height:  50,
		version: "0.4.1",
		styles:  s,
	}
	m.heroLines = heroLines
	m.recalcLayout()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		renderCard(m.innerWidth, m.heroRows, i%frameWrap, s, heroLines, m.staticCard, m.heroFrameStyle, m.boxStyle, m.panelStyle)
	}
}

func BenchmarkRenderRayCell(b *testing.B) {
	s := newCardStyles()
	fp := newFrameParams(80, 14, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		renderRayCell(80, 14, i%80, i%14, fp, s.rayChars)
	}
}

func TestRenderCardIncludesExpectedContent(t *testing.T) {
	s := newCardStyles()
	heroLines := []heroContentLine{
		{text: "Lightfield CLI v0.4.1", style: s.eyebrow},
		{text: "The agent-native CRM platform.", style: s.hero},
	}
	m := Model{
		width:     100,
		height:    50,
		version:   "0.4.1",
		styles:    s,
		heroLines: heroLines,
	}
	m.recalcLayout()

	card := renderCard(m.innerWidth, m.heroRows, 0, s, heroLines, m.staticCard, m.heroFrameStyle, m.boxStyle, m.panelStyle)

	for _, tc := range []struct {
		name string
		want string
	}{
		{"eyebrow with version", "Lightfield CLI v0.4.1"},
		{"tagline", "agent-native CRM platform"},
		{"section header", "Start here"},
		{"step 1 command", `export LIGHTFIELD_API_KEY="sk_lf_..."`},
		{"step 2 command", `lightfield account list --api-key "$LIGHTFIELD_API_KEY" --limit 1`},
		{"step 3 command", "lightfield --help"},
		{"docs home link", DocsHomeURL},
		{"quickstart link", QuickstartURL},
		{"cli reference link", CLIReferenceURL},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if !strings.Contains(card, tc.want) {
				t.Errorf("card missing %q", tc.want)
			}
		})
	}
}

func TestCardWidth(t *testing.T) {
	for _, tc := range []struct {
		name      string
		terminal  int
		wantWidth int
	}{
		{"zero defaults to 92", 0, 92},
		{"narrow clamps to terminal", 70, 70},
		{"very narrow clamps to terminal width", 40, 40},
		{"medium constrained by terminal", 80, 74},
		{"wide caps at default 92", 200, 92},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := cardWidth(tc.terminal); got != tc.wantWidth {
				t.Errorf("cardWidth(%d) = %d, want %d", tc.terminal, got, tc.wantWidth)
			}
		})
	}
}

func TestSanitizeOSC8(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"clean URL passes through", "https://example.com/path", "https://example.com/path"},
		{"strips ESC", "https://evil\x1bsite", "https://evilsite"},
		{"strips BEL", "https://evil\x07site", "https://evilsite"},
		{"strips newlines", "https://evil\n\rsite", "https://evilsite"},
		{"strips semicolons", "id;inject;extra", "idinjectextra"},
		{"empty string", "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := sanitizeOSC8(tc.input); got != tc.want {
				t.Errorf("sanitizeOSC8(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestModelUpdate_KeyQuit(t *testing.T) {
	m := NewModel("1.0.0")
	m.width = 100
	m.height = 50
	m.recalcLayout()

	keys := []struct {
		name string
		msg  tea.KeyMsg
	}{
		{"enter", tea.KeyMsg(tea.Key{Type: tea.KeyEnter})},
		{"esc", tea.KeyMsg(tea.Key{Type: tea.KeyEsc})},
		{"q", tea.KeyMsg(tea.Key{Type: tea.KeyRunes, Runes: []rune{'q'}})},
		{"ctrl+c", tea.KeyMsg(tea.Key{Type: tea.KeyCtrlC})},
	}

	for _, tc := range keys {
		t.Run(tc.name, func(t *testing.T) {
			_, cmd := m.Update(tc.msg)
			if cmd == nil {
				t.Fatalf("expected a command for key %q", tc.name)
			}
			msg := cmd()
			if _, ok := msg.(tea.QuitMsg); !ok {
				t.Errorf("key %q: expected tea.QuitMsg, got %T", tc.name, msg)
			}
		})
	}
}

func TestModelUpdate_Resize(t *testing.T) {
	m := NewModel("1.0.0")

	updated, cmd := m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	model := updated.(Model)

	if model.width != 120 || model.height != 40 {
		t.Errorf("resize: got %dx%d, want 120x40", model.width, model.height)
	}
	if model.innerWidth == 0 {
		t.Error("resize: innerWidth was not recomputed")
	}
	if model.staticCard == "" {
		t.Error("resize: staticCard was not recomputed")
	}
	if cmd == nil {
		t.Error("resize: expected ClearScreen command")
	}
}

func TestModelUpdate_Pulse(t *testing.T) {
	m := NewModel("1.0.0")
	m.width = 100
	m.height = 50
	m.frame = 0

	updated, cmd := m.Update(pulseMsg{})
	model := updated.(Model)

	if model.frame != 1 {
		t.Errorf("pulse: frame = %d, want 1", model.frame)
	}
	if cmd == nil {
		t.Error("pulse: expected tickPulse command")
	}
}

func TestModelUpdate_PulseWraps(t *testing.T) {
	m := NewModel("1.0.0")
	m.frame = frameWrap - 1

	updated, _ := m.Update(pulseMsg{})
	model := updated.(Model)

	if model.frame != 0 {
		t.Errorf("pulse wrap: frame = %d, want 0", model.frame)
	}
}

func TestModelUpdate_UnhandledKey(t *testing.T) {
	m := NewModel("1.0.0")
	m.width = 100
	m.height = 50
	frameBefore := m.frame

	updated, cmd := m.Update(tea.KeyMsg(tea.Key{Type: tea.KeyRunes, Runes: []rune{'x'}}))
	model := updated.(Model)

	if model.frame != frameBefore {
		t.Error("unhandled key should not change frame")
	}
	if cmd != nil {
		t.Error("unhandled key should return nil command")
	}
}

func TestHeroDecorationRows(t *testing.T) {
	for _, tc := range []struct {
		name     string
		height   int
		wantRows int
	}{
		{"zero height defaults to 6", 0, 6},
		{"short terminal clamps to 1", 28, 1},
		{"medium terminal", 40, 6},
		{"tall terminal caps at 6", 80, 6},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := heroDecorationRows(tc.height, 25, 2); got != tc.wantRows {
				t.Errorf("heroDecorationRows(%d, 25, 2) = %d, want %d", tc.height, got, tc.wantRows)
			}
		})
	}
}
