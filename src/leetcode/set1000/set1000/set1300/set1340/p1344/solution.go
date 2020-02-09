package p1344

func checkIfExist(arr []int) bool {
	mem := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		x := arr[i]
		if x&1 == 0 && mem[x>>1] {
			return true
		}
		if mem[2*x] {
			return true
		}
		mem[x] = true
	}

	return false
}

func minSteps(s string, t string) int {
	cnt := make([]int, 26)

	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		cnt[x]++
	}

	// cnt2 := make([]int, 26)

	for i := 0; i < len(t); i++ {
		x := int(t[i] - 'a')
		cnt[x]--
	}
	var add, remove int
	for i := 0; i < 26; i++ {
		if cnt[i] == 0 {
			continue
		}
		if cnt[i] < 0 {
			//t has more i
			remove -= cnt[i]
		} else {
			add += cnt[i]
		}
	}

	return abs(remove-add) + min(remove, add)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
