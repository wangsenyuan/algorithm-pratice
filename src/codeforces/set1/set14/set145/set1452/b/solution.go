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
		buf.WriteString(fmt.Sprintf("%d\n", solve(A)))
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

func solve(A []int) int64 {
	n := len(A)
	if n <= 2 {
		return 0
	}
	// (sum - A[i]) % (n - 1) = 0
	// (sum - A[i] )/ (n - 1) >= max(A/i)
	// 要取最大值
	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(A[i])
	}

	arr := make([]int, 2*n)
	copy(arr[n:], A)
	for i := n - 1; i > 0; i-- {
		arr[i] = max(arr[2*i], arr[2*i+1])
	}

	get := func(l int, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}
	var ans int64
	for i := 0; i < n; i++ {
		avg := (sum + int64(n-2)) / int64(n-1)
		x := max(get(0, i), get(i+1, n))
		if avg < int64(x) {
			avg = int64(x)
		}

		need := avg*int64(n-1) - sum

		if ans < need {
			ans = need
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
