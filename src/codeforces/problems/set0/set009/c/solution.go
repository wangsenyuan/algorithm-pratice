package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	fmt.Println(solve(n))
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

func solve(n int) int {
	xs := make([]int, 0, 10)
	num := n
	var i int
	for num > 0 {
		i++
		xs = append(xs, num%10)
		num /= 10
	}

	if xs[i-1] > 1 {
		return pow(2, i) - 1
	}
	var res int
	// 112 => 100, 110, 111, 101
	// 102 => 100, 101
	for i > 0 {
		if xs[i-1] > 1 {
			res += pow(2, i) - 1
			break
		}
		if xs[i-1] == 1 {
			// put 0 at i-1
			res += pow(2, i-1)
		}
		// put 1 or 0
		i--
	}

	return res
}

func pow(a, b int) int {
	var res = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a)
		}
		a *= a
		b >>= 1
	}

	return res
}
