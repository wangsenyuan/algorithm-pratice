package main

import "testing"

func runSample(t *testing.T, statements []string, expect string) {
	res := solve(statements)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	statements := []string{
		"try",
		"try",
		"throw ( AE )",
		"catch ( BE, \"BE in line 3\")",
		"",
		"try",
		"catch(AE, \"AE in line 5\")",
		"catch(AE,\"AE somewhere\")",
	}
	expect := "AE somewhere"
	runSample(t, statements, expect)
}

func TestSample2(t *testing.T) {
	statements := []string{
		"try",
		"try",
		"throw ( AE )",
		"catch ( AE, \"AE in line 3\")",

		"try",
		"catch(BE, \"BE in line 5\")",
		"catch(AE,\"AE somewhere\")",
	}
	expect := "AE in line 3"
	runSample(t, statements, expect)
}
