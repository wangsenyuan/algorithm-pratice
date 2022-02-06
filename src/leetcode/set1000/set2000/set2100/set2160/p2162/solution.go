package p2162

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	// ts = a * 60 + b
	// where a >= 0 && a <= 99, b >= 0 && b <= 99
	best := 1 << 30
	for a := 0; a <= 99; a++ {
		if a*60 > targetSeconds {
			break
		}
		b := targetSeconds - a*60
		if b > 99 {
			continue
		}
		best = min(best, calc(a, b, startAt, moveCost, pushCost))
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func calc(a int, b int, start int, move int, push int) int {
	var res int
	if a > 0 {
		if a >= 10 {
			res += moveAndPush(start, a/10, move, push)
			start = a / 10
		}
		res += moveAndPush(start, a%10, move, push)
		start = a % 10
		res += moveAndPush(start, b/10, move, push)
		start = b / 10
		b = b % 10
	}
	if b >= 10 {
		res += moveAndPush(start, b/10, move, push)
		start = b / 10
	}
	res += moveAndPush(start, b%10, move, push)
	return res
}

func moveAndPush(x, y int, move int, push int) int {
	if x == y {
		return push
	}
	return move + push
}
