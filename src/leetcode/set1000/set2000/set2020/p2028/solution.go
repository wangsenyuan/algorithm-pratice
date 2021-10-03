package p2028

func stoneGameIX(stones []int) bool {
	n := len(stones)
	if n == 1 {
		return false
	}

	// n <= 10**5
	// stones[i] <= 10**4
	cnt := make([]int, 3)
	for i := 0; i < n; i++ {
		cnt[stones[i]%3]++
	}

	cnt2 := []int{cnt[0], cnt[2], cnt[1]}

	return check(cnt) || check(cnt2)
}

func check(arr []int) bool {
	if arr[1] == 0 {
		return false
	}
	arr[1]--
	turn := 1 + min(arr[1], arr[2])*2 + arr[0]
	if arr[1] > arr[2] {
		turn++
		arr[1]--
	}
	return turn%2 == 1 && arr[1] != arr[2]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
