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
		c, d, x := readThreeNums(reader)
		res := solve(c, d, x)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MAX_X = 20000010

var V [MAX_X]int

func init() {
	mind := make([]int, MAX_X)
	for i := 0; i < MAX_X; i++ {
		mind[i] = -1
	}
	for i := 2; i < MAX_X; i++ {
		if mind[i] < 0 {
			for j := i; j < MAX_X; j += i {
				if mind[j] < 0 {
					mind[j] = i
				}
			}
		}
	}
	for i := 2; i < MAX_X; i++ {
		j := i / mind[i]
		V[i] = V[j]
		if mind[i] != mind[j] {
			V[i]++
		}
	}
}

func solve(c, d, x int) int {
	var res int
	for i := 1; i*i <= x; i++ {
		if x%i != 0 {
			continue
		}
		k1 := x/i + d
		if k1%c == 0 {
			res += 1 << V[k1/c]
		}
		if i*i == x {
			continue
		}
		k2 := i + d
		if k2%c == 0 {
			res += 1 << V[k2/c]
		}
	}

	return res
}
