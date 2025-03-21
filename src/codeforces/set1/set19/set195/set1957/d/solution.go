package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

const D = 30

func solve(a []int) int {

	var sum int
	cnt := make([][2]int, D)

	for _, num := range a {
		sum ^= num
		for i := D - 1; i >= 0; i-- {
			x := (sum >> i) & 1
			cnt[i][x]++
		}
	}
	var res int
	sum = 0

	prev := make([][2]int, D)

	for _, num := range a {
		for i := D - 1; i >= 0; i-- {
			x := (sum >> i) & 1
			prev[i][x]++
		}
		j := bits.Len(uint(num)) - 1
		res += prev[j][0]*cnt[j][0] + prev[j][1]*cnt[j][1]

		sum ^= num
		for i := D - 1; i >= 0; i-- {
			x := (sum >> i) & 1
			cnt[i][x]--
		}
	}

	return res
}

func solve1(a []int) int {
	var res int

	// dp[i] = a[y][i] 时的 g(x-1)[i] = 它相反数时的计数
	dp := make([][2]int, D)
	gp := make([][2]int, D)

	for i := 0; i < D; i++ {
		gp[i][0]++
	}

	var sum int
	for _, num := range a {
		sum ^= num

		j := 31 - bits.LeadingZeros32(uint32(num))
		dp[j][0] += gp[j][0]
		dp[j][1] += gp[j][1]

		for i := D - 1; i >= 0; i-- {
			x := (sum >> i) & 1
			res += dp[i][x]
			gp[i][x]++
		}
	}

	return res
}
