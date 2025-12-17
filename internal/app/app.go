package app

import (
	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var AppStyle = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63"))

var SelectedStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("229")).
	Background(lipgloss.Color("63"))

var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("250"))

var TitleUnderline = lipgloss.NewStyle().
	Bold(true).Underline(true).
	Foreground(lipgloss.Color("250"))

var NormalStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("250"))

type screen int

const (
	screenRoot screen = iota
	screenCreateConfig
	screenCreateConfigConfirmation
	screenSelectConfig
)

type model struct {
	screen screen

	root                           root
	createConfigScreen             createConfigScreen
	createConfigConfirmationScreen createConfigConfirmationScreen
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
		/// Later adding data during runtime
		createConfigConfirmationScreen: createConfigConfirmationScreen{
			cursor: 0,
			choices: [2]string{
				"Go Back",
				"Create Config",
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
		m.createConfigScreen.Update(msg, func(screen screen, data models.CurrentDirectoryData) {
			if data != nil {
				m.createConfigConfirmationScreen.data = data
			}
			m.screen = screen
		})
	case screenCreateConfigConfirmation:
		m.createConfigConfirmationScreen.Update(msg, func(screen screen) {
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
		return AppStyle.Render(m.root.View())
	case screenCreateConfig:
		return AppStyle.Render(m.createConfigScreen.View())
	case screenCreateConfigConfirmation:
		return AppStyle.Render(m.createConfigConfirmationScreen.View())
	default:
		return "Unkown\n"
	}
}
