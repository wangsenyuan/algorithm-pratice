package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nums := readNNums(reader, 4)
	res := solve(nums[0], nums[1], nums[2], nums[3])

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

func solve(a int, b int, f int, k int) int {
	var pos int
	var fuel = b

	var res int

	for i := 0; i < k; i++ {
		if fuel < 0 {
			return -1
		}
		if pos == 0 {
			if fuel < f {
				return -1
			}
			if fuel >= a {
				if i+1 == k {
					break
				}
				// 可以到达位置a
				if fuel-a >= a-f {
					fuel -= a
				} else {
					// 无法到达下一次加油站
					res++
					fuel = b - (a - f)
				}
			} else {
				res++
				fuel = b - (a - f)
			}
			pos = a
		} else {
			// pos = a
			if fuel < a-f {
				return -1
			}
			if fuel >= a {
				if i+1 == k {
					break
				}
				if fuel-a >= f {
					fuel -= a
				} else {
					res++
					fuel = b - f
				}
			} else {
				res++
				fuel = b - f
			}
			pos = 0
		}
	}

	if fuel < 0 {
		return -1
	}

	return res
}
