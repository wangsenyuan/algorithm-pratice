package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	B := readNNums(reader, n)

	fmt.Println(solve(n, B))
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

const MOD = 1000000007
const MAX = 51
const MAX2 = MAX * MAX

var pre [MAX][MAX]int

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func init() {
	for x := 0; x < MAX; x++ {
		for y := 0; y < MAX; y++ {
			pre[x][y] = gcd(x, y)
		}
	}
}

func solve(n int, B []int) int {
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, MAX)
		for j := 0; j < MAX; j++ {
			dp[i][j] = make([]int, MAX2)
		}
	}

	for x := 0; x <= B[0]; x++ {
		dp[0][x][x] = 1
	}

	for pos := 0; pos < n-1; pos++ {
		for x := 0; x < MAX; x++ {
			for y := 0; y <= B[pos+1]; y++ {
				for s := 0; s+y < MAX2; s++ {
					z := pre[x][y]
					modAdd(&dp[pos+1][z][s+y], dp[pos][x][s])
				}
			}
		}
	}

	var ans int

	for g := 1; g < MAX; g++ {
		for s := 1; s*g < MAX2; s *= 2 {
			modAdd(&ans, dp[n-1][g][s*g])
		}
	}

	return ans
}

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}
