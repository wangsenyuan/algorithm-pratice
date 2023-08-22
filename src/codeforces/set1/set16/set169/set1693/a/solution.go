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

		if res {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(A []int) bool {
	// sum of A = 0
	// A[0] 只能增加， A[n-1] 只能减少
	n := len(A)
	// 考虑 [2, -1, -1, 0]
	// 考虑 A[0] = 2
	// 它表示的是，至少从 0 向右两次，
	// 所以，对于 1来说，相当于向左一次 + 最后一次返回，它至少有两次向左,至少是 -2
	// 因为 -1 - (-2) = 1, 说明它向右移动了一次
	//  同样的对于2来说，它需要至少向左移动1次
	// -1 - （-1） = 0， 说明它不需要在向右移动
	prev := int64(A[0])
	it := 0
	for prev > 0 && it+1 < n {
		cur := int64(A[it+1])
		// 只向左边移动时，-prev，不能比它更小，否则向左移动太多次了
		if cur < -prev {
			return false
		}
		// cur >= -prev
		cur += prev
		prev = cur
		it++
	}

	if prev != 0 {
		return false
	}

	for i := it + 1; i < n; i++ {
		if A[i] != 0 {
			return false
		}
	}

	return true
}
