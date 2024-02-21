Doremy lives in a country consisting of 𝑛
cities numbered from 1
to 𝑛
, with 𝑎𝑖
people living in the 𝑖
-th city. It can be modeled as an undirected graph with 𝑛
nodes.

Initially, there are no edges in the graph. Now Doremy wants to make the graph connected.

To do this, she can add an edge between 𝑖
and 𝑗
if

∑𝑘∈𝑆𝑎𝑘≥𝑖⋅𝑗⋅𝑐,
where 𝑆
is the set of all the nodes that are currently in the same connected component of either 𝑖
or 𝑗
, and 𝑐
is a given constant.

Can Doremy make the graph connected?

Two nodes (𝑖,𝑗)
are in the same connected component if there exists a path from 𝑖
to 𝑗
. A graph is connected if all its nodes are in the same connected component.

### thoughts

1. 假设(i, j)还没有联通，现在要联通它们，它们连个联通块的sum(a) >= i * j * c
2. sum(ai) + sum(aj) >= i * j * c
3. 我们考虑第一个时刻，必须存在一对(i, j)它们必须要能够连接起来
4. 假设存在一个这样的集合，包含s1 < s2 < s3...等元素，现在要加入j，那么j最优的选择是选择其中最小的s1
5. 如果是两个集合，也是选择其中最小的两个下标进行合并
6. 但是这个左边是加法，右边是乘法，要怎么变化一下
7. (i, ai) (j, aj), (k, ak)
8. i < j < k
9. 如果最后能够链接起来， 且考虑操作的顺序是(k + j) + i
10. aj + ak >= j * k * c
11. aj + ak + ai >= j * i * c
12. aj + ak >= j * k * c >= j * i * c
13. 如果 ai >= ak, 那么就可以先链接(j, i)
14. 这个观察是说，对于最小的i(也就是1来说)， 要找到右边第一个j，这个j应该是a[j]越大越好，但是j越小越好
15. 在i = 1的情况下
16. a[1] + a[j] >= j * c
17. 假设不存在这样子的，是否就一定没有答案呢？
18. a[j] + a[1] < j * c => a[j] < j * c - a[1]
19. a[2] + a[3] < 5 * c - 2 * a[1] < 6 * c
20. a[i] + a[j] < i * c + j * c - 2 * a[1] < (i + j) * c < i * j * c