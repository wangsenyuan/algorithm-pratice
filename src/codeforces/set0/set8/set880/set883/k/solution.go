package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)

	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = readNNums(reader, 2)
	}

	sum, res := solve(segs)
	if sum < 0 {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", sum))
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(segs [][]int) (int, []int) {
	n := len(segs)
	lo, hi := segs[0][0], segs[0][0]+segs[0][1]
	L := make([]int, n)
	L[0] = hi
	for i := 1; i < n; i++ {
		cur := segs[i]
		lo--
		hi++
		if cur[0]+cur[1] < lo || cur[0] > hi {
			// 无法达到
			return -1, nil
		}
		lo = max(lo, cur[0])
		hi = min(hi, cur[0]+cur[1])
		L[i] = hi
	}

	res := L[n-1] - segs[n-1][0]

	for i := n - 2; i >= 0; i-- {
		L[i] = min(L[i+1]+1, L[i])
		res += L[i] - segs[i][0]
	}

	return res, L
}
