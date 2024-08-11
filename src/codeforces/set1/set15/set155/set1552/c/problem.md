On a circle lie 2𝑛
 distinct points, with the following property: however you choose 3
 chords that connect 3
 disjoint pairs of points, no point strictly inside the circle belongs to all 3
 chords. The points are numbered 1,2,…,2𝑛
 in clockwise order.

Initially, 𝑘
 chords connect 𝑘
 pairs of points, in such a way that all the 2𝑘
 endpoints of these chords are distinct.

You want to draw 𝑛−𝑘
 additional chords that connect the remaining 2(𝑛−𝑘)
 points (each point must be an endpoint of exactly one chord).

In the end, let 𝑥
 be the total number of intersections among all 𝑛
 chords. Compute the maximum value that 𝑥
 can attain if you choose the 𝑛−𝑘
 chords optimally.

Note that the exact position of the 2𝑛
 points is not relevant, as long as the property stated in the first paragraph holds.


### ideas
1. however you choose 3
 chords that connect 3
 disjoint pairs of points, no point strictly inside the circle belongs to all 3
 chords 这个表示，没有三条弦共点
2. 如果不考虑这些已经连接好的chord，其他的，要怎么排列才能获得最大值呢？
3. 假设有 2 * x 条，貌似是 (1, x+1) (2, x + 2), ... (x, 2 * x)
4. 这样子，共产生 x * (x - 1 + 0) / 2 条个交点
5. 然后考虑这些已经链接的弦。 假设i, j,处没有连接，那么它的贡献= 开始节点在它前面，结束节点在它i后面，且在j前面的部分