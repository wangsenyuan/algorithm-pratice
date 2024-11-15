Given a rooted tree with the root at vertex 1
. For any vertex 𝑖
 (1<𝑖≤𝑛
) in the tree, there is an edge connecting vertices 𝑖
 and 𝑝𝑖
 (1≤𝑝𝑖<𝑖
), with a weight equal to 𝑡𝑖
.

Iris does not know the values of 𝑡𝑖
, but she knows that ∑𝑖=2𝑛𝑡𝑖=𝑤
 and each of the 𝑡𝑖
 is a non-negative integer.

The vertices of the tree are numbered in a special way: the numbers of the vertices in each subtree are consecutive integers. In other words, the vertices of the tree are numbered in the order of a depth-first search.

The tree in this picture satisfies the condition. For example, in the subtree of vertex 2
, the vertex numbers are 2,3,4,5
, which are consecutive integers.
The tree in this picture does not satisfy the condition, as in the subtree of vertex 2
, the vertex numbers 2
 and 4
 are not consecutive integers.
We define dist(𝑢,𝑣)
 as the length of the simple path between vertices 𝑢
 and 𝑣
 in the tree.

Next, there will be 𝑛−1
 events:

Iris is given integers 𝑥
 and 𝑦
, indicating that 𝑡𝑥=𝑦
.
After each event, Iris wants to know the maximum possible value of dist(𝑖,𝑖mod𝑛+1)
 independently for each 𝑖
 (1≤𝑖≤𝑛
). She only needs to know the sum of these 𝑛
 values. Please help Iris quickly get the answers.

Note that when calculating the maximum possible values of dist(𝑖,𝑖mod𝑛+1)
 and dist(𝑗,𝑗mod𝑛+1)
 for 𝑖≠𝑗
, the unknown edge weights may be different.

### ideas
1. 当n = 2 时，2 * t[x]
2. 在已经确定了weight的边后，剩余的量,应该尽可能的用到一条边上，
3. 比如 确定了t[2],t[3]后，那么剩余的w应该在计算4时，放在4上，在计算5的时候，放在5上。。类似这样
4. 理论上，这样会产生额外的 m * w 个结果
5. 但是问题出现在那些切换的点上，比如在题目中，的点5和点6上面
6. 当t3被确定后，dist[4,5]也需要加入t[3]的结果
7. t[3]只会影响到dist[4,5]
8. 在确定了t[2] = y的情况下，剩余w - t[2] 
9. dist[3, 4] = w - t[2]
10. dist[4, 5] = 2 * (w - t[2]) 如果放置在t[3]上面
11. 那么规律是，所有新的head，比如5，6，在它们未定的时候，*2
12. 在确定的时候，需要把路径上的都需要*2