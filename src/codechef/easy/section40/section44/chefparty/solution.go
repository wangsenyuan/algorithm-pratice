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
		sad, res := solve(n)
		buf.WriteString(fmt.Sprintf("%d\n", sad))
		for _, p := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", p[0], p[1]))
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const N = 200010

var lpf []int

func init() {
	lpf = make([]int, N)
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			if N/i < i {
				continue
			}
			for j := i * i; j < N; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func solve(n int) (sad int, res [][]int) {
	m := 2 * n
	var a []int
	a = append(a, 1)
	var b []int
	buf := make([]int, n+1)
	ok := make([]bool, m+1)

	sad = n

	for i := m; i >= 2; i-- {
		if lpf[i] == i {
			// a prime
			if i > n {
				a = append(a, i)
				continue
			}
			var it int
			for j := i; j <= m; j += i {
				if !ok[j] {
					ok[j] = true
					buf[it] = j
					it++
				}
			}

			var x int
			if it%2 == 1 {
				b = append(b, 2*i)
				buf[1] = i
				x = 1
			}

			for ; x < it; x += 2 {
				res = append(res, []int{buf[x], buf[x+1]})
				sad--
			}
		}
	}

	for i := 0; i < len(b); i += 2 {
		if i+1 < len(b) {
			res = append(res, []int{b[i], b[i+1]})
			sad--
		} else {
			a = append(a, b[i])
		}
	}

	for i := 0; i < len(a); i += 2 {
		res = append(res, []int{a[i], a[i+1]})
	}

	return sad, res
}
