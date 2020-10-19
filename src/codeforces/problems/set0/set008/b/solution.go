package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')

	ok := solve(normalize(s))

	if ok {
		fmt.Println("OK")
	} else {
		fmt.Println("BUG")
	}
}

const N = 110

var vis [2 * N][2 * N]int

func reset(n int) {
	for i := N - n; i <= N+n; i++ {
		for j := N - n; j <= N+n; j++ {
			vis[i][j] = 0
		}
	}
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s[:i]
		}
	}
	return ""
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(s string) bool {
	reset(len(s))

	var x, y = N, N
	vis[x][y] = 1
	for i := 0; i < len(s); i++ {
		steps := vis[x][y]
		cur := s[i]
		if cur == 'L' {
			x--
		} else if cur == 'R' {
			x++
		} else if cur == 'U' {
			y++
		} else {
			y--
		}

		if vis[x][y] > 0 {
			return false
		}

		for k := 0; k < 4; k++ {
			u, v := x+dd[k], y+dd[k+1]
			if vis[u][v] > 0 && vis[u][v] != steps {
				return false
			}
		}
		vis[x][y] = steps + 1
	}

	return true
}
