// https://www.codechef.com/problems/ALTDIA?tab=statement
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
		B, W := readTwoNums(reader)
		ok, s, E := solve(B, W)
		if !ok {
			buf.WriteString("-1\n")
			continue
		}
		buf.WriteString(s)
		buf.WriteByte('\n')
		for _, e := range E {
			buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(B int, W int) (bool, string, [][]int) {
	n := B + W
	if n == 1 {
		if B == 1 {
			return true, "B", nil
		}
		return true, "W", nil
	}
	// total n nodes,
	// 假设长为L的path满足条件
	// L = a + b, abs(a - b) <= 1
	// 在Li处连接边出去， 那么可以使用 min(i-1, L - i) 个节点 (不论黑白)
	if n == 2 {
		if B == 1 {
			return true, "BW", [][]int{{1, 2}}
		}
		return false, "", nil
	}

	// n > 2, 那么就是diameter = 3
	if B == 0 || W == 0 {
		return false, "", nil
	}
	// B > 0
	var buf []byte
	var edges [][]int
	if B > W {
		// put W at mid
		buf, edges = solve1(B, W)
	} else {
		buf, edges = solve1(W, B)
		for i := 0; i < len(buf); i++ {
			if buf[i] == 'B' {
				buf[i] = 'W'
			} else {
				buf[i] = 'B'
			}
		}
	}

	return true, string(buf), edges
}

func solve1(B int, W int) ([]byte, [][]int) {
	n := B + W
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = 'B'
	}
	var edges [][]int

	edges = append(edges, []int{1, 2})
	edges = append(edges, []int{2, 3})
	buf[1] = 'W'
	B -= 2
	W--
	i := 3
	for B > 0 {
		edges = append(edges, []int{2, i + 1})
		buf[i] = 'B'
		i++
		B--
	}
	for W > 0 {
		edges = append(edges, []int{2, i + 1})
		buf[i] = 'W'
		i++
		W--
	}
	return buf, edges
}
