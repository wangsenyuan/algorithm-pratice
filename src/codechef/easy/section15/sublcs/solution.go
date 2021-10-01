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
		A := readNNums(reader, n)
		res := solve(n, A)
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

func solve(n int, A []int) int {
	B := make([]int, n)
	copy(B, A)
	ILR, ILRC := longestNonDecreasing(B, false)
	IRL, IRLC := longestNonDecreasing(B, true)
	copy(B, A)

	for i := 0; i < n; i++ {
		B[i] *= -1
	}
	DLR, DLRC := longestNonDecreasing(B, false)
	DRL, DRLC := longestNonDecreasing(B, true)

	a := maxOf(ILR)
	b := maxOf(DLR)

	lnds := make([]int, n+1)
	dnds := make([]int, n+1)

	for i := 0; i < n; i++ {
		if ILR[i]+DRL[i]-1 == a {
			lnds[A[i]] = max(lnds[A[i]], ILRC[i]+DRLC[i]-1)
		}
		if DLR[i]+IRL[i]-1 == b {
			dnds[A[i]] = max(dnds[A[i]], IRLC[i]+DLRC[i]-1)
		}
	}

	var ans int

	for i := 1; i <= n; i++ {
		ans = max(ans, min(lnds[i], dnds[i]))
	}

	return ans
}

func min(a, b int) int {
	return a + b - max(a, b)
}

func maxOf(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = max(res, arr[i])
	}
	return res
}

func longestNonDecreasing(arr []int, rev bool) ([]int, []int) {
	if rev {
		reverse(arr)
	}
	n := len(arr)
	stack := make([]int, n+1)
	for i := 1; i <= n; i++ {
		stack[i] = math.MaxInt32
	}
	stack[0] = math.MinInt32

	I := make([]int, n)
	IC := make([]int, n)
	for i := 0; i < n; i++ {
		IC[i] = 1
	}
	last := make(map[int]int)

	for i := 0; i < n; i++ {
		j := search(n+1, func(j int) bool {
			return stack[j] > arr[i]
		})
		stack[j] = arr[i]

		I[i] = j

		if v, found := last[arr[i]]; found && I[v] == j-1 {
			IC[i] = IC[v] + 1
		}
		last[arr[i]] = i
	}

	if rev {
		reverse(I)
		reverse(IC)
	}

	return I, IC
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
