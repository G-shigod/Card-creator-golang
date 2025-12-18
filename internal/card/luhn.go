package card

func LuhnDigit(card []int) int {
	sum := 0
	parity := len(card) % 2 //false

	for i := 0; i < len(card); i++ {
		d := card[i]
		if i%2 != parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	checkDigit := (10 - sum%10) % 10
	return checkDigit
}

func LuhnCheck(card []int) bool {
	mySlice := []int{}
	sumSlc := 0

	for i := len(card) - 1; i >= 0; i-- {
		if i%2 != 0 {
			mySlice = append(mySlice, card[i])
		} else {
			ch := 0
			ch = card[i] * 2
			if ch > 9 {
				ch -= 9
				mySlice = append(mySlice, ch)
			} else {
				mySlice = append(mySlice, ch)
			}
		}
	}
	for i := 0; i < len(mySlice); i++ {
		sumSlc += mySlice[i]
	}
	remainder := sumSlc % 10
	return remainder == 0
}
