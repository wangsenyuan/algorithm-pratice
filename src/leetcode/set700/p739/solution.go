package p739

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := make([]int, n)
	p := 0

	for i := 0; i < n; i++ {
		for p > 0 && temperatures[stack[p-1]] < temperatures[i] {
			res[stack[p-1]] = i - stack[p-1]
			p--
		}
		stack[p] = i
		p++
	}

	return res
}
