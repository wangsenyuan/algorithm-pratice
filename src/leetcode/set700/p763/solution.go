package p763

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return []int{}
	}
	n := len(S)
	right := make([]int, 26)
	for i := 0; i < 26; i++ {
		right[i] = n
	}
	for i := 0; i < n; i++ {
		x := int(S[i] - 'a')
		right[x] = i
	}

	res := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		x := int(S[i] - 'a')
		j := right[x]
		for k := i + 1; k < j; k++ {
			y := int(S[k] - 'a')
			if right[y] > j {
				j = right[y]
			}
		}
		res = append(res, j-i+1)
		i = j
	}
	return res
}
