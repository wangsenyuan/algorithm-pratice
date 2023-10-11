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
		b := readNNums(reader, n/2)
		res := solve(b)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
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

func solve(b []int) []int {
	h := len(b)

	n := 2 * h
	ok := make([]bool, n)

	for i := 0; i < h; i++ {
		x := b[i] - 1
		if ok[x] {
			return nil
		}
		ok[x] = true
	}

	arr := make([]Pair, 2*n)
	for i := n; i < 2*n; i++ {
		if ok[i-n] {
			arr[i] = Pair{-1, i - n}
		} else {
			arr[i] = Pair{i - n, i - n}
		}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = max_pair(arr[2*i], arr[2*i+1])
	}

	set := func(p int, v Pair) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = max_pair(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) Pair {
		l += n
		r += n
		res := Pair{-1, -1}
		for l < r {
			if l&1 == 1 {
				res = max_pair(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max_pair(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	a := make([]int, h)

	for i := h - 1; i >= 0; i-- {
		//
		cur := b[i] - 1
		v := get(0, cur)
		if v.first < 0 {
			return nil
		}
		a[i] = v.first + 1
		set(v.first, Pair{-1, v.first})
	}

	ans := make([]int, n)
	for i := 0; i < n; i += 2 {
		ans[i] = a[i/2]
		ans[i+1] = b[i/2]
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

func max_pair(a, b Pair) Pair {
	if a.first > b.first {
		return a
	}
	return b
}
