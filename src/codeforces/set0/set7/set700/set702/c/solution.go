package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	res := solve(a, b)
	fmt.Println(res)
}

func solve(a []int, b []int) int {
	// 不需要重复的元素
	a = removeDuplicates(a)
	b = removeDuplicates(b)

	// 给定r是否能够覆盖所有的city
	check := func(r int) bool {
		var j int
		for i := 0; i < len(a); i++ {
			for j < len(b) && a[i]-b[j] > r {
				j++
			}
			if j == len(b) {
				// no tower to cover i
				return false
			}
			if b[j]-a[i] > r {
				// next is too far
				return false
			}
		}
		return true
	}

	l, r := 0, 1<<32
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func removeDuplicates(a []int) []int {
	var n int
	for i := 1; i <= len(a); i++ {
		if i == len(a) || a[i] > a[i-1] {
			a[n] = a[i-1]
			n++
		}
	}
	return a[:n]
}
