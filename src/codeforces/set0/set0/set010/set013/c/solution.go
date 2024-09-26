package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

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

const inf = 1 << 60

func solve(a []int) int {

	n := len(a)

	b := make([]int, n)
	copy(b, a)
	sort.Ints(b)

	// f[i][j] = 前i个元素，最后的元素是b[j]时的最小值
	f := make([]int, n)

	for i := 0; i < n; i++ {
		f[i] = abs(a[0] - b[i])
		if i > 0 {
			f[i] = min(f[i], f[i-1])
		}
	}

	for i := 1; i < n; i++ {
		// f[i][j] = abs(a[i] - b[j]) + abs(f[i-1][j])
		// 前i个最大为a[j], 那么前i-1个最大也要为b[j] (而不是b[j-1])
		f[0] += abs(a[i] - b[0])
		for j := 1; j < n; j++ {
			f[j] = min(f[j-1], f[j]+abs(a[i]-b[j]))
		}
	}
	best := inf

	for i := 0; i < n; i++ {
		best = min(best, f[i])
	}

	return best
}

func abs(num int) int {
	return max(num, -num)
}
