package card

import (
	"testing"
)

func TestCard(t *testing.T) {
	cfg := DebitCardConfig{
		CardBase: []int{4, 4, 0, 0, 3, 8},
		CardLen:  16,
	}

	for i := 0; i < 900; i++ {
		cardT, err := CreateCard(cfg)
		if err != nil {
			t.Fatalf("failed to generate card: #%v, with err %v", i, err)
		}
		if !LuhnCheck(cardT) {
			t.Fatalf("card Luhn failed: #%v", i)
		}
	}
}
