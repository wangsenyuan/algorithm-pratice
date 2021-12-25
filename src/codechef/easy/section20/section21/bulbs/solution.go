package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		S, _ := reader.ReadBytes('\n')
		res := solve(n, k, S)
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

func solve(n, k int, S []byte) int {
	var i int
	for i < n && S[i] == '0' {
		i++
	}
	if i == n {
		return 0
	}

	var j = n - 1
	for j >= 0 && S[j] == '0' {
		j--
	}

	a := solve1(k, S[i:j+1]) + i + n - j - 1
	b := solve1(k-1, S[i:j+1]) + min(i, n-j-1)
	c := solve1(k-2, S[i:j+1])
	return min3(a, b, c)
}

func solve1(k int, S []byte) int {
	if k < 0 {
		return math.MaxInt32
	}
	// S[0] == '1' && S[n-1] == '1'
	var V []int
	var cnt int
	for i := 0; i < len(S); i++ {
		if S[i] == '1' {
			if cnt > 0 {
				V = append(V, cnt)
			}
			cnt = 0
		} else {
			cnt++
		}
	}
	sort.Ints(V)
	var p = len(V)
	for k > 1 && p > 0 {
		p--
		k -= 2
	}
	var res int
	for p > 0 {
		res += V[p-1]
		p--
	}
	return res
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
