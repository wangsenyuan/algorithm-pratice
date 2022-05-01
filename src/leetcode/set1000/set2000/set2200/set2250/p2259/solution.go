package p2259

func removeDigit(number string, digit byte) string {

	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			j := i + 1
			for j < len(number) && number[j] == number[i] {
				j++
			}
			if j == len(number) || number[j] > number[i] {
				return number[:i] + number[i+1:]
			}
		}
	}

	for i := len(number) - 1; i >= 0; i-- {
		if number[i] == digit {
			return number[:i] + number[i+1:]
		}
	}
	return number
}
