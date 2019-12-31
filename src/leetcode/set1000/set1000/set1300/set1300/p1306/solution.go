package p1306

func canReach(arr []int, start int) bool {
	n := len(arr)
	vis := make([]bool, n)
	vis[start] = true
	que := make([]int, n)
	var end int
	que[end] = start
	end++
	var front int
	for front < end {
		cur := que[front]
		front++
		if arr[cur] == 0 {
			return true
		}
		x := cur + arr[cur]
		if x < n && !vis[x] {
			vis[x] = true
			que[end] = x
			end++
		}
		x = cur - arr[cur]
		if x >= 0 && !vis[x] {
			vis[x] = true
			que[end] = x
			end++
		}
	}
	return false
}
