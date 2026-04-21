package interactive

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/muesli/reflow/wordwrap"
	"github.com/urfave/cli/v3"
)

type Step int

const (
	stepWelcome Step = iota
	stepHelp
	stepResource
	stepAction
	stepFields
	stepEdit
	stepPreview
	stepResult
)

type welcomeChoice struct {
	Title       string
	Description string
}

type objectRow struct {
	Key   string
	Value string
}

type Model struct {
	ctx            context.Context
	root           *cli.Command
	catalog        Catalog
	theme          Theme
	banner         Banner
	step           Step
	previousStep   Step
	width          int
	height         int
	welcomeIndex   int
	resourceIndex  int
	actionIndex    int
	fieldIndex     int
	showOptional   bool
	state          SessionState
	filterInput    textinput.Model
	editorInput    textinput.Model
	editorArea     textarea.Model
	objectKeyInput textinput.Model
	objectValInput textinput.Model
	usingObjectUI  bool
	objectRows     []objectRow
	objectRowIndex int
	objectFocusVal bool
	editingField   *FieldSpec
	lastValidation string
	commandPreview string
	running        bool
	result         RunResult
	resultViewport viewport.Model
}

type bannerTickMsg struct{}

func Run(ctx context.Context, cmd *cli.Command) error {
	if !isInteractiveSession() {
		return cli.ShowAppHelp(cmd.Root())
	}

	model := NewModel(ctx, cmd.Root())
	program := tea.NewProgram(model, tea.WithAltScreen())
	_, err := program.Run()
	return err
}

func NewModel(ctx context.Context, root *cli.Command) Model {
	catalog := BuildCatalog(root)

	filterInput := textinput.New()
	filterInput.Prompt = "Filter: "
	filterInput.Placeholder = "Type to narrow choices"
	filterInput.Focus()
	filterInput.CharLimit = 256

	editorInput := textinput.New()
	editorInput.Prompt = "> "
	editorInput.CharLimit = 10000

	objectKeyInput := textinput.New()
	objectKeyInput.Prompt = "Field: "
	objectKeyInput.CharLimit = 1024

	objectValInput := textinput.New()
	objectValInput.Prompt = "Value: "
	objectValInput.CharLimit = 10000

	editorArea := textarea.New()
	editorArea.Placeholder = "{\n  \"$name\": \"Acme\"\n}"
	editorArea.Prompt = ""
	editorArea.CharLimit = 50000
	editorArea.SetHeight(10)

	state := SessionState{
		ResourceIdx: -1,
		ActionIdx:   -1,
		Values:      seedValues(root, catalog),
	}

	viewportModel := viewport.New(0, 0)
	viewportModel.SetContent("")

	return Model{
		ctx:            ctx,
		root:           root,
		catalog:        catalog,
		theme:          NewTheme(),
		banner:         NewBanner(reducedMotionRequested()),
		step:           stepWelcome,
		filterInput:    filterInput,
		editorInput:    editorInput,
		editorArea:     editorArea,
		objectKeyInput: objectKeyInput,
		objectValInput: objectValInput,
		objectRowIndex: -1,
		state:          state,
		resultViewport: viewportModel,
	}
}

func (m Model) Init() tea.Cmd {
	if m.banner.Animating() {
		return tea.Tick(bannerFrameDelay(), func(time.Time) tea.Msg {
			return bannerTickMsg{}
		})
	}
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.resultViewport.Width = max(20, msg.Width-8)
		m.resultViewport.Height = max(8, msg.Height-14)
		m.filterInput.Width = max(20, m.innerWidth()-lipgloss.Width(m.filterInput.Prompt)-1)
		m.editorInput.Width = max(20, m.innerWidth()-4)
		if m.usingObjectUI {
			objectInputWidth := max(20, m.innerWidth()-lipgloss.Width(m.objectValInput.Prompt)-2)
			m.objectKeyInput.Width = objectInputWidth
			m.objectValInput.Width = objectInputWidth
		}
		return m, nil

	case bannerTickMsg:
		if m.step != stepWelcome || !m.banner.Animating() {
			return m, nil
		}
		if !m.banner.Advance() {
			return m, nil
		}
		return m, tea.Tick(bannerFrameDelay(), func(time.Time) tea.Msg {
			return bannerTickMsg{}
		})

	case runFinishedMsg:
		m.running = false
		m.result = msg.Result
		m.resultViewport.SetContent(strings.TrimSpace(msg.Result.Output))
		m.step = stepResult
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if m.banner.Animating() {
			m.banner.Skip()
		}
	}

	switch m.step {
	case stepWelcome:
		return m.updateWelcome(msg)
	case stepHelp:
		return m.updateHelp(msg)
	case stepResource:
		return m.updateResource(msg)
	case stepAction:
		return m.updateAction(msg)
	case stepFields:
		return m.updateFields(msg)
	case stepEdit:
		return m.updateEdit(msg)
	case stepPreview:
		return m.updatePreview(msg)
	case stepResult:
		return m.updateResult(msg)
	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.step {
	case stepWelcome:
		return m.viewWelcome()
	case stepHelp:
		return m.viewHelp()
	case stepResource:
		return m.viewSelect("Choose a resource", m.filteredResources(), func(i int) (string, string) {
			resource := m.catalog.Resources[i]
			return resource.Name, resource.Usage
		}, m.resourceIndex, "Type to filter resources, then press Enter.")
	case stepAction:
		return m.viewActionSelect()
	case stepFields:
		return m.viewFields()
	case stepEdit:
		return m.viewEdit()
	case stepPreview:
		return m.viewPreview()
	case stepResult:
		return m.viewResult()
	default:
		return ""
	}
}

