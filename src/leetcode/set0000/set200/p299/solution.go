package p299

import "fmt"

func getHint(secret string, guess string) string {
	var bulls int
	a := make([]int, 10)
	b := make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			a[int(secret[i]-'0')]++
			b[int(guess[i]-'0')]++
		}
	}
	var cows int
	for i := 0; i < 10; i++ {
		cows += min(a[i], b[i])
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
