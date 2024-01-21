package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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
		m := readNum(reader)
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, queries)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
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

func solve(a []int, queries [][]int) []int {
	//n := len(a)
	m := len(queries)
	// a[i] + a[j] = b
	// a[i] * a[j] = c
	// x**2 - b * x + c = 0
	// x = (b +/- sqrt(D)) / 2
	// D = b * b - 4 * a * c

	freq := make(map[int]int)

	for _, num := range a {
		freq[num]++
	}

	ans := make([]int, m)

	for i, cur := range queries {
		b, c := cur[0], cur[1]
		D := b*b - 4*c
		if D < 0 {
			continue
		}
		if D > 0 {
			d := int(math.Sqrt(float64(D)))
			for j := 0; j < 2; j++ {
				if (d+j)*(d+j) == D {
					d += j
					break
				}
			}
			if d*d != D || (b+d)%2 != 0 {
				// no roots
				continue
			}
			x1 := (b + d) / 2
			x2 := (b - d) / 2
			ans[i] = freq[x1] * freq[x2]
		} else {
			x := b / 2
			y := freq[x]
			ans[i] = y * (y - 1) / 2
		}
	}

	return ans
}
