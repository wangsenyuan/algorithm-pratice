The map of Bertown can be represented as a set of 𝑛
 intersections, numbered from 1
 to 𝑛
 and connected by 𝑚
 one-way roads. It is possible to move along the roads from any intersection to any other intersection. The length of some path from one intersection to another is the number of roads that one has to traverse along the path. The shortest path from one intersection 𝑣
 to another intersection 𝑢
 is the path that starts in 𝑣
, ends in 𝑢
 and has the minimum length among all such paths.

Polycarp lives near the intersection 𝑠
 and works in a building near the intersection 𝑡
. Every day he gets from 𝑠
 to 𝑡
 by car. Today he has chosen the following path to his workplace: 𝑝1
, 𝑝2
, ..., 𝑝𝑘
, where 𝑝1=𝑠
, 𝑝𝑘=𝑡
, and all other elements of this sequence are the intermediate intersections, listed in the order Polycarp arrived at them. Polycarp never arrived at the same intersection twice, so all elements of this sequence are pairwise distinct. Note that you know Polycarp's path beforehand (it is fixed), and it is not necessarily one of the shortest paths from 𝑠
 to 𝑡
.

Polycarp's car has a complex navigation system installed in it. Let's describe how it works. When Polycarp starts his journey at the intersection 𝑠
, the system chooses some shortest path from 𝑠
 to 𝑡
 and shows it to Polycarp. Let's denote the next intersection in the chosen path as 𝑣
. If Polycarp chooses to drive along the road from 𝑠
 to 𝑣
, then the navigator shows him the same shortest path (obviously, starting from 𝑣
 as soon as he arrives at this intersection). However, if Polycarp chooses to drive to another intersection 𝑤
 instead, the navigator rebuilds the path: as soon as Polycarp arrives at 𝑤
, the navigation system chooses some shortest path from 𝑤
 to 𝑡
 and shows it to Polycarp. The same process continues until Polycarp arrives at 𝑡
: if Polycarp moves along the road recommended by the system, it maintains the shortest path it has already built; but if Polycarp chooses some other path, the system rebuilds the path by the same rules.

Here is an example. Suppose the map of Bertown looks as follows, and Polycarp drives along the path [1,2,3,4]
 (𝑠=1
, 𝑡=4
):

Check the picture by the link http://tk.codeforces.com/a.png

When Polycarp starts at 1
, the system chooses some shortest path from 1
 to 4
. There is only one such path, it is [1,5,4]
;
Polycarp chooses to drive to 2
, which is not along the path chosen by the system. When Polycarp arrives at 2
, the navigator rebuilds the path by choosing some shortest path from 2
 to 4
, for example, [2,6,4]
 (note that it could choose [2,3,4]
);
Polycarp chooses to drive to 3
, which is not along the path chosen by the system. When Polycarp arrives at 3
, the navigator rebuilds the path by choosing the only shortest path from 3
 to 4
, which is [3,4]
;
Polycarp arrives at 4
 along the road chosen by the navigator, so the system does not have to rebuild anything.
Overall, we get 2
 rebuilds in this scenario. Note that if the system chose [2,3,4]
 instead of [2,6,4]
 during the second step, there would be only 1
 rebuild (since Polycarp goes along the path, so the system maintains the path [3,4]
 during the third step).

The example shows us that the number of rebuilds can differ even if the map of Bertown and the path chosen by Polycarp stays the same. Given this information (the map and Polycarp's path), can you determine the minimum and the maximum number of rebuilds that could have happened during the journey?


### ideas
1. 对于任何一个节点，计算出它到t的最短距离
2. dp[u]表示最少rebuild的次数,fp[u]表示最大rebuild次数
3. 从u移动到v时，（玩家只能从u移动到v）
4. 如果v是u的路径上的唯一一个下一个路口，dp[u] = dp[v], fp[u] = fp[v]
5. 如果不是，那么就存在两种情况, （至少存在多个节点）
6. v是u最短路径上的下一个节点，但是存在一个w，不是最短路径上的下一个节点
7. v是u最短路径上的下一个节点，但是还存在一个w，也是最短路径上的下一个节点
8. v不是最短路径上的节点，存在一个w是最短路径的节点
9. 考虑dp[u]次数，在v不是最短路径节点时， dp[u] = dp[v] + 1 (比如会rebuild一次)
10.   如果v是最短路径上的节点时, dp[u] = dp[v] (不管是否存在其他的最短路径节点)
11. 考虑fp[u], 如果v不是最短路径的节点时, fp[u] = max(fp[w]) + 1 (任何一个最短路径节点的最大值) + 1
12. 如果v是最短路径的节点时, fp[u] = max(fp[v], max(fp[w]) + 1) w是其他的最短路径节点