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
		pairs := make([][]int, n)
		for i := 0; i < n; i++ {
			pairs[i] = readNNums(reader, 2)
		}
		res := solve(pairs)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

const inf = 1 << 31

func solve(pairs [][]int) int {
	// (a, b), 必须保留一个，且将其放入两堆中的一个
	// 放入哪个堆其实没关系，只要找出需要保留的那些，从中选择最大的两个，一个放入A，另外一个放入B即可
	// 把a放在为止i，把b放在为止 n + i
	// 然后找这个区间内最大的两个值
	n := len(pairs)

	type Pair struct {
		first  int
		second int
	}
	arr := make([]Pair, 2*n)
	var mx int
	for i := 0; i < n; i++ {
		mx = max(mx, min(pairs[i][0], pairs[i][1]))
		arr[i*2] = Pair{pairs[i][0], i}
		arr[i*2+1] = Pair{pairs[i][1], i}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first < arr[j].first
	})

	ans := inf

	for i := 1; i < 2*n; i++ {
		x := arr[i].first
		if x < mx {
			continue
		}
		// x >= mx
		y := arr[i-1].first
		if arr[i-1].second == arr[i].second {
			if i == 1 {
				continue
			}
			y = arr[i-2].first
		}
		ans = min(ans, x-y)
	}

	return ans
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
