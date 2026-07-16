package tui

import (
	tea "charm.land/bubbletea/v2"
)

func Run() error {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

type model struct {
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	return tea.NewView("rogue")
}
