package welcome

import (
	"fmt"
	"math"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	DocsHomeURL     = "https://docs.lightfield.app/"
	QuickstartURL   = "https://docs.lightfield.app/getting-started/cli-quickstart/"
	CLIReferenceURL = "https://docs.lightfield.app/api/cli"

	tickInterval    = 180 * time.Millisecond
	maxDecorRows    = 6
	frameWrap       = 10000
	phaseRate       = 0.42
	waveSpatialFreq = 0.45
	perRayPhaseOff  = 0.22
	fadeStart       = 0.18
	fadeRange       = 0.15
)

// rayBrightness indexes into the pre-rendered ray character table.
type rayBrightness int

const (
	rayFaint  rayBrightness = 0
	rayMid    rayBrightness = 1
	rayAccent rayBrightness = 2
)

// rayChar indexes the character shape within each brightness level.
type rayChar int

const (
	rayVert rayChar = 0
	rayBack rayChar = 1
	rayFwd  rayChar = 2
)

type cardStyles struct {
	eyebrow       lipgloss.Style
	hero          lipgloss.Style
	body          lipgloss.Style
	muted         lipgloss.Style
	link          lipgloss.Style
	stepNumber    lipgloss.Style
	stepLabel     lipgloss.Style
	commandPrompt lipgloss.Style
	section       lipgloss.Style
	command       lipgloss.Style
	heroFrame     lipgloss.Style
	commandBlock  lipgloss.Style
	border        lipgloss.Color

	// Pre-rendered ray characters: [brightness][charShape] → ANSI-styled string.
	// Computed once at init, eliminating ~1000 Render() calls per frame.
	rayChars [3][3]string
}

func newCardStyles() cardStyles {
	primary := lipgloss.Color("#DEDEDE")
	secondary := lipgloss.Color("#ABABAB")
	tertiary := lipgloss.Color("#6E6E6E")

	decorFaint := lipgloss.NewStyle().Foreground(lipgloss.Color("#2A2A2A"))
	decorMid := lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
	decorAccent := lipgloss.NewStyle().Foreground(lipgloss.Color("#ABABAB"))

	chars := [3]string{"│", "╲", "╱"}
	decorStyles := [3]lipgloss.Style{decorFaint, decorMid, decorAccent}
	var rayChars [3][3]string
	for b, s := range decorStyles {
		for c, ch := range chars {
			rayChars[b][c] = s.Render(ch)
		}
	}

	return cardStyles{
		border:        lipgloss.Color("#343434"),
		eyebrow:       lipgloss.NewStyle().Foreground(tertiary).Bold(true),
		hero:          lipgloss.NewStyle().Bold(true).Foreground(primary),
		body:          lipgloss.NewStyle().Foreground(secondary),
		muted:         lipgloss.NewStyle().Foreground(tertiary),
		link:          lipgloss.NewStyle().Foreground(secondary),
		stepNumber:    lipgloss.NewStyle().Foreground(secondary).Bold(true),
		stepLabel:     lipgloss.NewStyle().Foreground(primary),
		commandPrompt: lipgloss.NewStyle().Foreground(tertiary),
		section:       lipgloss.NewStyle().Foreground(primary).Bold(true),
		command:       lipgloss.NewStyle().Foreground(secondary),
		heroFrame:     lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")),
		commandBlock:  lipgloss.NewStyle().PaddingLeft(2),
		rayChars:      rayChars,
	}
}

// frameParams holds values that are constant for every cell within a single
// animation frame. Computing them once eliminates ~1000 redundant Sqrt and
// float64 conversions per frame.
type frameParams struct {
	cx, cy  float64
	maxDist float64
	phase   float64
}

func newFrameParams(width, totalRows, frame int) frameParams {
	cx := float64(width-1) / 2.0
	cy := float64(totalRows-1) / 2.0
	return frameParams{
		cx:      cx,
		cy:      cy,
		maxDist: math.Sqrt(cx*cx + cy*cy),
		phase:   float64(frame) * phaseRate,
	}
}

type Model struct {
	width   int
	height  int
	frame   int
	version string
	styles  cardStyles

	// Computed once at init (version + styles are immutable).
	heroLines []heroContentLine

	// Recomputed on resize only — these don't change per frame.
	innerWidth     int
	heroRows       int
	heroFrameStyle lipgloss.Style // heroFrame + Width + Align, composed once
	boxStyle       lipgloss.Style
	panelStyle     lipgloss.Style
	staticCard     string // pre-rendered steps + links + footer
}

func NewModel(version string) Model {
	s := newCardStyles()

	heroText := "Lightfield"
	if version != "" {
		heroText += " v" + version
	}

	return Model{
		version: version,
		styles:  s,
		heroLines: []heroContentLine{
			{text: "LIGHTFIELD CLI", style: s.eyebrow},
			{text: heroText, style: s.hero},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return tickPulse()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.recalcLayout()
		return m, tea.ClearScreen
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "esc", "q", "ctrl+c":
			return m, tea.Quit
		}
	case pulseMsg:
		m.frame = (m.frame + 1) % frameWrap
		return m, tickPulse()
	}

	return m, nil
}

