package p1235

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)

	jobs := make([]*Job, n)

	for i := 0; i < n; i++ {
		job := new(Job)
		job.start = startTime[i]
		job.end = endTime[i]
		job.profit = profit[i]
		jobs[i] = job
	}

	sort.Sort(Jobs(jobs))

	fp := make([]int, n)
	for i := 0; i < n; i++ {
		job := jobs[i]
		j := sort.Search(i, func(j1 int) bool {
			return jobs[j1].end > job.start
		})
		j--
		earn := job.profit
		if j >= 0 {
			// dp[j] <= job.start
			earn += fp[j]
		}
		if i > 0 {
			earn = max(earn, fp[i-1])
		}
		fp[i] = earn
	}

	return fp[n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Job struct {
	start  int
	end    int
	profit int
}

type Jobs []*Job

func (jobs Jobs) Len() int {
	return len(jobs)
}

func (jobs Jobs) Less(i, j int) bool {
	return jobs[i].end < jobs[j].end
}

func (jobs Jobs) Swap(i, j int) {
	jobs[i], jobs[j] = jobs[j], jobs[i]
}
