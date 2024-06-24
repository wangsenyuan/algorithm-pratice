Building bridges did not help Bernard, and he continued to be late everywhere. Then Rudolf decided to teach him how to use the subway.

Rudolf depicted the subway map as an undirected connected graph, without self-loops, where the vertices represent stations. There is at most one edge between any pair of vertices.

Two vertices are connected by an edge if it is possible to travel directly between the corresponding stations, bypassing other stations. The subway in the city where Rudolf and Bernard live has a color notation. This means that any edge between stations has a specific color. Edges of a specific color together form a subway line. A subway line cannot contain unconnected edges and forms a connected subgraph of the given subway graph.

An example of the subway map is shown in the figure.


Rudolf claims that the route will be optimal if it passes through the minimum number of subway lines.

Help Bernard determine this minimum number for the given departure and destination stations.

### ideas
1. 任意两个点之间的路径，必须考虑把它们连接在一起的颜色
2. 每条边都按照不同的颜色进行分组
3. 然后按照颜色进行处理？这样可以保证每条边只被处理一次。
4. 但是在组成一个新的compoent的时候，如何快速的知道，哪些query被更新了呢？
5. 以及怎么样被更新了？
6. 理解错了，是要计算，两个节点间，需要经过的最少的颜色
7. 而且b/e只有一组
8. dp[u][i] 表示在节点u时，最后一站是通过颜色i进入u时的最少的颜色个数
9. dp[v][j] = min(dp[u][i] + 1, dp[u][j])
10. 所以每个站点其实只需要知道两个数字，进入它最少的次数以及颜色，以及第二少的次数和颜色
11. 然后就可以据此更新外部的周围的节点