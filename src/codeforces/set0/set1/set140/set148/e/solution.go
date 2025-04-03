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
	n, m := readTwoNums(reader)
	shelves := make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k)
		pos++
		shelves[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &shelves[i][j]) + 1
		}
	}
	return solve(shelves, m)
}

const inf = 1 << 60

func solve(shelves [][]int, m int) int {
	var k int
	for _, row := range shelves {
		k = max(k, len(row))
	}
	fp := make([]int, min(k, m)+1)
	pref := make([]int, k+1)
	calc := func(row []int) {
		clear(fp)
		for i, x := range row {
			pref[i+1] = pref[i] + x
		}
		w := len(row)
		for x := 1; x <= min(w, m); x++ {
			var suf int
			for j := 0; j <= x; j++ {
				fp[x] = max(fp[x], suf+pref[x-j])
				if j < x {
					suf += row[w-1-j]
				}
			}
		}
	}
	n := len(shelves)
	dp := make([]int, m+1)
	for i := range m + 1 {
		dp[i] = -inf
	}
	dp[0] = 0
	pd := make([]int, m+1)
	var sum int
	for i := range n {
		calc(shelves[i])
		w := len(shelves[i])
		// 假设在这一行里两头选择了k个
		// 那么 dp[j] = max(dp[j-k] + fp[k])
		// 但问题是这样子，复杂性，变成了 1e8
		sum += w
		for k := 0; k <= min(w, m); k++ {
			for j := 0; j+k <= min(sum, m); j++ {
				pd[j+k] = max(pd[j+k], dp[j]+fp[k])
			}
		}
		copy(dp, pd)
	}

	return dp[m]
}
