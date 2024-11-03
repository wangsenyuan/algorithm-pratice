package p3340

func isBalanced(num string) bool {
	sum := make([]int, 2)
	for i, x := range num {
		v := int(x - '0')
		sum[i&1] += v
	}
	return sum[0] == sum[1]
}
