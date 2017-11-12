package main

import "fmt"

func main() {
	flights := [][]int{
		{0, 1, 1}, {1, 0, 1}, {1, 1, 0},
	}

	days := [][]int{
		{1, 3, 1}, {6, 0, 3}, {3, 3, 3},
	}

	/*flights := [][]int{
		{0, 0, 0}, {0, 0, 0}, {0, 0, 0},
	}

	days := [][]int{
		{1, 1, 1}, {7, 7, 7}, {7, 7, 7},
	}*/

	/*flights := [][]int{
		{0, 1, 1}, {1, 0, 1}, {1, 1, 0},
	}

	days := [][]int{
		{7, 0, 0}, {0, 7, 0}, {0, 0, 7},
	}*/

	fmt.Println(maxVacationDays(flights, days))
}

func maxVacationDays(flights [][]int, days [][]int) int {
	n := len(days)
	if n == 0 {
		return 0
	}
	k := len(days[0])
	if k == 0 {
		return 0
	}

	memo := make([][]int, n)

	for i := 0; i < n; i++ {
		memo[i] = make([]int, k)
		for j := 0; j < k; j++ {
			memo[i][j] = -1
		}
	}

	var visit func(city int, week int) int

	visit = func(city int, week int) int {
		if week == k {
			return 0
		}

		if memo[city][week] >= 0 {
			return memo[city][week]
		}

		res := 0
		for i := 0; i < n; i++ {
			if i == city || flights[city][i] == 1 {
				tmp := days[i][week] + visit(i, week+1)
				if tmp > res {
					res = tmp
				}
			}
		}
		memo[city][week] = res
		return res

	}

	return visit(0, 0)
}
