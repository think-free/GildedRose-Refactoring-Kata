package items_test

import (
	"context"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/internal/gildedrose/items"
	"github.com/emilybache/gildedrose-refactoring-kata/logging"
	"github.com/stretchr/testify/require"
)

const (
	nameShouldBeUnchanged = "name should be unchanged"
	qualityShouldBe       = "quality should be %d"
	sellInShouldBe        = "sellIn should be %d"
)

func TestAllItems(t *testing.T) {
	logging.Init("panic") // Disable logging

	t.Parallel()

	t.Run("item validity", testItemValidity)

	// Testing all items implementations

	t.Run("default item", testDefaultItemUpdate)
	t.Run("aged brie item", testAgedBrieItemUpdate)
	t.Run("backstage passes item", testBackstagePassesItemUpdate)
	t.Run("sulfuras item", testSulfurasItemUpdate)
	t.Run("conjured item", testConjuredItemUpdate)
}

// Test that the default item is updated correctly
func testDefaultItemUpdate(t *testing.T) {
	ctx := context.Background()

	// Create a new default item
	it := items.New(ctx, "foo", 50, 10)

	// Update the item 100 times
	for i := 100; i > 0; i-- {
		it.Update()
	}

	// Check the final result is correct
	require.Equal(t, "foo", it.GetName(), nameShouldBeUnchanged)
	require.Equal(t, 0, it.GetQuality(), qualityShouldBe, 0)
	require.Equal(t, -90, it.GetSellIn(), sellInShouldBe, -90)
}

// Test that the aged brie item is updated correctly
func testAgedBrieItemUpdate(t *testing.T) {
	ctx := context.Background()

	// Create a new aged brie item
	it := items.New(ctx, "Aged Brie", 10, 10)

	// Update the item 100 times
	for i := 100; i > 0; i-- {
		it.Update()
	}

	// Check the final result is correct
	require.Equal(t, "Aged Brie", it.GetName(), nameShouldBeUnchanged)
	require.Equal(t, 50, it.GetQuality(), qualityShouldBe, 50)
	require.Equal(t, -90, it.GetSellIn(), sellInShouldBe, -90)
}

// Test that backstage passes item is updated correctly
func testBackstagePassesItemUpdate(t *testing.T) {
	ctx := context.Background()

	// Create a new backstage passes item
	it := items.New(ctx, "Backstage passes to a TAFKAL80ETC concert", 10, 30)

	// Update the item 15 times
	for i := 15; i > 0; i-- {
		it.Update()
	}

	// Check the final result is correct
	require.Equal(t, "Backstage passes to a TAFKAL80ETC concert", it.GetName(), nameShouldBeUnchanged)
	require.Equal(t, 25, it.GetQuality(), qualityShouldBe, 25)
	require.Equal(t, 15, it.GetSellIn(), sellInShouldBe, 15)

	// Update the item 85 times
	for i := 85; i > 0; i-- {
		it.Update()
	}

	// Check the final result is correct
	require.Equal(t, "Backstage passes to a TAFKAL80ETC concert", it.GetName(), nameShouldBeUnchanged)
	require.Equal(t, 0, it.GetQuality(), qualityShouldBe, 0)
	require.Equal(t, -70, it.GetSellIn(), sellInShouldBe, -70)
}

// Test that the sulfuras item is updated correctly
func testSulfurasItemUpdate(t *testing.T) {
	ctx := context.Background()

	// Test sulfuras item validity
	i := items.New(ctx, "Sulfuras, Hand of Ragnaros", 50, 50)
	require.Nil(t, i, "item should be nil as it is invalid")

	// Create a new sulfuras item
	it := items.New(ctx, "Sulfuras, Hand of Ragnaros", 80, 60)
	require.NotNil(t, it, "item should not be nil as it is valid")

	// Update the item 100 times
	for i := 100; i > 0; i-- {
		it.Update()
	}

	// Check the final result is correct
	require.Equal(t, "Sulfuras, Hand of Ragnaros", it.GetName(), nameShouldBeUnchanged)
	require.Equal(t, 80, it.GetQuality(), qualityShouldBe, 80)
	require.Equal(t, 60, it.GetSellIn(), sellInShouldBe, 60)
}

// Test that the conjured item is updated correctly
func testConjuredItemUpdate(t *testing.T) {
	ctx := context.Background()

	// Create a new conjured item
	it := items.New(ctx, "Conjured", 10, 60)

	// Update the item 100 times
	for i := 4; i > 0; i-- {
		it.Update()
	}

	// Check the final result is correct
	require.Equal(t, "Conjured", it.GetName(), nameShouldBeUnchanged)
	require.Equal(t, 2, it.GetQuality(), qualityShouldBe, 0)
	require.Equal(t, 56, it.GetSellIn(), sellInShouldBe, 56)

	// Update the item 100 times
	for i := 96; i > 0; i-- {
		it.Update()
	}

	// Check the final result is correct
	require.Equal(t, "Conjured", it.GetName(), nameShouldBeUnchanged)
	require.Equal(t, 0, it.GetQuality(), qualityShouldBe, 0)
	require.Equal(t, -40, it.GetSellIn(), sellInShouldBe, -40)
}

// Test item validity
func testItemValidity(t *testing.T) {
	i := items.New(context.Background(), "foo", 60, 10)
	require.Nil(t, i, "item should be nil as it is invalid (quality > 50)")

	i = items.New(context.Background(), "foo", -10, 10)
	require.Nil(t, i, "item should be nil as it is invalid (quality < 0)")

	i = items.New(context.Background(), "foo", 50, -1)
	require.NotNil(t, i, "item should not be nil as it is valid")
}
