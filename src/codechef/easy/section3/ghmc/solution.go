package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		var N, K int
		pos := readInt(scanner.Bytes(), 0, &N)
		pos = readInt(scanner.Bytes(), pos+1, &K)
		var x, D int64
		pos = readInt64(scanner.Bytes(), pos+1, &x)
		readInt64(scanner.Bytes(), pos+1, &D)
		P := readNInt64Nums(scanner, K)
		res := solve(N, K, x, D, P)
		fmt.Println(res)
	}
}

func solve(N, K int, x int64, D int64, P []int64) int64 {
	if x-int64(N) < 0 {
		// can't have N distinct postive numbers
		return -1
	}

	sort.Sort(Int64Slice(P))

	for i := 1; i < K; i++ {
		if P[i] == P[i-1] {
			// not distinct
			return -1
		}
	}

	if P[K-1] > x {
		return -1
	}

	nums := make([]int64, N+10)
	var j int

	nums[j] = P[0]
	j++

	if K > 1 && P[0]+D < P[1] {
		nums[j] = P[0] + D
		j++
	}

	for i := 1; i < K; i++ {
		cur := P[i]
		if cur == nums[j-1] {
			// already in the list
			continue
		}
		if cur-nums[j-1] <= D {
			//safe
			if cur > nums[j-1] {
				nums[j] = cur
				j++
			} else {
				nums[j-1], nums[j] = cur, nums[j-1]
				j++
			}
		} else {
			nums[j] = cur
			j++
			if cur < x {
				nums[j] = min(x, cur+D)
				j++
			} else {
				nums[j-1], nums[j] = x-1, x
				j++
			}
		}
		if j > N {
			return -1
		}
	}
	if j == N {
		return sum(nums, j)
	}

	if N-j == 1 {
		return case1(nums, j, N, x, D)
	}

	return case2(nums, j, N, x, D)
}

func case1(nums []int64, j int, N int, x int64, D int64) int64 {
	y := min(x, nums[j-1]+D)
	i := j - 1
	for i >= 0 {
		if nums[i] != y {
			break
		}
		i--
		y--
	}
	return sum(nums, j) + y
}

func case2(nums []int64, j int, N int, x int64, D int64) int64 {
	if nums[j-1] < x {
		nums[j] = x
		j++
	}
	cnt := N - j
	res := sum(nums, j)

	for cnt > 0 && j > 1 {
		a, b := nums[j-1], nums[j-2]
		c := int(a - b - 1)

		res += sumBetween(b+1, a-1)

		if c > cnt {
			res -= sumBetween(b+1, b+int64(c-cnt))
		}

		cnt -= c
		j--
	}
	if j == 1 && cnt > 0 {
		y := nums[0] - 1
		if cnt > int(y) {
			return -1
		}
		// a number between [y - cnt + 1, ... y - cnt + cnt]
		res += sumBetween(y-int64(cnt)+1, y)
	}

	return res
}

func sumBetween(a, b int64) int64 {
	if a > b {
		return 0
	}
	x := b - a + 1
	return (a + b) * x / 2
}

func sum(nums []int64, n int) int64 {
	var res int64
	for i := 0; i < n; i++ {
		res += nums[i]
	}
	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Int64Slice []int64

func (slice Int64Slice) Len() int {
	return len(slice)
}

func (slice Int64Slice) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice Int64Slice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
