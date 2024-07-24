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

	n, v := readTwoNums(reader)
	boats := make([][]int, n)
	for i := 0; i < n; i++ {
		boats[i] = readNNums(reader, 2)
	}

	res, arr := solve(v, boats)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", res))
	for i := 0; i < len(arr); i++ {
		buf.WriteString(fmt.Sprintf("%d ", arr[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

type pair struct {
	first  int
	second int
}

func solve(v int, boats [][]int) (int, []int) {
	var kayaks []pair
	var catamaran []pair
	for i, boat := range boats {
		if boat[0] == 1 {
			kayaks = append(kayaks, pair{boat[1], i + 1})
		} else {
			catamaran = append(catamaran, pair{boat[1], i + 1})
		}
	}
	sort.Slice(kayaks, func(i, j int) bool {
		return kayaks[i].first < kayaks[j].first
	})
	sort.Slice(catamaran, func(i, j int) bool {
		return catamaran[i].first < catamaran[j].first
	})

	// 不对， 不能这样处理
	// 应该是在放置后面r个catamaran 的情况下，能够得到的最大值
	m := len(kayaks)
	suf := make([]int, m+1)
	n := len(catamaran)

	res := []int{m, n}

	var best int
	for i := m - 1; i >= 0; i-- {
		suf[i] = kayaks[i].first + suf[i+1]
		if m-i <= v {
			best = suf[i]
			// 从i开始选
			res[0] = i
		}
	}
	var sum int
	for i := n - 1; i >= 0; i-- {
		sum += catamaran[i].first
		rem := v - 2*(n-i)
		if rem < 0 {
			// 超了
			break
		}
		if rem >= m {
			// 全部kayaks都可以放入
			if sum+suf[0] > best {
				best = sum + suf[0]
				res[0] = 0
				res[1] = i
			}
		} else {
			if sum+suf[m-rem] > best {
				best = sum + suf[m-rem]
				res[0] = m - rem
				res[1] = i
			}
		}
	}

	var arr []int
	for i := res[0]; i < m; i++ {
		arr = append(arr, kayaks[i].second)
	}
	for i := res[1]; i < n; i++ {
		arr = append(arr, catamaran[i].second)
	}
	sort.Ints(arr)

	return best, arr
}