// recalcLayout recomputes all width/height-dependent values that don't change
// per animation frame. Called once on each resize.
func (m *Model) recalcLayout() {
	s := m.styles
	outerWidth := cardWidth(m.width)
	boxContentWidth := max(outerWidth-8, 40)

	m.innerWidth = max(boxContentWidth-4, 28)
	m.heroRows = heroDecorationRows(m.height, 25, len(m.heroLines))
	m.heroFrameStyle = s.heroFrame.Width(m.innerWidth).Align(lipgloss.Center)

	m.boxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(s.border).
		Width(boxContentWidth).
		Padding(1, 2)
	m.panelStyle = lipgloss.NewStyle().
		Padding(1, 1).
		MaxWidth(outerWidth)

	commandWidth := max(m.innerWidth-4, 20)

	startHere := strings.Join([]string{
		s.section.Render("Start here"),
		"",
		step(s.stepNumber, s.stepLabel, s.commandPrompt, s.command, s.commandBlock, commandWidth, "1", `Set your API key`, `export LIGHTFIELD_API_KEY="sk_lf_..."`),
		"",
		step(s.stepNumber, s.stepLabel, s.commandPrompt, s.command, s.commandBlock, commandWidth, "2", `Make your first request`, `lightfield account list --api-key "$LIGHTFIELD_API_KEY" --limit 1`),
		"",
		step(s.stepNumber, s.stepLabel, s.commandPrompt, s.command, s.commandBlock, commandWidth, "3", `Explore available commands`, `lightfield --help`),
	}, "\n")

	learnMore := strings.Join([]string{
		s.section.Render("Learn more"),
		linkRow(s.body, s.link, "docs", "Docs", DocsHomeURL),
		linkRow(s.body, s.link, "quickstart", "Quickstart", QuickstartURL),
		linkRow(s.body, s.link, "cli-ref", "CLI API reference", CLIReferenceURL),
	}, "\n")

	var b strings.Builder
	b.WriteString(startHere)
	b.WriteString("\n\n")
	b.WriteString(learnMore)
	b.WriteString("\n\n")
	b.WriteString(s.muted.Render("Press Enter to continue."))
	m.staticCard = b.String()
}

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	card := renderCard(m.innerWidth, m.heroRows, m.frame, m.styles, m.heroLines, m.staticCard, m.heroFrameStyle, m.boxStyle, m.panelStyle)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, card)
}

func renderCard(innerWidth, heroRows, frame int, s cardStyles, heroLines []heroContentLine, staticBelow string, heroFrameStyle, boxStyle, panelStyle lipgloss.Style) string {
	heroBlock := heroFrameStyle.Render(
		renderHeroFrame(innerWidth, heroRows, frame, s.rayChars, heroLines),
	)

	var b strings.Builder
	b.Grow(len(heroBlock) + len(staticBelow) + 4)
	b.WriteString(heroBlock)
	b.WriteString("\n\n")
	b.WriteString(staticBelow)

	return panelStyle.Render(boxStyle.Render(b.String()))
}

func cardWidth(width int) int {
	const defaultWidth = 92
	const minWidth = 72

	if width <= 0 {
		return defaultWidth
	}

	usable := width - 6
	if usable < minWidth {
		return minWidth
	}
	if usable < defaultWidth {
		return usable
	}
	return defaultWidth
}

func step(numberStyle, labelStyle, promptStyle, commandStyle, blockStyle lipgloss.Style, commandWidth int, number, label, command string) string {
	commandLine := lipgloss.JoinHorizontal(
		lipgloss.Top,
		promptStyle.Render("$ "),
		commandStyle.Width(commandWidth).Render(command),
	)

	return fmt.Sprintf("%s %s\n%s",
		numberStyle.Render(number+"."),
		labelStyle.Render(label),
		blockStyle.Render(commandLine),
	)
}

// linkRow emits an OSC 8 terminal hyperlink. The id parameter gives the
// terminal a stable identity so hover state survives redraws. URLs and IDs
// are sanitised to prevent escape-sequence injection (control chars, newlines,
// the ST terminator \033\\). This is defense-in-depth — the URLs are currently
// compile-time constants — but protects against future changes.
//
// Safety: because the TUI runs in the alt-screen buffer (tea.WithAltScreen),
// any partial writes from a mid-render signal cannot leak into the user's
// main scrollback.
func linkRow(labelStyle, linkStyle lipgloss.Style, id, label, url string) string {
	id = sanitizeOSC8(id)
	url = sanitizeOSC8(url)
	styledURL := linkStyle.Render(url)
	link := fmt.Sprintf("\033]8;id=%s;%s\033\\%s\033]8;;\033\\", id, url, styledURL)
	return fmt.Sprintf("%s %s",
		labelStyle.Render(label+":"),
		link,
	)
}

