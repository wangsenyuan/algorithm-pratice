You are given two integers, 𝑛
 and 𝑘
. There is a graph on 𝑛
 vertices, numbered from 1
 to 𝑛
, which initially has no edges.

You have to assign each vertex an integer; let 𝑎𝑖
 be the integer on the vertex 𝑖
. All 𝑎𝑖
 should be distinct integers from 1
 to 𝑛
.

After assigning integers, for every pair of vertices (𝑖,𝑗)
, you add an edge between them if |𝑖−𝑗|+|𝑎𝑖−𝑎𝑗|≤𝑘
.

Your goal is to create a graph which can be partitioned into the minimum possible (for the given values of 𝑛
 and 𝑘
) number of cliques. Each vertex of the graph should belong to exactly one clique. Recall that a clique is a set of vertices such that every pair of vertices in it are connected with an edge.

Since BledDest hasn't really brushed his programming skills up, he can't solve the problem "given a graph, partition it into the minimum number of cliques". So we also ask you to print the partition itself.

Input
The first line contains one integer 𝑡
 (1≤𝑡≤1600
) — the number of test cases.

Each test case consists of one line containing two integers 𝑛
 and 𝑘
 (2≤𝑛≤40
; 1≤𝑘≤2𝑛
).


### ideas
1. n <= 40
2. 对于i < j, 假设 a[i] < a[j] , (j - i) + a[j] - a[i] <= k
3. j + a[j] - (i + a[i]) <= k
4. 没啥想法，没有切入点～
5. 假设这里有两个clique（有链接的component), 现在要加入i
6. 其中一个有j代表，另外一个有k代表
7. 且 j < i < k 
8. 在一个分组内, 假设其中最大的为, 假设就有i, j组成
9. 且  i < j
10. (j - i) + abs(a[j] - a[i]) <= k
11. j - i + max(a[j], a[i]) - min(a[j], a[i]) <= k
12. 如果 a[i] < a[j]
13. j - i + a[j] - a[i] <= k
14. 如果 a[i] > a[j]
15. j - i + a[i] - a[j] <= k
16. 哪种组合更小？
17. j - a[j] 或者 j + a[j]
18. 假设有几个分组，是不是每个分组内的下标都是连续的？
19. 如果不是连续的，似乎就可以连起来
20. 1....i 是一个分组 (i + 1....j)是一个分组, ...
21. 然后在一个分组内，它们的分配也是连续的，但是跨组后，就不连续了
22. 还有个简单的观察，在分组内，就是i越大，貌似分配a[i]越小越好
23. 算了
24. 还是看看答案吧～
25. 看了solution，感觉很没意思。从上面的分析还是能推导出结果的

### solution

There are two main steps to solve the problem:

analyzing the maximum size of a clique;
showing a construction that always allows us to get a clique of the maximum possible size.
Firstly, the maximum size of a clique cannot exceed 𝑘
. If there are at least 𝑘+1
 vertices in the same clique, then at least two of them (call them 𝑖
 and 𝑗
) have |𝑖−𝑗|≥𝑘
. And since 𝑎𝑖≠𝑎𝑗
, then |𝑎𝑖−𝑎𝑗|≥1
. So, |𝑖−𝑗|+|𝑎𝑖−𝑎𝑗|
 is at least 𝑘+1
, so these two vertices won't have an edge connecting them (and cannot belong to the same clique).

Secondly, let's try to find a construction that always allows us to get cliques of size 𝑘
. To do this, try to solve the problem when 𝑘=𝑛
; and if 𝑛>𝑘
, we can split all vertices into ⌈𝑛𝑘⌉
 cliques as follows: for each clique, we assign a consecutive block of vertices and numbers that will be assigned to them (for example, vertices from 1
 to 𝑘
 and numbers from 1
 to 𝑘
 belong to the first clique, vertices from 𝑘+1
 to 2𝑘
 and numbers from 𝑘+1
 to 2𝑘
n belong to the second clique), and then use the solution for 𝑛=𝑘
 on each of these blocks.

To obtain a solution for 𝑛=𝑘
, you can either try bruteforcing it locally on, say, 𝑛≤10
 and analyzing the results. One of the possible constructions is as follows: let 𝑚=⌈𝑘2⌉
; split all vertices and numbers from 1
 to 𝑘
 into two blocks: [1,𝑚]
 and [𝑚+1,𝑘]
; and then, in each block, the greater the index of the vertex, the less the integer it gets. So it looks as follows: 𝑎1=𝑚,𝑎2=𝑚−1,…,𝑎𝑚=1,𝑎𝑚+1=𝑛,𝑎𝑚+2=𝑛−1,…,𝑎𝑛=𝑚+1
. We can show that the "distance" between any two vertices in different halves is exactly 𝑘
, and the distance between any two vertices in the same half is at most 2(𝑚−1)
, which never exceeds 𝑘
.