package items

import (
	"context"

	"github.com/emilybache/gildedrose-refactoring-kata/logging"
)

const ItemBackstagePasses = "Backstage passes to a TAFKAL80ETC concert"

type BackstagePasses struct {
	Item
}

func newBackstagePasses(quality, sellIn int) IItem {
	i := &BackstagePasses{
		Item: Item{
			Name:    ItemBackstagePasses,
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

func (i *BackstagePasses) Update() {
	i.updateQuality()
	i.updateSellIn()
}

func (i *BackstagePasses) updateQuality() {
	logging.L(context.Background()).Debugf("updating quality for backstage passes item '%s", i.Name)

	if i.SellIn > 0 && i.Quality < 50 {
		i.Quality++

		if i.SellIn < 6 {
			i.Quality = i.Quality + 2 // quality increases by 3 when there are 5 days or less
		} else if i.SellIn < 11 {
			i.Quality = i.Quality + 1 // quality increases by 2 when there are 10 days or less
		}
	}

	if i.Quality > 50 {
		i.Quality = 50
	}

	if i.SellIn <= 0 {
		i.Quality = 0
	}
}

func (i *BackstagePasses) updateSellIn() {
	logging.L(context.Background()).Debugf("updating sellIn for backstage passes item '%s", i.Name)

	i.SellIn--
}
