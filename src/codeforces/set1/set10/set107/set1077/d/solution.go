package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)

	res := solve(a, k)

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

func solve(s []int, k int) []int {
	sort.Ints(s)
	var arr []Pair
	n := len(s)
	for i := 0; i < n; {
		j := i
		for i < n && s[i] == s[j] {
			i++
		}
		arr = append(arr, Pair{i - j, s[j]})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first > arr[j].first
	})

	check := func(expect int) bool {
		// 是否能够找到k个数满足重复expect次
		var cnt int
		for i := 0; i < len(arr) && cnt < k; i++ {
			cnt += arr[i].first / expect
		}
		return cnt >= k
	}

	l, r := 1, n+1

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	r--

	ans := make([]int, k)

	for i, j := 0, 0; i < len(arr) && j < k; i++ {
		cnt := arr[i].first / r
		for j < k && cnt > 0 {
			ans[j] = arr[i].second
			j++
			cnt--
		}
	}

	return ans
}

type Pair struct {
	first  int
	second int
}
