package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
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

const D = 20

func solve(nums []int) int {
	x := slices.Max(nums)

	if x == 0 {
		return 1
	}
	n := len(nums)
	check := func(d int) int {
		var res int
		var cnt int
		for i := 0; i < n; i++ {
			if (nums[i]>>d)&1 == 0 {
				cnt++
			} else {
				res = max(res, cnt+1)
				cnt = 0
			}
		}

		if cnt == n {
			return 1
		}
		// 最后一段
		return max(res, cnt+1)
	}

	var res int

	for i := 0; i < D; i++ {
		res = max(res, check(i))
	}

	return res
}

func solve1(nums []int) int {
	n := len(nums)

	var or int
	for _, num := range nums {
		or |= num
	}

	cnt := make([]int, D)
	var cur int

	add := func(num int) {
		for i := 0; i < D; i++ {
			cnt[i] += (num >> i) & 1
			if cnt[i] > 0 {
				cur |= 1 << i
			}
		}
	}

	rem := func(num int) {
		for i := 0; i < D; i++ {
			cnt[i] -= (num >> i) & 1
			if cnt[i] == 0 && (cur>>i)&1 == 1 {
				cur ^= 1 << i
			}
		}
	}

	k := 1

	for j, i := 0, 0; j < n && i < n; j++ {
		for i < n && cur != or {
			add(nums[i])
			i++
		}
		if cur == or {
			k = max(k, i-j)
		}
		// else i == n
		rem(nums[j])
	}

	cur = 0
	clear(cnt)
	for j, i := n-1, n-1; j >= 0 && i >= 0; j-- {
		for i >= 0 && cur != or {
			add(nums[i])
			i--
		}
		if cur == or {
			k = max(k, j-i)
		}
		rem(nums[j])
	}

	return k
}