func (m Model) updateWelcome(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			m.welcomeIndex = max(0, m.welcomeIndex-1)
		case "down", "j":
			m.welcomeIndex = min(len(m.welcomeChoices())-1, m.welcomeIndex+1)
		case "?":
			m.previousStep = m.step
			m.step = stepHelp
		case "enter":
			switch m.welcomeIndex {
			case 0:
				m.step = stepResource
				m.resourceIndex = 0
				m.filterInput.SetValue("")
				m.filterInput.Focus()
			case 1:
				m.previousStep = m.step
				m.step = stepHelp
			case 2:
				return m, tea.Quit
			}
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) updateHelp(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "backspace", "enter", "q":
			m.step = m.previousStep
		}
	}
	return m, nil
}

func (m Model) updateResource(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.filterInput, cmd = m.filterInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		items := m.filteredResources()
		switch msg.String() {
		case "up", "k":
			m.resourceIndex = max(0, m.resourceIndex-1)
			return m, nil
		case "down", "j":
			m.resourceIndex = min(max(0, len(items)-1), m.resourceIndex+1)
			return m, nil
		case "esc":
			m.step = stepWelcome
			m.filterInput.SetValue("")
			return m, nil
		case "enter":
			if len(items) == 0 {
				return m, nil
			}
			m.state.ResourceIdx = items[min(m.resourceIndex, len(items)-1)]
			m.state.ActionIdx = -1
			m.actionIndex = 0
			m.filterInput.SetValue("")
			if len(m.catalog.Resources[m.state.ResourceIdx].Actions) == 1 {
				m.state.ActionIdx = 0
				m.step = stepFields
				m.fieldIndex = 0
				return m, nil
			}
			m.step = stepAction
			return m, nil
		}
	}

	return m, cmd
}

func (m Model) updateAction(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.filterInput, cmd = m.filterInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		items := m.filteredActions()
		switch msg.String() {
		case "up", "k":
			m.actionIndex = max(0, m.actionIndex-1)
			return m, nil
		case "down", "j":
			m.actionIndex = min(max(0, len(items)-1), m.actionIndex+1)
			return m, nil
		case "esc":
			m.step = stepResource
			m.filterInput.SetValue("")
			return m, nil
		case "enter":
			if len(items) == 0 {
				return m, nil
			}
			m.state.ActionIdx = items[min(m.actionIndex, len(items)-1)]
			m.step = stepFields
			m.fieldIndex = 0
			m.filterInput.SetValue("")
			return m, nil
		}
	}

	return m, cmd
}

func (m Model) updateFields(msg tea.Msg) (tea.Model, tea.Cmd) {
	fields := m.visibleFields()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			m.fieldIndex = max(0, m.fieldIndex-1)
		case "down", "j":
			m.fieldIndex = min(max(0, len(fields)-1), m.fieldIndex+1)
		case "o":
			m.showOptional = !m.showOptional
			m.fieldIndex = 0
		case "esc":
			m.step = stepAction
		case "p", "enter":
			if len(fields) == 0 {
				return m.moveToPreview()
			}
			field := fields[min(m.fieldIndex, len(fields)-1)]
			if msg.String() == "p" {
				return m.moveToPreview()
			}
			if field.Kind == FieldKindBool {
				if m.state.Values[field.Name] == "true" {
					delete(m.state.Values, field.Name)
				} else {
					m.state.Values[field.Name] = "true"
				}
				return m, nil
			}
			m.startEditing(field)
		}
	}
	return m, nil
}

func (m *Model) startEditing(field FieldSpec) {
	m.editingField = &field
	m.lastValidation = ""
	currentValue := m.state.Values[field.Name]
	m.usingObjectUI = field.Kind == FieldKindMultiline
	if m.usingObjectUI {
		m.objectRows = parseObjectRows(currentValue)
		m.objectRowIndex = -1
		m.objectFocusVal = false
		m.objectKeyInput.SetValue("")
		m.objectValInput.SetValue("")
		m.objectKeyInput.Placeholder = "e.g. $name"
		m.objectValInput.Placeholder = "e.g. Acme"
		m.objectKeyInput.Focus()
		m.objectValInput.Blur()
	} else {
		m.editorInput.SetValue(currentValue)
		m.editorInput.CursorEnd()
		m.editorInput.Focus()
		m.editorInput.Placeholder = field.Usage
	}
	m.step = stepEdit
}

