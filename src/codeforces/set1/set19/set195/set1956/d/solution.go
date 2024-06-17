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
	a := readNNums(reader, n)

	res, steps := solve(a)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %d\n", res, len(steps)))

	for _, cur := range steps {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

func solve(a []int) (int, [][]int) {
	n := len(a)

	// intersting
	var best_state int
	var best_ans int
	for state := 0; state < 1<<n; state++ {
		var tmp int
		for l := 0; l < n; l++ {
			if (state>>l)&1 == 1 {
				r := l
				for r+1 < n && (state>>(r+1))&1 == 1 {
					r++
				}
				// l...r 可以得到最大的 (r - l + 1)
				tmp += (r - l + 1) * (r - l + 1)

				l = r
			} else {
				tmp += a[l]
			}
		}
		if tmp > best_ans {
			best_ans = tmp
			best_state = state
		}
	}

	var oper func(l int, r int)
	var build func(l int, r int)

	build = func(l int, r int) {
		if l == r {
			if a[l] > 0 {
				oper(l, r)
			}
			return
		}
		build(l, r-1)
		if a[r] != r-l {
			oper(l, r)
			build(l, r-1)
		}
	}

	var steps [][]int

	cnt := make([]int, 2*n)

	oper = func(l int, r int) {
		clear(cnt)
		for i := l; i <= r; i++ {
			if a[i] <= n {
				cnt[a[i]]++
			}
		}
		var mex int
		for cnt[mex] > 0 {
			mex++
		}
		for i := l; i <= r; i++ {
			a[i] = mex
		}
		steps = append(steps, []int{l + 1, r + 1})
	}

	for l := 0; l < n; l++ {
		if (best_state>>l)&1 == 1 {
			r := l
			for r+1 < n && (best_state>>(r+1))&1 == 1 {
				r++
			}
			build(l, r)
			oper(l, r)
			l = r
		}
	}

	return best_ans, steps
}
