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
		// ignore
		readString(reader)
		a := readString(reader)
		b := readString(reader)
		res := solve(a, b)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(s string, x string) bool {

	n := len(s)
	set := make([]int, 2*n)
	for i := 0; i < len(set); i++ {
		set[i] = -1
	}
	// set 表示在区间[l..r)上选择过的最小值
	update := func(i int, v int) {
		i += n
		set[i] = v
		for i > 1 {
			set[i>>1] = max(set[i], set[i^1])
			i >>= 1
		}
	}

	get := func(l int, r int) int {
		l += n
		r += n
		res := -1
		for l < r {
			if l&1 == 1 {
				res = max(res, set[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, set[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	pos := make([][]int, 26)

	for i := 0; i < n; i++ {
		a := int(s[i] - 'a')
		if pos[a] == nil {
			pos[a] = make([]int, 0, 1)
		}
		pos[a] = append(pos[a], i)
	}

	it := make([]int, 26)

	for i := 0; i < len(x); i++ {
		a := int(x[i] - 'a')
		ok := false
		// 每个位置只会被访问一次，所以这里的复杂性是 n * lgn
		for it[a] < len(pos[a]) && !ok {
			k := pos[a][it[a]]
			it[a]++
			ok = get(k, n) < a
		}
		if !ok {
			// 没有合适的a，交换到这里来
			return false
		}
		update(pos[a][it[a]-1], a)
	}

	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
