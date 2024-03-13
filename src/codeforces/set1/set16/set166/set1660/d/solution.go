package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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
func solve(a []int) []int {
	n := len(a)
	// best[2] = 2的数量

	best := []int{n, 0, 0}
	//n := len(a)

	for i := 0; i < n; {
		if a[i] == 0 {
			i++
			continue
		}
		// a[i] != 0
		j := i
		for i < n && a[i] != 0 {
			i++
		}
		tmp := solve1(a[j:i])
		// tmp[2] 表示2的数量
		if tmp[2] > best[2] {
			// better
			best[0] = j + tmp[0]
			best[1] = n - i + tmp[1]
			best[2] = tmp[2]
		}
	}

	return best[:2]
}

func solve1(a []int) []int {
	// a不包含0
	n := len(a)
	// 如果要删除两头，貌似也只需要删除一头就可以了，而且就是某个负数的位置
	var cnt int
	// dp[i]表示到目前为止有多少个2
	res := make([]int, 3)
	res[1] = n
	res[2] = 0
	var neg int
	for i := 0; i < n; i++ {
		if a[i] < 0 {
			neg ^= 1
		}
		cnt += checkValue(abs(a[i]) == 2)
		if neg == 0 {
			// 把i后面的都删除掉
			res[1] = n - i - 1
			res[2] = cnt
		}
	}

	// 倒过来算一次即可
	neg = 0
	cnt = 0
	for i := n - 1; i >= 0; i-- {
		if a[i] < 0 {
			neg ^= 1
		}
		cnt += checkValue(abs(a[i]) == 2)
		if neg == 0 && cnt > res[2] {
			// 把i前面的都删除掉
			res[0] = i
			res[1] = 0
			res[2] = cnt
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func checkValue(b bool) int {
	if b {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
