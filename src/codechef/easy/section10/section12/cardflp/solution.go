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
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res, cnt := solve(n, A, B)
		buf.WriteString(fmt.Sprintf("%d %d\n", res, cnt))
	}
	fmt.Print(buf.String())
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

func solve(n int, A []int, B []int) (int, int) {
	// fp[i] = 0, no determined yet, 1 flip, 2 not flip
	fp := make([]int, n)

	var res int

	for b := 30; b >= 0; b-- {
		ok := true
		for i := 0; i < n && ok; i++ {
			if fp[i] == 1 {
				// check B[i][b]
				if (B[i]>>uint(b))&1 == 0 {
					ok = false
				}
			} else if fp[i] == 2 {
				// check A[i]
				if (A[i]>>uint(b))&1 == 0 {
					ok = false
				}
			} else {
				if (A[i]>>uint(b))&1 == 0 && (B[i]>>uint(b))&1 == 0 {
					ok = false
				}
			}
		}
		if !ok {
			// can't set this position, no change
			continue
		}
		// can set as 1
		res |= 1 << uint(b)

		for i := 0; i < n; i++ {
			if fp[i] == 0 {
				if (A[i]>>uint(b))&1 == 0 {
					fp[i] = 1
				} else if (B[i]>>uint(b))&1 == 0 {
					fp[i] = 2
				}
				// else still stay no determined yet
			}
		}
	}
	var cnt int
	for i := 0; i < n; i++ {
		if fp[i] == 1 {
			cnt++
		}
	}
	return res, cnt
}
