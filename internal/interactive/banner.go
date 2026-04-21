package interactive

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	bannerHeight = 9
	titleRow     = 3
	subtitleRow  = 4
)

type bannerPoint struct {
	xRatio float64
	y      int
	phase  int
}

var bannerPoints = []bannerPoint{
	{0.06, 0, 0}, {0.14, 1, 2}, {0.22, 0, 4}, {0.30, 1, 1}, {0.38, 0, 3},
	{0.62, 0, 2}, {0.70, 1, 4}, {0.78, 0, 1}, {0.86, 1, 3}, {0.94, 0, 0},
	{0.10, 2, 3}, {0.18, 2, 1}, {0.28, 2, 4}, {0.72, 2, 0}, {0.82, 2, 2}, {0.90, 2, 4},
	{0.04, 5, 1}, {0.12, 6, 3}, {0.20, 7, 0}, {0.28, 6, 2}, {0.36, 7, 4},
	{0.64, 7, 1}, {0.72, 6, 3}, {0.80, 7, 0}, {0.88, 6, 2}, {0.96, 5, 4},
	{0.08, 8, 2}, {0.18, 8, 4}, {0.28, 8, 1}, {0.72, 8, 3}, {0.82, 8, 0}, {0.92, 8, 2},
}

type Banner struct {
	frame         int
	animating     bool
	reducedMotion bool
}

func NewBanner(reducedMotion bool) Banner {
	if reducedMotion {
		return Banner{frame: bannerCycleLength() - 1, reducedMotion: true}
	}
	return Banner{animating: true}
}

func (b *Banner) Advance() bool {
	if !b.animating {
		return false
	}
	b.frame++
	if b.frame >= bannerCycleLength() {
		b.frame = 0
	}
	return true
}

func (b *Banner) Skip() {
	b.frame = bannerCycleLength() - 1
	b.animating = false
}

func (b Banner) Animating() bool {
	return b.animating
}

func (b Banner) View(theme Theme, width int) string {
	width = max(40, width)
	grid := make([][]cell, bannerHeight)
	for row := range grid {
		grid[row] = make([]cell, width)
		for col := range grid[row] {
			grid[row][col] = cell{ch: ' ', style: theme.BannerSubtle}
		}
	}

	for _, point := range bannerPoints {
		col := min(width-1, max(0, int(point.xRatio*float64(width-1))))
		row := min(bannerHeight-1, max(0, point.y))
		grid[row][col] = pulseCell(theme, point, b.frame)
	}

	markText(grid, titleRow, centeredStart(width, "L I G H T F I E L D"), "L I G H T F I E L D", theme.BannerText)
	markText(grid, subtitleRow, centeredStart(width, "customer memory"), "customer memory", theme.BannerSubtle)

	lines := make([]string, 0, len(grid))
	for _, row := range grid {
		var builder strings.Builder
		for _, cell := range row {
			builder.WriteString(cell.style.Render(string(cell.ch)))
		}
		lines = append(lines, builder.String())
	}
	return strings.Join(lines, "\n")
}

type cell struct {
	ch    rune
	style lipgloss.Style
}

func pulseCell(theme Theme, point bannerPoint, frame int) cell {
	pulse := (frame + point.phase) % bannerCycleLength()

	switch {
	case pulse == 0 || pulse == 1:
		return cell{ch: '*', style: theme.BannerGlow}
	case pulse == 2 || pulse == 3:
		return cell{ch: '·', style: theme.BannerRays}
	default:
		return cell{ch: '.', style: theme.BannerSubtle}
	}
}

func markText(grid [][]cell, row, start int, text string, style lipgloss.Style) {
	if row < 0 || row >= len(grid) {
		return
	}
	runes := []rune(text)
	for i, ch := range runes {
		col := start + i
		if col < 0 || col >= len(grid[row]) {
			continue
		}
		grid[row][col] = cell{ch: ch, style: style}
	}
}

func centeredStart(width int, text string) int {
	runes := []rune(text)
	return max(0, (width-len(runes))/2)
}

func bannerCycleLength() int {
	return 12
}
