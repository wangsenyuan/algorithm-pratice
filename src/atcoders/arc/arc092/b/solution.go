package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, n)
	res := solve(A, B)

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

const H = 30

func solve(a []int, b []int) int {
	n := len(a)
	var ans int
	for k := 0; k < 29; k++ {
		mask := 1<<(k+1) - 1
		sort.Slice(a, func(i, j int) bool { return a[i]&mask < a[j]&mask })
		sort.Slice(b, func(i, j int) bool { return b[i]&mask < b[j]&mask })
		cnt := 0
		i, j, p := n-1, n-1, n-1
		l1, r1, l2 := 1<<k, 1<<(k+1)-1, 3<<k
		for _, v := range a {
			for i >= 0 && v&mask+b[i]&mask >= l1 {
				i--
			}
			for j >= 0 && v&mask+b[j]&mask > r1 {
				j--
			}
			for p >= 0 && v&mask+b[p]&mask >= l2 {
				p--
			}
			cnt ^= i ^ j ^ p
		}
		ans |= cnt & 1 << k
	}

	return ans
}
