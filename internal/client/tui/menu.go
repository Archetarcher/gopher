package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type menu struct {
	options       []menuItem
	selectedIndex int
}

type menuItem struct {
	text    string
	onPress func() tea.Msg
}

func (m menu) Init() tea.Cmd { return nil }

func (m menu) View() string {
	var options []string
	for i, o := range m.options {
		if i == m.selectedIndex {
			options = append(options, fmt.Sprintf("-> %s", o.text))
		} else {
			options = append(options, fmt.Sprintf("   %s", o.text))
		}
	}
	return fmt.Sprintf(`%s

Press enter/return to select a list item, arrow keys to move, or Ctrl+C to exit.`,
		strings.Join(options, "\n"))
}
func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		case "down", "right", "up", "left":
			return m.moveCursor(msg.(tea.KeyMsg)), nil
		}
	}
	return m, nil
}

func (m menu) moveCursor(msg tea.KeyMsg) menu {
	switch msg.String() {
	case "up", "left":
		m.selectedIndex--
	case "down", "right":
		m.selectedIndex++
	default:
		// do nothing
	}

	optCount := len(m.options)
	m.selectedIndex = (m.selectedIndex + optCount) % optCount
	return m
}
