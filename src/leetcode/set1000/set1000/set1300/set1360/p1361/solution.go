package p1361

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	if n == 0 {
		return true
	}

	parent := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = -1
	}

	for i := 0; i < n; i++ {
		j := leftChild[i]
		if j == -1 {
			continue
		}
		if parent[j] >= 0 {
			return false
		}
		parent[j] = i
	}

	for i := 0; i < n; i++ {
		j := rightChild[i]
		if j == -1 {
			continue
		}
		if parent[j] >= 0 {
			return false
		}
		parent[j] = i
	}

	var rootCnt int

	for i := 0; i < n; i++ {
		if parent[i] < 0 {
			rootCnt++
		}
	}

	return rootCnt == 1
}
