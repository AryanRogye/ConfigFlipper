package app

import (
	"strconv"

	"github.com/AryanRogye/ConfigFlipper/internal/models"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type CreateConfigConfirmationScreen struct {
	cursor  int
	choices [2]string
	input   textinput.Model
	config  *models.UserConfig

	editingName    bool
	err            string
	errUpdateTick  int
	currentErrTick int
}

func NewCreateConfigConfirmationScreen(config *models.UserConfig) *CreateConfigConfirmationScreen {
	c := CreateConfigConfirmationScreen{
		cursor: 0,
		choices: [2]string{
			"[ Go Back ]",
			"[ Create Config ]",
		},
		config:         config,
		errUpdateTick:  5,
		currentErrTick: 0,
	}

	input := textinput.New()
	input.CharLimit = 64
	input.Width = 30
	input.Cursor.Style = FocusedStyle

	input.Placeholder = "Enter Config Name"

	input.PlaceholderStyle = DimBlueStyle
	input.PromptStyle = DimBlueStyle

	input.TextStyle = FocusedStyle
	c.input = input

	return &c
}

func (cc *CreateConfigConfirmationScreen) View() string {
	var ret string

	ret += TitleStyle.Render("Are You Sure You Want to Create This Config?")
	ret += "\n"

	ret += TitleUnderline.Render(cc.config.Data.Name())
	ret += "\n\n"

	/// Render Back First

	if cc.cursor == 0 {
		ret += SelectedStyle.Render(cc.choices[0])
		ret += "\n"
	} else {
		ret += NormalStyle.Render(cc.choices[0])
		ret += "\n"
	}

	ret += "\n"

	if cc.cursor == 1 {
		if !cc.editingName {
			ret += DimBlueStyle.Render("Press Enter To Start Typing")
		} else {
			ret += DimBlueStyle.Render("Press Esc To Stop Typing")
		}
	}
	ret += "\n"
	ret += cc.input.View()

	if cc.cursor == 2 {
		ret += "\n"
		ret += SelectedStyle.Render(cc.choices[1])
		ret += "\n"
	} else {
		ret += "\n"
		ret += NormalStyle.Render(cc.choices[1])
		ret += "\n"
	}

	if cc.err != "" {
		ret += ErrorStyle.Render("[ ")
		ret += ErrorStyle.Render(cc.err)
		ret += ErrorStyle.Render(" ]")
	}

	return ret
}

func (cc *CreateConfigConfirmationScreen) Update(msg tea.Msg, onSetScreen func(screen screen)) {

	if cc.err != "" {
		cc.currentErrTick++
		if cc.currentErrTick >= cc.errUpdateTick {
			cc.err = ""
			cc.currentErrTick = 0
		}
	}

	choices_len := len(cc.choices)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			cc.editingName = false
			cc.input.PlaceholderStyle = DimBlueStyle
			cc.input.PromptStyle = DimBlueStyle
			cc.input.Blur()
		case "enter":
			switch cc.cursor {
			case 0:
				onSetScreen(screenCreateConfig)
			case 1:
				cc.editingName = true
				cc.input.PlaceholderStyle = BitDimBlueStyle
				cc.input.PromptStyle = BlueStyle
				cc.input.Focus()
			case 2:
				name := cc.input.Value()
				if name == "" {
					cc.err = "Config Name Cannot Be Empty"
					cc.currentErrTick = 0
					return
				}
				count := cc.config.GetNConfigCount(name)
				if count > 0 {
					count += 1
					name = name + "(" + strconv.Itoa(count) + ")"
				}
				cc.config.CreateConfig(name)
				//if err != nil {
				//	cc.err = err.Error()
				//} else {
				//	onSetScreen(screenRoot)
				//}
			default:
				break
			}
		case "j":
			if cc.editingName {
				break
			}
			if cc.cursor < choices_len {
				cc.cursor++
			}
		case "k":
			if cc.editingName {
				break
			}
			if cc.cursor > 0 {
				cc.cursor--
			}
		}
	}
	cc.input, _ = cc.input.Update(msg)
}
