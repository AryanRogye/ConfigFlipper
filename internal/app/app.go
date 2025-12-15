package app

import(
	"github.com/rivo/tview"
)

var ConfigFlipperOptions = [3]string{
	"Create New Configuration",
	"Select Configuration",
	"Open Config Folder",
}

func createNewConfigScreen(pages *tview.Pages) tview.Primitive {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Create New Configuration")
	backButton := tview.NewButton("Back").SetSelectedFunc(func() {
		pages.SwitchToPage("root")
	})

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(backButton, 1, 0, true).
		AddItem(box, 3, 0, false)

	return flex
}

func selectNewConfigurationScreen(pages *tview.Pages) tview.Primitive {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Select Configuration")

	backButton := tview.NewButton("Back").SetSelectedFunc(func() {
		pages.SwitchToPage("root")
	})

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(backButton, 1, 0, true).
		AddItem(box, 3, 0, false)
	return flex
}

func rootView(pages *tview.Pages) tview.Primitive {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("ConfigFlipper")

	list := tview.NewList().
		AddItem(ConfigFlipperOptions[0], "", 'a', func() {
			pages.SwitchToPage("createNewConfigScreen")
		}).
		AddItem(ConfigFlipperOptions[1], "Select a new configuration, or create a new copy", 'b', func() {
			pages.SwitchToPage("selectNewConfigurationScreen")
		}).
		AddItem(ConfigFlipperOptions[2], "", 'c', func() {})

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(box, 3, 0, false).
		AddItem(list, 6, 0, true)

	return flex
}

func Init() error {
	app := tview.NewApplication()
	pages := tview.NewPages()

	root := rootView(pages)
	createNewConfigScreen := createNewConfigScreen(pages)
	selectNewConfigurationScreen := selectNewConfigurationScreen(pages)

	pages.AddPage("root", root, true, true)
	pages.AddPage("selectNewConfigurationScreen", selectNewConfigurationScreen, true, false)
	pages.AddPage("createNewConfigScreen", createNewConfigScreen, true, false)

	return app.SetRoot(pages, true).Run()
}
