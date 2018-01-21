package p769

func maxChunksToSorted(arr []int) int {
	var sum int
	n := len(arr)
	var ans int
	for i := 0; i < n; i++ {
		sum += arr[i]
		if sum == (i+1)*i/2 {
			ans++
		}
	}
	return ans
}
