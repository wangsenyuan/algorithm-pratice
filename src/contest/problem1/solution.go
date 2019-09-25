package problem1

func game(guess []int, answer []int) int {
	var ans int

	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			ans++
		}
	}
	return ans
}
