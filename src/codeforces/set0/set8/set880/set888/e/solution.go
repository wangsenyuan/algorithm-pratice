package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(m, a)

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

func solve(m int, a []int) int {
	if m == 1 {
		return 0
	}
	n := len(a)

	if n == 1 {
		return a[0] % m
	}

	h := n / 2
	first := bruteForce(a[:h], m)
	second := bruteForce(a[h:], m)

	reverse(second)

	res := max(first[len(first)-1], second[0])

	// first是升序的, second 是降序的
	// first[i] 来说，如果 first[i] + second[j] >= m, 那么 first[i+1] + second[j] >= m
	// 但是 first[i+1] + second[j] 有可能是 < m
	// a < m, b < m, a + b < 2 * m

	for i, j := 0, 0; i < len(first); i++ {
		for j < len(second) && first[i]+second[j] >= m {
			res = max(res, (first[i]+second[j])%m)
			j++
		}
		res = max(res, (first[i]+second[0])%m)
		if j < len(second) {
			res = max(res, first[i]+second[j])
		}
	}

	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func bruteForce(a []int, m int) []int {
	n := len(a)
	res := make(map[int]int)
	res[0]++
	for state := 1; state < 1<<n; state++ {
		var sum int
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				sum += a[i]
			}
		}
		res[sum%m]++
	}
	arr := make([]int, 0, len(res))
	for k := range res {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	return arr
}
