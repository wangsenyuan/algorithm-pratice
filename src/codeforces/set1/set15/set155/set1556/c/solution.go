package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	c := readNNums(reader, n)
	res := solve(c)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(a []int) int {
	n := len(a)

	// second是深度
	type pair struct {
		first  int
		second int
	}
	stack := make([]pair, n+1)
	top := 1

	var res int

	var cur int

	for i := 0; i+1 < n; i += 2 {
		cur += a[i] - a[i+1]
		res += min(a[i+1], cur+a[i+1]-stack[0].first)

		for top > 0 && stack[top-1].first > cur {
			res += stack[top-1].second
			top--
		}

		if top > 0 && stack[top-1].first == cur {
			res += stack[top-1].second
			stack[top-1] = pair{cur, stack[top-1].second + 1}
		} else {
			if top == 0 {
				stack[top] = pair{cur, 0}
			} else {
				stack[top] = pair{cur, 1}
			}
			top++
		}
	}

	return res
}

func solve1(a []int) int {
	n := len(a)

	var ans int

	for i := 0; i+1 < n; i += 2 {
		sum := a[i]
		ub := a[i]
		for j := i + 1; j < n; j++ {
			if j&1 == 1 {
				next := sum - a[j]
				ans += max(0, min(ub, sum)-max(0, next))
				ub = min(ub, next+1)
				sum = next
			} else {
				sum += a[j]
			}
		}
	}
	return ans
}
