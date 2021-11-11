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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
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

const H = 30

func solve(A []int) []int {
	n := len(A)
	L := make([][]int, n)
	R := make([][]int, n)
	for i := 0; i < n; i++ {
		L[i] = make([]int, H)
		R[i] = make([]int, H)
	}

	for h := 0; h < H; h++ {
		for i := 0; i < n; i++ {
			if (A[i]>>uint(h))&1 == 1 {
				L[i][h] = 1
				if i > 0 {
					L[i][h] += L[i-1][h]
				}
			}
		}
		for i := n - 1; i >= 0; i-- {
			if (A[i]>>uint(h))&1 == 1 {
				R[i][h] = 1
				if i+1 < n {
					R[i][h] += R[i+1][h]
				}
			}
		}
	}

	res := make([]int, n)

	for i := 0; i < n; i++ {
		a, b := 0, n-1
		var cur int
		for h := H - 1; h >= 0; h-- {
			if (A[i]>>uint(h))&1 == 0 {
				// no contribution
				continue
			}
			l, r := i-L[i][h]+1, i+R[i][h]-1

			// if equal, means size = 1
			if max(a, l) < min(b, r) {
				cur |= 1 << uint(h)
				a = max(a, l)
				b = min(b, r)
			}
		}
		res[i] = cur
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
