package p907

const MOD = 1000000007

func sumSubarrayMins(A []int) int {
	n := len(A)
	stack := make([]int, n)
	var p int
	left := make([]int, n)

	for i := 0; i < n; i++ {
		for p > 0 && A[stack[p-1]] >= A[i] {
			p--
		}
		if p > 0 {
			left[i] = stack[p-1]
		} else {
			left[i] = -1
		}
		stack[p] = i
		p++
	}

	right := make([]int, n)

	p = 0
	for i := n - 1; i >= 0; i-- {
		for p > 0 && A[stack[p-1]] > A[i] {
			p--
		}
		if p > 0 {
			right[i] = stack[p-1]
		} else {
			right[i] = n
		}
		stack[p] = i
		p++
	}

	var res int64

	for i := 0; i < n; i++ {
		a := int64(i - left[i])
		b := int64(right[i] - i)
		c := (((a * b) % MOD) * int64(A[i])) % MOD
		res += c
		if res >= MOD {
			res -= MOD
		}
	}

	return int(res)
}

/*func pow(a, n int) int64 {
	A := int64(a)
	res := int64(1)

	for n > 0 {
		if n&1 == 1 {
			res = (res * A) % MOD
		}
		A = (A * A) % MOD
		n >>= 1
	}
	return res
}
*/
