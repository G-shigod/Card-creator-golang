package main

import (
	"fmt"
	"os"

	"github.com/G-shigod/banking_emulator/internal/card"
)

func main() {
	NewCard, err := card.CreateCard()
	if err != nil {
		fmt.Println("failed to create card")
		os.Exit(1)
	}

	checker := card.LuhnCheck(NewCard)
	fmt.Println(checker)
}
