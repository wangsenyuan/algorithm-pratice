package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("./C-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var tc int
	fmt.Fscanf(f, "%d", &tc)

	for i := 1; i <= tc; i++ {
		var N, m, X, Y, Z int
		fmt.Fscanf(f, "%d %d %d %d %d", &N, &m, &X, &Y, &Z)
		A := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(f, "%d", &A[j])
		}
		fmt.Printf("Case #%d: %d\n", i, solve(N, m, X, Y, Z, A))
	}
}

const MOD = 1000000007

func solve(N, m, X, Y, Z int, A []int) int64 {
	B := make([]int, N)
	for i := 0; i < N; i++ {
		B[i] = A[i%m]
		A[i%m] = int(((int64(X) * int64(A[i%m])) + int64(Y)*int64(i+1)) % int64(Z))
	}
	C := make([]int, N)
	copy(C, B)
	sort.Ints(C)

	ii := make(map[int]int)
	var j int
	for i := 1; i <= N; i++ {
		if i < N && C[i] == C[i-1] {
			continue
		}
		ii[C[i-1]] = j
		j++
	}

	bit := CreateBIT(j)

	var ans int64
	for i := 0; i < N; i++ {
		pos := ii[B[i]]
		tmp := bit.Get(pos - 1)
		tmp++
		if tmp >= MOD {
			tmp -= MOD
		}
		bit.Set(pos, tmp)
		ans += tmp
		if ans >= MOD {
			ans -= MOD
		}
	}

	return ans
}

type BIT struct {
	arr  []int64
	size int
}

func CreateBIT(n int) BIT {
	arr := make([]int64, n+1)
	return BIT{arr, n}
}

func (bit *BIT) Set(pos int, val int64) {
	pos++
	for pos <= bit.size {
		bit.arr[pos] += val
		if bit.arr[pos] >= MOD {
			bit.arr[pos] -= MOD
		}
		pos += pos & (-pos)
	}
}

func (bit BIT) Get(pos int) int64 {
	var ans int64
	pos++
	for pos > 0 {
		ans += bit.arr[pos]
		if ans >= MOD {
			ans -= MOD
		}
		pos -= pos & (-pos)
	}
	return ans
}
