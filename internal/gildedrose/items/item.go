package items

// Item is the base struct for all items
// It is embedded in all other Item structs and not be used directly
// If you are looking for the default Item struct, see item_default.go
type Item struct { // Should be 'item' but goblin would complains if we touch it
	Name            string
	SellIn, Quality int
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) GetQuality() int {
	return i.Quality
}

func (i *Item) GetSellIn() int {
	return i.SellIn
}

func IsValid(it IItem) bool {
	return it.GetQuality() >= 0 && it.GetQuality() <= 50
}
