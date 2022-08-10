package main

func main() {

}

func solve(n int, blocks [][]int) []int {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	for _, blk := range blocks {
		x, y := blk[0], blk[1]
		x--
		y--
		board[x][y] = -1
	}

	safe := func(r, c int) bool {
		// put a queue at position (r, c)
		// no need to check same column or column after c

	}
}
