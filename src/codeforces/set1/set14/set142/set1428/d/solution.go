package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	ok, res := solve(n, A)

	if !ok {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func solve(n int, A []int) (bool, [][]int) {
	ones := make([][]int, n)
	for i := 0; i < n; i++ {
		ones[i] = []int{-1, -1}
	}
	h := 1
	var p int
	var res [][]int
	three := []int{-1, -1}
	for i := n - 1; i >= 0; i-- {
		if A[i] == 0 {
			continue
		}
		if A[i] == 1 {
			res = append(res, []int{h, i})
			ones[p] = []int{h, i}
			p++
			h++
		} else if A[i] == 2 {
			if p == 0 {
				return false, nil
			}
			p--
			res = append(res, []int{ones[p][0], i})
			three[0] = i
			three[1] = ones[p][0]
		} else {
			if three[0] == -1 {
				if p == 0 {
					return false, nil
				}
				p--
				three = []int{ones[p][1], ones[p][0]}
			}
			res = append(res, []int{h, i})
			res = append(res, []int{h, three[0]})
			three = []int{i, h}
			h++
		}
	}

	for i := 0; i < len(res); i++ {
		res[i][0] = n - res[i][0] + 1
		res[i][1]++
	}

	return true, res
}

func solve1(n int, A []int) (bool, [][]int) {
	one := make([]int, 0, n)
	two := make([]int, 0, n)
	three := make([]int, 0, n)
	next := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		next[i] = -1
		if A[i] == 0 {
			continue
		}
		if A[i] == 1 {
			one = append(one, i)
			continue
		}
		if A[i] == 2 {
			if len(one) == 0 {
				return false, nil
			}
			next[i] = one[len(one)-1]
			one = one[:len(one)-1]
			two = append(two, i)
		} else {
			// A[i] == 3
			if len(three) > 0 {
				next[i] = three[len(three)-1]
				three[len(three)-1] = i
			} else if len(two) > 0 {
				next[i] = two[len(two)-1]
				two = two[:len(two)-1]
				three = append(three, i)
			} else if len(one) > 0 {
				next[i] = one[len(one)-1]
				one = one[:len(one)-1]
				three = append(three, i)
			} else {
				return false, nil
			}
		}
	}
	var res [][]int
	var r int
	for _, i := range one {
		res = append(res, []int{r + 1, i + 1})
		r++
	}
	for _, i := range two {
		res = append(res, []int{r + 1, i + 1})
		j := next[i]
		res = append(res, []int{r + 1, j + 1})
		r++
	}

	for _, i := range three {
		j := i
		for A[j] == 3 {
			res = append(res, []int{r + 1, j + 1})
			j = next[j]
			res = append(res, []int{r + 1, j + 1})
			r++
		}
		res = append(res, []int{r + 1, j + 1})
		if A[j] == 2 {
			res = append(res, []int{r + 1, next[j] + 1})
			r++
		}
	}

	return true, res
}
