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
)

type Model struct {
	width   int
	height  int
	frame   int
	version string
}

func NewModel(version string) Model {
	return Model{version: version}
}

func (m Model) Init() tea.Cmd {
	return tickPulse()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, tea.ClearScreen
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "esc", "q", "ctrl+c":
			return m, tea.Quit
		}
	case pulseMsg:
		m.frame++
		return m, tickPulse()
	}

	return m, nil
}

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	card := renderCard(m.version, m.width, m.height, m.frame)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, card)
}

func renderCard(version string, width int, height int, frame int) string {
	primary := lipgloss.Color("#DEDEDE")
	secondary := lipgloss.Color("#ABABAB")
	tertiary := lipgloss.Color("#6E6E6E")
	border := lipgloss.Color("#343434")

	eyebrowStyle := lipgloss.NewStyle().
		Foreground(tertiary).
		Bold(true)
	heroStyle := lipgloss.NewStyle().
		Bold(true)
	bodyStyle := lipgloss.NewStyle().
		Foreground(secondary)
	mutedStyle := lipgloss.NewStyle().Foreground(tertiary)
	decorationFaintStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#2A2A2A"))
	decorationMidStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
	decorationAccentStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ABABAB"))
	linkStyle := lipgloss.NewStyle().
		Foreground(secondary)
	stepNumberStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Bold(true)
	stepLabelStyle := lipgloss.NewStyle().
		Foreground(primary)
	commandPromptStyle := lipgloss.NewStyle().
		Foreground(tertiary)
	sectionStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true)
	commandStyle := lipgloss.NewStyle().
		Foreground(secondary)
	heroFrameStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#555555"))
	outerWidth := cardWidth(width)
	boxContentWidth := max(outerWidth-8, 40)
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(border).
		Width(boxContentWidth).
		Padding(1, 2)
	panelStyle := lipgloss.NewStyle().
		Padding(1, 1).
		MaxWidth(outerWidth)

	innerWidth := max(boxContentWidth-4, 28)
	commandWidth := max(innerWidth-4, 20)

	heroText := "Lightfield"
	if version != "" {
		heroText += " v" + version
	}

	// Fixed content below the hero uses ~25 lines (steps, links, footer, box
	// chrome, panel padding). Reserve that, then give the rest to the hero.
	const fixedRows = 25
	heroRows := heroDecorationRows(height, fixedRows, len([]heroContentLine{
		{text: "LIGHTFIELD CLI", style: eyebrowStyle},
		{text: heroText, style: heroStyle.Foreground(primary)},
	}))

	heroBlock := heroFrameStyle.Width(innerWidth).Align(lipgloss.Center).Render(
		renderHeroFrame(innerWidth, heroRows, frame, decorationFaintStyle, decorationMidStyle, decorationAccentStyle, []heroContentLine{
			{text: "LIGHTFIELD CLI", style: eyebrowStyle},
			{text: heroText, style: heroStyle.Foreground(primary)},
		}),
	)

	startHere := strings.Join([]string{
		sectionStyle.Render("Start here"),
		"",
		step(stepNumberStyle, stepLabelStyle, commandPromptStyle, commandStyle, commandWidth, "1", `Set your API key`, `export LIGHTFIELD_API_KEY="sk_lf_..."`),
		"",
		step(stepNumberStyle, stepLabelStyle, commandPromptStyle, commandStyle, commandWidth, "2", `Make your first request`, `lightfield account list --api-key "$LIGHTFIELD_API_KEY" --limit 1`),
		"",
		step(stepNumberStyle, stepLabelStyle, commandPromptStyle, commandStyle, commandWidth, "3", `Explore available commands`, `lightfield --help`),
	}, "\n")

	learnMore := strings.Join([]string{
		sectionStyle.Render("Learn more"),
		linkRow(bodyStyle, linkStyle, "docs", "Docs", DocsHomeURL),
		linkRow(bodyStyle, linkStyle, "quickstart", "Quickstart", QuickstartURL),
		linkRow(bodyStyle, linkStyle, "cli-ref", "CLI API reference", CLIReferenceURL),
	}, "\n")

	var b strings.Builder
	b.WriteString(heroBlock)
	b.WriteString("\n\n")
	b.WriteString(startHere)
	b.WriteString("\n\n")
	b.WriteString(learnMore)
	b.WriteString("\n\n")
	b.WriteString(mutedStyle.Render("Press Enter to continue."))

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

func step(numberStyle, labelStyle, promptStyle, commandStyle lipgloss.Style, commandWidth int, number, label, command string) string {
	commandLine := lipgloss.JoinHorizontal(
		lipgloss.Top,
		promptStyle.Render("$ "),
		commandStyle.Width(commandWidth).Render(command),
	)
	commandBlock := lipgloss.NewStyle().
		PaddingLeft(2).
		Render(commandLine)

	return fmt.Sprintf("%s %s\n%s",
		numberStyle.Render(number+"."),
		labelStyle.Render(label),
		commandBlock,
	)
}

