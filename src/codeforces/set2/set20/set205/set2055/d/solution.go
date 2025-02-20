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

	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) int {
	n, k, l := readThreeNums(reader)
	a := readNNums(reader, n)
	return solve(k, l, a)
}

func solve(k int, l int, a []int) int {
	k *= 2
	l *= 2
	for i := 0; i < len(a); i++ {
		a[i] *= 2
	}
	var res int
	if a[0] > 0 {
		res = a[0]
		// a[0] = 0
	}
	pos := k
	// n := len(a)
	for i := 1; i < len(a) && pos < l; i++ {
		if a[i]+res < pos {
			pos = max(pos, a[i]+res+k)
			continue
		}
		if a[i]-res <= pos {
			pos += k
			//move for free
		} else {
			// a[i] - res > pos
			mid := pos + (a[i]-res-pos)/2
			if l <= mid {
				res += l - pos
				pos = l
				break
			}
			res += mid - pos
			pos = mid + k
		}
	}
	if pos < l {
		res += l - pos
	}

	return res
}
