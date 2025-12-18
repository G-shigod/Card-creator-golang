package card

import (
	"testing"
)

func TestCard(t *testing.T) {
	for i := 0; i < 30; i++ {
		cardT, err := CreateCard()
		if err != nil {
			t.Fatalf("failed to generate card: #%v, with err %v", i, err)
		}
		if !LuhnCheck(cardT) {
			t.Fatalf("card Luhn failed: #%v", i)
		}
	}
}
