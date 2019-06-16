package main

//Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	result := make([]Interval, 0, len(intervals))

	processed := false

	for _, interval := range intervals {
		if processed {
			result = append(result, interval)
			continue
		}

		if interval.Start > newInterval.End {
			processed = true
			result = append(result, newInterval)
			result = append(result, interval)
			continue
		}

		if interval.End < newInterval.Start {
			result = append(result, interval)
			continue
		}

		if interval.Start < newInterval.Start {
			newInterval.Start = interval.Start
		}

		if interval.End > newInterval.End {
			newInterval.End = interval.End
		}
	}

	if !processed {
		result = append(result, newInterval)
	}

	return result
}
