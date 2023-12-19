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
		n, m := readTwoNums(reader)
		l := readNNums(reader, n)
		res := solve(n, m, l)
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

func solve(n int, m int, arr []int) int {
	// x[i] == -1 => seat from right
	// x[i] == -2 => seat from left
	// x[i] > 0 => take the seat only
	// 一时不知所措了
	// 假设我们先把 x[i] > 0 的填上去，然后再处理 x[i] < 0 的部分，会怎么样？
	// 假设 x[i] = -1 的很多，我们希望它们占据足够多的部分，可以选择最右边的先填上去
	// 当然它后面的也可以用x[i] = -1的填上去（但这里假设还没有）
	// 所以必然要从中间一个地方开始
	cnt := make([]int, 2)
	var pos []int
	for i := 0; i < n; i++ {
		if arr[i] == -1 {
			cnt[0]++
		} else if arr[i] == -2 {
			cnt[1]++
		} else {
			pos = append(pos, arr[i])
		}
	}
	// remove duplicates
	sort.Ints(pos)
	var it int
	for i := 1; i <= len(pos); i++ {
		if i == len(pos) || pos[i] > pos[i-1] {
			pos[it] = pos[i-1]
			it++
		}
	}
	pos = pos[:it]
	best := min(m, max(cnt[0], cnt[1]))

	for i := 0; i < it; i++ {
		p := pos[i]
		// 如果p是第一个放进去的
		// p的前面一共有i个位置 + -1的部分
		tmp := i + min(p-1-i, cnt[0])
		tmp += it - i - 1 + min(m-p-(it-i-1), cnt[1])
		tmp++
		best = max(best, tmp)
	}

	// -1从最右边开始
	best = max(best, it+min(m-it, cnt[0]))
	// -2从最左边开始
	best = max(best, it+min(m-it, cnt[1]))

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
