package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(minDistance("sea", "eat"))
}

func minDistance(word1 string, word2 string) int {
	a, b := len(word1), len(word2)

	f := make([][]int, a+1)
	for i := 0; i <= a; i++ {
		f[i] = make([]int, b+1)
		f[i][0] = i
		for j := 1; j <= b; j++ {
			f[i][j] = math.MaxInt32
		}
	}

	for j := 0; j <= b; j++ {
		f[0][j] = j
	}

	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
				continue
			}

			// remove both
			tmp := 2 + f[i-1][j-1]
			if tmp > f[i-1][j]+1 {
				//remove word1[i - 1]
				tmp = f[i-1][j] + 1
			}
			if tmp > f[i][j-1]+1 {
				//remove word2[j - 1]
				tmp = f[i][j-1] + 1
			}
			f[i][j] = tmp
		}
	}

	return f[a][b]
}
