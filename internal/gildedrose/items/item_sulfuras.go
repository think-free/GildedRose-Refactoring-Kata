package items

import (
	"context"

	"github.com/emilybache/gildedrose-refactoring-kata/logging"
)

const ItemSulfuras = "Sulfuras, Hand of Ragnaros"

type Sulfuras struct {
	Item
}

func newSulfuras(quality, sellIn int) IItem {
	if quality != 80 { // legendary item, quality is always 80
		logging.L(context.Background()).Errorf("invalid item '%s' : quality -> %d, sellin -> %d, quality should always be '80'", ItemSulfuras, quality, sellIn)
		return nil
	}
	i := &Sulfuras{
		Item: Item{
			Name:    ItemSulfuras,
			Quality: quality,
			SellIn:  sellIn,
		},
	}

	return i
}

func (i *Sulfuras) Update() {
	i.updateQuality()
	i.updateSellIn()
}

func (i *Sulfuras) updateQuality() {
	logging.L(context.Background()).Debugf("updating quality for sulfuras item '%s", i.Name)
}

func (i *Sulfuras) updateSellIn() {
	logging.L(context.Background()).Debugf("updating sellIn for sulfuras item '%s", i.Name)
}
