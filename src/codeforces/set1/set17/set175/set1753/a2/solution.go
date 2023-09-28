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
		a := readNNums(reader, n)
		res := solve(a)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func solve(a []int) [][]int {
	n := len(a)
	var sum int
	for i := 0; i < n; i++ {
		sum += a[i]
	}

	if sum%2 != 0 {
		return nil
	}
	var res [][]int

	for i := 0; i < n; {
		first := -1
		second := -1
		var x int
		j := i
		for i < n && second < 0 {
			if a[i] != 0 {
				if first < 0 {
					first = i
				} else {
					second = i
				}
			}
			if (i-j)&1 == 0 {
				x += a[i]
			} else {
				x -= a[i]
			}
			i++
		}
		if x == 0 {
			// i == n
			res = append(res, []int{j + 1, i})
			continue
		}
		// x = 2 or x = -2
		ln := i - j
		if ln%2 == 0 {
			// flip last one
			res = append(res, []int{j + 1, i - 1})
			res = append(res, []int{i, i})
			continue
		} else {
			// last one is in the positive position
			if first == j {
				res = append(res, []int{j + 1, j + 1})
				res = append(res, []int{j + 2, i})
			} else {
				// a[j] == 0
				res = append(res, []int{j + 1, j + 1})
				j++
				res = append(res, []int{j + 1, i - 1})
				res = append(res, []int{i, i})
			}
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}
