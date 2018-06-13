package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		islands := make([][]int, n)

		for j := 0; j < n; j++ {
			var x, y int
			fmt.Scanf("%d %d", &x, &y)
			islands[j] = []int{x, y}
		}

		res := solve(n, islands)
		for j := 0; j < n; j++ {
			fmt.Println(res[j])
		}
	}
}

func solve(n int, islands [][]int) []string {
	// boundary
	a, b, c, d := islands[0][0], islands[0][0], islands[0][1], islands[0][1]
	for i := 1; i < n; i++ {
		x, y := islands[i][0], islands[i][1]

		if x < a {
			a = x
		}
		if x > b {
			b = x
		}

		if y < c {
			c = y
		}
		if y > d {
			d = y
		}
	}

	ans1, dir1, found1 := findOnePointSolution(n, islands, a, b, c, d)

	if found1 {
		return []string{fmt.Sprintf("%d %s", ans1, dir1)}
	}

	ans2, dir2, ans3, dir3 := findTwoPointsSolution(n, islands, a, b)

	return []string{fmt.Sprintf("%d %s", ans2, dir2), fmt.Sprintf("%d %s", ans3, dir3)}
}

func findOnePointSolution(n int, islands [][]int, a int, b int, c int, d int) (int, string, bool) {
	for i := 0; i < n; i++ {
		x, y := islands[i][0], islands[i][1]

		if x == a && y == d {
			//top-left
			return i, "SE", true
		}

		if x == a && y == c {
			//bottom-left
			return i, "NE", true
		}

		if x == b && y == d {
			//top right
			return i, "SW", true
		}

		if x == b && y == c {
			//bottom right
			return i, "NW", true
		}
	}

	return -1, "", false
}

func findTwoPointsSolution(n int, islands [][]int, a int, b int) (int, string, int, string) {
	i := 0
	for i < n {
		if islands[i][0] == a {
			break
		}
		i++
	}

	j := 0
	for j < n {
		if islands[j][0] == b {
			break
		}
		j++
	}

	if islands[i][1] <= islands[j][1] {
		return i, "NE", j, "SW"
	}
	return i, "SE", j, "NW"
}
