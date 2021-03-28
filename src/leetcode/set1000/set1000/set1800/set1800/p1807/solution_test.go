package p1807

import "testing"

func runSample(t *testing.T, word string, know [][]string, expect string) {
	res := evaluate(word, know)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word := "(name)is(age)yearsold"
	know := [][]string{
		{"name", "bob"}, {"age", "two"},
	}
	expect := "bobistwoyearsold"
	runSample(t, word, know, expect)
}
