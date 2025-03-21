package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	h, m := readTwoNums(reader)

	res := solve(h, m)

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

func solve(h int, m int) int {
	// 1, 2, .... 1 << n

	// h=1是，有两层, 回到顶点要走两步
	// 一共是h+1层节点
	// 0, 1, 2, .... h
	// 在偶数层，也就是 0, 2, 4,... 它们是先左后右
	// 在奇数层，是先右后左的

	var res int
	side := -1
	for i := h; i > 0; i-- {
		if m > 1<<(i-1) {
			// 必须走右边
			if i == h || side != 0 {
				res += 1<<i - 1
			}
			side = 1
			m -= 1 << (i - 1)
		} else {
			// 必须走左边
			if i < h && side != 1 {
				res += 1<<i - 1
			}
			side = 0
		}
	}

	return res + h
}