// sanitizeOSC8 strips characters that could break an OSC 8 escape sequence:
// ESC (0x1B), BEL (0x07), newlines, and the semicolon field separator.
func sanitizeOSC8(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '\x1b', '\x07', '\n', '\r':
			return -1
		default:
			return r
		}
	}, s)
}

type pulseMsg time.Time

func tickPulse() tea.Cmd {
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return pulseMsg(t)
	})
}

type heroContentLine struct {
	text  string
	style lipgloss.Style
}

// renderHeroFrame writes the full hero animation block into a single Builder,
// avoiding per-row intermediate string allocations.
func renderHeroFrame(width, decorRows, frame int, rc [3][3]string, lines []heroContentLine) string {
	if width <= 0 {
		return ""
	}

	totalRows := decorRows*2 + len(lines)
	fp := newFrameParams(width, totalRows, frame)

	var buf strings.Builder
	buf.Grow(totalRows * width * 8)

	for r := 0; r < decorRows; r++ {
		if r > 0 {
			buf.WriteByte('\n')
		}
		writeRayRow(&buf, width, totalRows, r, fp, rc)
	}
	for i, line := range lines {
		if decorRows > 0 || i > 0 {
			buf.WriteByte('\n')
		}
		writeHeroLine(&buf, width, totalRows, decorRows+i, fp, rc, line)
	}
	for i := 0; i < decorRows; i++ {
		buf.WriteByte('\n')
		writeRayRow(&buf, width, totalRows, decorRows+len(lines)+i, fp, rc)
	}

	return buf.String()
}

// heroDecorationRows calculates how many ray rows fit above/below the hero text
// given the terminal height and the fixed content that sits below the hero.
func heroDecorationRows(termHeight, fixedRows, textLines int) int {
	if termHeight <= 0 {
		return maxDecorRows
	}
	available := termHeight - fixedRows - textLines
	perSide := available / 2
	if perSide < 1 {
		return 1
	}
	if perSide > maxDecorRows {
		return maxDecorRows
	}
	return perSide
}

func writeHeroLine(buf *strings.Builder, width, totalRows, row int, fp frameParams, rc [3][3]string, line heroContentLine) {
	const gap = 3

	textWidth := lipgloss.Width(line.text)
	sideTotal := max(width-textWidth-gap*2, 0)
	leftWidth := sideTotal / 2
	rightWidth := sideTotal - leftWidth
	rightStart := leftWidth + textWidth + gap*2

	for c := 0; c < leftWidth; c++ {
		buf.WriteString(renderRayCell(width, totalRows, c, row, fp, rc))
	}
	for i := 0; i < gap; i++ {
		buf.WriteByte(' ')
	}
	buf.WriteString(line.style.Render(line.text))
	for i := 0; i < gap; i++ {
		buf.WriteByte(' ')
	}
	for c := 0; c < rightWidth; c++ {
		buf.WriteString(renderRayCell(width, totalRows, rightStart+c, row, fp, rc))
	}
}

func writeRayRow(buf *strings.Builder, width, totalRows, row int, fp frameParams, rc [3][3]string) {
	for x := 0; x < width; x++ {
		buf.WriteString(renderRayCell(width, totalRows, x, row, fp, rc))
	}
}

// raySlopes defines the set of light rays radiating from the hero text.
// Wider-angle slopes ensure rays reach the box edges even at moderate row counts.
var raySlopes = []float64{0.0, 0.4, 0.9, 1.6, 2.8, 4.5, 7.5}

func renderRayCell(width, totalRows, x, row int, fp frameParams, rc [3][3]string) string {
	dc := float64(x) - fp.cx
	dr := float64(row) - fp.cy
	absDr := math.Abs(dr)

	if absDr < 0.45 {
		return " "
	}

	tolerance := 0.55 + absDr*0.06

	bestIdx := -1
	bestResidual := tolerance
	absDc := math.Abs(dc)
	for i, slope := range raySlopes {
		residual := math.Abs(absDc - slope*absDr)
		if residual < bestResidual {
			bestIdx = i
			bestResidual = residual
		}
	}

	if bestIdx == -1 {
		return " "
	}

	distance := math.Sqrt(dc*dc + dr*dr)

	proximity := distance / math.Max(fp.maxDist, 1)
	if proximity < fadeStart {
		return " "
	}
	fadeIn := math.Min((proximity-fadeStart)/fadeRange, 1.0)

	wave := 0.5 + 0.5*math.Sin(fp.phase+distance*waveSpatialFreq+float64(bestIdx)*perRayPhaseOff)
	wave *= fadeIn

	var ci rayChar
	switch {
	case raySlopes[bestIdx] < 0.1:
		ci = rayVert
	case dr*dc >= 0:
		ci = rayBack
	default:
		ci = rayFwd
	}

	switch {
	case wave > 0.82:
		return rc[rayAccent][ci]
	case wave > 0.55:
		return rc[rayMid][ci]
	case wave > 0.22:
		return rc[rayFaint][ci]
	default:
		return " "
	}
}
