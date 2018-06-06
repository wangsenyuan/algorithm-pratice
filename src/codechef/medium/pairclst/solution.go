package main

func main() {

}

func solve(N, M, K int, special []int, edges [][]int) int {
	graph := make([]map[int]int, N)

	for i := 0; i < N; i++ {
		graph[i] = make(map[int]int)
	}

	for _, edge := range edges {
		u, v, w := edge[0]-1, edge[1]-1, edge[2]
		graph[u][v] = w
		graph[v][u] = w
	}

	INF := 1 << 30

	DD := make([][]int, N)
	for i := 0; i < N; i++ {
		DD[i] = make([]int, N)
		for j := 0; j < N; j++ {
			DD[i][j] = INF
		}
		DD[i][i] = 0
	}

	dd := make([]int, N)

	que := make([]int, N)
	bfs := func(x int, label int) {
		front, tail := 0, 0
		que[tail] = x
		tail++

		for front < tail {
			tt := tail

			for front < tt {
				u := que[front]
				front++
				for v, w := range graph[u] {

				}
			}
		}
	}
}
