package main

import "fmt"

func main() {
	var N int

	fmt.Scanf("%d", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	var Q int
	fmt.Scanf("%d", &Q)
	queries := make([][]int, Q)
	for i := 0; i < Q; i++ {
		var action string
		var a, b int
		fmt.Scanf("%s %d %d", &action, &a, &b)
		if action == "change" {
			queries[i] = []int{1, a, b}
		} else {
			queries[i] = []int{2, a, b}
		}
	}
	res := solve(N, A, Q, queries)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

const MOD = 3046201

var facts []int64

func init() {
	facts = make([]int64, MOD)
	facts[0] = 1
	for i := 1; i < MOD; i++ {
		facts[i] = (int64(i) * int64(facts[i-1])) % MOD
	}
}

func solve(N int, A []int, Q int, queries [][]int) []int {
	bit := NewBit(N)

	for i := 0; i < N; i++ {
		bit.Update(i, A[i])
	}

	calculate := func(sum int, k int) int {
		x := sum / k
		i := sum % k

		res := facts[k]
		res = divide(res, facts[i])
		res = divide(res, facts[k-i])
		res = (res * facts[sum]) % MOD
		res = divide(res, pow(facts[x+1], i))
		res = divide(res, pow(facts[x], k-i))

		return int(res)
	}

	res := make([]int, Q)
	var j int

	for i := 0; i < Q; i++ {
		query := queries[i]
		if query[0] == 1 {
			// change
			pos, val := query[1]-1, query[2]
			bit.Update(pos, val-A[pos])
			A[pos] = val
		} else {
			// query
			l, r := query[1]-1, query[2]-1
			if l > r {
				l, r = r, l
			}
			sum := bit.Get(r) - bit.Get(l-1)
			res[j] = calculate(sum, r-l+1)
			j++
		}
	}
	return res[:j]
}

func pow(a int64, n int) int64 {
	var res int64 = 1
	for n > 0 {
		if n&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		n >>= 1
	}
	return res
}

func divide(a, b int64) int64 {
	return (a * inverse(b)) % MOD
}

func inverse(x int64) int64 {
	return pow(x, MOD-2)
}

type BIT []int

func NewBit(n int) BIT {
	arr := make([]int, n+1)
	return BIT(arr)
}

func (bit BIT) Update(pos int, val int) {
	pos++
	n := len(bit) - 1
	for pos <= n {
		bit[pos] += val
		pos += pos & (-pos)
	}
}

func (bit BIT) Get(pos int) int {
	pos++
	var res int
	for pos > 0 {
		res += bit[pos]
		pos -= pos & (-pos)
	}
	return res
}
