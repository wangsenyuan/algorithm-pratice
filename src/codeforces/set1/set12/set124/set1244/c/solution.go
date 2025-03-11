package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(res[0], res[1], res[2])
	}
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

func process(reader *bufio.Reader) (nums []int, res []int) {
	nums = readNNums(reader, 4)
	res = solve(nums)
	return
}

func solve(nums []int) []int {
	n, p, w, d := nums[0], nums[1], nums[2], nums[3]

	// x * w + y * d = p
	for y := 0; y <= min(n, w-1); y++ {
		// x + y + z = n
		// y is the number of draws
		if p-y*d < 0 {
			break
		}
		// p - y * d >= 0
		if (p-y*d)%w == 0 {
			x := (p - y*d) / w
			if x > n-y {
				continue
			}
			return []int{x, y, n - x - y}
		}
	}

	return nil
}

func extgcd(a, b int) (g int, x int, y int) {
	if b == 0 {
		return a, 1, 0
	}
	g, y, x = extgcd(b, a%b)

	y -= x * (a / b)

	return
}
