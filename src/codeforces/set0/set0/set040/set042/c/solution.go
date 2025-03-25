package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	nums := readNNums(bufio.NewReader(os.Stdin), 4)
	res := solve(nums)
	var buf bytes.Buffer
	for _, op := range res {
		buf.WriteString(op)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(nums []int) []string {
	var ans []string

	prev := func(i int) int {
		return nums[(i+3)%4]
	}

	next := func(i int) int {
		return nums[(i+1)%4]
	}

	add := func(i int) {
		ans = append(ans, fmt.Sprintf("+%d", i+1))
		nums[i]++
		nums[(i+1)%4]++
	}

	div := func(i int) {
		ans = append(ans, fmt.Sprintf("/%d", i+1))
		nums[i] /= 2
		nums[(i+1)%4] /= 2
	}

	for slices.Max(nums) > 1 {
		x := slices.Max(nums)
		if x <= 2 {
			ans = append(ans, smart(nums)...)
			break
		}
		i := slices.Index(nums, x)
		if x%2 == 0 {
			if prev(i)%2 == 0 {
				div((i + 3) % 4)
				continue
			}
			if next(i)%2 == 0 {
				div(i)
				continue
			}
			// 前后都是奇数
			add((i + 1) % 4)
			div(i)
			continue
		}
		// x 是奇数
		add(i)
	}

	return ans
}

func smart(nums []int) []string {
	// 1 2 1 2类似这样子
	// 这时候不需要到1
	// 1 3 2 2 =》2 3 2 3 =》 2 4 3 3 =》 1 2 3 3
	// 2 2 3 4 =》 1 2 3 2 =》 2 3 3 2 =》 3 3 3 3
	M := 4 * 4 * 4 * 4
	vis := make([]bool, M)

	convert := func(arr []int) int {
		var res int
		for _, v := range arr {
			res = res*4 + v - 1
		}
		return res
	}

	var dfs func(state int) bool

	var res []string

	dfs = func(state int) bool {
		if vis[state] {
			return false
		}
		vis[state] = true
		arr := make([]int, 4)
		for num, i := state, 3; i >= 0; i-- {
			arr[i] = num % 4
			arr[i]++
			num /= 4
		}
		if slices.Max(arr) == 1 {
			return true
		}
		// +1 or /2
		for i := 0; i < len(arr); i++ {
			if arr[i]%2 == 0 && arr[(i+1)%4]%2 == 0 {
				arr[i] /= 2
				arr[(i+1)%4] /= 2
				if dfs(convert(arr)) {
					res = append(res, fmt.Sprintf("/%d", i+1))
					return true
				}
				arr[i] *= 2
				arr[(i+1)%4] *= 2
			}

			if arr[i]+1 <= 4 && arr[(i+1)%4]+1 <= 4 {
				arr[i]++
				arr[(i+1)%4]++
				if dfs(convert(arr)) {
					res = append(res, fmt.Sprintf("+%d", i+1))
					return true
				}
				arr[i]--
				arr[(i+1)%4]--
			}
		}
		return false
	}
	dfs(convert(nums))

	reverse(res)

	return res
}

func reverse[T any](arr []T) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
