package ui

// Package ui provides Bubble Tea TUI components for the desktop automation CLI.
// This includes interactive menus, forms, and displays for automation tasks.

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the main TUI model
type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	title    string
}

// NewModel creates a new TUI model
func NewModel(title string, choices []string) Model {
	return Model{
		choices:  choices,
		selected: make(map[int]struct{}),
		title:    title,
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

// View renders the model
func (m Model) View() string {
	s := fmt.Sprintf("%s\n\n", m.title)

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit, enter/space to select.\n"
	return s
}

// MenuModel represents a simple menu model
type MenuModel struct {
	title  string
	items  []MenuItem
	cursor int
	quit   bool
}

// MenuItem represents a menu item
type MenuItem struct {
	Title       string
	Description string
	Action      func() tea.Cmd
}

// NewMenuModel creates a new menu model
func NewMenuModel(title string, items []MenuItem) MenuModel {
	return MenuModel{
		title: title,
		items: items,
	}
}

// Init initializes the menu model
func (m MenuModel) Init() tea.Cmd {
	return nil
}

// Update handles menu updates
func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			if m.cursor < len(m.items) && m.items[m.cursor].Action != nil {
				return m, m.items[m.cursor].Action()
			}
		}
	}
	return m, nil
}

// View renders the menu
func (m MenuModel) View() string {
	if m.quit {
		return "Goodbye!\n"
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s\n\n", m.title))

	for i, item := range m.items {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}

		b.WriteString(fmt.Sprintf("%s%s\n", cursor, item.Title))
		if item.Description != "" {
			b.WriteString(fmt.Sprintf("   %s\n", item.Description))
		}
		b.WriteString("\n")
	}

	b.WriteString("Use arrow keys to navigate, enter to select, q to quit.\n")
	return b.String()
}

// RunInteractiveMenu starts an interactive menu
func RunInteractiveMenu(title string, items []MenuItem) error {
	p := tea.NewProgram(NewMenuModel(title, items))
	_, err := p.Run()
	return err
}
