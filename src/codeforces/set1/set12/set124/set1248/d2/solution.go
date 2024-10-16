package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)
	s := readString(reader)

	res, swaps := solve(s)

	fmt.Printf("%d\n%d %d\n", res, swaps[0], swaps[1])
}

func solve(s string) (int, []int) {
	n := len(s)

	a := make([]int, n+1)

	var level int
	var mn int
	var pos int
	for i, x := range []byte(s) {
		if x == '(' {
			a[i+1] = 1
		} else {
			a[i+1] = -1
		}
		level += a[i+1]
		if level < mn {
			mn = level
			pos = i + 1
		}
	}

	shift(a[1:], pos)

	for i := 1; i <= n; i++ {
		a[i] += a[i-1]
	}

	if a[n] != 0 {
		return 0, []int{1, 1}
	}

	var c2, c0, mx, last int
	var tl, tr int
	for i := 1; i <= n; i++ {
		if a[i] < 2 {
			c2 = 0
			last = i
		}
		if a[i] == 2 {
			c2++
		}
		if a[i] == 0 {
			c0++
		}
		if c2 > mx {
			tl = last
			tr = i
			mx = c2
		}
	}

	last = 0
	var c1 int
	for i := 1; i <= n; i++ {
		if a[i] < 1 {
			c1 = 0
			last = i
		}
		if a[i] == 1 {
			c1++
		}
		if c1-c0 > mx {
			mx = c1 - c0
			tl = last
			tr = i
		}
	}

	return c0 + mx, []int{(tl+pos)%n + 1, (tr+pos)%n + 1}
}

func shift[T any](arr []T, k int) {
	reverse(arr[:k])
	reverse(arr[k:])
	reverse(arr)
}

func reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
