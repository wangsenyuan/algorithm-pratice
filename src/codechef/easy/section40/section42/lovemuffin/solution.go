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

	k, q := readTwoNums(reader)

	Q := make([][]int64, q)
	for i := 0; i < q; i++ {
		Q[i] = readNInt64s(reader, 2)
	}
	res := solve(k, Q)

	var buf bytes.Buffer

	for i := 0; i < q; i++ {
		if res[i] < 0 {
			buf.WriteString("Thrown off the roof.\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const M = 1000000000000000000

func solve(k int, Q [][]int64) []int64 {
	var a []int64

	K := int64(k)

	var x int64 = 1
	for x <= M {
		a = append(a, x)
		p := K * (x / (K - 1))
		q := K*(x%(K-1)) + 2*K - 1
		x = p + q/(K-1)
	}

	res := make([]int64, len(Q))

	for i, cur := range Q {
		n, m := cur[0], cur[1]
		if K == 2 && n <= 3 {
			if n == 1 || n == 3 {
				res[i] = m
			} else if n == 2 {
				res[i] = -1
			}
			continue
		}
		var y int64
		if k == 2 && n%2 == 1 {
			y = 1
		}
		if m >= n/K+y {
			res[i] = m - n/K - y
			continue
		}
		tmp := n - m*K - K
		j := sort.Search(len(a), func(i int) bool {
			return a[i] >= tmp
		})
		if j < len(a) && a[j] == tmp {
			res[i] = 0
		} else {
			res[i] = -1
		}
	}

	return res
}
