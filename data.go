package main

import "fyne.io/fyne/v2"

// For TODO
func addTODO(todo string, p fyne.Preferences) {
	todos = append(todos, todo)
	list.Refresh()
	saveTODOs(todos, p)
}

func deleteTODO(todo string, p fyne.Preferences) {
	for i, text := range todos {
		if text != todo {
			continue
		}

		todos = append(todos[:i], todos[i+1:]...)
		break
	}
	list.Refresh()
	saveTODOs(todos, p)
}

// For Stock
func addSTOCK(stock string, p fyne.Preferences) {
	stocks = append(stocks, stock)
	listS.Refresh()
	saveSTOCKs(stocks, p)
}

func deleteSTOCK(stock string, p fyne.Preferences) {
	for i, text := range stocks {
		if text != stock {
			continue
		}

		stocks = append(stocks[:i], stocks[i+1:]...)
		break
	}
	listS.Refresh()
	saveSTOCKs(stocks, p)
}