func (m Model) updateEdit(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.editingField == nil {
		m.step = stepFields
		return m, nil
	}

	if m.usingObjectUI {
		return m.updateObjectEdit(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.step = stepFields
			m.editingField = nil
			m.lastValidation = ""
			return m, nil
		case "enter":
			value := strings.TrimSpace(m.editorInput.Value())
			if err := validateField(*m.editingField, value); err != nil {
				m.lastValidation = err.Error()
				return m, nil
			}
			if value == "" {
				delete(m.state.Values, m.editingField.Name)
			} else {
				m.state.Values[m.editingField.Name] = value
			}
			m.step = stepFields
			m.editingField = nil
			m.lastValidation = ""
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.editorInput, cmd = m.editorInput.Update(msg)
	return m, cmd
}

func (m Model) updateObjectEdit(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.step = stepFields
			m.editingField = nil
			m.lastValidation = ""
			m.objectRows = nil
			m.objectRowIndex = -1
			return m, nil
		case "tab":
			m.objectFocusVal = !m.objectFocusVal
			m.syncObjectInputFocus()
			return m, nil
		case "shift+tab":
			m.objectFocusVal = !m.objectFocusVal
			m.syncObjectInputFocus()
			return m, nil
		case "up":
			if len(m.objectRows) == 0 {
				return m, nil
			}
			if m.objectRowIndex <= 0 {
				m.objectRowIndex = len(m.objectRows) - 1
			} else {
				m.objectRowIndex--
			}
			m.loadObjectRowSelection()
			return m, nil
		case "down":
			if len(m.objectRows) == 0 {
				return m, nil
			}
			if m.objectRowIndex >= len(m.objectRows)-1 {
				m.objectRowIndex = 0
			} else {
				m.objectRowIndex++
			}
			m.loadObjectRowSelection()
			return m, nil
		case "ctrl+n":
			m.objectRowIndex = -1
			m.objectKeyInput.SetValue("")
			m.objectValInput.SetValue("")
			m.objectFocusVal = false
			m.syncObjectInputFocus()
			return m, nil
		case "ctrl+d", "ctrl+x", "delete":
			if m.objectRowIndex < 0 || m.objectRowIndex >= len(m.objectRows) {
				return m, nil
			}
			m.objectRows = append(m.objectRows[:m.objectRowIndex], m.objectRows[m.objectRowIndex+1:]...)
			if len(m.objectRows) == 0 {
				m.objectRowIndex = -1
				m.objectKeyInput.SetValue("")
				m.objectValInput.SetValue("")
			} else {
				m.objectRowIndex = min(m.objectRowIndex, len(m.objectRows)-1)
				m.loadObjectRowSelection()
			}
			return m, nil
		case "enter":
			if !m.objectFocusVal {
				if strings.TrimSpace(m.objectKeyInput.Value()) == "" {
					m.lastValidation = "Enter a field name first"
					return m, nil
				}
				m.objectFocusVal = true
				m.syncObjectInputFocus()
				return m, nil
			}
			if err := m.saveObjectRow(); err != nil {
				m.lastValidation = err.Error()
				return m, nil
			}
			m.lastValidation = ""
			return m, nil
		case "ctrl+s":
			if m.objectInputsDirty() {
				if err := m.saveObjectRow(); err != nil {
					m.lastValidation = err.Error()
					return m, nil
				}
			}
			value, err := serializeObjectRows(m.objectRows)
			if err != nil {
				m.lastValidation = err.Error()
				return m, nil
			}
			if err := validateField(*m.editingField, value); err != nil {
				m.lastValidation = err.Error()
				return m, nil
			}
			if value == "" {
				delete(m.state.Values, m.editingField.Name)
			} else {
				m.state.Values[m.editingField.Name] = value
			}
			m.step = stepFields
			m.editingField = nil
			m.lastValidation = ""
			m.objectRows = nil
			m.objectRowIndex = -1
			return m, nil
		}
	}

	var cmd tea.Cmd
	if m.objectFocusVal {
		m.objectValInput, cmd = m.objectValInput.Update(msg)
		return m, cmd
	}
	m.objectKeyInput, cmd = m.objectKeyInput.Update(msg)
	return m, cmd
}

func (m Model) updatePreview(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "backspace":
			m.step = stepFields
			return m, nil
		case "enter", "r":
			args, display, err := BuildArgs(m.catalog, m.state)
			if err != nil {
				m.lastValidation = err.Error()
				return m, nil
			}
			m.commandPreview = display
			m.running = true
			m.result = RunResult{}
			return m, RunCommandCmd(m.ctx, args, display)
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) updateResult(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.resultViewport, cmd = m.resultViewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "b":
			m.step = stepFields
			return m, nil
		case "r":
			args, display, err := BuildArgs(m.catalog, m.state)
			if err != nil {
				m.lastValidation = err.Error()
				return m, nil
			}
			m.commandPreview = display
			m.running = true
			return m, RunCommandCmd(m.ctx, args, display)
		case "h", "enter":
			m.step = stepWelcome
			m.lastValidation = ""
			return m, nil
		case "q":
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m Model) moveToPreview() (tea.Model, tea.Cmd) {
	fields := CollectFields(m.catalog, m.state)
	missing := MissingRequiredFields(fields, m.state.Values)
	if len(missing) > 0 {
		m.lastValidation = "Complete required fields before reviewing the command."
		if idx := indexOfField(fields, missing[0].Name); idx >= 0 {
			m.fieldIndex = idx
		}
		return m, nil
	}
	_, display, err := BuildArgs(m.catalog, m.state)
	if err != nil {
		m.lastValidation = err.Error()
		return m, nil
	}
	m.commandPreview = display
	m.lastValidation = ""
	m.step = stepPreview
	return m, nil
}

