package gildedrose_test

import (
	"context"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/internal/gildedrose/items"
	"github.com/emilybache/gildedrose-refactoring-kata/logging"
	"github.com/stretchr/testify/require"
)

type Scenario struct {
	name            string
	quality         int
	sellIn          int
	expectedResults []ExpectedResults
}

type ExpectedResults struct {
	quality int
	sellIn  int
}

func TestGildedrose(t *testing.T) {
	ctx := context.Background()
	logging.Init("panic") // Disable logging

	scenarios := []Scenario{
		{
			name:    "normal",
			quality: 40,
			sellIn:  9,
			expectedResults: []ExpectedResults{
				{39, 8}, {38, 7}, {37, 6}, {36, 5}, {35, 4}, {34, 3}, {33, 2}, {32, 1}, {31, 0}, {30, -1}, {28, -2}, {26, -3}, {24, -4}, {22, -5}, {20, -6}, {18, -7}, {16, -8}, {14, -9}, {12, -10},
			},
		},
		{
			name:    "Aged Brie",
			quality: 40,
			sellIn:  9,
			expectedResults: []ExpectedResults{
				{41, 8}, {42, 7}, {43, 6}, {44, 5}, {45, 4}, {46, 3}, {47, 2}, {48, 1}, {49, 0}, {50, -1}, {50, -2}, {50, -3}, {50, -4}, {50, -5}, {50, -6}, {50, -7}, {50, -8}, {50, -9}, {50, -10},
			},
		},
		{
			name:    "Backstage passes to a TAFKAL80ETC concert",
			quality: 5,
			sellIn:  15,
			expectedResults: []ExpectedResults{
				{6, 14}, {7, 13}, {8, 12}, {9, 11}, {10, 10}, {12, 9}, {14, 8}, {16, 7}, {18, 6}, {20, 5}, {23, 4}, {26, 3}, {29, 2}, {32, 1}, {35, 0}, {0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5},
			},
		},
		{
			name:    "Conjured",
			quality: 40,
			sellIn:  3,
			expectedResults: []ExpectedResults{
				{38, 2}, {36, 1}, {34, 0}, {32, -1}, {28, -2}, {24, -3}, {20, -4}, {16, -5}, {12, -6}, {8, -7}, {4, -8}, {0, -9}, {0, -10}, {0, -11}, {0, -12}, {0, -13}, {0, -14}, {0, -15}, {0, -16}, {0, -17},
			},
		},
		{
			name:    "Sulfuras, Hand of Ragnaros",
			quality: 80,
			sellIn:  3,
			expectedResults: []ExpectedResults{
				{80, 3}, {80, 3}, {80, 3},
			},
		},
	}

	for i := range scenarios {
		s := scenarios[i]
		t.Run(s.name, func(t *testing.T) {
			t.Parallel()
			item := items.New(ctx, s.name, s.quality, s.sellIn)
			for day := 0; day < len(s.expectedResults); day++ {
				item.Update()
				require.Equal(t, s.expectedResults[day].quality, item.GetQuality(), "quality should be %d for day %d", s.expectedResults[day].quality, day)
				require.Equal(t, s.expectedResults[day].sellIn, item.GetSellIn(), "sellIn should be %d for day %d", s.expectedResults[day].sellIn, day)
			}
		})
	}
}
