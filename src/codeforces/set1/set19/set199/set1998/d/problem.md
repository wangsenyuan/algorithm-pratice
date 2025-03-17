Two of Farmer John's cows, Bessie and Elsie, are planning to race on 𝑛
 islands. There are 𝑛−1
 main bridges, connecting island 𝑖
 to island 𝑖+1
 for all 1≤𝑖≤𝑛−1
. Additionally, there are 𝑚
 alternative bridges. Elsie can use both main and alternative bridges, while Bessie can only use main bridges. All bridges are one way and can only be used to travel from an island with a lower index to an island with a higher index.

Initially, Elsie starts on island 1
, and Bessie starts on island 𝑠
. The cows alternate turns, with Bessie making the first turn. Suppose the cow is on island 𝑖
. During a cow's turn, if there are any bridges connecting island 𝑖
 to island 𝑗
, then the cow can move to island 𝑗
. Then, island 𝑖
 collapses, and all bridges connecting to island 𝑖
 also collapse. Also, note the following:

If there are no bridges connecting island 𝑖
 to another island, then island 𝑖
 collapses, and this cow is eliminated from the race.
If the other cow is also on island 𝑖
, then after this cow moves to another island, the island collapses, and the other cow is eliminated from the race.
After an island or bridge collapses, no cows may use them.
If a cow is eliminated, their turn is skipped for the rest of the race.
The race ends once either cow reaches island 𝑛
. It can be shown that regardless of the cows' strategies, at least one cow reaches island 𝑛
. Bessie wins if and only if she reaches island 𝑛
 first.

For each 1≤𝑠≤𝑛−1
, determine whether Bessie wins if she starts the race on island 𝑠
. Assume both cows follow optimal strategies to ensure their own respective victories.

### ideas
1. 如果Bessie在s处出发，如果Elsie在某个节点跑到它前面了，那么Bessie就输了
2. 假设经过t步后， Elsie到达了节点u
3. s + t < u成立
4. 那么u必须存在一个节点v, v < s, 且 dist[u] <= t 才行
5. => dist[v] <= t - 1
6. 也就是说，从1...s这个区间，找到一个节点v， dist[v] 表示从1到v的最短距离
7. 然后从v存在一个点u，u - s> dist[v] + 1 = dist[u]
8. u > s + dist[v] + 1
9. 对于任何一个点v，应该尽可能的找下一个最大的点u
10. 对于u来说，它应该找最小的v，然后它会有一个范围
11. 当 u - s <= dist[u] => s >= u - dist[u]时，u就失效了
12. 还差一个东西，就是v < s 这个没有体现出来