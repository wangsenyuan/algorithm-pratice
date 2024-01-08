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
		segs := make([][]int, n)
		for i := 0; i < n; i++ {
			segs[i] = readNNums(reader, 2)
		}
		res := solve(segs)
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
func solve(segs [][]int) int {
	ids := make(map[int]int)
	ids[-1] = 1
	for _, seg := range segs {
		ids[seg[0]]++
		ids[seg[1]]++
	}
	arr := make([]int, 0, len(segs))
	for k := range ids {
		arr = append(arr, k)
	}
	sort.Ints(arr)

	m := len(arr)

	for i := 0; i < m; i++ {
		ids[arr[i]] = i
	}

	dp := make([]int, m)

	for i := 0; i < len(segs); i++ {
		l, r := segs[i][0], segs[i][1]
		segs[i][0] = ids[l]
		segs[i][1] = ids[r]
	}

	sort.Slice(segs, func(i, j int) bool {
		return segs[i][1] < segs[j][1] || segs[i][1] == segs[j][1] && segs[i][0] < segs[j][0]
	})
	j := 1
	for i := 0; i < len(segs); i++ {
		dp[segs[i][1]] = max(dp[segs[i][1]-1], dp[segs[i][1]])
		for k := i - 1; k >= 0; k-- {
			if segs[k][1] < segs[i][0] {
				break
			}
			dp[segs[i][1]] = max(dp[segs[i][1]], dp[min(segs[i][0], segs[k][0])-1]+1)
		}
		for j <= segs[i][1] {
			dp[j] = max(dp[j], dp[j-1])
			j++
		}
	}

	return len(segs) - 2*dp[j-1]
}

type Pair struct {
	first  int
	second int
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
