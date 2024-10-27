package p3335

func lengthAfterTransformations(s string, t int) int {
	nums := make([]int, 26)
	for i := 0; i < 26; i++ {
		nums[i] = 1
	}
	nums[25]++
	
	a := NewMatrix(26, 26)
	for i := 0; i < 26; i++ {
		for j := i + 1; j <= i+nums[i]; j++ {
			a[i][j%26] = 1
		}
	}
	a = Pow(a, t)

	v := NewMatrix(26, 1)
	for i := 0; i < 26; i++ {
		v[i][0] = 1
	}

	a = a.Mul(v)
	var res int

	for _, x := range []byte(s) {
		i := int(x - 'a')
		res = add(res, a[i][0])
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
	return a * b % mod
}

type matrix [][]int

func NewMatrix(n int, m int) matrix {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	return arr
}

func identity(n int) matrix {
	a := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		a[i][i] = 1
	}
	return a
}

func (a matrix) Mul(b matrix) matrix {
	n := len(a)
	m := len(b)
	k := len(b[0])

	c := NewMatrix(n, k)

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for u := 0; u < m; u++ {
				c[i][j] = add(c[i][j], mul(a[i][u], b[u][j]))
			}
		}
	}
	return c
}

func Pow(a matrix, n int) matrix {
	res := identity(len(a))

	for n > 0 {
		if n&1 == 1 {
			res = res.Mul(a)
		}
		a = a.Mul(a)
		n >>= 1
	}
	return res

}
