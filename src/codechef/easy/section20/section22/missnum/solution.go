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
		X := readNNums(reader, 4)
		res := solve(X)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
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

func solve(X []int) []int {
	// a + b, a - b, a * b, a / b
	// ops := []int{0, 1, 2, 3}

	check := func(i, j int) []int {
		// x is for addition, y is for substraction
		tmp := X[i] + X[j]
		if tmp%2 == 1 {
			return nil
		}
		a := tmp / 2
		b := X[i] - a
		if a <= 0 || a > 10000 || b <= 0 || b > 10000 {
			return nil
		}

		for l := 0; l < 4; l++ {
			if l == i || l == j {
				continue
			}
			if a*b == X[l] {
				k := 6 - i - j - l
				if a/b == X[k] {
					return []int{a, b}
				}
			}
		}

		return nil
	}

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if x == y {
				continue
			}
			res := check(x, y)
			if len(res) > 0 {
				return res
			}
		}
	}
	return []int{-1, -1}
}
