package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	S := readString(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		var tp int
		line, _ := reader.ReadBytes('\n')
		pos := readInt(line, 0, &tp)
		var l, r int
		pos = readInt(line, pos+1, &l)
		pos = readInt(line, pos+1, &r)
		if tp == 1 {
			Q[i] = []int{tp, l, r}
		} else {
			var x int
			readInt(line, pos+1, &x)
			Q[i] = []int{tp, l, r, x}
		}
	}
	res := solve(n, S, Q)
	var buf bytes.Buffer
	for _, ans := range res {
		if ans {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(n int, S string, Q [][]int) []bool {

	arr := make([]int, 2*n)

	INF := 2 * n

	for i := 0; i < len(arr); i++ {
		arr[i] = INF
	}

	set := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 0 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		l += n
		r += n
		res := INF
		for l < r {
			if l&1 == 1 {
				res = min(arr[l], res)
				l++
			}
			if r&1 == 1 {
				r--
				res = min(arr[r], res)
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	for i := 1; i < n; i++ {
		if S[i] != S[i-1] {
			set(i, i)
		}
	}

	check := func(l, r, x int) bool {
		b := 1
		it := get(l+1, n)
		keep := x - 1
		for it <= r && b < 3 {
			keep = it
			it = get(it+1, n)
			b++
		}
		if b == 3 {
			return true
		}
		if b == 2 && x != l && x == keep {
			return true
		}
		if b == 1 && x == l {
			return true
		}
		return false
	}

	var ans []bool

	for _, cur := range Q {
		tp, l, r := cur[0], cur[1], cur[2]
		l--
		r--
		if tp == 1 {
			// flip
			for _, x := range []int{l, r + 1} {
				if 0 < x && x < n {
					if get(x, x+1) == x {
						set(x, INF)
					} else {
						set(x, x)
					}
				}
			}
		} else {
			x := cur[3]
			x--
			ans = append(ans, check(l, r, x))
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
