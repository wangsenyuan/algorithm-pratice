The Fair Nut is going to travel to the Tree Country, in which there are 𝑛
cities. Most of the land of this country is covered by forest. Furthermore, the local road system forms a tree (
connected graph without cycles). Nut wants to rent a car in the city 𝑢
and go by a simple path to city 𝑣
. He hasn't determined the path, so it's time to do it. Note that chosen path can consist of only one vertex.

A filling station is located in every city. Because of strange law, Nut can buy only 𝑤𝑖
liters of gasoline in the 𝑖
-th city. We can assume, that he has infinite money. Each road has a length, and as soon as Nut drives through this
road, the amount of gasoline decreases by length. Of course, Nut can't choose a path, which consists of roads, where he
runs out of gasoline. He can buy gasoline in every visited city, even in the first and the last.

He also wants to find the maximum amount of gasoline that he can have at the end of the path. Help him: count it.

Input
The first line contains a single integer 𝑛
(1≤𝑛≤3⋅105
) — the number of cities.

The second line contains 𝑛
integers 𝑤1,𝑤2,…,𝑤𝑛
(0≤𝑤𝑖≤109
) — the maximum amounts of liters of gasoline that Nut can buy in cities.

Each of the next 𝑛−1
lines describes road and contains three integers 𝑢
, 𝑣
, 𝑐
(1≤𝑢,𝑣≤𝑛
, 1≤𝑐≤109
, 𝑢≠𝑣
), where 𝑢
and 𝑣
— cities that are connected by this road and 𝑐
— its length.

It is guaranteed that graph of road connectivity is a tree.

### ideas

1. 假设访问了path(u -> v) 那么答案 = sum of w(u -> v) - sum of edge length(u -> v)
2. dp[u][0] 表示从u出发(或者到达u)的最大收获
3. dp[u][1] 表示第二大的收获
4. res = max(dp[u][0] + dp[u][1] - w[u]) (通过u的两条路径获得最大值)
5. 