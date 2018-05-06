package p830

func largeGroupPositions(S string) [][]int {
	ans := make([][]int, 0, len(S)/3)

	for i, j := 1, 0; i <= len(S); i++ {
		if i < len(S) && S[i] == S[i-1] {
			continue
		}
		if i-j >= 3 {
			ans = append(ans, []int{j, i - 1})
		}
		j = i
	}
	return ans
}
