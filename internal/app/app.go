package app

import (
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
	screenCreate
	screenSelect
)

type model struct {
	screen screen

	root root
}

func InitialModel() model {
	return model{
		screen: screenRoot,
		/// Creating Root Screen
		root: root {
			cursor: 0,
			choices: [3]string{
				"Create New Configuration",
				"Select Configuration",
				"Open Config Folder",
			},
		},
		/// Other Screens Here
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
		m.root.Update(msg)
	default: break
	}


	return m, nil
}

func (m model) View() string {
	switch m.screen {
		case screenRoot:
			return appStyle.Render(m.root.View())
		default:
			return "Unkown\n"
	}
}

// func createNewConfigScreen(pages *tview.Pages) tview.Primitive {
//
// 	backButton := tview.NewButton("Back").SetSelectedFunc(func() {
// 		pages.SwitchToPage("root")
// 	})
//
// 	flex := tview.NewFlex().SetDirection(tview.FlexRow).
// 		AddItem(backButton, 1, 0, true)
//
// 	frame := tview.NewFrame(flex).
// 		SetBorder(true).
// 		SetTitle("Create New Configuration")
//
// 	return frame
// }
//
// func selectNewConfigurationScreen(pages *tview.Pages) tview.Primitive {
//
// 	backButton := tview.NewButton("Back").SetSelectedFunc(func() {
// 		pages.SwitchToPage("root")
// 	})
//
// 	frame := tview.NewFrame(backButton).
// 		SetBorder(true).
// 		SetTitle("Select Configuration")
//
// 	flex := tview.NewFlex().SetDirection(tview.FlexRow).
// 		AddItem(backButton, 1, 0, true).
// 		AddItem(frame, 3, 0, false)
// 	return flex
// }
//
// func rootView(pages *tview.Pages) tview.Primitive {
//
// 	list := tview.NewList().
// 		AddItem(ConfigFlipperOptions[0], "", 'a', func() {
// 			pages.SwitchToPage("createNewConfigScreen")
// 		}).
// 		AddItem(ConfigFlipperOptions[1], "Select a new configuration, or create a new copy", 'b', func() {
// 			pages.SwitchToPage("selectNewConfigurationScreen")
// 		}).
// 		AddItem(ConfigFlipperOptions[2], "", 'c', func() {})
//
// 	flex := tview.NewFlex().SetDirection(tview.FlexRow).
// 		AddItem(list, 0, 1, true)
//
// 	return flex
// }
//
// func Init() error {
// 	app := tview.NewApplication()
// 	pages := tview.NewPages()
//
// 	root := rootView(pages)
// 	createNewConfigScreen := createNewConfigScreen(pages)
// 	selectNewConfigurationScreen := selectNewConfigurationScreen(pages)
//
// 	pages.AddPage("root", root, true, true)
// 	pages.AddPage("selectNewConfigurationScreen", selectNewConfigurationScreen, true, false)
// 	pages.AddPage("createNewConfigScreen", createNewConfigScreen, true, false)
//
// 	return app.SetRoot(pages, true).Run()
// }
