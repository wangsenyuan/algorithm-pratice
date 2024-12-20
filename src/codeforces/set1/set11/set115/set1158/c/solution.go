package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		res := process(reader)
		if len(res) == 0 {
			buf.WriteString("-1\n")
			continue
		}
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	next := readNNums(reader, n)
	return solve(n, next)
}

func solve(n int, next []int) []int {
	g := make([][]int, n+1)

	type pair struct {
		first  int
		second int
	}

	var arr []pair

	for i := 0; i < n; i++ {
		j := next[i]
		if j > 0 {
			g[j-1] = append(g[j-1], i)
			arr = append(arr, pair{i, j - 1})
		} else {
			// next[i] = -1, make it as i + 1
			g[i+1] = append(g[i+1], i)
		}
	}

	stack := make([]int, n)
	var top int

	for i := 0; i < len(arr); i++ {
		for top > 0 {
			j := stack[top-1]
			if arr[i].first < arr[j].second && arr[j].second < arr[i].second {
				// conflicts
				return nil
			}
			if arr[i].second < arr[j].second {
				break
			}
			top--
		}
		stack[top] = i
		top++
	}

	ans := make([]int, n+1)
	id := n + 1
	var dfs func(u int)
	dfs = func(u int) {
		ans[u] = id
		id--
		for _, v := range g[u] {
			dfs(v)
		}
	}

	dfs(n)

	check := make([]int, n+1)
	top = 0
	for i := n - 1; i >= 0; i-- {
		for top > 0 && ans[stack[top-1]] < ans[i] {
			top--
		}
		if top == 0 {
			check[i] = n + 1
		} else {
			check[i] = stack[top-1] + 1
		}
		stack[top] = i
		top++
	}

	for i := 0; i < n; i++ {
		if next[i] < 0 {
			continue
		}
		if check[i] != next[i] {
			return nil
		}
	}

	return ans[:n]
}
