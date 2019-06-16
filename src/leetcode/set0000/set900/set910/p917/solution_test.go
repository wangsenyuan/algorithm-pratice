package p917

import "testing"

func runSample(t *testing.T, S string, expect string) {
	res := reverseOnlyLetters(S)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ab-cd", "dc-ba")
}

func TestSample2(t *testing.T) {
	runSample(t, "a-bC-dEf-ghIj", "j-Ih-gfE-dCba")
}

func TestSample3(t *testing.T) {
	runSample(t, "Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!")
}
