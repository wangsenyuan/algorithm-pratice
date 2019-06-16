package p920

import "testing"

func runSample(t *testing.T, N int, L int, K int, expect int) {
	res := numMusicPlaylists(N, L, K)
	if res != expect {
		t.Errorf("Sample %d %d %d, expect %d, but got %d", N, L, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 1, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 0, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, 1, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 2, 0, 1)
}
