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

	nums := readNNums(reader, n)

	fmt.Println(solve(n, nums))

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func solve(n int, nums []int) int {
	sort.Ints(nums)
	can := make([]bool, 10)

	for _, num := range nums {
		can[num] = true
	}

	var res int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				num := nums[i]*100 + nums[j]*10 + nums[k]

				for a := 0; a < n; a++ {
					x := num * nums[a]

					if x >= 1000 {
						break
					}
					if !can[x/100] || !can[x/10%10] || !can[x%10] {
						continue
					}

					for b := 0; b < n; b++ {
						x = num * nums[b]

						if num*nums[b] >= 1000 {
							break
						}

						if !can[x/100] || !can[x/10%10] || !can[x%10] {
							continue
						}

						num2 := nums[a]*10 + nums[b]
						if num*num2 >= 10000 {
							break
						}

						x = num * num2

						if !can[x/1000] || !can[x/100%10] || !can[x/10%10] || !can[x%10] {
							continue
						}

						res++
					}
				}
			}
		}
	}

	return res
}
