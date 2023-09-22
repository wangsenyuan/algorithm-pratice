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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(a []int) int {
	n := len(a)
	// 第一个观测是，被打破的两块区域的距离不会超过2
	//  (超过2的单独处理）
	// 如果是a, b => 最快的话
	// 当 a > b时，攻击a,
	first := 0
	second := 0
	for i := 1; i < n; i++ {
		if a[first] > a[i] {
			second = first
			first = i
		} else if a[second] > a[i] {
			second = i
		}
	}
	ans := (a[first]+1)/2 + (a[second]+1)/2

	process := func(x, y int) int {
		// 假设对a 进行了x次操作
		// 对b进行了y次操作
		// 2 * x + y == a
		// x + 2 * y == b
		// 30, 10 => 15
		// 30, 20 => 15 + 3 = 18
		if x < y {
			x, y = y, x
		}
		cnt := min(x-y, (x+1)/2)
		cur := cnt
		x -= 2 * cnt
		y -= cnt
		if x > 0 && y > 0 {
			cur += (x + y + 2) / 3
		}
		return cur
	}

	process2 := func(a, b, c int) int {
		return (a + c + 1) / 2
	}

	for i := 0; i+1 < n; i++ {
		// try destroy i, and i + 1
		tmp := process(a[i], a[i+1])
		ans = min(tmp, ans)
		if i > 0 {
			tmp = process2(a[i-1], a[i], a[i+1])
			ans = min(tmp, ans)
		}
	}

	return ans
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
