package main

import (
	"strings"

	"fyne.io/fyne/v2"
)

const joiner = "|"

// For TODO
func loadTODOs(p fyne.Preferences) []string {
	// StringWithFallback looks up a string value and returns the given fallback if not found
	all := p.StringWithFallback("items", "Do this item"+joiner+"Learn Fyne!"+joiner+"Build an app")
	if all == "" {
		return []string{}
	}
	return strings.Split(all, joiner)
}

func saveTODOs(items []string, p fyne.Preferences) {
	allItems := strings.Join(items, joiner)
	p.SetString("items", allItems)
}

// For Stock
func loadSTOCKs(p fyne.Preferences) []string {
	all := p.StringWithFallback("stocks", "Fiscal-period Listed-issue Return-rate span"+joiner+"Sep AAPL ~~% 1.2023--9.2025"+joiner+"Dec GOOG ~~% 3.2020--12.2026")
	if all == "" {
		return []string{}
	}
	return strings.Split(all, joiner)
}

func saveSTOCKs(stocks []string, p fyne.Preferences) {
	allItems := strings.Join(stocks, joiner)
	p.SetString("stocks", allItems)
}
