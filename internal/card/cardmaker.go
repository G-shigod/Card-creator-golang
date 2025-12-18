package card

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type DebitCardConfig struct {
	CardBase []int
	CardLen  int
}

func CreateCard(cfg DebitCardConfig) ([]int, error) {
	if len(cfg.CardBase) >= cfg.CardLen {
		return nil, fmt.Errorf("base length must be less than card length")
	}
	for i := 0; i < len(cfg.CardBase); i++ {
		if cfg.CardBase[i] > 9 || cfg.CardBase[i] < 0 {
			return nil, fmt.Errorf("card base digit must be 0-9")
		}
	}

	checkDigitIndex := cfg.CardLen - 1
	cardNumber := make([]int, 0, cfg.CardLen)
	cardNumber = append(cardNumber, cfg.CardBase...)

	for i := 0; i < checkDigitIndex-len(cfg.CardBase); i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return nil, err
		}
		cardNumber = append(cardNumber, int(r.Int64()))
	}
	dg := LuhnDigit(cardNumber)
	cardNumber = append(cardNumber, dg)
	return cardNumber, nil

}
