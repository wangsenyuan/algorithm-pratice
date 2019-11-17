package p1257

import "testing"

func runSample(t *testing.T, regions [][]string, region1 string, region2 string, expect string) {
	res := findSmallestRegion(regions, region1, region2)
	if res != expect {
		t.Errorf("Sample %v %s %s, expect %s, but got %s", regions, region1, region2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	regions := [][]string{
		{"Earth", "North America", "South America"},
		{"North America", "United States", "Canada"},
		{"United States", "New York", "Boston"},
		{"Canada", "Ontario", "Quebec"},
		{"South America", "Brazil"},
	}
	region1 := "Quebec"
	region2 := "New York"
	expect := "North America"
	runSample(t, regions, region1, region2, expect)
}
