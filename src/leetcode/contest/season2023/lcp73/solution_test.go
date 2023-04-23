package lcp73

import "testing"

func runSample(t *testing.T, expeditions []string, expect int) {
	res := adventureCamp(expeditions)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	expeditions := []string{
		"Alice->Dex", "", "Dex",
	}
	expect := -1
	runSample(t, expeditions, expect)
}
