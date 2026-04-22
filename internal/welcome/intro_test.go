package welcome

import (
	"strings"
	"testing"
)

func BenchmarkRenderCard(b *testing.B) {
	s := newCardStyles()
	heroLines := []heroContentLine{
		{text: "LIGHTFIELD CLI", style: s.eyebrow},
		{text: "Lightfield v0.4.1", style: s.hero},
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
		{text: "LIGHTFIELD CLI", style: s.eyebrow},
		{text: "Lightfield v0.4.1", style: s.hero},
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
		{"eyebrow", "LIGHTFIELD CLI"},
		{"product name", "Lightfield"},
		{"version", "v0.4.1"},
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
		{"narrow clamps to min 72", 70, 72},
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
		{"strips ESC", "https://evil\x1b]8;;hack\x1b\\", "https://evil]8;;hack\\"},
		{"strips BEL", "https://evil\x07site", "https://evilsite"},
		{"strips newlines", "https://evil\n\rsite", "https://evilsite"},
		{"empty string", "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := sanitizeOSC8(tc.input); got != tc.want {
				t.Errorf("sanitizeOSC8(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
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
