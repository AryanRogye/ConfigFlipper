package app

import (
	"fmt"

	"github.com/AryanRogye/ConfigFlipper/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

type root struct {
	cursor  int
	choices [3]string
	config  models.UserConfig
}

func (r *root) View() string {
	var ret string

	ret += TitleStyle.Render("ConfigFlipper")
	ret += "\n\n"

	for i, choice := range r.choices {
		if i == r.cursor {
			line := fmt.Sprintf("[x] %s", choice)
			ret += SelectedStyle.Render(line)
			ret += "\n"
		} else {
			line := fmt.Sprintf("[ ] %s", choice)
			ret += NormalStyle.Render(line)
			ret += "\n"
		}
	}
	return ret
}

func (r *root) Update(msg tea.Msg, onSetScreen func(screen screen)) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if r.cursor < len(r.choices)-1 {
				r.cursor++
			}
		case "k":
			if r.cursor > 0 {
				r.cursor--
			}
		case "enter":

			choice := r.choices[r.cursor]
			if choice == "Open Config Folder" {
				r.config.OpenConfigFolder()
			} else if choice == "Create New Configuration" {
				onSetScreen(screenCreateConfig)
			}
		}
	}
}
