package p1773

import "testing"

func runSample(t *testing.T, items [][]string, ruleKey string, ruleValue string, expect int) {
	res := countMatches(items, ruleKey, ruleValue)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	items := [][]string{
		{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"},
	}
	ruleKey := "type"
	ruleValue := "phone"
	expect := 2
	runSample(t, items, ruleKey, ruleValue, expect)
}
