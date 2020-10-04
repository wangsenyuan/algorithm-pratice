package p1496

func canArrange(arr []int, k int) bool {
	n := len(arr)

	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		cnt[(arr[i]%k+k)%k]++
	}

	for x := range cnt {
		y := ((k-x)%k + k) % k

		if x != y && cnt[y] != cnt[x] {
			return false
		} else if x == y && cnt[y]%2 != 0 {
			return false
		}
	}

	return true
}
