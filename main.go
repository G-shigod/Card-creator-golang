package main

import (
	"fmt"
	"os"

	"github.com/G-shigod/Card-creator-golang/internal/card"
)

func main() {
	type CardConfig struct {
		CardBase []int
		CardLen  int
	}
	abc := CardConfig{
		CardBase: []int{4, 4, 0, 0, 2, 9},
		CardLen:  16,
	}

	NewCard, err := card.CreateCard(card.DebitCardConfig(abc))
	if err != nil {
		fmt.Println("failed to create card")
		os.Exit(1)
	}
	fmt.Println(NewCard, len(NewCard))
	checker := card.LuhnCheck(NewCard)
	fmt.Println(checker)
}
