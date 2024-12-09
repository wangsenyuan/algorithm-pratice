package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n := readNum(reader)

	a := readNNums(reader, n)
	b := readNNums(reader, n)

	return solve(a, b)
}

func solve(a []int, b []int) int {
	type pair struct {
		first  int
		second int
	}
	n := len(a)
	students := make([]pair, n)

	for i := range n {
		students[i] = pair{a[i], b[i]}
	}

	slices.SortFunc(students, func(x, y pair) int {
		return x.first - y.first
	})

	// 如果相同a的人，有两个以上，他们可以组成一个分组。
	var res int

	var sum2 int

	var arr []pair

	for i := 0; i < n; {
		j := i
		var sum int
		for i < n && students[i].first == students[j].first {
			sum += students[i].second
			i++
		}
		if i-j >= 2 {
			// 必须至少有两个人一组
			var next []pair
			for _, x := range arr {
				if x.first&students[j].first == x.first {
					sum += x.second
				} else {
					next = append(next, x)
				}
			}
			sum2 += sum
			res = max(res, sum2)
			arr = next
		} else {
			arr = append(arr, students[j])
		}
	}
	return res
}
