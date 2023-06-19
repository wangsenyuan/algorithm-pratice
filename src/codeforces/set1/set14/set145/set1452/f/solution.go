package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	cnt := readNInt64s(reader, n)
	Q := make([][]int64, m)
	for i := 0; i < m; i++ {
		Q[i] = readNInt64s(reader, 3)
	}
	res := solve(cnt, Q)
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(cnt []int64, Q [][]int64) []int64 {
	var res []int64
	n := len(cnt)
	for _, qu := range Q {
		pos, val := int(qu[1]), qu[2]
		if qu[0] == 1 {
			cnt[pos] = val
			continue
		}
		var small, cur int64
		for i := 0; i <= pos; i++ {
			small += cnt[i] * (int64(1<<i) - 1)
			val -= cnt[i]
		}
		if val <= 0 {
			res = append(res, 0)
			continue
		}
		id := pos + 1
		for id < n {
			add := int64(1 << (id - pos))
			need := min(val/add, cnt[id])
			cur += need * (add - 1)
			val -= need * add
			small += need * add * int64((1<<pos)-1)
			if need == cnt[id] {
				id++
			} else {
				break
			}
		}
		if val <= 0 {
			res = append(res, cur)
			continue
		}
		if id >= n {
			if val > small {
				res = append(res, -1)
			} else {
				res = append(res, cur+val)
			}
			continue
		}
		var ans int64 = math.MaxInt64

		for id > pos {
			if small >= val {
				ans = min(ans, cur+val)
			}
			cur++
			id--
			add := int64(1 << (id - pos))
			if val >= add {
				cur += add - 1
				val -= add
				small += add * int64((1<<pos)-1)
			}
		}
		res = append(res, min(cur, ans))
	}

	return res
}

type Num interface {
	int | int64
}

func min[T Num](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
