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

	fmt.Println(res[0], res[1])
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

type pair struct {
	first  int
	second int
}

func construct(lectures [][]int, id int) []pair {
	n := len(lectures)
	res := make([]pair, n)
	for i := 0; i < n; i++ {
		res[i] = pair{lectures[i][id*2], lectures[i][id*2+1]}
	}

	return res
}

func getDistinct(arr []pair) []int {
	var nums []int
	for _, cur := range arr {
		nums = append(nums, cur.first, cur.second)
	}
	sort.Ints(nums)

	var res []int

	for i := 0; i < len(nums); {
		res = append(res, nums[i])
		j := i
		for i < len(nums) && nums[i] == nums[j] {
			i++
		}
	}

	return res
}

func solve(x []int) []int {
	sort.Ints(x)

	return []int{getMin(x), getMax(x)}
}

func getMin(x []int) int {
	var res int

	prev := -1

	for i := 0; i < len(x); {
		j := i
		for i < len(x) && x[i] == x[j] {
			i++
		}
		// 它们一起移动
		if x[j]-1 == prev || x[j] == prev {
			continue
		}
		prev = x[j] + 1
		res++
	}

	return res
}

func getMax(x []int) int {
	n := len(x)
	cnt := make([]int, n+1)
	for _, i := range x {
		cnt[i]++
	}
	free := make([]bool, n+2)
	for i := 0; i < len(free); i++ {
		free[i] = true
	}

	for i := 1; i <= n; i++ {
		if cnt[i] > 0 && free[i-1] {
			free[i-1] = false
			cnt[i]--
		}
		if cnt[i] > 0 && free[i] {
			free[i] = false
			cnt[i]--
		}

		if cnt[i] > 0 {
			free[i+1] = false
			cnt[i]--
		}
	}
	var ans int
	for i := 0; i < n+2; i++ {
		if !free[i] {
			ans++
		}
	}
	return ans
}
