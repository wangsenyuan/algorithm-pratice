Jill loves having good grades in university, so she never misses deadlines for her homework assignments. But even more, she loves watching the series and discussing it with her best friend Johnny. And unfortunately, today she needs to choose between these two activities!

Jill needs to complete 𝑛
 homework tasks. The 𝑖
-th task would require 𝑎𝑖
 minutes to complete and needs to be submitted to the teacher at most 𝑑𝑖
 minutes from now. Also, there are 𝑚
 new episodes of the series that Johnny and Jill want to discuss. The 𝑗
-th episode lasts 𝑙𝑗
 minutes. Jill can complete tasks in any order, but she needs to watch the episodes in the order they come. Neither completing a homework task nor watching an episode can be interrupted after starting.

Johnny and Jill need to agree on a time 𝑡𝑘
 when they would have a call to discuss the series. They are not sure yet which time to choose. For each possible time, compute the maximum number of episodes Jill could watch before that time while still being able to complete all 𝑛
 homework tasks in time.

Note that for the purpose of this problem we assume that discussing the series with Johnny at time 𝑡𝑘
 does not consume significant time from Jill and can happen even if she is in the middle of completing any of her homework tasks.

Input
There are several test cases in the input. The input begins with the number of test cases 𝑇
 (1≤𝑇≤1000
).

Each test case starts with a line with three integers 𝑛
 (1≤𝑛≤200000
) — the number of homework tasks, 𝑚
 (1≤𝑚≤200000
) — the number of episodes, and 𝑞
 (1≤𝑞≤200000
) — the number of possible times for the call with Jill.

The second line contains 𝑛
 integers 𝑎𝑖
 (1≤𝑎𝑖≤109
) — the number of minutes it takes to complete the task. The next line contains 𝑛
 integers 𝑑𝑖
 (1≤𝑑𝑖≤1015
) — the deadline before which this task must be completed. The next line contains 𝑚
 integers 𝑙𝑗
 (1≤𝑙𝑗≤109
) — the length of episodes in the order they need to be watched. The next line contains 𝑞
 integers 𝑡𝑘
 (1≤𝑡𝑘≤1015
) — the possible times of call with Jill.

It is possible to complete all tasks within their respective deadlines.

The sum of each of 𝑛
, 𝑚
, 𝑞
 over all test cases in input doesn't exceed 200000
.

Output
For each test case output a single line with 𝑞
 integers — for each possible time 𝑡𝑘
 the maximum number of episodes Jill can watch.

 ### ideas
 1. 作业似乎留到最后去完成。
 2. 但是可能存在这样一种情况，作业i,i+1中间的gap，对于episode?有点小，但如果将i前移一段时间，就足够了
 3. q个tk是不是独立的，如果是独立的话，答案似乎有所不同
 4. 对于给定的t，就是尽量将作业留到后面（保证能够全部按时完成的前提下，在前面看电视剧）
 5. 在前面的时间内，尽量使用短时间的电视剧
 6. 那些最晚开始时间(d-a)在t后面的，可以不用考虑
 7. 这样子感觉不大行。因为前面这部分作业和电视剧时间，不同的排列会造成不同的结果
 8. 假设当前是某个开始时刻（0），到下一个开始时间，（d[1] - a[1])
 9. 之间，如果能够放入足够多的（升序的电视剧）就放入，假设出现了gap（那么就调整下一个时刻任务的开始时间， 迁移）
 10. 