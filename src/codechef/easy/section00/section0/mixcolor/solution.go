package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewReader(os.Stdin)

	t := readNum(scanner)
	var buf bytes.Buffer

	for t > 0 {
		n := readNum(scanner)
		A := readNNums(scanner, n)
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
		t--
	}
	fmt.Print(buf.String())
}

func solve1(n int, A []int) int {
	nums := make(map[int]bool)

	for _, a := range A {
		nums[a] = true
	}

	return len(A) - len(nums)
}

func solve(n int, A []int) int {
	//mergeSort(n, A)
	sort.Ints(A)
	var j int
	var cnt int
	for i := 1; i <= n; i++ {
		if i < n && A[i] == A[j] {
			continue
		}
		j = i
		cnt++
	}
	return n - cnt
}

func mergeSort(n int, A []int) {

	var rec func(left, right int)
	help := make([]int, n)

	rec = func(left, right int) {
		if left+1 >= right {
			return
		}

		mid := left + (right-left)/2
		rec(left, mid)
		rec(mid, right)

		i, j, k := left, mid, left
		for i < mid && j < right {
			if A[i] <= A[j] {
				help[k] = A[i]
				i++
				k++
			} else {
				help[k] = A[j]
				j++
				k++
			}
		}

		for i < mid {
			help[k] = A[i]
			i++
			k++
		}
		for j < right {
			help[k] = A[j]
			j++
			k++
		}
		copy(A[left:right], help[left:right])
	}
	rec(0, n)
}
