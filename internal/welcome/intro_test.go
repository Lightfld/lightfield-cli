package welcome

import (
	"strings"
	"testing"
)

func TestRenderCardIncludesQuickstartAndNextSteps(t *testing.T) {
	card := renderCard("0.4.1", 100, 50, 0)

	for _, want := range []string{
		"LIGHTFIELD CLI",
		"Lightfield",
		"v0.4.1",
		"Start here",
		`export LIGHTFIELD_API_KEY="sk_lf_..."`,
		`lightfield account list --api-key "$LIGHTFIELD_API_KEY" --limit 1`,
		"lightfield --help",
		DocsHomeURL,
		QuickstartURL,
		CLIReferenceURL,
	} {
		if !strings.Contains(card, want) {
			t.Fatalf("expected card to contain %q", want)
		}
	}
}

func TestCardWidthRespectsTerminalBounds(t *testing.T) {
	if got := cardWidth(0); got != 92 {
		t.Fatalf("expected default width of 92, got %d", got)
	}
	if got := cardWidth(70); got != 72 {
		t.Fatalf("expected minimum width of 72, got %d", got)
	}
	if got := cardWidth(80); got != 74 {
		t.Fatalf("expected width constrained by terminal, got %d", got)
	}
}
