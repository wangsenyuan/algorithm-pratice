package main

import "fmt"

func main() {
	var tc int
	fmt.Scanf("%d", &tc)
	for tc > 0 {
		tc--
		var l, r, g int64
		fmt.Scanf("%d %d %d", &l, &r, &g)
		fmt.Println(solve(l, r, g))
	}
}

func solve(l, r, g int64) int64 {
	d := r/g - (l-1)/g
	if d != 1 {
		return d
	}
	if g >= l && g <= r {
		return 1
	}

	return 0
}
