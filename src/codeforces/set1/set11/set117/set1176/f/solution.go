package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	n := readNum(reader)
	blocks := make([][][]int, n)
	for i := range n {
		k := readNum(reader)
		blocks[i] = make([][]int, k)
		for j := 0; j < k; j++ {
			blocks[i][j] = readNNums(reader, 2)
		}
	}
	return solve(blocks)
}

const inf = 1 << 60

func solve(blocks [][][]int) int {

	dp := make([]int, 10)
	fp := make([]int, 10)

	for i := 0; i < 10; i++ {
		dp[i] = -inf
		fp[i] = -inf
	}

	dp[0] = 0

	pick := func(arr []int, num int, exp int) []int {
		for i, x := range arr {
			if num > x {
				arr[i], num = num, x
			}
		}
		if len(arr) < exp {
			arr = append(arr, num)
		}
		return arr
	}

	want := []int{3, 1, 1}

	for _, cards := range blocks {
		arr := make([][]int, 3)
		for _, card := range cards {
			c, s := card[0], card[1]
			c--
			arr[c] = pick(arr[c], s, want[c])
		}

		for x := 0; x < 10; x++ {
			for j := 0; j < 3; j++ {
				if len(arr[j]) > 0 {
					best := arr[j][0]
					y := (x + 1) % 10
					if y == 0 {
						best *= 2
					}
					fp[y] = max(fp[y], dp[x]+best)
				}
			}
			// 放置2个
			if len(arr[0]) > 0 && len(arr[1]) > 0 {
				best := arr[0][0] + arr[1][0]
				y := (x + 2) % 10
				if y <= 1 {
					best += max(arr[0][0], arr[1][0])
				}
				fp[y] = max(fp[y], dp[x]+best)
			}
			// 全部使用arr[0], 可以放置3个
			var sum int
			for j := 0; j < len(arr[0]); j++ {
				sum += arr[0][j]
				y := (x + j + 1) % 10
				tmp := sum
				if j == 0 && y == 0 || j == 1 && y <= 1 || j == 2 && y <= 2 {
					tmp += arr[0][0]
				}
				fp[y] = max(fp[y], dp[x]+tmp)
			}
		}

		for x := 0; x < 10; x++ {
			dp[x] = max(dp[x], fp[x])
			fp[x] = -inf
		}
	}

	return slices.Max(dp)
}