func (m Model) viewWelcome() string {
	choices := m.welcomeChoices()
	lines := make([]string, 0, len(choices))
	for i, choice := range choices {
		cursor := "  "
		style := m.theme.Secondary
		if i == m.welcomeIndex {
			cursor = "› "
			style = m.theme.Selected
		}
		lines = append(lines, style.Render(cursor+choice.Title))
		lines = append(lines, "  "+m.theme.Muted.Render(choice.Description))
		lines = append(lines, "")
	}

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		m.banner.View(m.theme, m.innerWidth()),
		"",
		m.theme.Title.Render("Guided command builder for the Lightfield API"),
		m.renderWrapped("Build a request, review the exact CLI command, then run it with confidence.", m.theme.Subtitle),
		"",
		strings.Join(lines, "\n"),
		m.renderWrapped("enter select • ? help • q quit", m.theme.Footer),
	)

	return m.renderPanel(body)
}

func (m Model) viewHelp() string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		m.theme.Title.Render("Help"),
		"",
		"• Use arrow keys or j/k to move through lists.",
		"• Press Enter to select the focused item.",
		"• Press Esc to go back.",
		"• Press Ctrl+C to quit from anywhere.",
		"",
		"Explicit commands still work exactly as before:",
		m.theme.Command.Render("lightfield account retrieve --api-key '...' --id acct_123"),
		"",
		m.theme.Footer.Render("Press Enter or Esc to return."),
	)
	return m.renderPanel(content)
}

func (m Model) viewSelect(title string, items []int, render func(int) (string, string), cursor int, footer string) string {
	lines := make([]string, 0, len(items))
	if len(items) == 0 {
		lines = append(lines, m.theme.Muted.Render("No matching results."))
	}
	itemWidth := max(20, m.innerWidth()-4)
	for idx, item := range items {
		prefix := "  "
		style := m.theme.Secondary
		if idx == min(cursor, max(0, len(items)-1)) {
			prefix = "› "
			style = m.theme.Selected
		}
		name, usage := render(item)
		lines = append(lines, style.Render(prefix+name))
		if strings.TrimSpace(usage) != "" {
			for _, wrapped := range wrapLines(singleLine(usage), itemWidth) {
				lines = append(lines, "  "+m.theme.Muted.Render(wrapped))
			}
		}
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		m.theme.Title.Render(title),
		"",
		m.filterInput.View(),
		"",
		strings.Join(lines, "\n"),
		"",
		m.renderWrapped(footer, m.theme.Footer),
	)
	return m.renderPanel(content)
}

func (m Model) viewActionSelect() string {
	if m.state.ResourceIdx < 0 || m.state.ResourceIdx >= len(m.catalog.Resources) {
		return m.renderPanel(m.theme.Muted.Render("No resource selected."))
	}

	resource := m.catalog.Resources[m.state.ResourceIdx]
	items := m.filteredActions()
	selectedIdx := min(m.actionIndex, max(0, len(items)-1))

	listLines := make([]string, 0, len(items))
	if len(items) == 0 {
		listLines = append(listLines, m.theme.Muted.Render("No matching actions."))
	}
	for idx, item := range items {
		action := resource.Actions[item]
		prefix := "  "
		style := m.theme.Secondary
		if idx == selectedIdx {
			prefix = "› "
			style = m.theme.Selected
		}
		listLines = append(listLines, style.Render(prefix+action.Name))
	}

	detailBlock := m.theme.Muted.Render("Use the list to choose an action.")
	if len(items) > 0 {
		action := resource.Actions[items[selectedIdx]]
		required := make([]string, 0)
		for _, field := range action.Fields {
			if field.Required {
				required = append(required, "--"+field.Name)
			}
		}

		detailWidth := max(20, m.innerWidth()-8)
		detailParts := []string{
			m.theme.Title.Render(action.Name),
			m.renderWrappedWidth(action.Usage, m.theme.Subtitle, detailWidth),
		}
		if len(required) > 0 {
			detailParts = append(detailParts, "", m.theme.Accent.Render("Required flags"))
			for _, name := range required {
				detailParts = append(detailParts, m.renderWrappedWidth("• "+name, m.theme.Secondary, detailWidth))
			}
		}
		detailBlock = lipgloss.NewStyle().PaddingLeft(1).Render(
			m.theme.Panel.Width(detailWidth).Render(strings.Join(detailParts, "\n")),
		)
	}

	header := lipgloss.JoinVertical(
		lipgloss.Left,
		m.theme.Title.Render("Choose an action"),
		m.renderWrapped("Resource: "+resource.Name, m.theme.Subtitle),
		"",
		m.filterInput.View(),
	)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		m.theme.Primary.Render("Actions"),
		strings.Join(listLines, "\n"),
		"",
		detailBlock,
		"",
		m.renderWrapped("enter select • esc back • type to filter", m.theme.Footer),
	)

	return m.renderPanel(content)
}

