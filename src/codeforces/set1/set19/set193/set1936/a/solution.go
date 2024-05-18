package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(a, b, c, d int) string {
		fmt.Printf("? %d %d %d %d\n", a, b, c, d)
		return readString(reader)
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		i, j := solve(n, ask)

		fmt.Printf("! %d %d\n", i, j)
	}
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

func solve(n int, ask func(int, int, int, int) string) (int, int) {
	var ans1 int

	for i := 1; i < n; i++ {
		res := ask(ans1, ans1, i, i)
		if res == "<" {
			ans1 = i
		}
	}
	// p[ans1] = n - 1
	var mx int
	v := []int{0}
	for i := 1; i < n; i++ {
		res := ask(mx, ans1, i, ans1)
		if res == "<" {
			mx = i
			v = []int{mx}
		} else if res == "=" {
			v = append(v, i)
		}
	}
	// p[mx] | （n - 1) is min
	ans2 := v[0]

	for i := 1; i < len(v); i++ {
		res := ask(ans2, ans2, v[i], v[i])
		if res == ">" {
			ans2 = v[i]
		}
	}

	return ans1, ans2
}
