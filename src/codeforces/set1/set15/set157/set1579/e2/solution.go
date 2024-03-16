package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(a []int) int {
	arr := compact(a)
	n := len(arr)

	set := make([]int, 2*n)

	update := func(p int, v int) {
		for i := p + n; i > 0; i >>= 1 {
			set[i] += v
		}
	}

	get := func(l int, r int) int {
		var res int
		for i, j := l+n, r+n; i < j; i, j = i>>1, j>>1 {
			if i&1 == 1 {
				res += set[i]
				i++
			}
			if j&1 == 1 {
				j--
				res += set[j]
			}
		}
		return res
	}

	var res int

	for i := 0; i < n; i++ {
		num := arr[i]
		// 比num小的数
		x := get(0, num)
		// 比num大的数
		y := get(num+1, n)
		// 如果x <= y, 就将num放在头部
		// else 将num放在尾部
		res += min(x, y)
		update(num, 1)
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func compact(a []int) []int {
	nums := make(map[int]int)
	for i, num := range a {
		nums[num] = i
	}
	var arr []int
	for num := range nums {
		arr = append(arr, num)
	}
	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		nums[arr[i]] = i
	}
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = nums[a[i]]
	}
	return res
}
