package prototype

import (
	"errors"
	"fmt"
)

// ShirtCloner TODO
type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

// TODO
const (
	White = 1
	Black = 2
	Blue  = 3
)

// ShirtCache TODO
type ShirtCache struct{}

// GetClone returns cached shirts sample by cloning them
func (s *ShirtCache) GetClone(d int) (ItemInfoGetter, error) {
	switch d {
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not reconized")
	}

}

// ItemInfoGetter TODO
type ItemInfoGetter interface {
	GetInfo() string
}

// ShirtColor Describes shirts color
type ShirtColor byte

// Shirt TODO
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// GetInfo TODO
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and color id %d that cost '%f'", s.SKU, s.Color, s.Price)
}

// GetPrice TODO
func (s *Shirt) GetPrice() float32 {
	return s.Price
}

// GetShirtsCloner TODO
func GetShirtsCloner() ShirtCloner {
	return new(ShirtCache)
}

var whitePrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 16.00,
	SKU:   "Empty",
	Color: Black,
}

var bluePrototype = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}
