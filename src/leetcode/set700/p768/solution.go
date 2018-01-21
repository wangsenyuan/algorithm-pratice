package p768

func maxChunksToSorted(arr []int) int {
	n := len(arr)
	var ans int

	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		f[i] = i
		if i < n-1 && arr[i] >= arr[f[i+1]] {
			f[i] = f[i+1]
		}
	}

	var j int
	for j < n {
		hat := arr[j]
		i := j + 1
		for i < n && arr[f[i]] < hat {
			if arr[i] > hat {
				hat = arr[i]
			}
			i++
		}
		ans++
		j = i
	}
	return ans
}
