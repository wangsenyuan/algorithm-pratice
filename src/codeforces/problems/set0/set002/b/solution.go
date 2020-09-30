package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1000

var dp [N][N]int
var cnt [N][N]int

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	M := make([][]int, n)
	for i := 0; i < n; i++ {
		M[i] = readNNums(reader, n)
	}
	best, path := solve(n, M)
	fmt.Println(best)
	fmt.Println(path)
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

const INF = 1 << 30

func solve(n int, M [][]int) (int, string) {
	var hasZero bool
	u, v := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 0 {
				hasZero = true
				u, v = i, j
				M[i][j] = 10
			}
		}
	}

	var limit = INF
	if hasZero {
		limit = 1
	}

	a, x := getMinCountPath(n, M, 2, limit)

	if a == 0 {
		return a, x
	}

	b, y := getMinCountPath(n, M, 5, a)

	if b == 0 {
		return b, y
	}

	if hasZero {
		// choose the 0, will give answer 1
		path := make([]byte, 0, 2*n-1)
		for i := 0; i < u; i++ {
			path = append(path, 'D')
		}
		for j := 0; j < v; j++ {
			path = append(path, 'R')
		}
		for i := u + 1; i < n; i++ {
			path = append(path, 'D')
		}
		for j := v + 1; j < n; j++ {
			path = append(path, 'R')
		}
		return 1, string(path)
	}

	if a <= b {
		return a, x
	}
	return b, y
}

func getMinCountPath(n int, M [][]int, div int, limit int) (int, string) {
	//dp := make([][]int, n)
	//cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		//dp[i] = make([]int, n)
		//cnt[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cnt[i][j] = count(M[i][j], div)
			//dp[i][j] = 0
		}
	}

	get := func(i, j int) int {
		if i < 0 || j < 0 {
			return INF
		}
		return dp[i][j]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = cnt[i][j]
			if i > 0 || j > 0 {
				dp[i][j] += min(get(i-1, j), get(i, j-1))
			}
		}
	}

	best := dp[n-1][n-1]

	if best >= limit {
		return best, ""
	}

	res := make([]byte, 0, 2*n-1)

	x, y := n-1, n-1

	for x >= 0 && y >= 0 {
		cur := dp[x][y]
		if x > 0 && dp[x-1][y] == cur-cnt[x][y] {
			x--
			res = append(res, 'D')
		} else {
			y--
			res = append(res, 'R')
		}
		if x == 0 && y == 0 {
			break
		}
	}

	reverse(res)

	return best, string(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func count(num int, div int) int {
	var res int
	for num%div == 0 {
		num /= div
		res++
	}
	return res
}

func reverse(buf []byte) {
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
}
