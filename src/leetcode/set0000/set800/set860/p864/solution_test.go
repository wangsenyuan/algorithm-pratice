package p864

import "testing"

func runSample(t *testing.T, N int, blackList []int, pickTimes int) {
	solution := Constructor(N, blackList)
	blacked := make(map[int]bool)
	for _, num := range blackList {
		blacked[num] = true
	}

	cnt := make(map[int]int)
	for i := 0; i < pickTimes; i++ {
		res := solution.Pick()
		if blacked[res] {
			t.Errorf("Sample pick a black number %d at time %d", res, i)
			return
		}
		cnt[res]++
	}

	t.Logf("%v\n", cnt)
}

func TestSample1(t *testing.T) {
	N := 1
	runSample(t, N, nil, 3)
}

func TestSample2(t *testing.T) {
	N := 2
	runSample(t, N, nil, 3)
}

func TestSample3(t *testing.T) {
	N := 3
	black := []int{1}
	runSample(t, N, black, 3)
}

func TestSample4(t *testing.T) {
	N := 4
	black := []int{2}
	runSample(t, N, black, 3)
}

func TestSample5(t *testing.T) {
	N := 4
	black := []int{0, 1}
	runSample(t, N, black, 20)
}

func TestSample6(t *testing.T) {
	N := 4
	black := []int{0, 2}
	runSample(t, N, black, 20)
}
