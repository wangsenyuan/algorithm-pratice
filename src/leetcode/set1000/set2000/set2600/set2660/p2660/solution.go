package p2660

func isWinner(player1 []int, player2 []int) int {

	var a, b int
	a += player1[0]
	b += player2[0]

	for i := 1; i < len(player1); i++ {
		if player1[i-1] == 10 || i > 1 && player1[i-2] == 10 {
			a += 2 * player1[i]
		} else {
			a += player1[i]
		}
		if player2[i-1] == 10 || i > 1 && player2[i-2] == 10 {
			b += 2 * player2[i]
		} else {
			b += player2[i]
		}
	}

	if a > b {
		return 1
	}
	if a < b {
		return 2
	}
	return 0
}
