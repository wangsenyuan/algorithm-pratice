package p2281

const MOD = 1000000007

func totalStrength(strength []int) int {

	// let fix i, and it is the min between [l...r]
	// then we want to get A[i] * total_sum_of arr[l...r] that having i as min
	// for j between [i..r]
	// for every k between [l..i-1]
	// sum_of  sum[j] - sum[k]  = (i - 1 - l + 1) * sum[j] - (sum[l-1] + sum[l] ... sum[i-1])
	// let S1 = sum[l-1] + sum[l] + ... sum[i-1]
	// (i - l) * sum[j] for every j between [i...r]
	// (i - l) * (sum[i] + sum[i+1] + ... sum[r])
	// let S2 =  sum[i] + sum[i+1] + ... sum[r]
	// (i - l) * S2 - (r - i + 1) * S1
	n := len(strength)
	sum := make([]int64, n)
	sum_of_sum := make([]int64, n)

	for i := 0; i < n; i++ {
		sum[i] = int64(strength[i])
		if i > 0 {
			sum[i] = modAdd(sum[i], sum[i-1])
		}
		sum_of_sum[i] = sum[i]
		if i > 0 {
			sum_of_sum[i] = modAdd(sum_of_sum[i], sum_of_sum[i-1])
		}
	}
	L := make([]int, n)
	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		for p > 0 && strength[stack[p-1]] > strength[i] {
			p--
		}

		L[i] = -1
		if p > 0 {
			L[i] = stack[p-1]
		}
		L[i]++

		stack[p] = i
		p++
	}

	p = 0

	var res int64

	for i := n - 1; i >= 0; i-- {
		for p > 0 && strength[stack[p-1]] >= strength[i] {
			p--
		}
		r := n
		if p > 0 {
			r = stack[p-1]
		}
		r--
		l := L[i]
		// [l...r] is the range having str[i] as min
		var S1, S2 int64
		if i > 0 {
			S1 = sum_of_sum[i-1]
			if l > 1 {
				S1 = modSub(S1, sum_of_sum[l-2])
			}
		}
		S1 = modMul(S1, int64(r-i+1))

		S2 = sum_of_sum[r]
		if i > 0 {
			S2 = modSub(S2, sum_of_sum[i-1])
		}

		S2 = modMul(S2, int64(i-l+1))

		tmp := modSub(S2, S1)

		tmp = modMul(tmp, int64(strength[i]))

		res = modAdd(res, tmp)

		stack[p] = i
		p++
	}

	return int(res)
}

func modAdd(a, b int64) int64 {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int64) int64 {
	return modAdd(a, MOD-b)
}

func modMul(a, b int64) int64 {
	a *= b
	a %= MOD
	return a
}
