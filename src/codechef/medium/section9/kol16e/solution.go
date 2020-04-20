package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for i := 1; i <= tc; i++ {
		P := readNum(reader)
		ans := solve(P)
		if ans%2 == 0 {
			fmt.Printf("Case %d: %d-altersum\n", i, ans/2)
		} else {
			fmt.Printf("Case %d: %d-sum\n", i, ans/2)
		}
	}
}

func solve(P int) int {
	p := int64(P - 1)
	var i int64 = 1
	for i*i <= p {
		if p%i == 0 {
			j := p / i
			if pow(10, int(i), P) == 1 {
				return int(i)
			}
			if pow(10, int(j), P) == 1 {
				return int(j)
			}
		}
		i++
	}

	return int(i)
}

func pow(a, b, mod int) int {
	A := int64(a)
	var R int64 = 1
	M := int64(mod)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % M
		}
		A = (A * A) % M
		b >>= 1
	}
	return int(R)
}
