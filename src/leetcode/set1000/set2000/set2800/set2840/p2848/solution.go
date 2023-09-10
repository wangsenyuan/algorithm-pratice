package p2848

func numberOfWays(s string, t string, k int64) int {
	pi := kmp(t)

	cnt := countOccurrence(s+s, t, pi)

	mat := NewMatrix(2)
	mat[0][0] = cnt - 1
	mat[0][1] = cnt
	mat[1][0] = len(s) - cnt
	mat[1][1] = len(s) - 1 - cnt

	res := PowMatrix(mat, k)

	if s == t {
		return res[0][0]
	}
	return res[0][1]
}

type matrix [][]int

func NewMatrix(n int) matrix {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	return res
}

func IdentityMatrix(n int) matrix {
	res := NewMatrix(n)
	for i := 0; i < n; i++ {
		res[i][i] = 1
	}
	return res
}

func MulMatrix(a, b matrix) matrix {
	n := len(a)
	c := NewMatrix(n)
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				c[i][j] = add(c[i][j], mul(a[i][k], b[k][j]))
			}
		}
	}
	return c
}

func PowMatrix(a matrix, b int64) matrix {
	res := IdentityMatrix(len(a))
	for b > 0 {
		if b&1 == 1 {
			res = MulMatrix(res, a)
		}
		a = MulMatrix(a, a)
		b >>= 1
	}
	return res
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func countOccurrence(s string, t string, lps []int) int {
	var res int
	var j int

	for i := 0; i < len(s); i++ {
		for j > 0 && s[i] != t[j] {
			j = lps[j-1]
		}
		if s[i] == t[j] {
			j++
		}
		if j == len(t) {
			if i-len(t)+1 < len(s)/2 {
				res++
			}
			j = lps[j-1]
		}
	}
	return res
}

func kmp(pat string) []int {
	n := len(pat)
	dp := make([]int, n)

	for i := 1; i < n; i++ {
		j := dp[i-1]
		for j > 0 && pat[i] != pat[j] {
			j = dp[j-1]
		}
		if pat[i] == pat[j] {
			j++
		}
		dp[i] = j
	}
	return dp
}
