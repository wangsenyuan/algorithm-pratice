package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k, l := readThreeNums(reader)
	a := readNNums(reader, n*k)
	fmt.Println(solve(a, k, l))
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

func solve(a []int, k int, l int) int {
	sort.Ints(a)

	n := len(a)

	vis := make([]bool, n)

	it := sort.Search(n, func(i int) bool {
		return a[i]-a[0] > l
	})
	if it < n/k {
		return 0
	}
	it--
	// a[it] - a[0] <= l
	var res int
	r := n - 1
	// stage 1
	for i := it; i >= 0; i-- {
		if r-it < k-1 {
			break
		}
		var cnt int
		for cnt < k-1 {
			vis[r] = true
			cnt++
			r--
		}
		res += a[i]
		vis[i] = true
	}

	for i := n - 1; i >= 0; {
		var cnt int
		for i >= 0 && cnt < k {
			if !vis[i] {
				cnt++
			}
			i--
		}
		if cnt == k {
			res += a[i+1]
		}
	}
	return res
}
