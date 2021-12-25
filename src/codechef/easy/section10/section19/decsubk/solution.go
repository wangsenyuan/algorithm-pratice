package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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
		n, K := readTwoNums(reader)
		A := readNNums(reader, n)
		B := solve(A, K)
		if len(B) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", B[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
}

func solve(A []int, K int) []int {
	sort.Ints(A)
	if K == len(A) {
		return A
	}
	reverse(A)
	sp := make([]int, len(A))

	if lis(A, sp) > K {
		return nil
	}

	B := make([]int, 0, len(A))

	for len(A) > 0 {
		i := len(A) - 1
		for ; i >= 0; i-- {
			tmp := B
			tmp = append(tmp, A[i])
			for j := 0; j < len(A); j++ {
				if j != i {
					// non-incr order
					tmp = append(tmp, A[j])
				}
			}
			if lis(tmp, sp) <= K {
				break
			}
		}
		B = append(B, A[i])
		copy(A[i:], A[i+1:])
		A = A[:len(A)-1]
	}

	return B
}

func lis(arr []int, sp []int) int {
	var p int
	for i := 0; i < len(arr); i++ {
		j := search(p, func(j int) bool {
			return sp[j] > arr[i]
		})
		sp[j] = arr[i]
		if j == p {
			p++
		}
	}
	return p
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
