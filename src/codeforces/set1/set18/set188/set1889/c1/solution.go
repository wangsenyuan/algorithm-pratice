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
		n, m, _ := readThreeNums(reader)
		days := make([][]int, m)
		for i := 0; i < m; i++ {
			days[i] = readNNums(reader, 2)
		}
		res := solve(n, days)
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

func solve(n int, days [][]int) int {
	m := len(days)
	add := make([][]int, n+2)
	del := make([][]int, n+2)

	update := func(arr [][]int, p int, v int) {
		if arr[p] == nil {
			arr[p] = make([]int, 0, 1)
		}
		arr[p] = append(arr[p], v)
	}

	for i, day := range days {
		l, r := day[0], day[1]
		update(add, l, i)
		update(del, r+1, i)

	}

	var dry int

	freq1 := make([]int, m+1)
	freq2 := make(map[Key]int)

	ban := make([]bool, m+1)
	var cur []int
	var cnt int
	for i := 1; i <= n; i++ {
		for _, j := range add[i] {
			cur = append(cur, j)
			cnt++
		}
		for _, j := range del[i] {
			ban[j] = true
			cnt--
		}
		if cnt > 2 {
			// too many
			continue
		}
		if cnt == 0 {
			dry++
			cur = cur[:0]
			continue
		}
		// cnt <= 2
		var tmp []int
		for _, j := range cur {
			if !ban[j] {
				tmp = append(tmp, j)
			}
			if len(tmp) == cnt {
				break
			}
		}
		if len(tmp) == 1 {
			freq1[tmp[0]]++
		} else {
			freq2[Key{min(tmp[0], tmp[1]), max(tmp[0], tmp[1])}]++
		}

		cur = tmp
	}

	count := func(k Key) int {
		return freq2[k] + freq1[k.first] + freq1[k.second]
	}

	var first, second int

	for _, v := range freq1 {
		if v > first {
			second = first
			first = v
		} else if v > second {
			second = v
		}
	}

	var best int
	for k := range freq2 {
		best = max(best, count(k))
	}

	best = max(best, first+second)

	return best + dry
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

type Key struct {
	first  int
	second int
}
