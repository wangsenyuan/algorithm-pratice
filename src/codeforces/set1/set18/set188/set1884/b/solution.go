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
		s := readString(reader)[:n]
		res := solve(s)

		s = fmt.Sprintf("%v", res)
		s = s[1 : len(s)-1]

		buf.WriteString(fmt.Sprintf("%s\n", s))
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

func solve(s string) []int {
	n := len(s)

	stack := make([]int, n)
	var top int

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			stack[top] = i
			top++
		}
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if top == 0 {
			ans[i] = -1
			continue
		}
		if i > 0 {
			ans[i] = ans[i-1]
		}
		j := n - 1 - i
		ans[i] += j - stack[top-1]
		top--
	}

	return ans
}

func solve1(s string) []int {
	n := len(s)

	// 对于i，就是要把它右边的1移动到左边去
	// 且后面的i已经被移走了

	arr := make([]int, 2*n)

	set := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			set(i, i+1)
		}
	}
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			ans[i] = ans[i-1]
			if ans[i] < 0 {
				continue
			}
		}
		j := n - 1 - i
		p := get(0, j+1) - 1
		if p < 0 {
			ans[i] = -1
		} else {
			ans[i] += j - p
			if p < j {
				set(p, 0)
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
