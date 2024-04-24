package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	res := solve(n)

	s := fmt.Sprintf("%v", res)

	fmt.Println(s[1 : len(s)-1])
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

func solve(n int) []int {
	res := make([]int, n)
	var it int

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	var loop func(cnt int, mul int)

	loop = func(cnt int, mul int) {
		if cnt == 1 {
			res[it] = mul
			it++
		} else if cnt == 2 {
			res[it] = mul
			it++
			res[it] = mul * 2
			it++
		} else if cnt == 3 {
			res[it] = mul
			it++
			res[it] = mul
			it++
			res[it] = mul * 3
			it++
		} else {
			for i := 0; i < cnt; i++ {
				if arr[i]&1 == 1 {
					res[it] = mul
					it++
				}
			}
			for i := 0; i < cnt/2; i++ {
				arr[i] = arr[2*i+1] / 2
			}
			loop(cnt/2, mul*2)
		}
	}

	loop(n, 1)

	return res
}

func solve1(n int) []int {
	vis := make([]bool, n+1)
	h := 1
	for h*2 <= n {
		h *= 2
	}

	var res []int

	for i := h; i > 0; i >>= 1 {
		for j := i; j <= n; j += i {
			if !vis[j] {
				vis[j] = true
				res = append(res, i)
			}
		}
	}

	if n == 1 {
		return res
	}

	if res[0] > res[1] {
		for i := n; i > res[0]; i-- {
			if i%res[1] == 0 {
				res[0] = i
				break
			}
		}
	}

	reverse(res)

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
