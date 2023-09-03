package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// For TODO
func loadUI(a fyne.App) fyne.CanvasObject {
	p := a.Preferences()
	list = widget.NewList(
		func() int {
			return len(todos)
		},
		func() fyne.CanvasObject {
			return widget.NewCheck("TODO Item", func(bool) {})
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			check := o.(*widget.Check)
			check.SetChecked(false)
			check.Text = todos[id]
			check.Refresh()
			check.OnChanged = func(done bool) {
				if !done {
					return
				}

				deleteTODO(check.Text, p)
			}
		})

	input := widget.NewEntry()
	input.SetPlaceHolder("New item")
	add := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		addTODO(input.Text, p)
		input.SetText("")
	})
	input.OnSubmitted = func(item string) {
		addTODO(item, p)
		input.SetText("")
	}
	head := container.NewBorder(nil, nil, nil, add, input)

	btn := widget.NewButton("STOCK", func() {
		w2 := a.NewWindow("STOCK")

		w2.SetContent(loadStockUI(a))
		w2.Resize(fyne.NewSize(400, 500))
		w2.Show()
	})

	return container.NewBorder(head, btn, nil, nil, list)
}

// For Stock
func loadStockUI(a fyne.App) fyne.CanvasObject {
	p := a.Preferences()
	listS = widget.NewList(
		func() int {
			return len(stocks)
		},
		func() fyne.CanvasObject {
			return widget.NewCheck("STOCK Item", func(bool) {})
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			check := o.(*widget.Check)
			check.SetChecked(false)
			check.Text = stocks[id]
			check.Refresh()
			check.OnChanged = func(done bool) {
				if !done {
					return
				}

				deleteSTOCK(check.Text, p)
			}
		})

	input := widget.NewEntry()
	input.SetPlaceHolder("New Stock")
	add := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		addSTOCK(input.Text, p)
		input.SetText("")
	})
	input.OnSubmitted = func(stock string) {
		addSTOCK(stock, p)
		input.SetText("")
	}
	head := container.NewBorder(nil, nil, nil, add, input)

	/*
			for Japan ¥Yen
		url, _ := url.Parse("https://keisan.casio.jp/exec/system/1248928927")
	*/

	// for $dollar
	url, _ := url.Parse("https://www.calculator.net/investment-calculator.html?ctype=returnrate&ctargetamountv=&cstartingprinciplev=&cyearsv=&cinterestratev=6&ccompound=annually&ccontributeamountv=0&cadditionat1=end&ciadditionat1=annually&printit=0&x=Calculate#calresult")

	link := widget.NewHyperlink("Calculate interest rate", url)
	return container.NewBorder(head, link, nil, nil, listS)
}
