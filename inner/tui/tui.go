package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/tentharmy/rogue/inner/app"
	"github.com/tentharmy/rogue/inner/geom"
)

func Run(serv *app.GameService) error {
	p := tea.NewProgram(initialModel(serv))
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

type model struct {
	serv *app.GameService
}

func initialModel(serv *app.GameService) model {
	return model{
		serv: serv,
	}
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
		case "up":
			m.serv.MovePlayer(geom.Vec2{X: 0, Y: -1})
		case "down":
			m.serv.MovePlayer(geom.Vec2{X: 0, Y: 1})
		case "left":
			m.serv.MovePlayer(geom.Vec2{X: -1, Y: 0})
		case "right":
			m.serv.MovePlayer(geom.Vec2{X: 1, Y: 0})
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	snapshot := m.serv.GetSnapshot()
	v := tea.NewView(fmt.Sprintf("Player: %v", snapshot.PlayerPos))
	v.AltScreen = true
	return v
}
