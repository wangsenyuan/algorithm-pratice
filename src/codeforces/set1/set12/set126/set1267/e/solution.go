package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = readNNums(reader, n)
	}
	res := solve(a)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')

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

func solve(a [][]int) []int {
	m := len(a)
	n := len(a[0])

	type pair struct {
		first  int
		second int
	}

	arr := make([]pair, m)

	check := func(col int) []int {
		var sum int
		for i := 0; i < m; i++ {
			arr[i] = pair{a[i][n-1] - a[i][col], i}
			sum += arr[i].first
		}

		slices.SortFunc(arr, func(x, y pair) int {
			return y.first - x.first
		})

		if sum <= 0 {
			return nil
		}

		var ans []int

		for _, x := range arr {
			sum -= x.first
			ans = append(ans, x.second+1)
			if sum <= 0 {
				return ans
			}
		}

		return ans
	}
	var best []int

	for i := 0; i < n-1; i++ {
		tmp := check(i)
		if i == 0 {
			best = tmp
		} else if len(best) > len(tmp) {
			best = tmp
		}
	}

	return best
}
