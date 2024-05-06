Ivan is a student at Berland State University (BSU). There are n days in Berland week, and each of these days Ivan might
have some classes at the university.

There are m working hours during each Berland day, and each lesson at the university lasts exactly one hour. If at some
day Ivan's first lesson is during i-th hour, and last lesson is during j-th hour, then he spends j - i + 1 hours in the
university during this day. If there are no lessons during some day, then Ivan stays at home and therefore spends 0
hours in the university.

Ivan doesn't like to spend a lot of time in the university, so he has decided to skip some lessons. He cannot skip more
than k lessons during the week. After deciding which lessons he should skip and which he should attend, every day Ivan
will enter the university right before the start of the first lesson he does not skip, and leave it after the end of the
last lesson he decides to attend. If Ivan skips all lessons during some day, he doesn't go to the university that day at
all.

Given n, m, k and Ivan's timetable, can you determine the minimum number of hours he has to spend in the university
during one week, if he cannot skip more than k lessons?

### ideas

1. dp[i][j] 表示到第i天，且skip j个课程时的最小在校时间
2. dp[i+1][j] = dp[i][j] + duration at i+1 day
3.            = dp[i-1][j-x] + 在这一天skip课程后的最小时间
4. 问题出在skip x个课程，首先，可以确定的是，最好skip两头的课程
5. fp[i][x] 表示在第i天skip x 个课程的最小时间
6. = 剩余cnt - x 课程最小的值