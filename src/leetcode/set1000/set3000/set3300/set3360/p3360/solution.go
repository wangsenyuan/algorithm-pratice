package p3360

func canAliceWin(n int) bool {
	cnt := 10
	var player int

	for n >= cnt && cnt > 0 {
		n -= cnt
		cnt--
		player ^= 1
	}

	return player == 1
}
