As the name of the task implies, you are asked to do some work with segments and trees.

Recall that a tree is a connected undirected graph such that there is exactly one simple path between every pair of its vertices.

You are given 𝑛
 segments [𝑙1,𝑟1],[𝑙2,𝑟2],…,[𝑙𝑛,𝑟𝑛]
, 𝑙𝑖<𝑟𝑖
 for every 𝑖
. It is guaranteed that all segments' endpoints are integers, and all endpoints are unique — there is no pair of segments such that they start in the same point, end in the same point or one starts in the same point the other one ends.

Let's generate a graph with 𝑛
 vertices from these segments. Vertices 𝑣
 and 𝑢
 are connected by an edge if and only if segments [𝑙𝑣,𝑟𝑣]
 and [𝑙𝑢,𝑟𝑢]
 intersect and neither of it lies fully inside the other one.

For example, pairs ([1,3],[2,4])
 and ([5,10],[3,7])
 will induce the edges but pairs ([1,2],[3,4])
 and ([5,7],[3,10])
 will not.

Determine if the resulting graph is a tree or not.


### ideas
1. 类似的题目貌似遇到过
2. 一个点，不能出现在3段中，否则就产生了圈
3. 否则似乎就一定是可以的？
4. 如果就是这样子，是不是太简单了？
5. 这里有个陷阱，就是 如果 u包含v，那么u和v之间是没有边的
6. 那可不可以反过来，统计没有边的节点对有多少？
7. 如果r[i] <= l[j], 那么这些可以算入
8. 如果在内部的也可以算入
9. 好像简单点