func (m Model) viewFields() string {
	fields := m.visibleFields()
	listLines := make([]string, 0, len(fields)+2)
	listLines = append(listLines, m.theme.Primary.Render(m.fieldsModeLabel()))
	listLines = append(listLines, m.theme.Muted.Render(strings.Repeat("─", max(20, m.innerWidth()-2))))

	if len(fields) == 0 {
		listLines = append(listLines, m.theme.Muted.Render("No fields in this view. Press o to show optional inputs."))
	}

	nameWidth := m.fieldNameColumnWidth(fields)
	statusWidth := 11
	valueWidth := max(12, m.innerWidth()-nameWidth-statusWidth-8)

	header := fmt.Sprintf("  %-*s  %-*s  %s", nameWidth, "Field", valueWidth, "Value", "Status")
	listLines = append(listLines, m.theme.Muted.Render(header))

	for i, field := range fields {
		cursor := "  "
		style := m.theme.Secondary
		if i == min(m.fieldIndex, max(0, len(fields)-1)) {
			cursor = "› "
			style = m.theme.Selected
		}

		value := fieldTableValue(field, m.state.Values[field.Name], valueWidth)
		status := fieldStatusLabel(field)
		line := fmt.Sprintf("%s %-*s  %-*s  %s", cursor, nameWidth, field.Name, valueWidth, value, status)
		listLines = append(listLines, style.Render(line))
	}

	detailBlock := m.theme.Muted.Render("Select a field to edit its value.")
	if len(fields) > 0 {
		field := fields[min(m.fieldIndex, len(fields)-1)]
		detailWidth := max(20, m.innerWidth()-8)
		detailParts := []string{
			m.theme.Title.Render("--" + field.Name),
		}

		meta := make([]string, 0, 3)
		meta = append(meta, fieldStatusLabel(field))
		if field.TypeName != "" {
			meta = append(meta, field.TypeName)
		}
		if field.RequestHint != "" {
			meta = append(meta, field.RequestHint)
		}
		if field.Global {
			meta = append(meta, "global")
		}
		detailParts = append(detailParts, m.renderWrappedWidth(strings.Join(meta, " • "), m.theme.Subtitle, detailWidth))

		if strings.TrimSpace(field.Usage) != "" {
			detailParts = append(detailParts, "", m.renderWrappedWidth(field.Usage, m.theme.Secondary, detailWidth))
		}

		current := fieldDetailValue(field, m.state.Values[field.Name])
		detailParts = append(detailParts, "", m.theme.Accent.Render("Current value"), m.renderWrappedWidth(current, m.theme.Secondary, detailWidth))

		if field.Required {
			detailParts = append(detailParts, "", m.theme.Accent.Render("Why it matters"), m.renderWrappedWidth("This field is required before you can review and run the command.", m.theme.Secondary, detailWidth))
		}

		detailBlock = lipgloss.NewStyle().PaddingLeft(1).Render(
			m.theme.Panel.Width(detailWidth).Render(strings.Join(detailParts, "\n")),
		)
	}

	progress := progressSummary(fields, m.state.Values)
	footer := "enter edit • o toggle optional • p review • esc back • ctrl+c quit"
	if len(fields) == 0 {
		footer = "o toggle optional • p review • esc back • ctrl+c quit"
	}

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		m.theme.Title.Render("Fill in command details"),
		m.renderWrapped(progress, m.theme.Subtitle),
		"",
		strings.Join(listLines, "\n"),
		"",
		detailBlock,
	)

	if m.lastValidation != "" {
		body = lipgloss.JoinVertical(lipgloss.Left, body, "", m.theme.Error.Render(m.lastValidation))
	}

	body = lipgloss.JoinVertical(lipgloss.Left, body, "", m.renderWrapped(footer, m.theme.Footer))
	return m.renderPanel(body)
}

func (m Model) visibleFields() []FieldSpec {
	fields := CollectFields(m.catalog, m.state)
	if m.showOptional {
		return fields
	}

	visible := make([]FieldSpec, 0, len(fields))
	for _, field := range fields {
		if field.Required || field.Name == "api-key" {
			visible = append(visible, field)
		}
	}
	return visible
}

func (m *Model) syncObjectInputFocus() {
	if m.objectFocusVal {
		m.objectKeyInput.Blur()
		m.objectValInput.Focus()
		return
	}
	m.objectValInput.Blur()
	m.objectKeyInput.Focus()
}

func (m *Model) loadObjectRowSelection() {
	if m.objectRowIndex < 0 || m.objectRowIndex >= len(m.objectRows) {
		return
	}
	row := m.objectRows[m.objectRowIndex]
	m.objectKeyInput.SetValue(row.Key)
	m.objectValInput.SetValue(row.Value)
	m.objectKeyInput.CursorEnd()
	m.objectValInput.CursorEnd()
}

func (m *Model) saveObjectRow() error {
	key := strings.TrimSpace(m.objectKeyInput.Value())
	value := strings.TrimSpace(m.objectValInput.Value())
	if key == "" {
		return fmt.Errorf("enter a field name before saving this row")
	}

	row := objectRow{Key: key, Value: value}
	if m.objectRowIndex >= 0 && m.objectRowIndex < len(m.objectRows) {
		m.objectRows[m.objectRowIndex] = row
	} else {
		m.objectRows = append(m.objectRows, row)
	}

	m.objectRowIndex = -1
	m.objectKeyInput.SetValue("")
	m.objectValInput.SetValue("")
	m.objectFocusVal = false
	m.syncObjectInputFocus()
	return nil
}

func (m Model) objectInputsDirty() bool {
	return strings.TrimSpace(m.objectKeyInput.Value()) != "" || strings.TrimSpace(m.objectValInput.Value()) != ""
}

func (m Model) fieldsModeLabel() string {
	if m.showOptional {
		return "All inputs"
	}
	return "Required inputs"
}

