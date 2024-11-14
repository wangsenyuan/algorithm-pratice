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
		res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &k) + 1
		a[i] = make([]int, k)
		for j := 0; j < k; j++ {
			pos = readInt(s, pos, &a[i][j]) + 1
		}
	}
	return solve(a, m)
}

type pair struct {
	first  int
	second int
}

func solve(a [][]int, m int) int {

	var k int
	for _, cur := range a {
		tmp := getMex(cur)
		k = max(k, tmp.second)
	}

	if m <= k {
		return (m + 1) * k
	}

	return k*k + (m+k)*(m-k+1)/2
}

func getMex(arr []int) pair {
	n := len(arr)
	var x int
	cnt := make([]int, n+2)

	for _, v := range arr {
		if v > n {
			// 不可能达到这个mex
			continue
		}
		cnt[v]++
		for cnt[x] > 0 {
			x++
		}
	}
	// 如果我们提供了x
	y := x + 1

	for cnt[y] > 0 {
		y++
	}

	return pair{x, y}
}
