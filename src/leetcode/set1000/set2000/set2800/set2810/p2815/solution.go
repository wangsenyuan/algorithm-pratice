package p2815

import "sort"

func maximumScore(nums []int, k int) int {
	x := max_of(nums)
	pf := make([]int, x+1)
	for i := 2; i <= x; i++ {
		if pf[i] == 0 {
			pf[i] = i
			if x/i < i {
				continue
			}
			for j := i * i; j <= x; j += i {
				if pf[j] == 0 {
					pf[j] = i
				}
			}
		}
	}
	n := len(nums)
	lfc := make([]int, n)

	for i, cur := range nums {
		for cur > 1 {
			p := pf[cur]
			for cur%p == 0 {
				cur /= p
			}
			lfc[i]++
		}
	}
	stack := make([]int, n)
	var p int
	L := make([]int, n)

	for i := 0; i < n; i++ {
		for p > 0 && lfc[stack[p-1]] < lfc[i] {
			p--
		}
		if p == 0 {
			L[i] = -1
		} else {
			L[i] = stack[p-1]
		}
		stack[p] = i
		p++
	}
	p = 0
	R := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for p > 0 && lfc[stack[p-1]] <= lfc[i] {
			p--
		}
		if p == 0 {
			R[i] = n
		} else {
			R[i] = stack[p-1]
		}
		stack[p] = i
		p++
	}
	type Num struct {
		val int
		cnt int64
	}
	arr := make([]Num, n)
	for i := 0; i < n; i++ {
		arr[i] = Num{nums[i], int64(i-L[i]) * int64(R[i]-i)}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val > arr[j].val
	})

	res := 1

	for i := 0; i < n && k > 0; i++ {
		cur := arr[i]
		if cur.cnt <= int64(k) {
			res = mul(res, pow(cur.val, int(cur.cnt)))
			k -= int(cur.cnt)
		} else {
			// cur.cnt > k
			res = mul(res, pow(cur.val, k))
			k = 0
		}
	}

	return res
}

const MOD = 1e9 + 7

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func max_of(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
