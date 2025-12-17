package app

import (
	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type createConfigConfirmationScreen struct {
	data    models.CurrentDirectoryData
	cursor  int
	choices [2]string
}

func (cc *createConfigConfirmationScreen) View() string {
	var ret string

	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("63"))

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("250"))
	titleUnderline := lipgloss.NewStyle().
		Bold(true).Underline(true).
		Foreground(lipgloss.Color("250"))

	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("250"))

	ret += titleStyle.Render("Are You Sure You Want to Create This Config?")
	ret += "\n"

	ret += titleUnderline.Render(cc.data.Name())
	ret += "\n\n"

	for i, choice := range cc.choices {
		if i == cc.cursor {
			ret += selectedStyle.Render("[x] ")
			ret += selectedStyle.Render(choice)
		} else {
			ret += normalStyle.Render("[ ] ")
			ret += normalStyle.Render(choice)
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