func (m Model) fieldNameColumnWidth(fields []FieldSpec) int {
	width := 12
	for _, field := range fields {
		width = max(width, len(field.Name))
	}
	return min(20, width)
}

func (m Model) viewEdit() string {
	if m.editingField == nil {
		return ""
	}

	field := *m.editingField
	if m.usingObjectUI {
		return m.viewObjectEdit(field)
	}

	kindHint := "Press Enter to save."
	editorView := m.editorInput.View()

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		m.theme.Title.Render("--"+field.Name),
		m.renderWrapped(singleLine(field.Usage), m.theme.Subtitle),
		"",
		editorView,
		"",
		m.renderWrapped(kindHint, m.theme.Footer),
	)
	if m.lastValidation != "" {
		body = lipgloss.JoinVertical(lipgloss.Left, body, "", m.theme.Error.Render(m.lastValidation))
	}
	return m.renderPanel(body)
}

func (m Model) viewObjectEdit(field FieldSpec) string {
	keyWidth := min(22, max(14, m.innerWidth()/3))
	valueWidth := max(18, m.innerWidth()-keyWidth-7)
	tableWidth := keyWidth + valueWidth + 2
	rowLines := []string{
		m.theme.Muted.Render(fmt.Sprintf("  %-*s  %-*s", keyWidth, "Field", valueWidth, "Value")),
		m.theme.Muted.Render("  " + strings.Repeat("─", max(12, tableWidth))),
	}
	if len(m.objectRows) == 0 {
		rowLines = append(rowLines, m.theme.Muted.Render("  No values added yet."))
	}
	for i, row := range m.objectRows {
		cursor := "  "
		style := m.theme.Secondary
		if i == m.objectRowIndex {
			cursor = "› "
			style = m.theme.Selected
		}
		value := row.Value
		if strings.TrimSpace(value) == "" {
			value = "[empty]"
		}
		line := fmt.Sprintf("%s %-*s  %s", cursor, keyWidth, truncate(row.Key, keyWidth), truncate(value, valueWidth))
		rowLines = append(rowLines, style.Render(line))
	}
	for len(rowLines) < 8 {
		rowLines = append(rowLines, "")
	}

	status := "Adding a new row"
	if m.objectRowIndex >= 0 && m.objectRowIndex < len(m.objectRows) {
		status = "Editing selected row"
	}

	helperTitle, helperText, helperExample := objectEditorCopy(field)
	formHint := "Enter moves key -> value -> add row"
	if m.objectFocusVal {
		formHint = "Press Enter to add this row"
	}

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		m.theme.Title.Render(helperTitle),
		m.renderWrapped(helperText, m.theme.Subtitle),
		"",
		m.theme.Accent.Render("Example"),
		m.renderWrapped(helperExample, m.theme.Secondary),
		"",
		m.theme.Accent.Render(status),
		m.objectKeyInput.View(),
		m.objectValInput.View(),
		m.renderWrapped(formHint, m.theme.Footer),
		"",
		m.theme.Accent.Render(fmt.Sprintf("Saved values (%d)", len(m.objectRows))),
		strings.Join(rowLines, "\n"),
		"",
		m.renderWrapped("tab switch input • ↑↓ load saved row • delete/ctrl+x delete • ctrl+n new row • ctrl+s save field • esc cancel", m.theme.Footer),
	)
	if m.lastValidation != "" {
		body = lipgloss.JoinVertical(lipgloss.Left, body, "", m.theme.Error.Render(m.lastValidation))
	}
	return m.renderPanel(body)
}

func (m Model) viewPreview() string {
	body := lipgloss.JoinVertical(
		lipgloss.Left,
		m.theme.Title.Render("Review command"),
		m.renderWrapped("The command below is exactly what Lightfield will run.", m.theme.Subtitle),
		"",
		m.theme.Command.Width(max(20, m.innerWidth()-4)).Render(m.commandPreview),
		"",
		m.renderWrapped("enter run • esc back • q quit", m.theme.Footer),
	)
	if m.lastValidation != "" {
		body = lipgloss.JoinVertical(lipgloss.Left, body, "", m.theme.Error.Render(m.lastValidation))
	}
	return m.renderPanel(body)
}

func (m Model) viewResult() string {
	status := m.theme.Success.Render("Request completed.")
	if m.running {
		status = m.theme.Accent.Render("Running command…")
	} else if m.result.Err != nil {
		status = m.theme.Error.Render("Command failed.")
	}

	output := strings.TrimSpace(m.result.Output)
	if output == "" && m.result.Err == nil {
		output = "Command completed without additional output."
	}
	if output == "" && m.result.Err != nil {
		output = m.result.Err.Error()
	}

	body := lipgloss.JoinVertical(
		lipgloss.Left,
		status,
		"",
		m.renderWrapped(m.result.Command, m.theme.Muted),
		"",
		m.resultViewport.View(),
		"",
		m.renderWrapped("b back • r run again • enter home • q quit", m.theme.Footer),
	)
	if !m.running && output != m.resultViewport.View() {
		m.resultViewport.SetContent(output)
	}
	return m.renderPanel(body)
}

