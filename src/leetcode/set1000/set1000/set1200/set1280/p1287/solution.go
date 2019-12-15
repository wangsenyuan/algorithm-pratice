
func findSpecialInteger(arr []int) int {
	n := len(arr)

	var j int

	for i := 1; i <= n; i++ {
		if i == n || arr[i] > arr[i-1] {
			cnt := i - j
			if 4*cnt > n {
				return arr[i-1]
			}
			j = i
		}
	}

	return -1
}
package p1287
