There are 𝑛
 cities and 𝑛−1
 undirected roads connecting pairs of cities. Citizens of any city can reach any other city traveling by these roads. Cities are numbered from 1
 to 𝑛
 and the city 1
 is a capital. In other words, the country has a tree structure.

There are 𝑚
 citizens living in the country. A 𝑝𝑖
 people live in the 𝑖
-th city but all of them are working in the capital. At evening all citizens return to their home cities using the shortest paths.

Every person has its own mood: somebody leaves his workplace in good mood but somebody are already in bad mood. Moreover any person can ruin his mood on the way to the hometown. If person is in bad mood he won't improve it.

Happiness detectors are installed in each city to monitor the happiness of each person who visits the city. The detector in the 𝑖
-th city calculates a happiness index ℎ𝑖
 as the number of people in good mood minus the number of people in bad mood. Let's say for the simplicity that mood of a person doesn't change inside the city.

Happiness detector is still in development, so there is a probability of a mistake in judging a person's happiness. One late evening, when all citizens successfully returned home, the government asked uncle Bogdan (the best programmer of the country) to check the correctness of the collected happiness indexes.

Uncle Bogdan successfully solved the problem. Can you do the same?

More formally, You need to check: "Is it possible that, after all people return home, for each city 𝑖
 the happiness index will be equal exactly to ℎ𝑖
".

### ideas
1. 有点懵的题目
2. 假设到h[u]都是对的，那么如何判断h[v]是否对的（v是u的子节点）
3. 很显然的是，到达节点v的额外的人是 sz[v] (包括v的整棵树)
4. 如果 h[v] > h[u] + sz[v], 或者 h[v] < h[u] - sz[v] => false
5. 假设在这个范围内，那么 diff = h[v] - h[u]
6. 如果diff 是整数，表示，至少有这么多人是进入h[v]是高兴进入的 （他们同样也进入了u）
7. 还是应该反过来考虑， 对于叶子节点 h[u] >= -a[u] <= a[u] (全部不高兴，或者全部高兴)
8. 那么这个时候，其实是可以知道有哪些人不高兴，哪些人高兴的 
9. x[u] + y[u] = a[u]
10. x[u] - y[u] = -h[u] 
11. 然后就可以往上处理了（这个结果是确定的）
12. 汗，这里有个问题在与，那些bad心情的，是可以在路上产生的，但是一旦bad，就不能改变了