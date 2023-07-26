package gildedrose

import "github.com/emilybache/gildedrose-refactoring-kata/internal/gildedrose/items"

func UpdateQuality(items []items.IItem) {
	for i := range items {
		items[i].Update()
	}
}
