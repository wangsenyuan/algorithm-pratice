package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line := readNNums(reader, 4)
	n, m, b, mod := line[0], line[1], line[2], line[3]
	a := readNNums(reader, n)
	res := solve(m, b, mod, a)

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

func solve(m int, b int, mod int, a []int) int {
	add := func(u int, v int) int {
		u += v
		if u >= mod {
			u -= mod
		}
		return u
	}

	fp := make([][]int, m+1)
	// dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		fp[i] = make([]int, b+1)
		// dp[i] = make([]int, b+1)
	}
	fp[0][0] = 1

	for _, x := range a {
		// 问题会出现x比较小的时候
		for j := 1; j <= m; j++ {
			for v := x; v <= b; v++ {
				// fp[j][v] = dp[j][v]
				fp[j][v] = add(fp[j][v], fp[j-1][v-x])
			}
		}
		// for j := 0; j <= m; j++ {
		// 	for v := 0; v <= b; v++ {
		// 		dp[j][v] = fp[j][v]
		// 		fp[j][v] = 0
		// 	}
		// }
	}
	var res int
	for v := 0; v <= b; v++ {
		res = add(res, fp[m][v])
	}

	return res
}
