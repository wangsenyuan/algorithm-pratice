package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var a, b, c, d, e, f, L, R int
		fmt.Scanf("%d %d %d %d %d %d %d %d", &a, &b, &c, &d, &e, &f, &L, &R)
		found, res := solve(a, b, c, d, e, f, L, R)
		if found {
			for _, coor := range res {
				fmt.Printf("%d %d\n", coor[0], coor[1])
			}
		} else {
			fmt.Println(-1)
		}
		t--
	}
}

func solve(a, b, c, d, e, f int, L, R int) (bool, [][]int) {
	// c & f is not useful
	x, y, z := area(a, b, d, e)

	var found bool
	res := make([][]int, 6)
	if (x <= L && y <= R) || (y <= L && x <= R) {
		res[0] = []int{0, b}
		res[1] = []int{a, 0}
		res[2] = []int{0, 0}
		res[3] = []int{0, d}
		res[4] = []int{e, 0}
		res[5] = []int{e, d}
		found = true
	}
	xx, yy, zz := area(b, a, d, e)

	if (xx <= L && yy <= R) || (y <= L && x <= R) {
		if !found || zz < z {
			res[0] = []int{0, 0}
			res[1] = []int{b, a}
			res[2] = []int{b, 0}
			res[3] = []int{0, d}
			res[4] = []int{e, 0}
			res[5] = []int{e, d}
			found = true
		}
	}
	if found {
		return true, res
	}
	return false, nil
}

func area(a, b, c, d int) (x int, y int, z int64) {
	x = max(a, c)
	y = max(b, d)
	z = int64(x) * int64(y)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
