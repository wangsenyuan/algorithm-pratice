package p0001

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 */
type BinaryMatrix interface {
	Get(int, int) int
	Dimensions() []int
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	m, n := dim[0], dim[1]

	if m == 0 || n == 0 {
		return -1
	}

	i, j := 0, n-1

	for i < m && j >= 0 {
		x := binaryMatrix.Get(i, j)
		if x == 1 {
			j--
		} else {
			i++
		}
	}

	if i == m {
		if j == n-1 {
			j = -1
		} else {
			j++
		}
	} else {
		j++
	}

	return j
}
