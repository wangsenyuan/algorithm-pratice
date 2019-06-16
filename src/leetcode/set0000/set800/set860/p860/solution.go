package p860

func lemonadeChange(bills []int) bool {
	var five, ten, twenty int

	for i := 0; i < len(bills); i++ {
		bill := bills[i]
		if bill == 5 {
			five++
		} else if bill == 10 {
			// change five
			if five <= 0 {
				return false
			}
			five--
			ten++
		} else {
			// 20
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
			twenty++
		}
	}
	return true
}
