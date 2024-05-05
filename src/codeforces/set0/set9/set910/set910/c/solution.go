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
	nums := make([]string, n)
	for i := 0; i < n; i++ {
		nums[i] = readString(reader)
	}

	res := solve(nums)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

const oo = 1 << 60

func solve(nums []string) int {
	val := make([]int, 10)
	id := make([]int, 10)
	cnt := make([][]int, 10)
	ml := make([]int, 10)

	for i := 0; i < 10; i++ {
		val[i] = -1
		id[i] = i
		cnt[i] = make([]int, 6)
	}
	for i := 0; i < 6; i++ {
		cnt[i] = make([]int, 10)
	}

	for _, num := range nums {
		ln := len(num)
		for i := ln - 1; i >= 0; i-- {
			x := int(num[i] - 'a')
			cnt[x][ln-1-i]++
		}
		ml[int(num[0]-'a')] = max(ml[int(num[0]-'a')], len(num))
	}

	getVal := func(x int) int {
		var sum int
		base := 1
		for i := 0; i < 6; i++ {
			sum += cnt[x][i] * base
			base *= 10
		}
		return sum
	}

	sort.Slice(id, func(i, j int) bool {
		x := getVal(id[i])
		y := getVal(id[j])
		return x > y
	})

	for i := 0; i < 10; i++ {
		if ml[id[i]] == 0 {
			// positive integer, 所以只要出现在首位的都不能为0
			val[id[i]] = 0
			break
		}
	}

	for i, j := 0, 1; i < 10; i++ {
		if val[id[i]] >= 0 {
			continue
		}
		val[id[i]] = j
		j++
	}

	var res int

	for _, num := range nums {
		var tmp int
		for i := 0; i < len(num); i++ {
			x := int(num[i] - 'a')
			tmp = tmp*10 + val[x]
		}
		res += tmp
	}
	return res
}
