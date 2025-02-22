package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	H := make([][]int, 3)
	H[0] = readNNums(reader, 3)
	H[1] = readNNums(reader, 2)
	H[2] = readNNums(reader, 1)
	res := solve(H)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	fmt.Println(len(res[0]))
	for _, x := range res {
		fmt.Println(x)
	}
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(H [][]int) []string {
	a, b, c := H[0][0], H[0][1], H[0][2]
	d, e := H[1][0], H[1][1]
	f := H[2][0]
	if (f-c+d-a)%2 != 0 || (b-a+f-e)%2 != 0 || (d-c+e-b)%2 != 0 || (c-a+e)%2 != 0 {
		return nil
	}

	write := func(x string, cnt int, res []*bytes.Buffer) {
		for i := 0; i < cnt; i++ {
			for j := 0; j < 4; j++ {
				res[j].WriteByte(x[j])
			}
		}
	}

	for x2 := 0; x2 < 100000; x2++ {
		x7 := x2 - (f-c+d-a)/2
		x4 := x2 - (b-a+f-e)/2
		x1 := x2 - (d-c-e+b)/2
		x3 := (c-a+e)/2 - x1
		x6 := (f-c-(d-a))/2 + x3
		x5 := x3 - (b-a-(f-e))/2

		if x1 < 0 || x2 < 0 || x3 < 0 || x4 < 0 || x5 < 0 || x6 < 0 || x7 < 0 {
			continue
		}
		buf := make([]*bytes.Buffer, 4)
		for i := 0; i < 4; i++ {
			buf[i] = new(bytes.Buffer)
		}
		write("aaab", x1, buf)
		write("aaba", x2, buf)
		write("aabb", x3, buf)
		write("abaa", x4, buf)
		write("abab", x5, buf)
		write("abba", x6, buf)
		write("abbb", x7, buf)
		res := make([]string, 4)
		for i := 0; i < 4; i++ {
			res[i] = buf[i].String()
		}

		return res
	}

	return nil
}
