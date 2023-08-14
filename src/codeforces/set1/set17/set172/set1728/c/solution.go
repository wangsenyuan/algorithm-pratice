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
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(A []int, B []int) int {
	// 任何一个数，如果要改变它
	// 一次后，可以得到 x < 10 (最多9位), 再一次变成1
	// 两次后都变成了1
	// 考虑它们组成一颗树
	// 每个节点上有两个数，cnt_a, cnt_b (a中的个数，b中的个数）
	// 往上移动一次计算一次
	cnt := make([]map[int]int, 2)
	for i := 0; i < 2; i++ {
		cnt[i] = make(map[int]int)
	}

	var nums []int
	vis := make(map[int]bool)

	add := func(num int) {
		for num > 1 {
			if !vis[num] {
				vis[num] = true
				nums = append(nums, num)
			}
			num = change(num)
		}
		if !vis[num] {
			nums = append(nums, num)
			vis[num] = true
		}
	}
	for _, a := range A {
		cnt[0][a]++
		add(a)
	}

	for _, b := range B {
		cnt[1][b]++
		add(b)
	}

	sort.Ints(nums)
	reverse(nums)

	var ans int

	for _, num := range nums {
		if cnt[0][num] == cnt[1][num] {
			continue
		}
		if cnt[0][num] > cnt[1][num] {
			x := cnt[0][num] - cnt[1][num]
			ans += x
			cnt[0][change(num)] += x
		} else {
			x := cnt[1][num] - cnt[0][num]
			ans += x
			cnt[1][change(num)] += x
		}
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func change(num int) int {
	var res int
	for num > 0 {
		res++
		num /= 10
	}
	return res
}
