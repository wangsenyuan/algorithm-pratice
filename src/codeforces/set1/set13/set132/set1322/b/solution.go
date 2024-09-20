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
	a := readNNums(reader, n)
	res := solve(a)
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

func solve(a []int) int {
	n := len(a)
	var ans int

	var all int
	for i := 0; i < n; i++ {
		all ^= i
	}

	for k := 0; k < 25; k++ {
		mask := 1<<(k+1) - 1
		sort.Slice(a, func(i, j int) bool { return a[i]&mask < a[j]&mask })
		f := func(high int) (cnt int) {
			i, j := 0, n-1
			for i < j {
				if a[i]&mask+a[j]&mask < high {
					cnt ^= i ^ j
					i++
				} else {
					j--
				}
			}
			return
		}
		ans |= (f(1<<k) ^ f(1<<(k+1)) ^ f(3<<k) ^ all) & 1 << k
	}
	return ans
}

func search(l, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func bruteForce(a []int) int {
	var res int
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res ^= (a[i] + a[j])
		}
	}

	return res
}
