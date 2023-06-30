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
		A := readNNums(reader, n)
		m := readNum(reader)
		Q := make([]int64, m)
		for i := 0; i < m; i++ {
			Q[i] = readNInt64s(reader, 1)[0]
		}
		res := solve(A, Q)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}

	fmt.Print(buf.String())
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

const MAX_A = 1_000_010

var spf [MAX_A]int

func init() {
	for i := 2; i < MAX_A; i++ {
		if spf[i] == 0 {
			spf[i] = i
			if MAX_A/i < i {
				continue
			}
			for j := i * i; j < MAX_A; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}
}

func solve(A []int, Q []int64) []int {
	freq := make(map[int]int)

	for _, a := range A {
		tmp := 1
		for a > 1 {
			if tmp%spf[a] == 0 {
				tmp /= spf[a]
			} else {
				tmp *= spf[a]
			}
			a /= spf[a]
		}
		freq[tmp]++
	}

	var a int
	var b int
	for k, v := range freq {
		if k == 1 || v%2 == 0 {
			b += v
		}
		a = max(a, v)
	}

	ans := make([]int, len(Q))

	for i, w := range Q {
		if w == 0 {
			ans[i] = a
		} else {
			ans[i] = max(a, b)
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
