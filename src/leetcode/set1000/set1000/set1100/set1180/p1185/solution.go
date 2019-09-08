package p1185

var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

var DAY_OF_WEEK = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func dayOfTheWeek(day int, month int, year int) string {
	// today is 2019-09-08, sunday
	// and we need to know 2019-01-01
	var diff int
	if year < 2019 || (year == 2019 && month < 9) || (year == 2019 && month == 9 && day <= 8) {
		diff = calDaysBetween([]int{day, month, year}, []int{8, 9, 2019})
		diff %= 7
		if diff != 0 {
			diff = 7 - diff
		}
	} else if year > 2019 {
		diff = calDaysBetween([]int{8, 9, 2019}, []int{day, month, year})
		diff %= 7
	}

	return DAY_OF_WEEK[diff]
}

func calDaysBetween(from []int, to []int) int {
	day := from[0]
	month := from[1]
	year := from[2]

	daysDiff := -day
	day = 1

	if year < to[2] {
		// go to year + 1, 01, 01
		if isLeapYear(year) && month == 2 {
			daysDiff += 29 - day + 1
		} else {
			daysDiff += monthDays[month-1] - day + 1

		}
		month++

		day = 1

		for month <= 12 {
			daysDiff += monthDays[month-1]
			if month == 2 && isLeapYear(year) {
				daysDiff++
			}
			month++
		}

		month = 1
		year++

		for year < to[2] {
			daysDiff += 365
			if isLeapYear(year) {
				daysDiff++
			}
			year++
		}
	}

	// 2019-01-01
	for month < to[1] {
		daysDiff += monthDays[month-1]
		if month == 2 && isLeapYear(year) {
			daysDiff++
		}
		month++
	}
	// 2019-09-01

	daysDiff += to[0] - day + 1

	return daysDiff
}

func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 != 0 {
		return true
	}

	return year%400 == 0
}
