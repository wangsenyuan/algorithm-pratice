package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func solve(a []int) bool {
	n := len(a)

	type pair struct {
		first  int
		second int
	}

	stack := make([]pair, n)
	var top int
	for i := 0; i < n; i++ {
		var used bool
		for top > 0 && stack[top-1].first <= a[i] {
			if stack[top-1].first == a[i] {
				stack[top-1].second++
				used = true
				break
			}
			// stack[top-1].first < a[i]
			if stack[top-1].second%2 == 1 {
				return false
			}
			top--
		}
		if !used {
			stack[top] = pair{a[i], 1}
			top++
		}
	}

	for top > 1 && stack[top-1].second%2 == 0 {
		top--
	}

	return top <= 1
}
