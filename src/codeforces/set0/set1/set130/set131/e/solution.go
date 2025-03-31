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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	queens := make([][]int, m)
	for i := range m {
		queens[i] = readNNums(reader, 2)
	}
	return solve(n, queens)
}

func solve(n int, queens [][]int) []int {
	slices.SortFunc(queens, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	m := len(queens)
	cnt := make([]int, m)
	col := make(map[int]int)
	d1 := make(map[int]int)
	d2 := make(map[int]int)

	for i := 0; i < m; {
		j := i
		for i < m && queens[i][0] == queens[j][0] {
			// 同一行
			r, c := queens[i][0], queens[i][1]
			if i > j {
				// 同一行的前后
				cnt[i-1]++
				cnt[i]++
			}
			if v, ok := col[c]; ok {
				// 同一列的上下
				cnt[v]++
				cnt[i]++
			}
			col[c] = i
			if v, ok := d1[r-c]; ok {
				cnt[v]++
				cnt[i]++
			}
			d1[r-c] = i
			if v, ok := d2[r+c]; ok {
				cnt[v]++
				cnt[i]++
			}
			d2[r+c] = i
			i++
		}
	}
	ans := make([]int, 9)
	for i := range m {
		ans[cnt[i]]++
	}
	return ans
}
