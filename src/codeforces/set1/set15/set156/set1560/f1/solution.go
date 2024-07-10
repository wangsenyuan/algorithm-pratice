package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		res := solve(n, k)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, k int) string {
	var arr []int
	for num := n; num > 0; num /= 10 {
		arr = append(arr, num%10)
	}
	m := len(arr)

	res := solve1(arr)
	if k == 1 || m == 1 {
		return res
	}

	res = min_num(res, "1"+strings.Repeat("0", m))
	// k == 2
	reverse(arr)
	// 第一个数要们是 arr[0], 要么是arr[0] + 1
	if arr[0]+1 <= 9 {
		res = min_num(res, fmt.Sprintf("%d", arr[0]+1)+strings.Repeat("0", m-1))
	}
	// 然后就是第一数如果是 arr[0], 那么另外一个数是其他的
	// 一个是arr[0], 另外一个是y
	for y := 0; y <= 9; y++ {
		ok, tmp := solve2(min(arr[0], y), max(arr[0], y), arr)
		if ok {
			res = min_num(res, tmp)
		}
	}

	return res
}

func solve2(x int, y int, arr []int) (bool, string) {
	// 只使用x,y是否能得到比arr 表示的number更大的数
	// x <= y
	buf := make([]byte, len(arr))

	for i := 0; i < len(arr); i++ {
		buf[i] = byte(x + '0')
	}
	// arr[0] = x or y already known
	buf[0] = byte(arr[0] + '0')

	m := len(arr)

	var dfs func(i int) bool

	dfs = func(i int) bool {
		if i == m {
			return true
		}
		if arr[i] < x {
			return true
		}
		if arr[i] == x {
			if dfs(i + 1) {
				return true
			}
		}
		buf[i] = byte(y + '0')
		if arr[i] < y {
			return true
		}
		if arr[i] == y {
			if dfs(i + 1) {
				return true
			}
		}
		buf[i] = byte(x + '0')
		return false
	}

	ok := dfs(1)

	if !ok {
		return false, ""
	}

	return true, string(buf)
}

func min_num(a, b string) string {
	if len(a) < len(b) || len(a) == len(b) && a < b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func solve1(arr []int) string {

	m := len(arr)

	r := m - 1
	for r >= 0 && arr[r] == arr[m-1] {
		r--
	}

	if r < 0 || arr[r] < arr[m-1] {
		// 所有的都和arr[m-1相同, 或者第一个不同的，小于 arr[m-1]
		return strings.Repeat(fmt.Sprintf("%d", arr[m-1]), m)
	}
	if arr[m-1]+1 <= 9 {
		return strings.Repeat(fmt.Sprintf("%d", arr[m-1]+1), m)
	}
	// 1111
	return strings.Repeat("1", m+1)
}
