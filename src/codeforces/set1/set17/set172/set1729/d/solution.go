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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		x := readNNums(reader, n)
		y := readNNums(reader, n)
		res := solve(x, y)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(x []int, y []int) int {
	type User struct {
		x  int
		y  int
		id int
	}

	n := len(x)
	users := make([]User, n)
	for i := 0; i < n; i++ {
		users[i] = User{x[i], y[i], i}
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].x-users[i].y < users[j].x-users[j].y
	})

	// 左边是 x - y 最小的，右边是 x - y 最大的，也就是 y - x最小的
	// 只需要考虑2人group
	// x[a] + x[b] <= y[a] + y[b]
	// => x[a] - y[a] <= y[b] - x[b]

	var ans int

	for l, r := 0, n-1; l < r; {
		a := users[l].x - users[l].y
		b := users[r].y - users[r].x
		if a <= b {
			ans++
			l++
		}
		// a > b
		// r 向左移动时，b会变大
		// l 向右移动时，a会变大
		// 移动l更加不会获得答案
		r--
	}
	return ans
}
