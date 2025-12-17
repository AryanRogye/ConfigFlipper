package app

import (
	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

type createConfigConfirmationScreen struct {
	data    models.CurrentDirectoryData
	cursor  int
	choices [2]string
}

func (cc *createConfigConfirmationScreen) View() string {
	var ret string

	ret += TitleStyle.Render("Are You Sure You Want to Create This Config?")
	ret += "\n"

	ret += TitleUnderline.Render(cc.data.Name())
	ret += "\n\n"

	for i, choice := range cc.choices {
		if i == cc.cursor {
			ret += SelectedStyle.Render("[x] ")
			ret += SelectedStyle.Render(choice)
		} else {
			ret += NormalStyle.Render("[ ] ")
			ret += NormalStyle.Render(choice)
		}
		ret += "\n"
	}
	return ret
}

func (cc *createConfigConfirmationScreen) Update(msg tea.Msg, onSetScreen func(screen screen)) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			choice := cc.choices[cc.cursor]
			if choice == "Go Back" {
				onSetScreen(screenCreateConfig)
			}
			if choice == "Create Config" {
				/// Screen Root For Now
				onSetScreen(screenRoot)
			}
		case "j":
			if cc.cursor < len(cc.choices)-1 {
				cc.cursor++
			}
		case "k":
			if cc.cursor > 0 {
				cc.cursor--
			}
		}
	}
}
