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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	query := func(x uint64) uint64 {
		fmt.Printf("1 %d\n", x)
		b, _ := reader.ReadBytes('\n')
		var v uint64
		readUint64(b, 0, &v)
		return v
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n, query)
		fmt.Printf("2 ")
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", res[i])
		}
		fmt.Println()
		readNum(reader)
	}
}

func solve(n int, query func(uint64) uint64) []uint64 {
	res := make([]uint64, n)
	res[0] = query(0)

	for i := 1; i < n; i++ {
		cur := res[i-1]

		for b := 0; b < 64; b++ {
			nxt := cur ^ (1 << uint64(b))

			if nxt > cur {
				// cur b-th digit is 0
				y := query(nxt)
				if y > res[i-1] {
					res[i] = y
					break
				}
			} else {
				// cur b-th digit is 1, and reset it to 0
				cur = nxt
			}
		}
	}

	return res
}
