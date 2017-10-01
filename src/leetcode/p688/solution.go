package main

import "fmt"

func knightProbability(N int, K int, r int, c int) float64 {
	pos := make([][][]float64, N)

	for i := 0; i < N; i++ {
		pos[i] = make([][]float64, N)
		for j := 0; j < N; j++ {
			pos[i][j] = make([]float64, K+1)
			for k := 0; k <= K; k++ {
				pos[i][j][k] = 0.0
			}
		}
	}

	pos[r][c][0] = 1.0

	for k := 1; k <= K; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if i-2 >= 0 && j-1 >= 0 {
					pos[i-2][j-1][k] += pos[i][j][k-1] * .125
				}
				if i-2 >= 0 && j+1 < N {
					pos[i-2][j+1][k] += pos[i][j][k-1] * .125
				}
				if i-1 >= 0 && j+2 < N {
					pos[i-1][j+2][k] += pos[i][j][k-1] * .125
				}
				if i+1 < N && j+2 < N {
					pos[i+1][j+2][k] += pos[i][j][k-1] * .125
				}
				if i+2 < N && j+1 < N {
					pos[i+2][j+1][k] += pos[i][j][k-1] * .125
				}
				if i+2 < N && j-1 >= 0 {
					pos[i+2][j-1][k] += pos[i][j][k-1] * .125
				}
				if i+1 < N && j-2 >= 0 {
					pos[i+1][j-2][k] += pos[i][j][k-1] * .125
				}
				if i-1 >= 0 && j-2 >= 0 {
					pos[i-1][j-2][k] += pos[i][j][k-1] * .125
				}
			}
		}
		//fmt.Printf("%d, %d, %d, %f\n", k, r, c, pos[r][c][k])
	}

	var res float64
	for i := 0; i < N; i ++ {
		for j := 0; j < N; j ++ {
			res += pos[i][j][K]
		}
	}
	return res
}

func main() {
	//fmt.Println(knightProbability(3, 2, 0, 0))
	fmt.Println(knightProbability(3, 3, 1, 2))

}
