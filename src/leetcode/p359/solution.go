package main

//Logger is the solution
type Logger struct {
	records map[string]int
}

/** Constructor Initialize your data structure here. */
func Constructor() Logger {
	return Logger{make(map[string]int)}
}

/** Shouldprintmessage Returns true if the message should be printed in the given timestamp, otherwise returns false. The timestamp is in seconds granularity. */
func (logger *Logger) Shouldprintmessage(timestamp int, message string) bool {
	record := false
	if ts, found := logger.records[message]; !found {
		record = true
	} else if timestamp-ts >= 10 {
		record = true
	}
	if record {
		logger.records[message] = timestamp
	}
	return record
}
