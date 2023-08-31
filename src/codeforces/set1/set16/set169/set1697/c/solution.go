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
		n := readNum(reader)
		s := readString(reader)[:n]
		x := readString(reader)[:n]
		res := solve(s, x)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(s string, x string) bool {
	if count(s, 'b') != count(x, 'b') {
		return false
	}
	n := len(s)
	for i, j := 0, 0; i < n; i++ {
		if s[i] == 'b' {
			continue
		}
		for j < n && x[j] == 'b' {
			j++
		}
		// a move left, so the position in x for 'a should be less than position in s
		if s[i] != x[j] || (s[i] == 'a' && i > j) || (s[i] == 'c' && i < j) {
			return false
		}

		j++
	}

	return true
}

func count(s string, x byte) int {
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == x {
			res++
		}
	}
	return res
}

func solve1(s string, x string) bool {
	// ab => ba
	// bc => cb
	// ab+(a|c)
	// aa 不能移动了，但是如果后面遇到b，还是可以把b移动过来的
	// 如果 x[i] = 'b', s[i] = 'a', 那么只要下一个连续的a后面，有一个b
	// 就可以
	n := len(s)
	pos := make([][]int, 3)
	for i := 0; i < 3; i++ {
		pos[i] = make([]int, 0, 1)
	}
	for i := n - 1; i >= 0; i-- {
		j := int(s[i] - 'a')
		pos[j] = append(pos[j], i)
	}

	id := make([]int, 3)
	for i := 0; i < 3; i++ {
		id[i] = len(pos[i])
	}

	get_min_of := func(i int) int {
		ans := pos[i][id[i]-1]
		for j := 0; j < 3; j++ {
			if id[j] > 0 {
				ans = min(ans, pos[j][id[j]-1])
			}
		}

		return ans
	}

	for i := 0; i < n; i++ {
		it := int(x[i] - 'a')

		if id[it] == 0 {
			return false
		}

		j := get_min_of(it)

		if s[j] == x[i] {
			id[it]--
			continue
		}

		if x[i] == 'b' && s[j] == 'a' {
			if id[2] > 0 && pos[2][id[2]-1] < pos[1][id[1]-1] {
				return false
			}
			id[1]--
		} else if x[i] == 'c' && s[j] == 'b' {
			if id[0] > 0 && pos[0][id[0]-1] < pos[2][id[2]-1] {
				return false
			}
			id[2]--
		} else {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		if id[i] > 0 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
