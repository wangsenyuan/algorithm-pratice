Recently Monocarp got a job. His working day lasts exactly 𝑚
minutes. During work, Monocarp wants to drink coffee at certain moments: there are 𝑛
minutes 𝑎1,𝑎2,…,𝑎𝑛
, when he is able and willing to take a coffee break (for the sake of simplicity let's consider that each coffee break
lasts exactly one minute).

However, Monocarp's boss doesn't like when Monocarp takes his coffee breaks too often. So for the given coffee break
that is going to be on minute 𝑎𝑖
, Monocarp must choose the day in which he will drink coffee during the said minute, so that every day at least 𝑑
minutes pass between any two coffee breaks. Monocarp also wants to take these 𝑛
coffee breaks in a minimum possible number of working days (he doesn't count days when he is not at work, and he doesn't
take coffee breaks on such days). Take into account that more than 𝑑
minutes pass between the end of any working day and the start of the following working day.

For each of the 𝑛
given minutes determine the day, during which Monocarp should take a coffee break in this minute. You have to minimize
the number of days spent.

### ideas

1. 非常难理解这个题目
2. 这n个coffee break的时间，必须都安排掉，同一天内两个break之间的时间 > d
3. 且要求使用的天数要少
4. 可以二分查找，如果n天内能完成，那肯定可以在n+1天内完成
5. 对于给定的n天，要怎么判断呢？应该尽可能的早的安排下一个break