package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(a, Q)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

const inf = 1 << 30

func solve(a []int, Q [][]int) []int {
	// 如果一个序列存在 a[i] >= a[i+1] >= a[i+2]
	// 那么就是一个不好的序列
	// 假设出现了，那是不是删除掉最后一个，就是最优的方案呢？
	// 假设在这个序列中出现一个连续的递减序列，长度为l
	// 如果 l >= 3, 那是不是删到长度为2就好了？只保留最大值和最小值？
	// 也就是这里可以把整个数组分段
	// 类似 (i, x) 表示从i开始的，长度为x的递增数组，
	//     (i, x) 表示从i开始的长度为x的递减数组
	// 它们交替进行
	// 找到l所在的区间，如果它是递增的，找到下一个递减的
	n := len(a)
	ps := make([]Pair, n+1)
	var top int
	// 连续的递减数组，只能保留最大值和最小值
	for i := 0; i < n; i++ {
		for top > 1 && ps[top-2].first >= ps[top-1].first && ps[top-1].first >= a[i] {
			top--
		}
		ps[top] = Pair{a[i], i}
		top++
	}
	//
	//ps[top] = Pair{inf, n}
	//top++

	ans := make([]int, len(Q))

	for i, cur := range Q {
		l, r := cur[0]-1, cur[1]-1
		j := search(top, func(j int) bool {
			return ps[j].second > r
		})
		k := search(top, func(k int) bool {
			return ps[k].second >= l
		})
		ans[i] = j - k
		if r > ps[j-1].second {
			ans[i]++
		}
		if l != r && ps[k].second > l && k > 0 && ps[k-1].second < l {
			ans[i]++
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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

type Pair struct {
	first  int
	second int
}
