package app

import (
	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var appStyle = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63"))

type screen int

const (
	screenRoot screen = iota
	screenCreateConfig
	screenSelectConfig
)

type model struct {
	screen screen

	root               root
	createConfigScreen createConfigScreen
}

func InitialModel(config models.UserConfig) model {
	return model{
		screen: screenRoot,
		/// Creating Root Screen
		root: root{
			cursor: 0,
			choices: [3]string{
				"Create New Configuration",
				"Select Configuration",
				"Open Config Folder",
			},
			config: config,
		},
		/// Other Screens Here
		createConfigScreen: createConfigScreen{
			config: config,
			cursor: 0,
			choices: [1]string{
				"Go Back",
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	/// Global Keybinds
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	switch m.screen {
	case screenRoot:
		m.root.Update(msg, func(screen screen) {
			m.screen = screen
		})
	case screenCreateConfig:
		m.createConfigScreen.Update(msg, func(screen screen) {
			m.screen = screen
		})
	default:
		break
	}

	return m, nil
}

func (m model) View() string {
	switch m.screen {
	case screenRoot:
		return appStyle.Render(m.root.View())
	case screenCreateConfig:
		return appStyle.Render(m.createConfigScreen.View())
	default:
		return "Unkown\n"
	}
}
