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
		readString(reader)
		n, m := readTwoNums(reader)
		field := make([]string, n)
		for i := 0; i < n; i++ {
			field[i] = readString(reader)[:m]
		}
		res := solve(field)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(field []string) int {
	n := len(field)
	m := len(field[0])
	// 这个题目，我没法把它抽象出来

	vs := make([][]int, 2)
	for i := 0; i < 2; i++ {
		vs[i] = make([]int, 0, 1)
	}
	prev := make([]int, 2)
	var res int

	for d := 0; d < m+n; d++ {
		// i + j = d
		li := max(0, d-m+1)
		ri := min(n-1, d)
		if li > ri {
			continue
		}

		var cur []int
		for i := li; i <= ri; i++ {
			j := d - i

			if field[i][j] == '1' {
				cur = append(cur, i)
			}
		}
		ad := d / 2
		x := d & 1
		for prev[x] < len(vs[x]) && vs[x][prev[x]]+ad > ri {
			prev[x]++
		}
		p := len(cur)
		// vs[x][prev] + ad <= ri
		for i := prev[x]; i < len(vs[x]); i++ {
			new_val := vs[x][i]
			for p > 0 && cur[p-1]-ad >= vs[x][i] {
				new_val = max(new_val, cur[p-1]-ad)
				p--
			}
			vs[x][i] = new_val
		}
		if p > 0 {
			res++
			vs[x] = append(vs[x], cur[p-1]-ad)
		}
	}

	return res
}

func last(arr []int) int {
	return arr[len(arr)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
