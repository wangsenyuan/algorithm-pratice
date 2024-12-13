package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	ok, res := solve(n, k)

	if !ok {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("YES\n%d\n", len(res)))
	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
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

func solve(n int, k int) (bool, [][]int) {
	if n == k {
		return n == 1, nil
	}

	if n == 2 {
		// k == 1
		return k == 1, [][]int{{1, 2}}
	}

	if k == 1 {
		var res [][]int
		for i := 0; i < n; i++ {
			res = append(res, []int{i + 1, (i+1)%n + 1})
		}
		return true, res
	}

	var res [][]int
	var build func(n int, k int) int

	build = func(n int, k int) int {
		if k+1 < n {
			r := build(n-1, k)
			res = append(res, []int{r, n})
			return r
		} else {
			if n <= 4 {
				res = append(res, []int{1, 2})
				res = append(res, []int{2, 3})
			}
			if n == 3 {
				return 2
			}
			if n == 4 {
				res = append(res, []int{1, 4})
				res = append(res, []int{1, 3})
				return 1
			}
			build(n-2, k-2)

			for i := 1; i < n; i++ {
				res = append(res, []int{i, n})
			}

			return n
		}
	}

	build(n, k)

	return true, res
}
