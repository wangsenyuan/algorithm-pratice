package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		ok, res := solve(A)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, vs := range res {
				for i := 0; i < n; i++ {
					buf.WriteString(fmt.Sprintf("%d ", vs[i]))
				}
				buf.WriteByte('\n')
			}
		}
	}

	fmt.Print(buf.String())
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
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

const INF = 1 << 30

func solve(A []int) (bool, [][]int) {
	n := len(A)
	for i := 0; i < n; i++ {
		A[i]--
	}

	var ans [][]int

	buf := make([]int, n)

	apply := func(x []int) {
		ans = append(ans, x)
		for i := 0; i < n; i++ {
			buf[x[(i+1)%n]] = A[x[i]]
		}
		copy(A, buf)
	}

	for step := 0; step < 3; step++ {
		c := cycles(A)
		var num int
		for _, vs := range c {
			if len(vs) > 1 {
				if len(vs)%2 == 0 {
					num = INF
					break
				} else {
					num++
				}
			}
		}
		if num < 2 {
			break
		}
		if step == 2 {
			return false, nil
		}
		b := make([]int, 0, n)
		for _, vs := range c {
			b = append(b, vs...)
		}
		apply(b)
	}

	c := cycles(A)
	var b []int
	for _, vs := range c {
		if len(vs) > 1 {
			b = vs
		}
	}
	reverse(b)

	for _, vs := range c {
		if len(vs) == 1 {
			b = append(b, vs[0])
		}
	}

	apply(b)

	c = cycles(A)
	apply(c[0])

	for _, vs := range ans {
		for i := 0; i < len(vs); i++ {
			vs[i]++
		}
	}

	return true, ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func cycles(A []int) [][]int {
	var res [][]int
	vis := make([]bool, len(A))

	for i := 0; i < len(A); i++ {
		if !vis[i] {
			var cur []int
			for x := i; !vis[x]; x = A[x] {
				cur = append(cur, x)
				vis[x] = true
			}
			res = append(res, cur)
		}
	}
	return res
}
