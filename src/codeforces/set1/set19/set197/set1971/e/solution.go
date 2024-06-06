package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k, q := readThreeNums(reader)
		a := readNNums(reader, k)
		b := readNNums(reader, k)
		d := make([]int, q)
		for i := 0; i < q; i++ {
			d[i] = readNum(reader)
		}
		res := solve(n, a, b, d)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
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
func solve(n int, a []int, b []int, d []int) []int {
	id := make([]int, len(d))
	for i := 0; i < len(d); i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return d[id[i]] < d[id[j]]
	})
	ans := make([]int, len(d))
	k := len(a)
	prev_pos := 0
	prev_time := 0
	for i, j := 0, 0; i < k; i++ {
		for j < len(d) && d[id[j]] <= a[i] {
			if d[id[j]] == prev_pos {
				ans[id[j]] = prev_time
			} else {
				// d[id[j]] > prev_pos
				// (t - prev_time) / (d - prev_pos) = (b[i] - prev_time) / (a[i] - prev_pos)
				tmp := (b[i] - prev_time) * (d[id[j]] - prev_pos) / (a[i] - prev_pos)
				ans[id[j]] = tmp + prev_time
			}
			j++
		}

		prev_time = b[i]
		prev_pos = a[i]
	}

	return ans
}
