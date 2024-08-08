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
		s, n := readTwoNums(reader)
		res := solve(s, n)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(s int, n int) []int {
	arr := []int{s}

	pw := make(map[int]bool)

	for i := 1; i <= s; i *= 10 {
		pw[i] = true
	}

	for len(arr) < n {
		x := -1
		for i, y := range arr {
			if y != 1 && !pw[y] {
				x = y
				copy(arr[i:], arr[i+1:])
				arr = arr[:len(arr)-1]
				break
			}
		}
		if x == -1 {
			for _, y := range arr {
				if y == 1 {
					continue
				}
				if x == -1 || x > y {
					x = y
				}
			}
			for i := 0; i < len(arr); i++ {
				if arr[i] == x {
					copy(arr[i:], arr[i+1:])
					arr = arr[:len(arr)-1]
					break
				}
			}
		}
		y := split(x)
		arr = append(arr, y)
		arr = append(arr, x-y)
	}

	return arr
}

func split(num int) int {
	for i := 1; ; i *= 10 {
		if i >= num {
			return i / 10
		}
	}
}
