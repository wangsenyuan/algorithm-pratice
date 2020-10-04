package p1599

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	var people int
	var best int
	var cnt = -1
	var benefit int
	for i := 0; i < len(customers); i++ {
		people += customers[i]

		x := min(4, people)
		people -= x
		benefit += x * boardingCost
		benefit -= runningCost
		if benefit > best {
			best = benefit
			cnt = i + 1
		}
	}

	if people > 0 {
		x, r := people/4, people%4

		if 4*boardingCost > runningCost {
			benefit += x * (4*boardingCost - runningCost)
			if benefit >= best {
				best = benefit
				cnt = len(customers) + x
				if r > 0 && r*boardingCost > runningCost {
					cnt++
				}
			}
		}
	}

	return cnt
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
