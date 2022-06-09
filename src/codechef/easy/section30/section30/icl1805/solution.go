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

	n, q := readTwoNums(reader)
	Q := make([][]int, q)

	for i := 0; i < q; i++ {
		s, _ := reader.ReadBytes('\n')
		var t int
		pos := readInt(s, 0, &t) + 1
		if t == 1 {
			Q[i] = make([]int, 3)
			Q[i][0] = 1
		} else {
			Q[i] = make([]int, 4)
			Q[i][0] = 2
		}
		for j := 1; j < len(Q[i]); j++ {
			pos = readInt(s, pos, &Q[i][j]) + 1
		}
	}

	res := solve(n, Q)

	for i := 0; i < q; i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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
func solve(n int, Q [][]int) []int {
	n2 := 1
	var bit int
	for n2 < n {
		bit++
		n2 *= 2
	}

	pre := make([][]int, n2+1)

	for day := 1; day <= n2; day++ {
		pre[day] = gen(n, day, bit)
	}

	get := func(l, r int) int {
		r %= n2
		if r == 0 {
			r = n2
		}
		var lo, hi = 0, len(pre[r]) - 1
		ans := -1
		for lo <= hi {
			mid := (lo + hi) / 2
			if pre[r][mid]+l-1 <= n {
				ans = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		return ans
	}

	ans := make([]int, len(Q))

	for i, cur := range Q {
		if cur[0] == 1 {
			l, r := cur[1], cur[2]
			ans[i] = get(l, r)
			ans[i] = n - ans[i] - 1
		} else {
			l, r, k := cur[1], cur[2], cur[3]
			tmp := get(l, r)
			if k > tmp+1 {
				ans[i] = -1
			} else {
				ans[i] = pre[r][k-1] + l - 1
			}
		}
	}
	return ans
}

func gen(n int, idx int, bit int) []int {
	var res []int
	if bit == 0 {
		res = append(res, 1)
		if idx == 1 {
			res = append(res, 2)
		}
		return res
	}
	mid := 1 << uint(bit)
	if idx <= mid {
		res = gen(n, idx, bit-1)
		sz := len(res)
		for i := 0; i < sz; i++ {
			if res[i]+mid <= n {
				res = append(res, res[i]+mid)
			} else {
				break
			}
		}
	} else {
		res = gen(n, idx-mid, bit-1)
	}
	return res
}
