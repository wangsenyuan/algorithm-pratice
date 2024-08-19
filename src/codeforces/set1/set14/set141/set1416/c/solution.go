package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	sum, x := solve(a)
	fmt.Println(sum, x)
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

const D = 30

func solve(a []int) (int, int) {

	// 到d位为止，arr都是一样的，在d位设置0/1,会产生inversions
	// 选择 inversion少的那部分
	var dfs func(d int, arr []int)

	inv := make([][]int, D)
	for i := 0; i < D; i++ {
		inv[i] = make([]int, 2)
	}

	dfs = func(d int, arr []int) {
		if d < 0 || len(arr) < 2 {
			return
		}
		// 如果在d位设置为0
		cnt := []int{0, 0}
		var zero []int
		var one []int
		for _, num := range arr {
			if (num>>d)&1 == 0 {
				inv[d][0] += cnt[1]
				cnt[0]++
				zero = append(zero, num)
			} else {
				inv[d][1] += cnt[0]
				cnt[1]++
				one = append(one, num)
			}
		}
		dfs(d-1, zero)
		dfs(d-1, one)
	}

	dfs(D-1, a)

	var x int
	var sum int
	for i := D - 1; i >= 0; i-- {
		if inv[i][0] > inv[i][1] {
			x |= 1 << i
		}
		sum += min(inv[i][0], inv[i][1])
	}

	return sum, x
}
