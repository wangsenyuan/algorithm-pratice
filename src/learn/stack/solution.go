package stack

func isPopOrder(push []int, pop []int) bool {
	n := len(push)
	stack := make([]int, n)
	var p, j int
	for i := 0; i < n; i++ {
		for p > 0 && j < n && stack[p-1] == pop[j] {
			p--
			j++
		}
		stack[p] = push[i]
		p++
	}
	for p > 0 && j < n && stack[p-1] == pop[j] {
		p--
		j++
	}

	return p == 0 && j == n
}
