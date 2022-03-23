package p2212

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	//bob获得最高分的方式
	tot := 1 << 12

	var best_score int
	var best_state int
	bob := make([]int, 12)
	var sum int

	for state := 1; state < tot; state++ {
		// bob needs to win state
		var score int
		for i := 0; i < 12; i++ {
			if (state>>i)&1 == 1 {
				bob[i] = aliceArrows[i] + 1
				score += i
			} else {
				bob[i] = 0
			}
		}

		sum = 0
		for i := 0; i < 12; i++ {
			sum += bob[i]
		}

		if sum > numArrows {
			// not a valid option
			continue
		}
		if best_score < score {
			best_score = score
			best_state = state
		}
	}
	for i := 0; i < 12; i++ {
		bob[i] = 0
	}
	if best_score == 0 {
		// no win at all
		bob[0] = numArrows
		return bob
	}
	sum = 0
	for i := 0; i < 12; i++ {
		if (best_state>>i)&1 == 1 {
			bob[i] = aliceArrows[i] + 1
			sum += bob[i]
		}
	}
	bob[0] += numArrows - sum
	return bob
}
