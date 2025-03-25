package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Println(buf.String())
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, k)
}

func solve(a []int, k int) int {

	sort.Ints(a)

	var ans int
	n := len(a)
	// 策略1
	marked := make([]bool, n)
	for l, r, i := 0, n-1, n-1; l+1 < r && l < i; l++ {
		for r > l && marked[r] {
			r--
		}
		if l == r {
			break
		}
		// 至少能选择3个
		// a[l] + a[r] >= k
		i = min(i, r-1)
		for i > l && a[l]+a[i] >= k {
			i--
		}
		if i == l {
			break
		}
		ans++
		marked[l] = true
		marked[r] = true
		marked[i] = true
		r--
		i--
	}
	var arr []int
	for i, v := range a {
		if !marked[i] {
			arr = append(arr, v)
		}
	}
	clear(marked)
	n = len(arr)

	for l, r := 0, n-1; l < r && arr[l] < k; l++ {
		ans++
		marked[l] = true
		marked[r] = true
		r--
	}
	for i := range len(arr) {
		if !marked[i] {
			ans++
		}
	}

	return ans - 1
}

type PQ struct {
	sort.IntSlice
}

func (pq PQ) Less(i, j int) bool {
	return pq.IntSlice[i] > pq.IntSlice[j]
}

func (pq *PQ) Push(v any) {
	pq.IntSlice = append(pq.IntSlice, v.(int))
}

func (pq *PQ) Pop() any {
	a := pq.IntSlice
	v := a[len(a)-1]
	pq.IntSlice = a[:len(a)-1]
	return v
}
