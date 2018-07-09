package main

import "fmt"

func main() {
	var tc int
	fmt.Scanf("%d", &tc)

	for tc > 0 {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		ps := make([][]float64, n)

		for i := 0; i < n; i++ {
			ps[i] = make([]float64, m)
			for j := 0; j < m; j++ {
				fmt.Scanf("%f", &ps[i][j])
			}
		}
		fmt.Printf("%.6f\n", solve(n, m, ps))
		tc--
	}
}

func solve(n int, m int, ps [][]float64) float64 {
	for j := 0; j < m; j++ {
		if ps[0][j] == 0.0 {
			return 0.0
		}
	}

	var ans float64 = 1

	for i := 1; i < n; i++ {
		var tmp float64 = 1
		for j := 0; j < m; j++ {
			tmp *= ps[i][j] / ps[0][j]
		}
		ans += tmp
	}
	return 1.0 / ans
}
