package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		m := readNum(reader)
		C := readNNums(reader, m)
		n := readNum(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, n)
		}
		res := solve(C, A)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}

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

const MOD = 998244353

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	A := int64(a) * int64(b)
	return int(A % MOD)
}

func solve(C []int, A [][]int) bool {
	n := len(A)
	D := make([]int, n)
	for t := 0; t < 5; t++ {
		B := make([]int, n)

		for i := 0; i < n; i++ {
			B[i] = rand.Intn(MOD)
			D[i] = 0
		}

		for i := 0; i < len(C); i++ {
			for j := 0; j < n; j++ {
				D[j] = modAdd(D[j], modMul(C[i], B[j]))
			}
			B = matMul(A, B)
		}
		for j := 0; j < n; j++ {
			if D[j] != 0 {
				return false
			}
		}
	}
	return true
}

func matMul(A [][]int, B []int) []int {
	n := len(B)
	C := make([]int, n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			C[r] = modAdd(C[r], modMul(A[r][c], B[c]))
		}
	}
	return C
}
