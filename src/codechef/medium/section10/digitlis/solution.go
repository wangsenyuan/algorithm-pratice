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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		L := readNNums(reader, n)
		fmt.Println(solve(n, L))
	}
}

const MAX_D = 10
const MAX_N = 1001
const MOD = 1000000007

var state [][]int
var dp [][]int

func init() {
	state = make([][]int, 1<<MAX_D)
	for i := 0; i < len(state); i++ {
		state[i] = make([]int, MAX_D+1)

		state[i][0] = -1

		for k := 1; k <= MAX_D; k++ {
			state[i][k] = 9
		}

		for x, j, k := i, 0, 1; x > 0; x, j = x>>1, j+1 {
			if x&1 == 1 {
				state[i][k] = j
				k++
			}
		}
	}

	dp = make([][]int, MAX_N)
	for i := 0; i < MAX_N; i++ {
		dp[i] = make([]int, 1<<MAX_D)
	}
}

func reset(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = 0
		}
	}
	for i := 0; i < MAX_D; i++ {
		if i > 0 || n == 1 {
			dp[0][1<<uint(i)] = 1
		} else {
			dp[0][1<<uint(i)] = 0
		}
	}
}

func add(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func solve(n int, L []int) int {
	reset(n)
	//first element has to be 1
	for i := 1; i < n; i++ {
		l := L[i]
		for mask := 0; mask < 1<<MAX_D; mask++ {
			for x := state[mask][l-1] + 1; x <= state[mask][l]; x++ {
				add(&dp[i][mask^(1<<uint(state[mask][l]))^(1<<uint(x))], dp[i-1][mask])
			}
		}
	}

	var res int

	for mask := 0; mask < 1<<MAX_D; mask++ {
		add(&res, dp[n-1][mask])
	}

	return res
}
