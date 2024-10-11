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
		jars := readNNums(reader, 2*n)
		res := solve(jars)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(jars []int) int {
	n := len(jars)
	for i := 0; i < n; i++ {
		if jars[i] == 2 {
			jars[i] = -1
		}
	}
	// n is even
	// pref[l] + suf[r] = 0
	pos := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		pos[i] = -1
	}
	res := n

	var sum int
	for i := 0; i < n/2; i++ {
		sum += jars[i]
		// 越靠近右边越好
		pos[sum+n] = i
		if sum == 0 {
			// 把右边的都删除掉
			res = min(res, n-i-1)
		}
	}

	var suf int
	for i := n - 1; i >= n/2; i-- {
		suf += jars[i]
		if suf == 0 {
			// 把前面的都删掉
			res = min(res, i)
		}
		// pref + suf = 0
		if pos[-suf+n] >= 0 {
			// 把中间的删掉
			res = min(res, i-pos[-suf+n]-1)
		}
	}

	return res
}
