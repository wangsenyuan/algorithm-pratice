package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, m)

	fmt.Println(solve(n, A, m, B))
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

func solve(n int, A []int, m int, B []int) int {
	//if n < m {
	//	return 0
	//}
	reverse(A)
	reverse(B)

	cur := A[0]
	var pos int

	for pos < n && cur > B[0] {
		pos++
		if pos < n {
			cur = min(cur, A[pos])
		} else {
			cur = -1
		}
	}

	if pos == n || cur < B[0] {
		return 0
	}

	var res int64 = 1
	var ii = 0

	for {
		if ii == m-1 {
			if minElement(A[pos:]) != B[ii] {
				return 0
			}
			break
		}
		var f = true
		npos := pos

		for npos < n && cur != B[ii+1] {
			npos++
			if npos < n {
				cur = min(cur, A[npos])
			} else {
				cur = -1
			}
			if f && cur < B[ii] {
				f = false
				res *= int64(npos - pos)
				res %= MOD
			}
		}

		if npos == n || cur != B[ii+1] {
			return 0
		}
		ii++
		pos = npos
	}

	return int(res)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minElement(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = min(res, arr[i])
	}
	return res
}
