package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var (
	todos  []string
	list   *widget.List
	stocks []string
	listS  *widget.List
)

func main() {
	a := app.NewWithID("TodoStock")
	w := a.NewWindow("TODO")

	todos = loadTODOs(a.Preferences())
	stocks = loadSTOCKs(a.Preferences())

	// desktop menu
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("TodoStock",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	w.SetContent(loadUI(a))
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	w.Resize(fyne.NewSize(300, 380))
	w.Show()
	a.Run()
}
