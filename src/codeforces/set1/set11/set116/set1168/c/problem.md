Toad Pimple has an array of integers 𝑎1,𝑎2,…,𝑎𝑛
.

We say that 𝑦
 is reachable from 𝑥
 if 𝑥<𝑦
 and there exists an integer array 𝑝
 such that 𝑥=𝑝1<𝑝2<…<𝑝𝑘=𝑦
, and 𝑎𝑝𝑖&𝑎𝑝𝑖+1>0
 for all integers 𝑖
 such that 1≤𝑖<𝑘
.

Here &
 denotes the bitwise AND operation.

You are given 𝑞
 pairs of indices, check reachability for each of them.

Input
The first line contains two integers 𝑛
 and 𝑞
 (2≤𝑛≤300000
, 1≤𝑞≤300000
) — the number of integers in the array and the number of queries you need to answer.

The second line contains 𝑛
 space-separated integers 𝑎1,𝑎2,…,𝑎𝑛
 (0≤𝑎𝑖≤300000
) — the given array.

The next 𝑞
 lines contain two integers each. The 𝑖
-th of them contains two space-separated integers 𝑥𝑖
 and 𝑦𝑖
 (1≤𝑥𝑖<𝑦𝑖≤𝑛
). You need to check if 𝑦𝑖
 is reachable from 𝑥𝑖
.

### ideas
1. x = p[1] < p[2] < .... < p[k] = y
2. a[p[i]] & a[p[i+1]] > 0
3. 在n个数中, 1, 2, 3.... n. 找出k个数, p[1], p[2], ... p[k]
4. 满足 a[p[i]] & a[p[i+1]] > 0
5. & 有个特点，就是参与的数越多，越容易变成0
6. 对于，任何一个 a[i], 记录p[i][j] (j是a[i]的第j为1的位)是多少
7. 那么理论上，沿着 p[i][j]就可以从y找到x
8.  但是这样子，肯定是太慢了。需要更快的处理
9.  假设从x能够到的最远的距离 = w （这个似乎没有用， 比如w-1的位置，等于0，其实是x无法到达的）
10. a[x] = 1, a[y] = 2, 
11. 如果中间有个位置 a[z] = 3, 那么yes
12. 能找到一棵树，把，x和y变成其中的某个节点吗？
13. 比如a[i] 的父节点 j = a[i] & a[j] > 0 
14. 假设0.... 20 是不同的楼层，a[x], a[y]本身在某些楼层上面（它们有的那些位）
15. 如果a[x]可以到达最近的某个楼层d的位置l, 小于 r（a[y]能够到达的最近的楼层d的位置)
16. 那么就ok
