package p1185

import "testing"

func runSample(t *testing.T, day int, month int, year int, expect string) {
	res := dayOfTheWeek(day, month, year)

	if res != expect {
		t.Errorf("Sample %d.%d.%d, expect %s, but got %s", day, month, year, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 31, 8, 2019, "Saturday")
}

func TestSample2(t *testing.T) {
	runSample(t, 18, 7, 1999, "Sunday")
}

func TestSample3(t *testing.T) {
	runSample(t, 15, 8, 1993, "Sunday")
}

func TestSample4(t *testing.T) {
	runSample(t, 29, 2, 2016, "Monday")
}

func TestSample5(t *testing.T) {
	runSample(t, 1, 2, 2014, "Saturday")
}
