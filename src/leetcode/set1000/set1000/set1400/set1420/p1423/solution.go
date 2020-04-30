package p1423

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	left := make([]int, k)

	for i := 0; i < k; i++ {
		left[i] = cardPoints[i]
		if i > 0 {
			left[i] += left[i-1]
		}
	}

	best := left[k-1]

	var sum int

	for i := n - 1; i >= 0; i-- {
		sum += cardPoints[i]

		if sum > best {
			best = sum
		}

		j := k - (n - i)

		if j < 1 {
			break
		}
		if sum+left[j-1] > best {
			best = sum + left[j-1]
		}
	}

	return best
}
