package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	blocks := make([][]int, n)
	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var m int
		pos := readInt(s, 0, &m)
		blocks[i] = make([]int, m)
		for j := 0; j < m; j++ {
			pos = readInt(s, pos+1, &blocks[i][j])
		}
	}
	split, merge := solve(blocks)
	fmt.Printf("%d %d\n", split, merge)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(blocks [][]int) (split int, merge int) {
	// 先split后merge
	// 如果两个数中间，存在其他的数，就必须split
	nums := make(map[int]int)
	for _, block := range blocks {
		for _, v := range block {
			nums[v]++
		}
	}
	arr := make([]int, 0, len(nums))
	for k := range nums {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	for _, block := range blocks {
		for i := 0; i+1 < len(block); i++ {
			a := block[i]
			b := block[i+1]
			if a == b {
				continue
			}
			if a > b {
				split++
				continue
			}
			// a < b
			j := search(len(arr), func(j int) bool {
				return arr[j] > a
			})
			j--
			// upper bound arr[j] == a
			k := search(len(arr), func(k int) bool {
				return arr[k] >= b
			})
			// lower bound
			if j+1 < k {
				split++
			}
		}
	}

	merge = len(blocks) + split - 1
	return
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
