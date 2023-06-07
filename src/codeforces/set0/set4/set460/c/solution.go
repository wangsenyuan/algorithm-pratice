package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m, w := readThreeNums(reader)
		A := readNNums(reader, n)
		res := solve(A, w, m)
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

const INF = 1e9 + 1e5 + 10

func solve(A []int, w int, m int) int {
	n := len(A)
	diff := make([]int, n+1)

	check := func(x int) bool {
		for i := 0; i < n; i++ {
			diff[i] = 0
		}
		tmp := m
		for i := 0; i < n; i++ {
			if i > 0 {
				diff[i] += diff[i-1]
			}
			cur := A[i] + diff[i]
			if cur < x {
				if tmp == 0 {
					return false
				}
				tmp -= x - cur

				if tmp < 0 {
					return false
				}

				diff[i] += x - cur
				j := min(n, i+w)
				diff[j] -= x - cur
			}
		}
		return true
	}
	var l, r int = min(A...), INF

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

func min(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < res {
			res = arr[i]
		}
	}
	return res
}
