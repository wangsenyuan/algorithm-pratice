package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n, k := readTwoNums(reader)
	pairs := make([][]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = readNNums(reader, 2)
	}
	return solve(k, pairs)
}

const H = 30

func solve(k int, pairs [][]int) int {
	// 在满足 or(a) < k 的情况下，选择尽可能多的pairs
	// var mask int
	var best int
	// n := len(pairs)

	for i := H - 1; i >= 0; i-- {
		if (k>>i)&1 == 1 {
			// 如果在这一位要求a[i]都为0是的最大值
			// 把从i开始的都变成0
			mask := (k >> (i + 1)) << (i + 1)
			var sum int
			for _, cur := range pairs {
				a, b := cur[0], cur[1]
				hi := (a >> i) << i
				// 对低位没有要求
				if hi&mask == hi {
					sum += b
				}
			}

			best = max(best, sum)
		}
	}

	var sum int
	for _, cur := range pairs {
		a, b := cur[0], cur[1]
		if a&k == a {
			sum += b
		}
	}

	return max(best, sum)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
