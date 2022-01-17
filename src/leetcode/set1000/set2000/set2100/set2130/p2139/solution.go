package p2139

func minMoves(target int, maxDoubles int) int {
	// 越后面double，越好
	if target == 1 {
		return 0
	}
	if maxDoubles == 0 {
		return target - 1
	}

	if target%2 == 0 {
		return 1 + minMoves(target/2, maxDoubles-1)
	}
	return 1 + minMoves(target-1, maxDoubles)
}
