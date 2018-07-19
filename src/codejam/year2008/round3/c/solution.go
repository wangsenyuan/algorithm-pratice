package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./C-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var M, N int
		fmt.Fscanf(f, "%d %d", &M, &N)
		room := make([][]byte, M)
		for j := 0; j < M; j++ {
			var s string
			fmt.Fscanf(f, "%s", &s)
			// fmt.Printf("[debug]%s\n", s)
			room[j] = []byte(s)
		}
		res := solve(M, N, room)
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

func solve(M, N int, room [][]byte) int {
	visited := make([]int, M*N)
	pair := make([]int, M*N)
	for i := 0; i < M*N; i++ {
		pair[i] = -1
	}

	var dfs func(u int, flag int) bool

	dfs = func(u int, flag int) bool {
		if visited[u] == flag {
			return false
		}
		visited[u] = flag

		x, y := u/N, u%N

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j += 2 {
				a, b := x+i, y+j
				if a >= 0 && a < M && b >= 0 && b < N && room[a][b] == '.' {
					v := a*N + b
					if pair[v] < 0 || dfs(pair[v], flag) {
						pair[v] = u
						return true
					}
				}
			}
		}

		return false
	}

	var ans, flag int

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if room[i][j] == '.' {
				ans++
				flag++
				if j%2 == 0 && dfs(i*N+j, flag) {
					ans--
				}
			}
		}
	}

	return ans
}
