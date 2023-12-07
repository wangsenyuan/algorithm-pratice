The famous magician Borya Budini traveled through the country 𝑋
, which consists of 𝑛
cities. However, an accident happened, and he was robbed in the city number 1
. Now Budini will have a hard way home to the city number 𝑛
.
He's going to get there by plane. In total, there are 𝑚
flights in the country, 𝑖
-th flies from city 𝑎𝑖
to city 𝑏𝑖
and costs 𝑠𝑖
coins. Note that the 𝑖
-th flight is one-way, so it can't be used to get from city 𝑏𝑖
to city 𝑎𝑖
. To use it, Borya must be in the city 𝑎𝑖
and have at least 𝑠𝑖
coins (which he will spend on the flight).

After the robbery, he has only 𝑝
coins left, but he does not despair! Being in the city 𝑖
, he can organize performances every day, each performance will bring him 𝑤𝑖
coins.

Help the magician find out if he will be able to get home, and what is the minimum number of performances he will have
to organize.

The first line contains three integers 𝑛
, 𝑚
and 𝑝
(2≤𝑛≤800
, 1≤𝑚≤3000
, 0≤𝑝≤109
) — the number of cities, the number of flights and the initial amount of coins.

### thoughts

1. 如果不连通的话，从1不能到达n，那么就是-1
2. 到达cityi的时候，如果有足够的钱，他可以继续前进
3. 如果他没有足够的钱，会有好几种选择，在当前城市表演，或者在他经过的某个城市进行表演
4. 一个贪心的想法是，这个城市里面的收入越高，他越应该在哪里表演
5. 当然前提是他要能够从1通过这个city到达i
6. 如果 f[i] 表示这个city的下标
7. dp[i] 表示到达这个cityi时，花费的最小表演次数
8. 这样还不够，必须知道这个i经过了那个最大的city
9. dp[i][j] 表示从1到达j，然后到达i，且j是这条路径上最大收益的那个city
10. 如果从city u要到达下一个city v
11. 迭代最大的city i， 如果 dp[u][i] < inf (存在)
12. dp[v][i] = dp[u][i] + (s + w[i] - 1) / w[i] 如果 w[i] >= w[v]
13. dp[v][v] = ....
14. 但这里还确少一个信息是到达cityu的时候，他剩余的钱
15. 上面的第4条不是很对，因为他可能为了到达那个最高收入的城市，花费也是最多的
16. 如果只有一条路径，上面的想法是没问题的。但现在有很多种选择时，就有问题了
17. 假设在city u，它就是目前路径上的最大值，我们考虑它最远可以到达哪里？
18. 如果到达了city v，但是w[v] < w[u], 那么更优的选择就是在city u这里挣钱，直到到达下一个x
19. w[x] >= w[u] 才是一个更优的方案，那这个是最优的方案吗？
20. dp[u][x] 表示到达city u时，在路径的某些位置，进行了共x次表演后，最多剩余的钱
21. 从u到v
22. 如果 dp[u][x] >= cost[u-v] => dp[v][x] = dp[u][x] - cost
23. dp[v][x+j] = dp[u][x] + w[u] * j - cost ? dp[u][x] + w[u] * j >= cost
24. 似乎是可行的 但是x的范围是什么？
25. w[i] * x >= cost
26. x >= cost / w[i]
27. 但是这个潜在的非常大
28. 先试试吧

editorial

Note that the show can be done "postponed". As soon as we don't have enough money to walk along the edge, we can do
several shows in advance among the peaks that we have already passed, so as to earn the maximum amount of money.

For the general case, you can write 𝑑𝑝[𝑣][𝑏𝑒𝑠𝑡]=(min show,max money)
, where 𝑣
is the number of the vertex where we are, and 𝑏𝑒𝑠𝑡
is the vertex with max. 𝑤𝑖
, which we have already passed through. It can be shown that it is optimal to minimize the number of shows first, and
then maximize the amount of money. This dynamics can be recalculated using Dijkstra's algorithm. Asymptotics of 𝑂(
𝑚𝑛log𝑛)