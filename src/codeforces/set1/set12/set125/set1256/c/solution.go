package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res, _, _, _ := process(reader)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
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

func process(reader *bufio.Reader) (res []int, n int, m int, d int) {
	n, m, d = readThreeNums(reader)
	p := readNNums(reader, m)
	res = solve(n, m, d, p)
	return
}

func solve(n int, m int, d int, p []int) []int {
	var sum int
	for _, x := range p {
		sum += x
	}

	ans := make([]int, n+2)
	var i, j int
	for i <= n && j < m {
		ni := min(i+d, n-sum+1)
		for u := ni; u < ni+p[j]; u++ {
			ans[u] = j + 1
		}
		i = ni + p[j] - 1
		sum -= p[j]
		j++
	}

	if i+d <= n {
		// no able to reach n+1
		return nil
	}

	return ans[1 : n+1]
}
