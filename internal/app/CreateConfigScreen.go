package app

import (
	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type createConfigScreen struct {
	/// list out the files where the user is currently
	config  models.UserConfig
	cursor  int
	choices [1]string
}

func (cc *createConfigScreen) View() string {
	var ret string

	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("63"))

	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("250"))

	ret += normalStyle.Render("Create Config Screen")
	ret += "\n\n"

	if cc.choices[0] == "Go Back" {
		if cc.cursor == 0 {
			ret += selectedStyle.Render("Go Back")
		} else {
			ret += normalStyle.Render("Go Back")
		}
		ret += "\n"
	}

	for i, entry := range cc.config.CurrentDirectory.Data {
		var selected bool

		if cc.cursor == i+1 {
			selected = true
		} else {
			selected = false
		}

		switch entry.(type) {
		case models.File:
			if selected {
				ret += selectedStyle.Render("󰂺 ")
				ret += selectedStyle.Render(entry.Name())
			} else {
				ret += normalStyle.Render("󰂺 ")
				ret += normalStyle.Render(entry.Name())
			}
		case models.Folder:
			if selected {
				ret += selectedStyle.Render(" ")
				ret += selectedStyle.Render(entry.Name())
			} else {
				ret += normalStyle.Render(" ")
				ret += normalStyle.Render(entry.Name())
			}
		}
		ret += "\n"
	}
	return ret
}

func (cc *createConfigScreen) Update(msg tea.Msg, onSetScreen func(screen screen, data models.CurrentDirectoryData)) {
	totalLength := len(cc.choices) + len(cc.config.CurrentDirectory.Data)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if cc.cursor == 0 {
				onSetScreen(screenRoot, nil)
			} else {
				data := cc.config.CurrentDirectory.Data[cc.cursor-1]
				onSetScreen(screenCreateConfigConfirmation, data)
			}
		case "j":
			if cc.cursor < totalLength-1 {
				cc.cursor++
			}
		case "k":
			if cc.cursor > 0 {
				cc.cursor--
			}
		}
	}
}
