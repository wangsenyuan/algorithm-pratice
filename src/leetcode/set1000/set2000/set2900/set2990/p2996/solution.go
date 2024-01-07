package p2996

func minimumOperationsToMakeEqual(x int, y int) int {
	mem := make(map[int]int)

	var dfs func(y int) int

	dfs = func(x int) int {
		if x <= y {
			return y - x
		}
		if v, ok := mem[x]; ok {
			return v
		}
		tmp := x - y
		tmp = min(tmp, 1+dfs(x/5)+x%5)
		// (x + z) % 5 = 0
		if x%5 != 0 {
			z := x/5*5 + 5
			tmp = min(tmp, z-x+1+dfs(z/5))
		}
		tmp = min(tmp, 1+dfs(x/11)+x%11)
		if x%11 != 0 {
			z := x/11*11 + 11
			tmp = min(tmp, z-x+1+dfs(z/11))
		}
		mem[x] = tmp
		return tmp
	}

	return dfs(x)
}
