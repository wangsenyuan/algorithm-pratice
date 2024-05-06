package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	first := readNNums(reader, 5)
	v := first[4]
	stairs := readNNums(reader, first[2])
	elevators := readNNums(reader, first[3])

	m := readNum(reader)
	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 4)
	}

	res := solve(first[0], first[1], v, stairs, elevators, queries)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

const inf = 1 << 60

func solve(n int, m int, v int, stairs []int, elevators []int, queries [][]int) []int {

	checkStair := func(x1, y1, x2, y2 int, p int) int {
		res := abs(p - y1)
		res += abs(p - y2)
		res += abs(x2 - x1)
		return res
	}

	checkElevator := func(x1, y1, x2, y2 int, p int) int {
		res := abs(p - y1)
		res += abs(p - y2)
		h := abs(x2 - x1)
		w := (h + v - 1) / v
		res += w
		return res
	}

	checkPos := func(p int, arr []int, f func(int) int) int {
		i := search(len(arr), func(i int) bool {
			return arr[i] > p
		})
		res := inf
		if i < len(arr) {
			res = f(arr[i])
		}
		if i > 0 {
			res = min(res, f(arr[i-1]))
		}
		return res
	}

	check := func(cur []int) int {
		x1, y1, x2, y2 := cur[0], cur[1], cur[2], cur[3]
		if x1 == x2 {
			// same floor
			return abs(y2 - y1)
		}
		// x1 < x2
		a := min(y1, y2)
		b := max(y1, y2)
		ra := checkPos(a, stairs, func(p int) int {
			return checkStair(x1, y1, x2, y2, p)
		})
		rb := checkPos(b, stairs, func(p int) int {
			return checkStair(x2, y2, x1, y1, p)
		})

		xa := checkPos(a, elevators, func(p int) int {
			return checkElevator(x1, y1, x2, y2, p)
		})

		xb := checkPos(b, elevators, func(p int) int {
			return checkElevator(x1, y1, x2, y2, p)
		})

		return min(ra, rb, xa, xb)
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = check(cur)
	}

	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
