package p842

import (
	"testing"
	"bytes"
	"strconv"
)

func runSample(t *testing.T, S string, can bool) {
	res := splitIntoFibonacci(S)

	if !can && len(res) == 0 {
		return
	}

	if can && len(res) == 0 {
		t.Errorf("Sample %s, expect %t, but got empty result", S, can)
	} else if can && len(res) < 3 {
		t.Errorf("Sample %s, result should have at least 3 elements, but got %v", S, res)
	} else if can {
		for i := 2; i < len(res); i++ {
			if res[i] != res[i-1]+res[i-2] {
				t.Errorf("Sample %s, result %v, not correct", S, res)
				break
			}
		}
		var buf bytes.Buffer
		for _, num := range res {
			buf.WriteString(strconv.Itoa(num))
		}

		if S != buf.String() {
			t.Errorf("Sample %s, result %v not correct", S, buf.String())
		}
	} else {
		t.Errorf("Sample %s, expect emtpy, but got %v", S, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "123456579", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "112358130", false)
}

func TestSample3(t *testing.T) {
	runSample(t, "0123", false)
}

func TestSample4(t *testing.T) {
	runSample(t, "1101111", true)
}

func TestSample5(t *testing.T) {
	runSample(t, "539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511", false)
}
