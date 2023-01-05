package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	board := make([]string, n)

	for i := 0; i < n; i++ {
		board[i] = readString(reader)[:m]
	}

	res := solve(n, m, board)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	for _, row := range res {
		fmt.Println(row)
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, m int, board []string) []string {
	// . for cut out
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			res[i][j] = '.'
		}
	}

	get := func(x, y int, dx int, dy int) [][]int {
		var pos [][]int
		pos = append(pos, []int{x, y})
		if x+dx < n && y+dy < m && board[x+dx][y+dy] == 'b' && res[x+dx][y+dy] == '.' {
			pos = append(pos, []int{x + dx, y + dy})
			if x+2*dx < n && y+2*dy < m && board[x+2*dx][y+2*dy] == 'w' && res[x+2*dx][y+2*dy] == '.' {
				pos = append(pos, []int{x + 2*dx, y + 2*dy})
				return pos
			}
		}
		return nil
	}

	getRight := func(x, y int) [][]int {
		return get(x, y, 0, 1)
	}

	getDown := func(x, y int) [][]int {
		return get(x, y, 1, 0)
	}

	check := func(pos [][]int, c byte) bool {
		for _, cell := range pos {
			x, y := cell[0], cell[1]
			if x > 0 && res[x-1][y] == c {
				return false
			}
			if x+1 < n && res[x+1][y] == c {
				return false
			}
			if y > 0 && res[x][y-1] == c {
				return false
			}
			if y+1 < m && res[x][y+1] == c {
				return false
			}
		}
		return true
	}

	fill := func(pos [][]int, c byte) {
		for _, cell := range pos {
			x, y := cell[0], cell[1]
			res[x][y] = c
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if res[i][j] == '.' {
				// not filled yet
				if board[i][j] == '.' {
					continue
				}
				if board[i][j] == 'b' {
					// bad case, must always starts with w
					return nil
				}
				res := getRight(i, j)
				if len(res) == 0 {
					res = getDown(i, j)
				}
				if len(res) == 0 {
					return nil
				}
				ok := false
				for x := 0; x < 4; x++ {
					c := byte(x + 'a')
					if check(res, c) {
						ok = true
						fill(res, c)
						break
					}
				}
				if !ok {
					return nil
				}
			}
		}
	}

	ans := make([]string, n)

	for i := 0; i < n; i++ {
		ans[i] = string(res[i])
	}

	return ans
}
