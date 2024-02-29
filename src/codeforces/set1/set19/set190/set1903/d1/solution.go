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
	q := make([]int, m)
	for i := 0; i < m; i++ {
		q[i] = readNum(reader)
	}

	res := solve(a, q)

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

const H = 60

func solve(a []int, q []int) []int {
	n := len(a)

	var rs int
	for i := 0; i < n; i++ {
		rs += (1 << 20) - a[i]
	}

	process := func(expect int, num int, d int) int {
		// 高位要和expect一致
		// (num + cnt) & expect == expect
		// 最小的cnt
		var res int
		for i := H - 1; i >= d; i-- {
			x := (expect >> i) & 1
			y := (num >> i) & 1

			if x == 0 || y == 1 {
				continue
			}
			// 101
			// 011
			// x == 1 && y == 0
			res += (1 << i) - (num & ((1 << i) - 1))
			// 剩下的部分都变成了0
			num = 0
		}
		return res
	}

	check := func(expect int, d int, k int) int {
		var cnt int
		for i := 0; i < n; i++ {
			if (a[i]>>d)&(expect>>d) == (expect >> d) {
				continue
			}
			// num & expect != expect
			cnt += process(expect, a[i], d)
			if cnt > k {
				break
			}
		}
		return cnt
	}

	ans := make([]int, len(q))

	for j, k := range q {
		if k >= rs {
			ans[j] = (1 << 20) + (k-rs)/n
			continue
		}

		var res int
		for i := H - 1; i >= 0; i-- {
			tmp := check(res|(1<<i), i, k)
			if k >= tmp {
				res |= 1 << i
			}
		}
		ans[j] = res
	}

	return ans
}
