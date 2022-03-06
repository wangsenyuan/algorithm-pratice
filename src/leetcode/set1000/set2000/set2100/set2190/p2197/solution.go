package p2197

func replaceNonCoprimes(nums []int) []int {
	n := len(nums)

	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		cur := nums[i]
		for p > 0 {
			g := gcd(cur, stack[p-1])
			if g == 1 {
				break
			}
			cur = cur / g * stack[p-1]
			p--
		}
		stack[p] = cur
		p++
	}
	return stack[:p]
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
