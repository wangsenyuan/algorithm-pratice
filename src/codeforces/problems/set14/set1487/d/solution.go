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
		res := solve(n)
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

//const MAX_N = 1000000007
//
//var cc []int
//
//func init() {
//	// c = b + 1
//	// a * a = 2 *b + 1
//
//	for a := 3; a <= (MAX_N*2)/a; a += 2 {
//		x := a * a
//		if x%2 == 1 {
//			b := (x - 1) / 2
//			c := b + 1
//			if c < MAX_N && right(a, b, c) {
//				cc = append(cc, c)
//			}
//		}
//	}
//	sort.Ints(cc)
//}

func right(a, b, c int) bool {
	A, B, C := int64(a), int64(b), int64(c)
	return A*A+B*B == C*C
}

func solve(n int) int {
	if n < 5 {
		return 0
	}
	x := int(math.Sqrt(float64(2*n) - 1))

	if x*x < 2*n-1 {
		if (x+1)*(x+1) == 2*n-1 {
			x++
		}
	}

	if x%2 == 0 {
		x--
	}

	return (x+1)/2 - 1
}
