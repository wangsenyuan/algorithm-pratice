package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	B, C := solve(n, A)

	var buf bytes.Buffer
	var buf2 bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", B[i]))
		buf2.WriteString(fmt.Sprintf("%d ", C[i]))
	}
	fmt.Println(buf.String())
	fmt.Println(buf2.String())
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

const MAX_A = 10000003

var minDiv [MAX_A]int

func init() {
	for x := 2; x < MAX_A; x++ {
		if minDiv[x] == 0 {
			for y := x; y < MAX_A; y += x {
				if minDiv[y] == 0 {
					minDiv[y] = x
				}
			}
		}
	}
}

func solve(n int, A []int) ([]int, []int) {
	B := make([]int, n)
	C := make([]int, n)

	for i := 0; i < n; i++ {
		fs := make(map[int]int)

		num := A[i]

		for num > 1 {
			x := minDiv[num]
			fs[x]++
			num /= x
		}
		if len(fs) == 1 {
			B[i] = -1
			C[i] = -1
			continue
		}
		for k := range fs {
			if B[i] == 0 {
				B[i] = k
				C[i] = 1
			} else {
				C[i] *= k
			}
		}
	}
	return B, C
}
