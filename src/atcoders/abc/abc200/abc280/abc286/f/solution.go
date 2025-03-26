package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ask := func(arr []int) []int {
		fmt.Println(len(arr))
		s := fmt.Sprintf("%v", arr)
		fmt.Println(s[1 : len(s)-1])
		return readNNums(reader, len(arr))
	}

	n := solve(ask)

	fmt.Println(n)

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

func solve(ask func([]int) []int) int {
	nums := []int{4, 9, 5, 7, 11, 13, 17, 19, 23}
	var arr []int

	for _, x := range nums {
		tmp := construct(x)
		for i := 0; i < len(tmp); i++ {
			tmp[i] += len(arr)
		}
		arr = append(arr, tmp...)
	}
	res := ask(arr)
	rem := make([]int, len(nums))
	var sum int
	for i, x := range nums {
		first := sum + 1
		// 2 3 4 1
		// 4 1 2 3
		cur := res[sum]
		rem[i] = (cur - first + x) % x
		sum += x
	}
	M := 1
	for _, x := range nums {
		M *= x
	}
	var ans int
	for i, x := range nums {
		r := rem[i]
		mi := M / x
		var inv int
		for j := 1; j < x; j++ {
			if mi*j%x == 1 {
				inv = j
				break
			}
		}
		ans += r * mi * inv
		ans %= M
	}
	return ans
}

func construct(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = (i + 1) % n
		arr[i]++
	}
	return arr
}
