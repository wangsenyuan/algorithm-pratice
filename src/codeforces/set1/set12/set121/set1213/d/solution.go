package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, k)
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

const inf = 1 << 60

func solve(a []int, k int) int {
	x := slices.Max(a) + 1

	sz := make([]int, x)

	for _, num := range a {
		sz[num]++
	}

	ans := inf

	merge := func(x, y []int) []int {
		if len(x) > len(y) {
			x, y = y, x
		}
		// len(x) < len(y)
		for i := 0; i < len(x); i++ {
			y[i] += x[i]
		}
		return y
	}

	// dfs(u) 返回对应层级的数量
	var dfs func(u int) []int

	dfs = func(u int) []int {
		var arr []int

		if 2*u < x {
			arr = merge(arr, dfs(2*u))
		}
		if 2*u+1 < x {
			arr = merge(arr, dfs(2*u+1))
		}

		// arr[0] = sz[u]
		tmp := make([]int, len(arr)+1)
		tmp[0] = sz[u]
		copy(tmp[1:], arr)
		var sum int
		var cost int
		for i := 0; i < len(tmp); i++ {
			if sum+tmp[i] >= k {
				// sum + tmp[i] >= k
				add := k - sum
				cost += add * i
				ans = min(ans, cost)
				break
			}
			sum += tmp[i]
			cost += tmp[i] * i
		}

		return tmp
	}

	dfs(1)

	return ans
}
