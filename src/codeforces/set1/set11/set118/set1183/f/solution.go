package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func solve(a []int) int {

	sort.Ints(a)

	var arr []int

	for i := 1; i <= len(a); i++ {
		if i == len(a) || a[i] > a[i-1] {
			arr = append(arr, a[i-1])
		}
	}

	a = arr

	n := len(a)
	marked := make([]bool, n)

	getBest := func(cnt int, n int) int {
		clear(marked)
		var res int

		for j := 0; j < cnt; j++ {
			i := n - 1
			for i >= 0 && marked[i] {
				i--
			}
			if i < 0 {
				break
			}
			x := a[i]
			res += x
			for i >= 0 {
				if x%a[i] == 0 {
					marked[i] = true
				}
				i--
			}
		}
		return res
	}

	// n > 3
	res := getBest(3, n)
	if n <= 3 || a[n-1]%a[n-2] != 0 || a[n-2]*2 < a[n-1] {
		return res
	}
	// a[n-2] *2 = a[n-1]
	res = max(res, getBest(3, n-1))

	return res
}
