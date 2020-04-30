package main

func main() {

}

func solve(n int, A []uint64, rgs [][]int, K int) []uint64 {
	// K is too large to enumerate
	mat := NewMatrix(n)

	for i := 0; i < n; i++ {
		l, r := rgs[i][0]-1, rgs[i][1]

		for j := l; j < r; j++ {
			mat.Set(i, j)
		}
	}

	mat = Pow(mat, K)

	ans := make([]uint64, n)

	for v := uint64(0); v < 60; v++ {
		vec := NewVector(n)

		for i := 0; i < n; i++ {
			if A[i]&(1<<v) > 0 {
				vec.Set(i, 0)
			}
		}

		vec = vec.Mul(mat)

		for i := 0; i < n; i++ {
			if vec.Get(i, 0) == 1 {
				ans[i] |= 1 << v
			}
		}
	}

	return ans
}

type Matrix struct {
	arr [][]int
	m   int
	n   int
}

func (mat *Matrix) Set(i, j int) {
	mat.arr[i][j] = 1
}

func (mat Matrix) Get(i, j int) int {
	return mat.arr[i][j]
}

func NewVector(n int) Matrix {
	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, 1)
	}

	return Matrix{arr, n, 1}
}

func NewMatrix(n int) Matrix {
	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}

	return Matrix{arr, n, n}
}

func NewUnitMatrix(n int) Matrix {
	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		arr[i][i] = 1
	}

	return Matrix{arr, n, n}
}

func (this Matrix) Mul(that Matrix) Matrix {
	m := this.m
	n := that.n
	//this.n == that.m
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// res[i][j] = sum(this[i][k] * that[k][j])
			for k := 0; k < this.n; k++ {
				res[i][j] += this.arr[i][k] * that.arr[k][j]
				// this is for bitwise xor
				res[i][j] &= 1
			}
		}
	}

	return Matrix{res, m, n}
}

func Pow(a Matrix, n int) Matrix {
	res := NewUnitMatrix(a.n)

	for n > 0 {
		if n&1 == 1 {
			res = res.Mul(a)
		}
		a = a.Mul(a)
		n >>= 1
	}

	return res
}
