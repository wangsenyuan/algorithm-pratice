package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, K := readTwoNums(reader)

	A := readNNums(reader, N)

	fmt.Println(solve(N, K, A))
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

func solve(N, K int, A []int) int64 {

	var loop func(l, r int) int64

	loop = func(l, r int) int64 {
		if r <= l {
			return 0
		}

		var mv int = A[l]

		for i := l + 1; i <= r; i++ {
			if A[i] < mv {
				mv = A[i]
			}
		}

		mid := (l + r) / 2
		var ind = mid
		for i := 0; mid-i >= l || mid+i <= r; i++ {
			if mid-i >= l && A[mid-i] == mv {
				ind = mid - i
				break
			}
			if mid+i <= r && A[mid+i] == mv {
				ind = mid + i
				break
			}
		}

		ans := loop(l, ind-1) + loop(ind+1, r)

		mem := make(map[int]int64)
		var sum int
		for i := ind - 1; i >= l; i-- {
			sum += A[i]
			sum %= K
			if sum == 0 {
				ans++
			}
			mem[sum]++
		}
		sum = 0
		for i := ind + 1; i <= r; i++ {
			sum += A[i]
			sum %= K
			if sum == 0 {
				ans++
			}
			ans += mem[(K-sum)%K]
		}

		return ans
	}

	return loop(0, N-1)
}
