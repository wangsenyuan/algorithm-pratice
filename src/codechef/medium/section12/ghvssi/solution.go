package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)

		if res {
			fmt.Println("Ghayeeth")
		} else {
			fmt.Println("Siroj")
		}
	}
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

const S = 311

var dp [2][S][S][S][4]bool
var vis [2][S][S][S][4]bool

func reset(a, b, c, d int) {
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				for u := 0; u <= d; u++ {
					for v := 0; v < 4; v++ {
						dp[i][j][k][u][v] = false
						vis[i][j][k][u][v] = false
					}
				}
			}
		}
	}
}

func solve(n int, A []int) bool {
	cnt := make([]int, 4)

	for i := 0; i < n; i++ {
		cnt[A[i]%4]++
	}

	cnt[0] %= 2

	reset(cnt[0], cnt[1], cnt[2], cnt[3])

	search(cnt[0], cnt[1], cnt[2], cnt[3], 0)

	return dp[cnt[0]][cnt[1]][cnt[2]][cnt[3]][0]
}

func search(i, j, k, l, p int) {
	if vis[i][j][k][l][p] {
		return
	}
	vis[i][j][k][l][p] = true
	dp[i][j][k][l][p] = false

	if i > 0 {
		search(i-1, j, k, l, p)
	}
	if j > 0 && (p+1)%4 != 2 {
		search(i, j-1, k, l, (p+1)%4)
	}
	if k > 0 && (p+2)%4 != 2 {
		search(i, j, k-1, l, (p+2)%4)
	}
	if l > 0 && (p+3)%4 != 2 {
		search(i, j, k, l-1, (p+3)%4)
	}

	if i > 0 && !dp[i-1][j][k][l][p] {
		dp[i][j][k][l][p] = true
	}

	if j > 0 && !dp[i][j-1][k][l][(p+1)%4] && (p+1)%4 != 2 {
		dp[i][j][k][l][p] = true
	}

	if k > 0 && !dp[i][j][k-1][l][(p+2)%4] && (p+2)%4 != 2 {
		dp[i][j][k][l][p] = true
	}

	if l > 0 && !dp[i][j][k][l-1][(p+3)%4] && (p+3)%4 != 2 {
		dp[i][j][k][l][p] = true
	}
}
