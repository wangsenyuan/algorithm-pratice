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

	cnt, res := solve(n)
	fmt.Println(cnt)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int) (int, string) {

	var res [][]string

	var process func(k int)

	process = func(k int) {
		var par []string

		for i := k; i+k <= n; i += 2 * k {
			par = append(par, fmt.Sprintf("%d+%d=%d", i, i+k, i+k))
		}

		res = append(res, par)

		if 4*k <= n {
			process(2 * k)
		}
		if 3*k <= n {
			var now []string
			for i := 3 * k; i <= n; i += 2 * k {
				now = append(now, fmt.Sprintf("%d+%d=%d", i-k, i, i))
			}
			res = append(res, now)
		}
	}

	process(1)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		if i > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(fmt.Sprintf("%d", len(res[i])))
		for j := 0; j < len(res[i]); j++ {
			buf.WriteString(fmt.Sprintf(" %s", res[i][j]))
		}
	}

	return len(res), buf.String()
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
