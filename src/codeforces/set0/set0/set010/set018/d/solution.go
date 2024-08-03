package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	days := make([]string, n)
	for i := 0; i < n; i++ {
		days[i] = readString(reader)
	}
	res := solve(days)

	fmt.Println(res)
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

const X = 2001

var pre_defined [X][]int

func init() {

	pre_defined[0] = []int{1}
	pre_defined[1] = []int{2}

	for i := 2; i < X; i++ {
		pre_defined[i] = add(pre_defined[i-1], pre_defined[i-1])
	}
}

func solve(days []string) string {

	win := make([]int, X)
	for i := 0; i < X; i++ {
		win[i] = -1
	}
	n := len(days)

	pairs := make([][]int, X)

	for i := 0; i < n; i++ {
		cur := days[i]
		if strings.HasPrefix(cur, "sell") {
			var x int
			readInt([]byte(cur), 5, &x)
			if win[x] >= 0 {
				pairs[x] = []int{win[x], i}
			}
		} else {
			var x int
			readInt([]byte(cur), 4, &x)
			win[x] = i
		}
	}

	var res []int
	var dfs func(l int, r int, best int)

	dfs = func(l int, r int, best int) {
		// 找这个区间内的最大值

		for best >= 0 {
			if len(pairs[best]) > 0 && pairs[best][0] >= l && pairs[best][1] <= r {
				break
			}
			best--
		}

		if best < 0 {
			return
		}
		pair := pairs[best]

		res = add(res, pre_defined[best])
		if pair[0] > l {
			dfs(l, pair[0]-1, best-1)
		}
		if pair[1] < r {
			dfs(pair[1]+1, r, best-1)
		}
	}

	dfs(0, n-1, 2000)

	if len(res) == 0 {
		return "0"
	}

	reverse(res)

	ans := make([]byte, len(res))

	for i := 0; i < len(res); i++ {
		ans[i] = byte(res[i] + '0')
	}

	return string(ans)
}

func add(a, b []int) []int {
	var carry int
	var res []int
	for i := 0; i < len(a) || i < len(b); i++ {
		var x, y int
		if i < len(a) {
			x = a[i]
		}
		if i < len(b) {
			y = b[i]
		}
		cur := x + y + carry
		res = append(res, cur%10)
		carry = cur / 10
	}
	if carry > 0 {
		res = append(res, carry)
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
