package gildedrose

import(
	"strings"
)

func UpdateQuality(items []*Item) {
	for _, item := range items {
		item.update()
	}
}

type inventoryItem interface {
	update()
}

type Item struct {
	Name            string
	SellIn, Quality int
}

func (i *Item) decrementSellIn() {
	i.SellIn = i.SellIn - 1
}

func (i *Item) changeQuality(qualityChange int) {
	i.Quality += qualityChange
	if i.Quality < 0 {
		i.Quality = 0
	}
	if i.Quality > 50 {
		i.Quality = 50
	}
}

const (
	sulfuras = "Sulfuras, Hand of Ragnaros"
	brie = "Aged Brie"
	pass = "Backstage passes to a TAFKAL80ETC concert"

	conjured = "Conjured"
)

func (i *Item) update() {
	var item inventoryItem
	switch {
	case i.Name == sulfuras:
		item = &legendaryItem{i}
	case i.Name == brie:
		item = &agedBrie{i}
	case i.Name == pass:
		item = &backstagePass{i}
	case strings.HasPrefix(i.Name, conjured):
		item = &conjuredItem{i}
	default:
		item = &regularItem{i}
	}
	item.update()
}

type agedBrie struct {
	*Item
}

func (b *agedBrie) update() {
	if b.SellIn <= 0 {
		b.changeQuality(2)
	} else {
		b.changeQuality(1)
	}
	b.decrementSellIn()
}


type legendaryItem struct {
	*Item
}

func (l *legendaryItem) update() {
	return
}


type backstagePass struct {
	*Item
}

func (b *backstagePass) update() {
	switch {
	case b.SellIn <= 0:
		b.changeQuality(-b.Quality)
	case b.SellIn <= 5:
		b.changeQuality(3)
	case b.SellIn <= 10:
		b.changeQuality(2)
	default:
		b.changeQuality(1)
	}
	b.decrementSellIn()
}


type conjuredItem struct {
	*Item
}

func (c *conjuredItem) update() {
	if c.SellIn <= 0 {
		c.changeQuality(-4)
	} else {
		c.changeQuality(-2)
	}
	c.decrementSellIn()
}


type regularItem struct {
	*Item
}

func (r regularItem) update() {
	if r.SellIn <= 0 {
		r.changeQuality(-2)
	} else {
		r.changeQuality(-1)
	}
	r.decrementSellIn()
}
