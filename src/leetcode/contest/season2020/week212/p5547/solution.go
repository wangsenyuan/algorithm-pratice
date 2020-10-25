package p5547

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n := len(nums)

	m := len(l)
	ans := make([]bool, m)
	tmp := make([]int, n)
	for i := 0; i < m; i++ {
		a, b := l[i], r[i]+1
		if b-a == 1 {
			ans[i] = false
			continue
		}
		copy(tmp, nums)
		tmp2 := tmp[a:b]
		sort.Ints(tmp2)
		ans[i] = true

		for j := 1; j < len(tmp2); j++ {
			if tmp2[j]-tmp2[j-1] != tmp2[1]-tmp2[0] {
				ans[i] = false
				break
			}
		}
	}
	return ans
}
