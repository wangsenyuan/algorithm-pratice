package p304

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	var n int
	if m > 0 {
		n = len(matrix[0])
	}
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		sum[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			sum[i+1][j+1] = matrix[i][j]
			sum[i+1][j+1] += sum[i][j+1]
			sum[i+1][j+1] += sum[i+1][j]
			sum[i+1][j+1] -= sum[i][j]
		}
	}
	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := this.sum[row2+1][col2+1]
	res -= this.sum[row1][col2+1]
	res -= this.sum[row2+1][col1]
	res += this.sum[row1][col1]
	return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
