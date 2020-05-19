package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		R, C := readTwoNums(reader)
		G := make([]string, R)
		for i := 0; i < R; i++ {
			G[i], _ = reader.ReadString('\n')
		}
		fmt.Println(solve(R, C, G))
	}
}

const MAX_N = 20

var sg [MAX_N][MAX_N][MAX_N][MAX_N]int
var monsters [MAX_N][MAX_N]bool

func solve(R, C int, G []string) string {

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			for u := 0; u < R; u++ {
				for v := 0; v < C; v++ {
					sg[i][j][u][v] = 0
				}
			}
		}
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			monsters[i][j] = false

			if G[i][j] == '#' {
				continue
			}

			if i-2 < 0 || G[i-1][j] == '#' || G[i-2][j] == '#' {
				continue
			}
			if i+2 >= R || G[i+1][j] == '#' || G[i+2][j] == '#' {
				continue
			}
			if j-2 < 0 || G[i][j-1] == '#' || G[i][j-2] == '#' {
				continue
			}
			if j+2 >= C || G[i][j+1] == '#' || G[i][j+2] == '#' {
				continue
			}
			monsters[i][j] = true
		}
	}

	for w := 0; w < C; w++ {
		for xl := 0; xl+w < C; xl++ {
			for h := 0; h < R; h++ {
				for yl := 0; yl+h < R; yl++ {
					seen := make(map[int]bool)
					for x := xl; x <= xl+w; x++ {
						for y := yl; y <= yl+h; y++ {
							if monsters[x][y] {
								tmp := sg[xl][yl][x-1][y-1] ^
									sg[xl][y+1][x-1][yl+w] ^
									sg[x+1][yl][xl+w][y-1] ^
									sg[x+1][y+1][xl+w][yl+w]

								seen[tmp] = true
							}
						}
					}
					var num int
					for seen[num] {
						num++
					}
					sg[xl][yl][xl+w][yl+h] = num
				}
			}
		}
	}

	if sg[0][0][R-1][C-1] > 0 {
		return "Asuna"
	}
	return "Kirito"
}
