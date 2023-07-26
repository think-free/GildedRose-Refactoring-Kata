package items

import (
	"context"

	"github.com/emilybache/gildedrose-refactoring-kata/logging"
)

// Default item is used for non specific cases
type Default struct {
	Item
}

func newDefault(name string, quality, sellIn int) IItem {
	i := &Default{
		Item: Item{
			Name:    name,
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

func (i *Default) Update() {
	i.updateQuality()
	i.updateSellIn()
}

func (i *Default) updateQuality() {
	logging.L(context.Background()).Debugf("updating quality for default item '%s' : %d", i.Name, i.Quality)

	if i.Quality > 0 {
		i.Quality--
	}
	if i.SellIn < 0 && i.Quality > 0 { // quality degrades twice as fast after sellIn date
		i.Quality--
	}
}

func (i *Default) updateSellIn() {
	logging.L(context.Background()).Debugf("updating sellIn for default item '%s'", i.Name)

	i.SellIn--
}
