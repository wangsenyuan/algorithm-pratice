package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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

func solve(A []int) int64 {
	sort.Ints(A)
	n := len(A)
	var i int
	for i < n && A[i] < 0 {
		i++
	}
	if i == n {
		// all less than 0, choose the ends
		res := int64(1)
		for k := 0; k < 5; k++ {
			res *= int64(A[n-1-k])
		}
		return res
	}
	// A[i] >= 0
	j := i
	for j < n && A[j] == 0 {
		j++
	}

	var ans int64

	for _, cnt := range []int{1, 3, 5} {
		if n-j < cnt {
			break
		}
		if 5-cnt > i {
			continue
		}
		// n - j >= cnt
		var cur int64 = 1
		for k := 0; k < 5-cnt; k++ {
			cur *= int64(A[k])
		}
		for k := 0; k < cnt; k++ {
			cur *= int64(A[n-1-k])
		}
		if cur > ans {
			ans = cur
		}
	}

	if ans == 0 {
		return solve1(A)
	}

	return ans
}

func solve1(A []int) int64 {
	sort.Slice(A, func(i, j int) bool {
		return abs(A[i]) < abs(A[j])
	})

	var res int64 = 1
	for i := 0; i < 5; i++ {
		res *= int64(A[i])
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
