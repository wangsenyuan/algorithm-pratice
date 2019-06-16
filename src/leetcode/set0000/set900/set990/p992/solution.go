package p992

func subarraysWithKDistinct(A []int, K int) int {
	var j = 0
	var k = 0

	window1 := InitWindow()
	window2 := InitWindow()

	var ans int
	for i := 0; i < len(A); i++ {
		window1.Add(A[i])
		window2.Add(A[i])

		for window1.Count() > K {
			window1.Remove(A[j])
			j++
		}
		for window2.Count() >= K {
			window2.Remove(A[k])
			k++
		}

		ans += k - j
	}
	return ans
}

type Window struct {
	nums map[int]int
}

func InitWindow() Window {
	return Window{make(map[int]int)}
}

func (wnd *Window) Add(x int) {
	wnd.nums[x]++
}

func (wnd *Window) Remove(x int) {
	if wnd.nums[x] == 1 {
		delete(wnd.nums, x)
	} else {
		wnd.nums[x]--
	}
}

func (wnd *Window) Count() int {
	return len(wnd.nums)
}
