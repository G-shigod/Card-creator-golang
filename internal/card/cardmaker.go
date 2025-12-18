package card

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const cardLength = 16
const checkDigitIndex = cardLength - 1

func CreateCard() ([]int, error) {
	cardNumber := make([]int, 0, cardLength)
	cardBody := []int{4, 1, 9, 4, 2}
	cardNumber = append(cardNumber, cardBody...)

	for i := 0; i < checkDigitIndex-len(cardBody); i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			fmt.Println("failed to generate random digits")
			return nil, err
		}
		cardNumber = append(cardNumber, int(r.Int64()))
	}
	dg := LuhnDigit(cardNumber)
	cardNumber = append(cardNumber, dg)

	fmt.Println(len(cardNumber), cardNumber)
	return cardNumber, nil

}
