package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		s1 := readString(reader)
		s2 := readString(reader)
		x := readString(reader)
		res := solve(s1, s2, x)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const N_INF = math.MinInt32

func solve(S1, S2, X string) int64 {

	z1 := zFunction(S1 + "#" + X)
	z2 := zFunction(S2 + "#" + X)

	z1 = z1[len(S1)+1:]
	z2 = z2[len(S2)+1:]

	n := len(X)

	rmq := make([][]int, n)

	for i := 0; i < n; i++ {
		rmq[i] = make([]int, 22)
		for j := 0; j < 22; j++ {
			rmq[i][j] = N_INF
		}
		rmq[i][0] = z2[i] + i
	}

	for j := 1; j < 22; j++ {
		for i := 0; i+(1<<uint(j)) <= n; i++ {
			rmq[i][j] = max(rmq[i][j-1], rmq[i+(1<<uint(j-1))][j-1])
		}
	}
	query := func(l, r int) int {
		j := lg(r - l + 1)
		return max(rmq[l][j], rmq[r-(1<<uint(j))+1][j])
	}

	var res int64
	for l := 0; l < n; l++ {
		r := query(l, l+z1[l])
		res += int64(r - l)
	}

	return res
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func lg(num int) int {
	var res int
	for num > 0 {
		num >>= 1
		res++
	}
	return res - 1
}
