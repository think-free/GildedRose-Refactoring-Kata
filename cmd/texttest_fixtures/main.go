package main

import (
	"context"

	"github.com/emilybache/gildedrose-refactoring-kata/internal/cliparams"
	"github.com/emilybache/gildedrose-refactoring-kata/internal/gildedrose"
	"github.com/emilybache/gildedrose-refactoring-kata/internal/gildedrose/items"

	"github.com/emilybache/gildedrose-refactoring-kata/logging"
)

func main() {
	ctx := context.Background()
	cp := cliparams.New()

	logging.Init(cp.LogLevel)

	var items = []items.IItem{
		items.New(ctx, "+5 Dexterity Vest", 20, 10),
		items.New(ctx, "Aged Brie", 0, 2),
		items.New(ctx, "Elixir of the Mongoose", 7, 5),
		items.New(ctx, "Sulfuras, Hand of Ragnaros", 80, 0),
		items.New(ctx, "Sulfuras, Hand of Ragnaros", 80, -1),
		items.New(ctx, "Backstage passes to a TAFKAL80ETC concert", 20, 15),
		items.New(ctx, "Backstage passes to a TAFKAL80ETC concert", 49, 10),
		items.New(ctx, "Backstage passes to a TAFKAL80ETC concert", 49, 5),
		items.New(ctx, "Conjured Mana Cake", 6, 3),
	}

	for i := 0; i < len(items); i++ {
		if items[i] == nil {
			logging.L(ctx).Fatal("there was an error creating the items")
		}
	}

	for day := 1; day < cp.Days; day++ {
		for i := 0; i < len(items); i++ {
			logging.L(ctx).Infof("day %d -> name : %s, sellIn : %d, quality : %d", day, items[i].GetName(), items[i].GetSellIn(), items[i].GetQuality())
		}
		gildedrose.UpdateQuality(items)
	}
}
