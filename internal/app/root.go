package app

import (
	"fmt"
	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type root struct {
	cursor  int
	choices [3]string
	config models.UserConfig
}
func (r *root) View() string {
	var ret string;

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Render("ConfigFlipper")

	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("63"))

	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("250"))

	ret += title
	ret += "\n\n"

	for i, choice := range(r.choices) {
		if i == r.cursor {
			line := fmt.Sprintf("[x] %s", choice)
			ret += selectedStyle.Render(line)
			ret += "\n"
		} else {
			line := fmt.Sprintf("[ ] %s", choice)
			ret += normalStyle.Render(line)
			ret += "\n"
		}
	}
	return ret
}

func (r *root) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if r.cursor < len(r.choices) - 1 {
				r.cursor++
			}
		case "k":
			if r.cursor > 0 {
				r.cursor--
			}
		case "enter":
			if r.choices[r.cursor] == "Open Config Folder" {
				r.config.OpenFile()
			}
		}
	}
}
