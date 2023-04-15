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
		n := readNum(reader)
		s := readString(reader)[:n]
		t := readString(reader)[:n]
		res := solve(s, t)
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

const INF = 1 << 30

func solve(s string, t string) int64 {
	n := len(s)

	pos := make([][]int, 26)

	addPos := func(x int, i int) {
		if pos[x] == nil {
			pos[x] = make([]int, 0, 1)
		}
		pos[x] = append(pos[x], i)
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		addPos(x, i)
	}

	rem := make([]int, 2*n)

	update := func(p int, v int) {
		p += n
		rem[p] += v
		for p > 1 {
			rem[p>>1] = rem[p] + rem[p^1]
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res += rem[l]
				l++
			}
			if r&1 == 1 {
				r--
				res += rem[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}
	cur := make([]int, 26)
	var best int64 = -1
	var sum int64

	mark := make([]bool, n)

	for i, j := 0, 0; i < n; i++ {
		// same until i, swap smaller ( < t[i]) char from suffix
		x := -1
		for k := 0; k < int(t[i]-'a'); k++ {
			if cur[k] == len(pos[k]) {
				continue
			}
			a := pos[k][cur[k]]
			if a < n {
				// 需要swap的距离等于两者下标的距离 - 中间已经被匹配掉的字符
				b := a - j - get(j, a)
				if x < 0 || x > b {
					x = b
				}
			}
		}
		if x >= 0 {
			tmp := sum + int64(x)
			if best < 0 || best > tmp {
				best = tmp
			}
		}
		k := int(t[i] - 'a')
		// try make s[i] = t[i]
		if cur[k] == len(pos[k]) {
			break
		}
		a := pos[k][cur[k]]
		cur[k]++
		sum += int64(a - j - get(j, a))
		mark[a] = true

		update(a, 1)

		for j < n && mark[j] {
			j++
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
