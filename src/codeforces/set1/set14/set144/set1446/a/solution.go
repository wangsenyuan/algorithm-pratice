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
		first := readNInt64s(reader, 2)
		n := int(first[0])
		W := first[1]
		A := readNNums(reader, n)
		res := solve(A, W)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
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

func solve(items []int, W int64) []int {
	// sum(picked) >= W / 2 and sum(picked) <= W

	type Item struct {
		weight int
		id     int
	}
	n := len(items)
	its := make([]Item, n)

	for i := 0; i < n; i++ {
		its[i] = Item{items[i], i}
	}

	sort.Slice(its, func(i, j int) bool {
		return its[i].weight < its[j].weight
	})

	for n > 0 && int64(its[n-1].weight) > W {
		n--
	}
	if n == 0 {
		return nil
	}
	// 如果所有的weight  > W, 那么就没有答案
	// 或者 sum(weight) < W / 2, 也没有答案
	// sum(weight) >= W / 2, 如果sum(weight) > W
	half := (W + 1) / 2
	var sum int64
	for i := n - 1; i >= 0; i-- {
		if int64(its[i].weight) >= half {
			return []int{its[i].id + 1}
		}
		sum += int64(its[i].weight)
	}
	// 所有的weight 都 < half
	if sum < half {
		return nil
	}
	if sum <= W {
		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = its[i].id + 1
		}
		return res
	}
	// sum > W, needs to de-select some to get a smaller one, 因为所有的都 < half, 应该放弃大的部分
	for i := n - 1; i > 0; i-- {
		sum -= int64(its[i].weight)
		if sum <= W {
			res := make([]int, i)
			for j := 0; j < i; j++ {
				res[j] = its[j].id + 1
			}
			return res
		}
	}
	return nil
}
