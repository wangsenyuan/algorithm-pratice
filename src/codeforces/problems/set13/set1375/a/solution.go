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

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		solve(n, A)
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", A[i]))
		}
		fmt.Println(buf.String())
	}
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

func solve(n int, A []int) {

	for i := 0; i < n; i += 2 {
		if A[i] < 0 {
			A[i] = -A[i]
		}
		if i+1 < n && A[i+1] > 0 {
			A[i+1] = -A[i+1]
		}
	}
}

func solve1(n int, A []int) {
	f := make([]int, n)
	f[0] = 1

	var a, b int
	for i := 1; i < n; i++ {
		f[i] = 1

		if A[i-1] == A[i] {
			continue
		}

		if A[i-1] < A[i] {
			if a > b {
				// 4, 5, 6 => -4, -5, 6
				// -3, -2, -1 => 3, 2, -1
				// -3, -2, 1 => 3, 2, 1
				if -A[i-1] < A[i] {
					// it is ok to just flip A[i-1]
					f[i-1] = -1
					a, b = b, a
				} else {
					// -A[i-1] > A[i], then after flip, b > a, and we need to incrase b, not correct, but we can flip before i - 2
					f[i-2] *= -1
					// to revert A[i-2] > A[i-1]
					a--
					a, b = b, a
					if f[i-2]*A[i-2] < A[i-1] {
						a++
					} else if f[i-2]*A[i-2] > A[i-1] {
						b++
					}
				}
			}
			a++
		} else {
			//A[i-1] > A[i]
			// 6, 5, 4
			// -1, -2, -3
			if b > a {
				if -A[i-1] > A[i] {
					f[i-1] = -1
					a, b = b, a
				} else {
					f[i-2] *= -1
					b--
					a, b = b, a
					if f[i-2]*A[i-2] < A[i-1] {
						a++
					} else if f[i-2]*A[i-2] > A[i-1] {
						b++
					}
				}
			}
			b++
		}
	}

	for i := n - 2; i >= 0; i-- {
		f[i] *= f[i+1]
		A[i] *= f[i]
	}
}
