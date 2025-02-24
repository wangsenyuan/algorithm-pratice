There are 𝑛
 cities in a row, numbered 1,2,…,𝑛
 left to right.

At time 1
, you conquer exactly one city, called the starting city.
At time 2,3,…,𝑛
, you can choose a city adjacent to the ones conquered so far and conquer it.
You win if, for each 𝑖
, you conquer city 𝑖
 at a time no later than 𝑎𝑖
. A winning strategy may or may not exist, also depending on the starting city. How many starting cities allow you to win?

### ideas
1. 假设从city i开始，获胜策略是什么呢？
2. 假设a[j]是所有中最小的，那么是不是应该先往j移动，然后再在未访问到的城市中的，最小的a[k]
3. 这里j、k应该在i的两边的（如果在同一边，不论j、k的顺序怎么样，往j移动，肯定是个更优的策略）
4. a[j] < a[k], 不管怎么样，在所有的策略中, k肯定在j的后面进行访问的
5. 所以尽快的访问j，同时对k来说也是最优的策略
6. 那是否能够快速的判断i是一个good starting place？
7. 如果存在a[i] = 1, 最多不会超过1（且必须从它开始）
8. 假设按照a[i]顺序处理
9. 如果到达位置i, 开始能够在合理的时间内，访问完所有 1....i-1,(按照上面的想法， 这些应该被优先访问到)
10. 然后才能去访问i+1, i+2.... 
11. 假设i只考虑 i+1, i+2的位置可行吗？（开始访问i+1的时候，前i-1个必须是访问到的）
12. 首先从i+1到后面的都可以访问到，然后考虑i，假设他们之间的距离是w, 那么
13. 好混乱呐
14. 首先检查从最小的i开始，能否完成，如果不行就是0
15. 然后检查，从i开始，如果从它开始最远的不能够在合理内到达的距离
16. 是不是符合二分的条件呢？就是如果从a[i]开始不能够完成，那么a[i+1]（排好序后）也肯定不行呢？
17. 如果符合这个条件的话，倒是可以二分
18. 看样子不符合2分的性质
19. 如果 a[r] or a[l] < r - l + 1, 那么就进行不小去了
20. 而l, r 是目前的边界
21. 所以对于i来说， (它的l。。。r是可以计算出来的，就是所有 < a[i]组成的区域)
22. i可以的前提是，l...r区域中的位置，
23. a[j] > (j - i) => 假设j在它的右边, a[j] > j - i => a[j] - j > -i
24. a[j] >= j - i => i >= j - a[j] => i < j - a[i] 时没有答案
25. 同样的左边 a[j] > i - j => a[j] + j > i （a[j] + j ）<= i