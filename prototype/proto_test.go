package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received Coloner was nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("Item cannot be equal to WhitePrototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type asertion for shirt1 cann't be done successfully")
	}
	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type asertion for shirt1 cann't be done successfully")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU for shirt 1 should be different from shirt2 SKU")
	}

	if shirt1 == shirt2 {
		t.Error("Shirt1 should be defference from shirt2")
	}

	var whiteShirtPrice float32 = 15.00
	if shirt1.GetPrice() == whiteShirtPrice {
		t.Errorf("Price of white shirt shuld be 15.00 and %f", whiteShirtPrice)
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())
	t.Logf("LOG: The memory positions of the shirts are different %p != %p\n\n", &shirt1, &shirt2)

}
