package main

import "testing"

func TestIsValidSerialization(t *testing.T) {
	str := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	if isValidSerialization(str) == false {
		t.Errorf("%v is valid\n", str)
	}
}
