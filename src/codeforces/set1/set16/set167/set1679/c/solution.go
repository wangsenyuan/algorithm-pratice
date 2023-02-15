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

	tc := 1

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			s, _ := reader.ReadBytes('\n')
			if s[0] == '1' {
				var x, y int
				pos := readInt(s, 2, &x)
				readInt(s, pos+1, &y)
				Q[i] = []int{1, x, y}
			} else if s[0] == '2' {
				var x, y int
				pos := readInt(s, 2, &x)
				readInt(s, pos+1, &y)
				Q[i] = []int{2, x, y}
			} else {
				var x1, y1, x2, y2 int
				pos := readInt(s, 2, &x1)
				pos = readInt(s, pos+1, &y1)
				pos = readInt(s, pos+1, &x2)
				readInt(s, pos+1, &y2)
				Q[i] = []int{3, x1, y1, x2, y2}
			}
		}

		res := solve(n, Q)
		for _, as := range res {
			if as {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, Q [][]int) []bool {

	set := func(arr []int, p int, v int) {
		p++
		for p < len(arr) {
			arr[p] += v
			p += p & -p
		}
	}

	get := func(arr []int, p int) int {
		p++
		var res int
		for p > 0 {
			res += arr[p]
			p -= p & -p
		}
		return res
	}

	//n := len(Q)
	arr1 := make([]int, n+1)
	arr2 := make([]int, n+1)

	row := make([]int, n)
	col := make([]int, n)

	add := func(arr []int, cnt []int, i int) {
		i--
		cnt[i]++
		if cnt[i] == 1 {
			set(arr, i, 1)
		}
	}

	rem := func(arr []int, cnt []int, i int) {
		i--
		cnt[i]--
		if cnt[i] == 0 {
			set(arr, i, -1)
		}
	}

	var res []bool

	for _, cur := range Q {
		if cur[0] == 1 {
			x, y := cur[1], cur[2]
			add(arr1, row, x)
			add(arr2, col, y)
		} else if cur[0] == 2 {
			x, y := cur[1], cur[2]
			rem(arr1, row, x)
			rem(arr2, col, y)
		} else {
			x1, y1, x2, y2 := cur[1], cur[2], cur[3], cur[4]
			x1--
			y1--
			x2--
			y2--
			cnt1 := get(arr1, x2) - get(arr1, x1-1)
			cnt2 := get(arr2, y2) - get(arr2, y1-1)
			if cnt1 == x2-x1+1 || cnt2 == y2-y1+1 {
				res = append(res, true)
			} else {
				res = append(res, false)
			}
		}
	}

	return res
}
