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
		n := readNum(reader)

		res := solve(n)

		if len(res) == 0 {
			buf.WriteString(fmt.Sprintf("-1\n"))
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
}

const N = 100010

var arr []int

func init() {
	arr = make([]int, N)

	copy(arr, []int{1, 3, 2, 6, 5, 4, 7})
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

func solve(n int) []int {
	if n == 0 {
		return []int{1}
	}
	if n&(n-1) == 0 {
		return nil
	}

	if n == 5 {
		return []int{2, 3, 1, 5, 4}
	}

	if n < 8 {
		return arr[:n]
	}

	i := 7

	for pw := uint(3); (1 << pw) <= n; pw++ {
		st := 1 << pw
		en := min(n+1, 1<<(pw+1))
		arr[i] = st + 1
		i++
		arr[i] = st
		i++
		for j := st + 2; j < en; j++ {
			arr[i] = j
			i++
		}
	}
	return arr[:n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
