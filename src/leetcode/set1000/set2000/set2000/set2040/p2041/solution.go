package p2041

func areNumbersAscending(s string) bool {
	prev := -1

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			var num int
			for i < len(s) && s[i] != ' ' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			// s[i] is ' '
			if num <= prev {
				return false
			}
			prev = num
		}
	}
	return true
}
