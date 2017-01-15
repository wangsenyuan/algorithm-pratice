package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	t := make([][]int, n)

	for i := 0; i < n; i++ {
		t[i] = make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &t[i][j])
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if t[i][k]+t[k][j] < t[i][j] {
					t[i][j] = t[i][k] + t[k][j]
				}
			}
		}
	}

	var m int
	fmt.Scanf("%d", &m)

	for m > 0 {
		var s, g, d int
		fmt.Scanf("%d %d %d", &s, &g, &d)

		fmt.Printf("%d %d\n", t[s][g]+t[g][d], t[s][g]+t[g][d]-t[s][d])

		m--
	}
}
