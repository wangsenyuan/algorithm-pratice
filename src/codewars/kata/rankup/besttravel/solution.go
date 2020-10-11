package main

import "fmt"

func ChooseBestSum(t, k int, ls []int) int {
	n := len(ls)
	var dfs func(i int, ans int, m int) int

	dfs = func(i int, ans int, m int) int {
		if m == k {
			if ans > t {
				return -1
			}
			return ans
		}

		if i == n {
			return -1
		}
		res := -1

		for j := i; j < n; j++ {
			tmp := dfs(j+1, ans+ls[j], m+1)
			if tmp > res {
				res = tmp
			}
		}

		return res
	}

	return dfs(0, 0, 0)
}

func main() {
	fmt.Println(ChooseBestSum(163, 3, []int{50, 55, 56, 57, 58}))
	fmt.Println(ChooseBestSum(163, 3, []int{50}))
	fmt.Println(ChooseBestSum(230, 3, []int{91, 74, 73, 85, 73, 81, 87}))
	fmt.Println(ChooseBestSum(331, 2, []int{91, 74, 73, 85, 73, 81, 87}))
	fmt.Println(ChooseBestSum(331, 4, []int{91, 74, 73, 85, 73, 81, 87}))
	fmt.Println(ChooseBestSum(331, 5, []int{91, 74, 73, 85, 73, 81, 87}))
}
