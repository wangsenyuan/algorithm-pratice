package p2600

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	}
	if k <= numOnes+numZeros {
		return numOnes
	}
	x := k - (numOnes + numZeros)
	return numOnes - x
}
