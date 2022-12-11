package p2496

func maxJump(stones []int) int {
	n := len(stones)
	if n == 2 {
		return stones[1] - stones[0]
	}
	// n <= 1e5
	// 假设青蛙落脚在i点，结果最小肯定是 stones[i+1] - stones[i]
	// 为了得到最小值，青蛙所有的点都需要跳到
	var res int
	i := 2
	for i < n {
		res = max(res, stones[i]-stones[i-2])
		i += 2
	}
	j := n - 2
	if i == n {
		res = max(res, stones[n-1]-stones[n-2])
		j = n - 3
	} else {
		res = max(res, stones[n-1]-stones[n-2])
		j -= 2
	}

	for j > 0 {
		res = max(res, stones[j+2]-stones[j])
		j -= 2
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
