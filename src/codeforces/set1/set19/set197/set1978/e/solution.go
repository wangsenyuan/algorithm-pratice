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
		readNum(reader)
		a := readString(reader)
		b := readString(reader)
		n := readNum(reader)
		queries := make([][]int, n)
		for i := 0; i < n; i++ {
			queries[i] = readNNums(reader, 2)
		}
		res := solve(a, b, queries)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
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

type Pair struct {
	first  int
	second int
}

func solve(a string, b string, queries [][]int) []int {
	n := len(a)

	get := func(i int) Pair {
		if a[i] == '1' {
			return Pair{i, i}
		}
		l, r := -1, -1
		if i > 0 && b[i-1] == '1' {
			l = i - 1
		} else if i > 1 && b[i-1] == '0' && a[i-2] == '0' {
			l = i - 2
		}
		if i+1 < n && b[i+1] == '1' {
			r = i + 1
		} else if i+2 < n && b[i+1] == '0' && a[i+2] == '0' {
			r = i + 2
		}
		if l == -1 {
			r = -1
		}
		if r == -1 {
			l = -1
		}
		return Pair{l, r}
	}

	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		tmp := get(i)
		pref[i+1] = pref[i]
		if tmp.first >= 0 {
			pref[i+1]++
		}
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		if r-l+1 <= 5 {
			for j := l - 1; j < r; j++ {
				tmp := get(j)
				if tmp.first >= l-1 && tmp.second < r {
					ans[i]++
				}
			}
		} else {
			ans[i] = pref[r] - pref[l-1]
			for _, j := range []int{l - 1, l, r - 2, r - 1} {
				tmp := get(j)
				if tmp.first >= 0 && tmp.first < l-1 || tmp.second >= r {
					ans[i]--
				}
			}
		}
	}
	return ans
}

func convert(s string) []int {
	n := len(s)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(s[i] - '0')
	}
	return arr
}
