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
		crickets := readNNums(reader, 6)
		x, y := readTwoNums(reader)
		res := solve(n, crickets, x, y)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
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

func solve(n int, crickets []int, x int, y int) bool {
	// every move, what not change?
	// 把board分成两个颜色， 黑色和白色，
	// 在黑色格子的cricket，始终在黑色，白色的始终在白色
	// 如果黑色的有两个，所有黑色的位置都可以到达
	// 但是白色的，只能隔一个才能到达（相近的不能到达）
	// 这里还有个问题，移动的方向，其实是有限制的
	// 如果正好在角落里，那么它只能沿着两条边运行
	center := findCenter(crickets)

	if (center[0] == 1 || center[0] == n) && (center[1] == 1 || center[1] == n) {
		return x == center[0] || y == center[1]
	}

	if (center[0]+center[1])%2 == (x+y)%2 {
		return center[0]%2 == x%2
	}
	return true
}

func findCenter(crickets []int) []int {
	x := crickets[0]
	if crickets[0] != crickets[2] {
		x = crickets[4]
	}
	y := crickets[1]
	if crickets[1] != crickets[3] {
		y = crickets[5]
	}
	return []int{x, y}
}

type Item struct {
	x int
	y int
}
