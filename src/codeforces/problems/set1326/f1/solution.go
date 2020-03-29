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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	P := make([]string, n)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadString('\n')
		P[i] = s[:len(s)-1]
	}
	res := solve(n, P)

	for i := 0; i < len(res); i++ {
		fmt.Printf("%d ", res[i])
	}
	fmt.Println()
}

func solve(n int, P []string) []int64 {
	N := 1 << n
	dp := make([][]map[int]int64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]map[int]int64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make(map[int]int64)
		}
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i][0] = 1
	}

	for state := 1; state < N; state++ {
		for i := 0; i < n; i++ {
			if state&(1<<i) == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if state&(1<<j) > 0 {
					continue
				}
				next := state | (1 << j)
				x := int(P[j][i] - '0')
				tmp := dp[state][i]

				for k, v := range tmp {
					dp[next][j][k<<1|x] += v
				}
			}
		}
	}

	ans := make([]int64, N/2)

	for i := 0; i < n; i++ {
		for k, v := range dp[N-1][i] {
			kk := reverse(k, n) / 2
			ans[kk] += v
		}
	}

	return ans
}

func reverse(num int, n int) int {
	var res int
	for n > 0 {
		res = res<<1 | num&1
		num >>= 1
		n--
	}
	return res
}
