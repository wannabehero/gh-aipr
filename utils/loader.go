package utils

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type loaderModel struct {
	spinner spinner.Model
	text    string
}

func newLoaderModel(text string) loaderModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return loaderModel{spinner: s, text: text}
}

func (m loaderModel) Init() tea.Cmd {
	return m.spinner.Tick
}

type doneMsg struct{}

func (m loaderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case doneMsg:
		return m, tea.Quit
	}
	return m, nil
}

func (m loaderModel) View() string {
	return fmt.Sprintf("%s %s", m.spinner.View(), m.text)
}

// StartLoader runs the spinner asynchronously and returns a function to stop it.
func StartLoader(text string) func() {
	program := tea.NewProgram(newLoaderModel(text))
	go func() {
		// ignore errors so loader doesn't block main flow
		_, _ = program.Run()
	}()
	return func() {
		program.Quit()
		fmt.Println()
	}
}