func (m Model) actionFooter() string {
	if m.state.ResourceIdx < 0 || m.state.ResourceIdx >= len(m.catalog.Resources) {
		return "Use Enter to choose an action."
	}
	items := m.filteredActions()
	if len(items) == 0 {
		return "Use Enter to choose an action."
	}
	action := m.catalog.Resources[m.state.ResourceIdx].Actions[items[min(m.actionIndex, len(items)-1)]]
	required := make([]string, 0)
	for _, field := range action.Fields {
		if field.Required {
			required = append(required, "--"+field.Name)
		}
	}
	if len(required) == 0 {
		return singleLine(action.Usage)
	}
	return singleLine(action.Usage) + "  Required: " + strings.Join(required, ", ")
}

func (m Model) wrap(content string) string {
	outerWidth := m.outerWidth()
	if m.width == 0 {
		return lipgloss.NewStyle().Width(outerWidth).Render(content)
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.NewStyle().Width(outerWidth).Render(content))
}

func (m Model) renderPanel(content string) string {
	contentBox := lipgloss.NewStyle().
		Width(m.innerWidth()).
		Height(m.innerHeight()).
		MaxHeight(m.innerHeight()).
		Align(lipgloss.Left, lipgloss.Top).
		Render(content)

	return m.wrap(
		m.theme.Panel.
			Width(m.innerWidth()).
			Height(m.innerHeight()).
			Render(contentBox),
	)
}

func (m Model) renderWrapped(text string, style lipgloss.Style) string {
	return renderWrappedText(singleLine(text), style, m.innerWidth())
}

func (m Model) renderWrappedWidth(text string, style lipgloss.Style, width int) string {
	return renderWrappedText(singleLine(text), style, width)
}

func (m Model) outerWidth() int {
	if m.width == 0 {
		return 96
	}
	return max(60, min(100, m.width-4))
}

func (m Model) innerWidth() int {
	return max(20, m.outerWidth()-6)
}

func (m Model) outerHeight() int {
	if m.height == 0 {
		return 30
	}
	return max(18, min(34, m.height-4))
}

func (m Model) innerHeight() int {
	return max(10, m.outerHeight()-4)
}

func (m Model) welcomeChoices() []welcomeChoice {
	return []welcomeChoice{
		{Title: "Build a command", Description: "Start with a resource, add flags, and review the exact CLI invocation."},
		{Title: "Open help", Description: "See keyboard shortcuts and how the guided flow relates to the normal CLI."},
		{Title: "Quit", Description: "Leave the assistant and return to the terminal."},
	}
}

func (m Model) filteredResources() []int {
	query := strings.ToLower(strings.TrimSpace(m.filterInput.Value()))
	indices := make([]int, 0, len(m.catalog.Resources))
	for i, resource := range m.catalog.Resources {
		haystack := strings.ToLower(resource.Name + " " + resource.Usage)
		if query == "" || strings.Contains(haystack, query) {
			indices = append(indices, i)
		}
	}
	if m.resourceIndex >= len(indices) {
		m.resourceIndex = max(0, len(indices)-1)
	}
	return indices
}

func (m Model) filteredActions() []int {
	if m.state.ResourceIdx < 0 || m.state.ResourceIdx >= len(m.catalog.Resources) {
		return nil
	}
	query := strings.ToLower(strings.TrimSpace(m.filterInput.Value()))
	actions := m.catalog.Resources[m.state.ResourceIdx].Actions
	indices := make([]int, 0, len(actions))
	for i, action := range actions {
		haystack := strings.ToLower(action.Name + " " + action.Usage)
		if query == "" || strings.Contains(haystack, query) {
			indices = append(indices, i)
		}
	}
	if m.actionIndex >= len(indices) {
		m.actionIndex = max(0, len(indices)-1)
	}
	return indices
}

func seedValues(root *cli.Command, catalog Catalog) map[string]string {
	values := make(map[string]string)
	for _, field := range catalog.GlobalFields {
		if !root.IsSet(field.Name) {
			continue
		}
		values[field.Name] = stringifyValue(root.Value(field.Name), field.Kind)
	}
	return values
}

func validateField(field FieldSpec, value string) error {
	if field.Required && strings.TrimSpace(value) == "" {
		return fmt.Errorf("--%s is required", field.Name)
	}
	if field.Kind == FieldKindNumber && strings.TrimSpace(value) != "" {
		if strings.Contains(value, ".") {
			var f float64
			if _, err := fmt.Sscanf(value, "%f", &f); err != nil {
				return fmt.Errorf("--%s expects a number", field.Name)
			}
		} else {
			var i int64
			if _, err := fmt.Sscanf(value, "%d", &i); err != nil {
				return fmt.Errorf("--%s expects a number", field.Name)
			}
		}
	}
	return nil
}

func currentValueSummary(field FieldSpec, value string) string {
	if strings.TrimSpace(value) == "" {
		if field.Required {
			return "missing"
		}
		return ""
	}
	switch field.Kind {
	case FieldKindBool:
		return "enabled"
	case FieldKindMultiline:
		return truncate(singleLine(value), 40)
	default:
		return truncate(value, 40)
	}
}

func fieldStatusLabel(field FieldSpec) string {
	switch {
	case field.Required:
		return "required"
	case field.Name == "api-key":
		return "auth"
	case field.Global:
		return "global"
	default:
		return "optional"
	}
}

