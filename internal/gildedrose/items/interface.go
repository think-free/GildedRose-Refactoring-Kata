package items

import (
	"context"

	"github.com/emilybache/gildedrose-refactoring-kata/logging"
)

type IItem interface {
	Update()

	GetName() string
	GetSellIn() int
	GetQuality() int
}

// Map all the specific items strings to their respective constructors
var specificUpdates = map[string]func(quality, sellIn int) IItem{
	ItemAgedBrie:        newAgedBrie,
	ItemBackstagePasses: newBackstagePasses,
	ItemSulfuras:        newSulfuras,
	ItemConjured:        newConjured,
}

// New returns a new item based on the name
func New(ctx context.Context, name string, quality, sellIn int) IItem {
	var it IItem
	if newSpecific, ok := specificUpdates[name]; ok {
		logging.L(ctx).Infof("creating specific item '%s'", name)
		it = newSpecific(quality, sellIn)
	} else {
		logging.L(ctx).Infof("creating default item '%s'", name)
		it = newDefault(name, quality, sellIn)
	}

	return it
}
