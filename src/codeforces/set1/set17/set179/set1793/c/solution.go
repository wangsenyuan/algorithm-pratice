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
		a := readNNums(reader, n)
		res := solve(a)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
		}
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

func solve(a []int) []int {
	n := len(a)
	// l not min/max of the segment [l...r]
	// 先找到n和1的位置，如果它们两边可以选，就找到答案了
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		a[i]--
		pos[a[i]] = i
	}

	u, v := 0, n-1
	l, r := 0, n-1
	for u < v {
		pu := pos[u]
		pv := pos[v]
		pu, pv = min(pu, pv), max(pu, pv)
		if pu > l && pv < r {
			// pu, pv 是最大/最小值，且它们在中间
			return []int{l + 1, r + 1}
		}
		if pu == l && pv == r {
			u++
			v--
			l++
			r--
			continue
		}
		if pu == l {
			if a[pu] == u {
				u++
			} else {
				// a[pu] = v
				v--
			}
			l++
		} else {
			// pv == r
			if a[pv] == v {
				v--
			} else {
				u++
			}
			r--
		}
	}

	return nil
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
