— This is not playing but duty as allies of justice, Nii-chan!

— Not allies but justice itself, Onii-chan!

With hands joined, go everywhere at a speed faster than our thoughts! This time, the Fire Sisters — Karen and Tsukihi —
is heading for somewhere they've never reached — water-surrounded islands!

There are three clusters of islands, conveniently coloured red, blue and purple. The clusters consist of a, b and c
distinct islands respectively.

Bridges have been built between some (possibly all or none) of the islands. A bridge bidirectionally connects two
different islands and has length 1. For any two islands of the same colour, either they shouldn't be reached from each
other through bridges, or the shortest distance between them is at least 3, apparently in order to prevent oddities from
spreading quickly inside a cluster.

The Fire Sisters are ready for the unknown, but they'd also like to test your courage. And you're here to figure out the
number of different ways to build all bridges under the constraints, and give the answer modulo 998 244 353. Two ways
are considered different if a pair of islands exist, such that there's a bridge between them in one of them, but not in
the other.

### observation

1. 任意不同颜色的两个节点间，也不能随便添加
2. 假设红色有a, b, 蓝色有 c, d
3. 如果 (a, c) 连接了，那么 (a, d)就不能连接
4. 所以，红蓝之间的方案数 = 第一个红色节点有 cnt(蓝色) + 1中，第二个红色节点有(如果以一个选择了，那么它有 cnt - 1 + 1个)
   如果没有连接那么它有 cnt + 1
5. dp[i][j] = 第i个红色节点，且当前有j个蓝色节点被连接
6. dp[i+1][j] = dp[i][j] + dp[i][j-1] * (cnt - j)
7. 不同颜色直接的方案数就出来了，同一个颜色直接，不能有连接