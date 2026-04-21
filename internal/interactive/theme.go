package interactive

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	App           lipgloss.Style
	Muted         lipgloss.Style
	Accent        lipgloss.Style
	AccentSoft    lipgloss.Style
	Primary       lipgloss.Style
	Secondary     lipgloss.Style
	Border        lipgloss.Style
	Selected      lipgloss.Style
	SelectedMuted lipgloss.Style
	BadgeRequired lipgloss.Style
	BadgeOptional lipgloss.Style
	Error         lipgloss.Style
	Success       lipgloss.Style
	Command       lipgloss.Style
	BannerText    lipgloss.Style
	BannerRays    lipgloss.Style
	BannerGlow    lipgloss.Style
	BannerSubtle  lipgloss.Style
	Footer        lipgloss.Style
	Panel         lipgloss.Style
	Title         lipgloss.Style
	Subtitle      lipgloss.Style
}

func NewTheme() Theme {
	accent := lipgloss.AdaptiveColor{Light: "#356AE6", Dark: "#7EA7FF"}
	accentSoft := lipgloss.AdaptiveColor{Light: "#DCE7FF", Dark: "#22345F"}
	accentGlow := lipgloss.AdaptiveColor{Light: "#7EA7FF", Dark: "#C9D9FF"}
	fg := lipgloss.AdaptiveColor{Light: "#1F2430", Dark: "#EEF2FF"}
	muted := lipgloss.AdaptiveColor{Light: "#637085", Dark: "#9EABC2"}
	border := lipgloss.AdaptiveColor{Light: "#D6DDEB", Dark: "#33415C"}
	panel := lipgloss.AdaptiveColor{Light: "#F6F8FC", Dark: "#161D2D"}
	errorColor := lipgloss.AdaptiveColor{Light: "#B42318", Dark: "#FF8A80"}
	successColor := lipgloss.AdaptiveColor{Light: "#067647", Dark: "#75E0A7"}

	return Theme{
		App:        lipgloss.NewStyle().Foreground(fg),
		Muted:      lipgloss.NewStyle().Foreground(muted),
		Accent:     lipgloss.NewStyle().Foreground(accent).Bold(true),
		AccentSoft: lipgloss.NewStyle().Foreground(accent).Background(accentSoft),
		Primary:    lipgloss.NewStyle().Foreground(fg).Bold(true),
		Secondary:  lipgloss.NewStyle().Foreground(muted),
		Border: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(border),
		Selected: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true),
		SelectedMuted: lipgloss.NewStyle().
			Foreground(accent),
		BadgeRequired: lipgloss.NewStyle().
			Foreground(accent).
			Background(accentSoft).
			Padding(0, 1),
		BadgeOptional: lipgloss.NewStyle().
			Foreground(muted).
			Background(panel).
			Padding(0, 1),
		Error: lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true),
		Success: lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true),
		Command: lipgloss.NewStyle().
			Foreground(fg).
			Background(panel).
			Padding(1, 2),
		BannerText: lipgloss.NewStyle().
			Foreground(fg).
			Bold(true),
		BannerRays: lipgloss.NewStyle().
			Foreground(accent).
			Bold(true),
		BannerGlow: lipgloss.NewStyle().
			Foreground(accentGlow).
			Bold(true),
		BannerSubtle: lipgloss.NewStyle().
			Foreground(muted),
		Footer: lipgloss.NewStyle().
			Foreground(muted),
		Panel: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(border).
			Padding(1, 2),
		Title: lipgloss.NewStyle().
			Foreground(fg).
			Bold(true),
		Subtitle: lipgloss.NewStyle().
			Foreground(muted),
	}
}
