package items

import (
	"context"
	"log"

	"github.com/emilybache/gildedrose-refactoring-kata/logging"
)

const ItemConjured = "Conjured"

type Conjured struct {
	Item
}

func newConjured(quality, sellIn int) IItem {
	i := &Conjured{
		Item: Item{
			Name:    ItemConjured,
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

func (i *Conjured) Update() {
	i.updateQuality()
	i.updateSellIn()
}

func (i *Conjured) updateQuality() {
	i.Quality = i.Quality - 2

	if i.SellIn < 0 {
		i.Quality = i.Quality - 2
	}

	if i.Quality < 0 {
		i.Quality = 0
	}
}

func (i *Conjured) updateSellIn() {
	log.Printf("updating sellIn for conjured item '%s", i.Name)

	i.SellIn--
}
