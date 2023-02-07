package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		segs := make([][]int, n)
		for i := 0; i < n; i++ {
			segs[i] = readNNums(reader, 3)
		}
		res := solve(n, segs)

		for _, num := range res {
			buf.WriteString(fmt.Sprintf("%d\n", num))
		}
	}
	fmt.Print(buf.String())
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

func solve(n int, segs [][]int) []int {
	//只有l,r 是关键
	res := make([]int, n)
	const INF = 1 << 30

	minL := INF
	costL := INF
	maxR := 0
	costR := INF
	maxLen := 0
	costLen := INF

	for i := 0; i < n; i++ {
		l, r, c := segs[i][0], segs[i][1], segs[i][2]
		if l < minL {
			minL = l
			costL = INF
		}
		if minL == l {
			costL = min(costL, c)
		}
		if maxR < r {
			maxR = r
			costR = INF
		}
		if maxR == r {
			costR = min(costR, c)
		}
		if maxLen < r-l+1 {
			maxLen = r - l + 1
			costLen = INF
		}
		if maxLen == r-l+1 {
			costLen = min(costLen, c)
		}
		tmp := costL + costR
		if maxLen == maxR-minL+1 {
			tmp = min(tmp, costLen)
		}
		res[i] = tmp
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
