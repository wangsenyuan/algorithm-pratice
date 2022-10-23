package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	res := solve(s)

	fmt.Println(res)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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

const MOD = 998244353

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func solve(s string) int {
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, 8)
		for j := 0; j < 8; j++ {
			dp[i][j] = make([]int, 8)
		}
	}
	dp[0][0][0] = 1

	n := len(s)

	for i := 1; i <= n; i++ {
		clear(dp[i&1])
		for mask1 := 0; mask1 < 8; mask1++ {
			for mask2 := 0; mask2 < 8; mask2++ {
				for m := 0; m < 8; m++ {
					var flag bool
					for j := 0; j < 3; j++ {
						if s[i-1] == '0' && (mask2>>j)&1 == 0 && (m>>j)&1 == 1 {
							flag = true
							break
						}
					}
					if flag {
						continue
					}
					tmp1 := mask1
					tmp2 := mask2

					for j := 0; j < 3; j++ {
						if int(s[i-1]-'0') != (m>>j)&1 {
							tmp2 |= 1 << j
						}
					}
					for j := 0; j < 3; j++ {
						if m == (1<<j) || m == 7-(1<<j) {
							tmp1 |= 1 << j
						}
					}
					dp[i&1][tmp1][tmp2] = modAdd(dp[i&1][tmp1][tmp2], dp[(i-1)&1][mask1][mask2])
				}
			}
		}
	}

	var ans int

	for i := 0; i < 8; i++ {
		ans = modAdd(ans, dp[n&1][7][i])
	}

	return ans
}

func clear(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] = 0
		}
	}
}
