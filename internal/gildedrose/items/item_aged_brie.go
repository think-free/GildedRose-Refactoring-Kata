package items

import (
	"context"

	"github.com/emilybache/gildedrose-refactoring-kata/logging"
)

const ItemAgedBrie = "Aged Brie"

type AgedBrie struct {
	Item
}

func newAgedBrie(quality, sellIn int) IItem {
	i := &AgedBrie{
		Item: Item{
			Name:    ItemAgedBrie,
			Quality: quality,
			SellIn:  sellIn,
		},
	}

	if !IsValid(i) {
		logging.L(context.Background()).Errorf("invalid item '%s' : quality -> %d, sellin -> %d", i.Name, quality, sellIn)
		return nil
	}

	return i
}

func (i *AgedBrie) Update() {
	i.updateQuality()
	i.updateSellIn()
}

func (i *AgedBrie) updateQuality() {
	logging.L(context.Background()).Debugf("updating quality for aged brie item '%s' to %d", i.Name, i.Quality)

	if i.Quality < 50 {
		i.Quality++
	}
}

func (i *AgedBrie) updateSellIn() {
	logging.L(context.Background()).Debugf("updating sellIn for aged brie item '%s' to %d", i.Name, i.SellIn)

	i.SellIn--
}