func parseObjectRows(raw string) []objectRow {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}

	var object map[string]any
	if err := json.Unmarshal([]byte(raw), &object); err != nil {
		return nil
	}

	keys := make([]string, 0, len(object))
	for key := range object {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	rows := make([]objectRow, 0, len(keys))
	for _, key := range keys {
		rows = append(rows, objectRow{
			Key:   key,
			Value: displayObjectValue(object[key]),
		})
	}
	return rows
}

func serializeObjectRows(rows []objectRow) (string, error) {
	if len(rows) == 0 {
		return "", nil
	}

	object := make(map[string]any, len(rows))
	for _, row := range rows {
		key := strings.TrimSpace(row.Key)
		if key == "" {
			continue
		}
		value, err := coerceObjectValue(row.Value)
		if err != nil {
			return "", fmt.Errorf("invalid value for %q: %w", key, err)
		}
		object[key] = value
	}

	if len(object) == 0 {
		return "", nil
	}

	encoded, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func displayObjectValue(value any) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return v
	default:
		encoded, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprint(v)
		}
		return string(encoded)
	}
}

func coerceObjectValue(value string) (any, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return "", nil
	}

	if strings.HasPrefix(trimmed, "{") || strings.HasPrefix(trimmed, "[") || strings.HasPrefix(trimmed, "\"") {
		var parsed any
		if err := json.Unmarshal([]byte(trimmed), &parsed); err != nil {
			return nil, err
		}
		return parsed, nil
	}

	switch trimmed {
	case "true":
		return true, nil
	case "false":
		return false, nil
	case "null":
		return nil, nil
	}

	if i, err := strconv.ParseInt(trimmed, 10, 64); err == nil {
		return i, nil
	}
	if f, err := strconv.ParseFloat(trimmed, 64); err == nil {
		return f, nil
	}

	return trimmed, nil
}

func objectEditorCopy(field FieldSpec) (title string, helper string, example string) {
	switch field.Name {
	case "fields":
		return "Set field values",
			"Add one field per row. Use the field slug on the left and the value on the right. Lightfield will turn the rows into JSON automatically when you save.",
			"Example: $name -> Acme, tier -> enterprise, employeeCount -> 42"
	default:
		label := "--" + field.Name
		helper = field.Usage
		if strings.TrimSpace(helper) == "" {
			helper = "Add one key/value pair per row. The rows will be saved as JSON automatically."
		}
		return label,
			singleLine(helper),
			"Example: key -> value"
	}
}

func fieldTableValue(field FieldSpec, value string, width int) string {
	if strings.TrimSpace(value) == "" {
		switch {
		case field.Required:
			return "[required]"
		case field.Name == "api-key":
			return "[recommended]"
		default:
			return "[empty]"
		}
	}

	switch field.Kind {
	case FieldKindBool:
		if value == "true" {
			return "enabled"
		}
		return "[empty]"
	case FieldKindMultiline:
		if rows := parseObjectRows(value); len(rows) > 0 {
			return truncate(fmt.Sprintf("[set • %d fields]", len(rows)), width)
		}
		return truncate(fmt.Sprintf("[set • %d chars]", len(value)), width)
	default:
		return truncate(value, width)
	}
}

func fieldDetailValue(field FieldSpec, value string) string {
	if strings.TrimSpace(value) == "" {
		switch {
		case field.Required:
			return "This field still needs a value."
		case field.Name == "api-key":
			return "No API key set yet. You can still review the command, but the request may fail without authentication."
		default:
			return "No value set yet."
		}
	}

	switch field.Kind {
	case FieldKindBool:
		if value == "true" {
			return "Enabled"
		}
		return "Disabled"
	case FieldKindMultiline:
		if rows := parseObjectRows(value); len(rows) > 0 {
			return fmt.Sprintf("%d field values configured.", len(rows))
		}
		return value
	default:
		return value
	}
}

func progressSummary(fields []FieldSpec, values map[string]string) string {
	required := 0
	completed := 0
	for _, field := range fields {
		if !field.Required {
			continue
		}
		required++
		if strings.TrimSpace(values[field.Name]) != "" {
			completed++
		}
	}
	if required == 0 {
		return "No required fields."
	}
	return fmt.Sprintf("%d of %d required fields completed", completed, required)
}

func truncate(value string, limit int) string {
	if len(value) <= limit {
		return value
	}
	return value[:limit-1] + "…"
}

func singleLine(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

func wrapLines(value string, width int) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	wrapped := wordwrap.String(value, width)
	return strings.Split(wrapped, "\n")
}

func renderWrappedText(text string, style lipgloss.Style, width int) string {
	lines := wrapLines(text, width)
	if len(lines) == 0 {
		return ""
	}

	rendered := make([]string, 0, len(lines))
	for _, line := range lines {
		rendered = append(rendered, style.Render(line))
	}

	return strings.Join(rendered, "\n")
}

func indexOfField(fields []FieldSpec, name string) int {
	for i, field := range fields {
		if field.Name == name {
			return i
		}
	}
	return -1
}

func bannerFrameDelay() time.Duration {
	return 70 * time.Millisecond
}

func reducedMotionRequested() bool {
	return os.Getenv("LIGHTFIELD_REDUCED_MOTION") == "1"
}

func isInteractiveSession() bool {
	return term.IsTerminal(os.Stdout.Fd()) && term.IsTerminal(os.Stdin.Fd())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
