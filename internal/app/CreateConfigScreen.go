package app

import (
	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

type CreateConfigScreen struct {
	/// list out the files where the user is currently
	config  *models.UserConfig
	cursor  int
	choices [1]string
}

func NewCreateConfigScreen(config *models.UserConfig) *CreateConfigScreen {
	return &CreateConfigScreen{
		config: config,
		cursor: 0,
		choices: [1]string{
			"Go Back",
		},
	}
}

func (cc *CreateConfigScreen) View() string {
	var ret string

	ret += NormalStyle.Render("Create Config Screen")
	ret += "\n\n"

	if cc.choices[0] == "Go Back" {
		if cc.cursor == 0 {
			ret += SelectedStyle.Render("Go Back")
		} else {
			ret += NormalStyle.Render("Go Back")
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
				ret += SelectedStyle.Render("󰂺 ")
				ret += SelectedStyle.Render(entry.Name())
			} else {
				ret += NormalStyle.Render("󰂺 ")
				ret += NormalStyle.Render(entry.Name())
			}
		case models.Folder:
			if selected {
				ret += SelectedStyle.Render(" ")
				ret += SelectedStyle.Render(entry.Name())
			} else {
				ret += NormalStyle.Render(" ")
				ret += NormalStyle.Render(entry.Name())
			}
		}
		ret += "\n"
	}
	return ret
}

func (cc *CreateConfigScreen) Update(msg tea.Msg, onSetScreen func(screen screen)) {
	totalLength := len(cc.choices) + len(cc.config.CurrentDirectory.Data)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if cc.cursor == 0 {
				onSetScreen(screenRoot)
			} else {
				data := cc.config.CurrentDirectory.Data[cc.cursor-1]
				cc.config.Data = data
				onSetScreen(screenCreateConfigConfirmation)
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
