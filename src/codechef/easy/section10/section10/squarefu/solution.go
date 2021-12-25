package main

import (
	"bufio"
	"bytes"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

const X = 1000010

var F []int
var cnt []int64
var dd []int
var bak []int64

func init() {
	F = make([]int, X)
	cnt = make([]int64, X)
	dd = make([]int, X)
	bak = make([]int64, X)
	for i := 2; i < X; i++ {
		if F[i] == 0 {
			F[i] = i
			for j := int64(i) * int64(i); j < X; j += int64(i) {
				if F[int(j)] == 0 {
					F[int(j)] = i
				}
			}
		}
	}
}

func solve(A []int) int64 {
	n := len(A)
	tot := int64(n) * int64(n-1) / 2

	//sort.Ints(A)

	// A[i] * A[j] = S * S
	// let A[i] = p ** x * q ** y
	// let A[j] = p ** a * w ** z
	// if (x + a) % 2 == 0 & y % 2 == 0 & z % 2 == 0

	copy(cnt, bak)

	for i := 0; i < n; i++ {
		cur := A[i]

		for cur > 1 {
			dd[F[cur]] ^= 1
			cur /= F[cur]
		}

		num := 1
		cur = A[i]

		for cur > 1 {
			if dd[F[cur]] == 1 {
				num *= F[cur]
				dd[F[cur]] = 0
			}
			cur /= F[cur]
		}

		tot -= cnt[num]
		cnt[num]++
	}

	return tot
}
