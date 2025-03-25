package p3495

func minOperations(queries [][]int) int64 {
	var res int
	for _, cur := range queries {
		res += solve(cur)
	}
	return int64(res)
}

func solve(query []int) int {
	l, r := query[0], query[1]
	sum := calc(r) - calc(l-1)
	mx := calc(r) - calc(r-1)
	return max(mx, (sum+1)/2)
}

func calc(x int) int {
	var res int
	p := 1
	for i := 1; p <= x; p *= 4 {
		cnt := min(p*4-1, x) - p + 1
		res += cnt * i
		i++
	}
	return res
}
