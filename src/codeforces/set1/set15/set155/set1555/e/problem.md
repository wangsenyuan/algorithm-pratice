You are given 𝑛
 segments on a number line, numbered from 1
 to 𝑛
. The 𝑖
-th segments covers all integer points from 𝑙𝑖
 to 𝑟𝑖
 and has a value 𝑤𝑖
.

You are asked to select a subset of these segments (possibly, all of them). Once the subset is selected, it's possible to travel between two integer points if there exists a selected segment that covers both of them. A subset is good if it's possible to reach point 𝑚
 starting from point 1
 in arbitrary number of moves.

The cost of the subset is the difference between the maximum and the minimum values of segments in it. Find the minimum cost of a good subset.

In every test there exists at least one good subset.

### ideas
1. cost = max - min
2. 在确定min的情况下，找到最小的max（或者在确定max的时候，找到最大的min）
3. 然后剩下就是在给定的区间（min， max中间的分段）如何判断它们是连通的，以及在删除、增加时如何快速的改变连通状态
4. range update + range query
5. 