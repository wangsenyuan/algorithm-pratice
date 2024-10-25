package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Print(res)
}

func process(reader *bufio.Reader) string {
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		l, r := readTwoNums(reader)
		res := solve(l, r)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	return buf.String()
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

func solve(l int, r int) int {
	if l == r {
		return l
	}
	if (r+1)&r == 0 {
		// 全部是1
		return r
	}
	n := bits.Len(uint(r))
	var res int
	for i := n - 1; i >= 0; i-- {
		if (l>>i)&1 != (r>>i)&1 {
			// 先把当前位包括进去
			mask := 1<<(i+1) - 1

			rem := r & mask

			if rem == mask {
				// r的后面全是1
				res |= mask
			} else {
				// 除了当前位，后面都是1
				res |= 1<<i - 1
			}

			break
		}

		if (l>>i)&1 == 1 {
			res |= 1 << i
		}
	}

	return res
}
