package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
func solve(n int, A []int) int64 {
	// A[i] * A[j] % (i * j) == 0
	dp := make(map[int]map[int]int64)

	var res int64

	ff := make(map[int]map[int]bool)

	getDivisors := func(n int) map[int]bool {

		if ff[n] != nil {
			return ff[n]
		}

		ff[n] = make(map[int]bool)

		for i := 1; i*i <= n; i++ {
			if n%i == 0 {
				ff[n][i] = true
				j := n / i
				if i != j {
					ff[n][j] = true
				}
			}
		}
		return ff[n]
	}

	for i := 1; i <= n; i++ {
		a, b := A[i-1], i
		g := gcd(a, b)
		a /= g
		b /= g
		fs := getDivisors(a)

		if dp[b] == nil {
			dp[b] = make(map[int]int64)
		}

		for f := range fs {
			res += dp[b][f]
		}

		for f := range fs {
			if dp[f] == nil {
				dp[f] = make(map[int]int64)
			}
			dp[f][b]++
		}
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
