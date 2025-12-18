package app

import (
	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	AppStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63"))

	SelectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("63"))

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("250"))

	TitleUnderline = lipgloss.NewStyle().
			Bold(true).Underline(true).
			Foreground(lipgloss.Color("250"))

	NormalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("250"))

	BlueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5555C2"))

	BitDimBlueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#30305C"))

	DimBlueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#242447"))
	FocusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("63"))
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000"))
)

type screen int

const (
	screenRoot screen = iota
	screenCreateConfig
	screenCreateConfigConfirmation
	screenSelectConfig
)

type Model struct {
	screen screen

	root                           *Root
	createConfigScreen             *CreateConfigScreen
	createConfigConfirmationScreen *CreateConfigConfirmationScreen
}

func InitialModel(config *models.UserConfig) Model {
	model := Model{
		screen: screenRoot,
		/// Creating Root Screen
		root:                           NewRoot(config),
		createConfigScreen:             NewCreateConfigScreen(config),
		createConfigConfirmationScreen: NewCreateConfigConfirmationScreen(config),
	}

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if screen == screenCreateConfigConfirmation {
				m.createConfigConfirmationScreen.input.SetValue("Config_" + m.createConfigScreen.config.Data.Name())
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

func (m Model) View() string {
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
