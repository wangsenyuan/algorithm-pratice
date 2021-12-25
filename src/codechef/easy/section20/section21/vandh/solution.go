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
		n, m := readTwoNums(reader)

		res := solve(n, m)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func solve(n, m int) string {
	if n >= m {
		return create(n, m)
	}
	res := create(m, n)
	return flip(res)
}

func flip(s string) string {
	buf := []byte(s)
	for i := 0; i < len(buf); i++ {
		if buf[i] == '0' {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}

	return string(buf)
}

func create(n, m int) string {
	// n >= m
	// vally >= hill
	// first and last are always 0
	var res []int
	if n > m {
		res = make([]int, 2*n-1+2)
	} else {
		// n == m
		res = make([]int, 2*n+2)
	}

	for i := 1; i < len(res); i += 2 {
		res[i] = 1
	}
	var cnt int
	for i := 1; i+1 < len(res); i++ {
		if res[i-1] > res[i] && res[i] < res[i+1] {
			cnt++
		}
	}
	pos := 2
	for cnt > m {
		res[pos] = 2
		pos += 2
		cnt--
	}

	var buf []byte

	buf = append(buf, '0')

	for i := 1; i < len(res); i++ {
		if res[i] == 1 {
			buf = append(buf, '1')
		} else if res[i] == 0 {
			buf = append(buf, '0')
		} else {
			// res[i] == 2
			buf = append(buf, '0')
			buf = append(buf, '0')
		}
	}

	return string(buf)
}
