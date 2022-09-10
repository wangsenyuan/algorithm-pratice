package main

import (
	"testing"
	"time"
)

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "2019:03:31"
	expect := 1
	runSample(t, s, expect)
}


func TestNextDay(t *testing.T) {
	now, _ := time.Parse(YYYYMMDD, "2022:09:10")

	year := now.Year()
	month := int(now.Month())
	day := now.Day()

	for i := 1; i <= 2000; i++ {
		new_year, new_month, new_day := nextDay(year, month, day)

		new_time := now.Add(time.Hour * 24)

		expect_year, expect_month, expect_day := new_time.Year(), new_time.Month(), new_time.Day()

		if new_year != expect_year || new_month != int(expect_month) || new_day != expect_day {
			t.Fatalf("Sample exepect %d:%d%d, but got %d:%d:%d after %d days", expect_year, expect_month, expect_day, new_year, new_month, new_day, i)
		}
		year, month, day = new_year, new_month, new_day
		now = new_time
	}
}
