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
	a := readNNums(reader, n)

	res := solve(a)

	fmt.Println(res)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(nums []int) int {
	n := len(nums)
	items := make([]Item, n)
	for i, num := range nums {
		items[i] = process(num)
	}
	sort.Slice(items, func(i, j int) bool {
		a := items[i].square_diff - items[i].no_square_diff
		b := items[j].square_diff - items[j].no_square_diff
		return a < b
	})
	// 前面的用来作为square number
	var res int
	for i := 0; i < n/2; i++ {
		res += items[i].square_diff
	}
	for i := n / 2; i < n; i++ {
		res += items[i].no_square_diff
	}

	return res
}

func process(num int) Item {
	if num == 0 {
		return Item{0, 0, 2}
	}
	if num == 1 {
		return Item{1, 0, 1}
	}

	i := sort.Search(num, func(i int) bool {
		return i*i >= num
	})
	a := i * i
	b := (i - 1) * (i - 1)
	diff := min(a-num, num-b)
	if diff == 0 {
		// num is a square number
		return Item{num, 0, 1}
	}
	// num is not a square number
	return Item{num, diff, 0}
}

type Item struct {
	val            int
	square_diff    int
	no_square_diff int
}
