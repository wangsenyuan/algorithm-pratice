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
		n, a, b := readThreeNums(reader)
		res := solve(n, a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, a int, b int) int64 {
	if n == 1 {
		return 0
	}
	// n => n / x => n / x
	var odd, even int

	cur := n

	for cur&1 == 0 {
		even++
		cur >>= 1
	}

	if cur > 1 {
		for x := 3; x <= cur/x; x += 2 {
			if cur%x == 0 {
				for cur%x == 0 {
					odd++
					cur /= x
				}
			}
		}
		if cur != 1 {
			odd++
		}
	}
	A := int64(a)
	B := int64(b)
	if a >= 0 && b >= 0 {
		return int64(even)*A + int64(odd)*B
	}
	if a >= 0 && b < 0 {
		if n%2 == 0 {
			return int64(even) * A
		}
		return B
	}
	if a <= 0 && b <= 0 {
		if n%2 == 0 {
			return A
		}
		return B
	}
	// a <= 0 && b > 0
	if n%2 == 0 {
		return int64(odd)*B + A
	}
	return int64(odd) * B
}
