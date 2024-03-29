ZS the Coder and Chris the Baboon has arrived at Udayland! They walked in the park where n trees grow. They decided to
be naughty and color the trees in the park. The trees are numbered with integers from 1 to n from left to right.

Initially, tree i has color ci. ZS the Coder and Chris the Baboon recognizes only m different colors, so 0 ≤ ci ≤ m,
where ci = 0 means that tree i is uncolored.

ZS the Coder and Chris the Baboon decides to color only the uncolored trees, i.e. the trees with ci = 0. They can color
each of them them in any of the m colors from 1 to m. Coloring the i-th tree with color j requires exactly pi, j litres
of paint.

The two friends define the beauty of a coloring of the trees as the minimum number of contiguous groups (each group
contains some subsegment of trees) you can split all the n trees into so that each group contains trees of the same
color. For example, if the colors of the trees from left to right are 2, 1, 1, 1, 3, 2, 2, 3, 1, 3, the beauty of the
coloring is 7, since we can partition the trees into 7 contiguous groups of the same color : {2}, {1, 1, 1}, {3}, {2,
2}, {3}, {1}, {3}.

ZS the Coder and Chris the Baboon wants to color all uncolored trees so that the beauty of the coloring is exactly k.
They need your help to determine the minimum amount of paint (in litres) needed to finish the job.

Please note that the friends can't color the trees that are already colored.

### ideas

1. n, m, k <= 100
2. 如果一开始beauty已经超过了k，那么 -1
3. 然后只能涂改c[i] = 0 的部分
4. 对于连续的这样一段 [l...r] dp[l..r][x] 表示将它们涂成x的beauty时的最小cost
5. 还需要知道开始的颜色和结束的颜色（这样可以和前后的进行比较，是否产生了了新的beauty)
6. dp[i][j][x] 表示第i位为颜色j时，且前面共有x个beauty时的最小值
7. dp[i+1][j][x] = dp[i][j][x] + price (和前面一致时(如果可以))