func linkRow(labelStyle, linkStyle lipgloss.Style, id, label, url string) string {
	styledURL := linkStyle.Render(url)
	link := fmt.Sprintf("\033]8;id=%s;%s\033\\%s\033]8;;\033\\", id, url, styledURL)
	return fmt.Sprintf("%s %s",
		labelStyle.Render(label+":"),
		link,
	)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pulseMsg time.Time

func tickPulse() tea.Cmd {
	return tea.Tick(180*time.Millisecond, func(t time.Time) tea.Msg {
		return pulseMsg(t)
	})
}

type heroContentLine struct {
	text  string
	style lipgloss.Style
}

func renderHeroFrame(width, decorRows, frame int, faintStyle, midStyle, accentStyle lipgloss.Style, lines []heroContentLine) string {
	if width <= 0 {
		return ""
	}

	totalRows := decorRows*2 + len(lines)

	parts := make([]string, 0, totalRows)
	for r := 0; r < decorRows; r++ {
		parts = append(parts, renderRayRow(width, totalRows, r, frame, faintStyle, midStyle, accentStyle))
	}
	for i, line := range lines {
		parts = append(parts, renderHeroLine(width, totalRows, decorRows+i, frame, faintStyle, midStyle, accentStyle, line))
	}
	for i := 0; i < decorRows; i++ {
		parts = append(parts, renderRayRow(width, totalRows, decorRows+len(lines)+i, frame, faintStyle, midStyle, accentStyle))
	}

	return strings.Join(parts, "\n")
}

// heroDecorationRows calculates how many ray rows fit above/below the hero text
// given the terminal height and the fixed content that sits below the hero.
func heroDecorationRows(termHeight, fixedRows, textLines int) int {
	if termHeight <= 0 {
		return 6
	}
	available := termHeight - fixedRows - textLines
	perSide := available / 2
	if perSide < 1 {
		return 1
	}
	if perSide > 6 {
		return 6
	}
	return perSide
}

func renderHeroLine(width, totalRows, row, frame int, faintStyle, midStyle, accentStyle lipgloss.Style, line heroContentLine) string {
	const gap = 3

	textWidth := lipgloss.Width(line.text)
	sideTotal := max(width-textWidth-gap*2, 0)
	leftWidth := sideTotal / 2
	rightWidth := sideTotal - leftWidth
	rightStart := leftWidth + textWidth + gap*2

	var leftB, rightB strings.Builder
	for c := 0; c < leftWidth; c++ {
		leftB.WriteString(renderRayCell(width, totalRows, c, row, frame, faintStyle, midStyle, accentStyle))
	}
	for c := 0; c < rightWidth; c++ {
		rightB.WriteString(renderRayCell(width, totalRows, rightStart+c, row, frame, faintStyle, midStyle, accentStyle))
	}

	return leftB.String() + strings.Repeat(" ", gap) + line.style.Render(line.text) + strings.Repeat(" ", gap) + rightB.String()
}

func renderRayRow(width, totalRows, row, frame int, faintStyle, midStyle, accentStyle lipgloss.Style) string {
	var line strings.Builder
	for x := 0; x < width; x++ {
		line.WriteString(renderRayCell(width, totalRows, x, row, frame, faintStyle, midStyle, accentStyle))
	}
	return line.String()
}

// raySlopes defines the set of light rays radiating from the hero text.
// Wider-angle slopes ensure rays reach the box edges even at moderate row counts.
var raySlopes = []float64{0.0, 0.4, 0.9, 1.6, 2.8, 4.5, 7.5}

func renderRayCell(width, totalRows, x, row, frame int, faintStyle, midStyle, accentStyle lipgloss.Style) string {
	cx := float64(width-1) / 2.0
	cy := float64(totalRows-1) / 2.0
	dc := float64(x) - cx
	dr := float64(row) - cy
	absDr := math.Abs(dr)

	if absDr < 0.45 {
		return " "
	}

	// Tolerance scales slightly with distance so outer rays remain visible at the edges.
	tolerance := 0.55 + absDr*0.06

	bestIdx := -1
	bestResidual := tolerance
	for i, slope := range raySlopes {
		residual := math.Abs(math.Abs(dc) - slope*absDr)
		if residual < bestResidual {
			bestIdx = i
			bestResidual = residual
		}
	}

	if bestIdx == -1 {
		return " "
	}

	distance := math.Sqrt(dc*dc + dr*dr)

	// Fade out near the center so the text has breathing room.
	maxDist := math.Sqrt(cx*cx + cy*cy)
	proximity := distance / math.Max(maxDist, 1)
	if proximity < 0.18 {
		return " "
	}
	fadeIn := math.Min((proximity-0.18)/0.15, 1.0)

	phase := float64(frame) * 0.42
	wave := 0.5 + 0.5*math.Sin(phase+distance*0.45+float64(bestIdx)*0.22)
	wave *= fadeIn

	var char string
	switch {
	case raySlopes[bestIdx] < 0.1:
		char = "│"
	case dr*dc >= 0:
		char = "╲"
	default:
		char = "╱"
	}

	switch {
	case wave > 0.82:
		return accentStyle.Render(char)
	case wave > 0.55:
		return midStyle.Render(char)
	case wave > 0.22:
		return faintStyle.Render(char)
	default:
		return " "
	}
}
