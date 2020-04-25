package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		solve(reader)
		readNum(reader)
	}
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

func output(x int) {
	fmt.Printf("3 %d\n", x)
}

func askType2(reader *bufio.Reader, x, d int) int {
	fmt.Printf("2 %d %d\n", x, d)
	return readNum(reader)
}

func askType1(reader *bufio.Reader, x int) int {
	fmt.Printf("1 %d\n", x)
	return readNum(reader)
}

func solve(reader *bufio.Reader) {
	h := readNum(reader)

	if h == 1 {
		output(1)
		return
	}
	x := askType2(reader, 1, h)
	// x is one of the leaf

	for {
		d := askType1(reader, x)

		if d == 0 {
			output(x)
			return
		}
		// d can only be 0, 2, 4, 8, ...2**h

		x = askType2(reader, x, d)

		if d == 2 {
			// only one leaf node with distance of 2
			output(x)
			return
		}
	}

}
