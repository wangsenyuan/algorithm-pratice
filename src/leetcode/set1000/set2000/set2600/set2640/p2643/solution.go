package p2643

func rowAndMaximumOnes(mat [][]int) []int {

	res := []int{0, 0}

	for i := 0; i < len(mat); i++ {
		cnt := countOnes(mat[i])
		if cnt > res[1] {
			res[0] = i
			res[1] = cnt
		}
	}
	return res
}

func countOnes(nums []int) int {
	var res int
	for _, num := range nums {
		res += num
	}
	return res
}
