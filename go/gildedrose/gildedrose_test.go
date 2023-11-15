package gildedrose_test

import (
	"testing"

	"gildedrose/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	items := []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"+5 Dexterity Vest", 0, 12},
		{"Aged Brie", 2, 0},
		{"Aged Brie", 3, 50},
		{"Aged Brie", 0, 40},
		{"Aged Brie", -1, 48},
		{"Elixir of the Mongoose", 5, 7},
		{"Elixir of the Mongoose", -1, 2},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 47},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 46},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Conjured Mana Cake", 3, 6},
		{"Conjured Mana Cake", 0, 4},
	}

	gildedrose.UpdateQuality(items)

	expected := []*gildedrose.Item{
		{"+5 Dexterity Vest", 9, 19},
		{"+5 Dexterity Vest", -1, 10},
		{"Aged Brie", 1, 1},
		{"Aged Brie", 2, 50},
		{"Aged Brie", -1, 42},
		{"Aged Brie", -2, 50},
		{"Elixir of the Mongoose", 4, 6},
		{"Elixir of the Mongoose", -2, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Conjured Mana Cake", 2, 4},
		{"Conjured Mana Cake", -1, 0},
	}

	for i, item := range items {
		if item.Name != expected[i].Name {
			t.Errorf("Name %d: Expected %s but got %s ", i, expected[i].Name, item.Name)
		}
		if item.SellIn != expected[i].SellIn {
			t.Errorf("SellIn %d: Expected %d but got %d ", i, expected[i].SellIn, item.SellIn)
		}
		if item.Quality != expected[i].Quality {
			t.Errorf("Quality %d: Expected %d but got %d ", i, expected[i].Quality, item.Quality)
		}
	}
	
}
