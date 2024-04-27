You are given a tree (an undirected connected graph without cycles) and an integer 𝑠
.

Vanya wants to put weights on all edges of the tree so that all weights are non-negative real numbers and their sum is
𝑠
. At the same time, he wants to make the diameter of the tree as small as possible.

Let's define the diameter of a weighed tree as the maximum sum of the weights of the edges lying on the path between two
some vertices of the tree. In other words, the diameter of a weighed tree is the length of the longest simple path in
the tree, where length of a path is equal to the sum of weights over all edges in the path.

Find the minimum possible diameter that Vanya can get.

### ideas

1. deg = 1 的节点定义为叶子节点, 假设有x个，那么找到一个center（root）， 然后让这个center到这些叶子节点的距离 = s / x
2. 那么dia = 2 * s / x (如果存在这样的点)
3. 首先这个肯定能够达到，只需要把叶子节点出来的线给赋值为 s/ x即可
4. 有没有更好的方案呢？
5. 假设一条边设置为 y > s / x, 其他的边都设置 z
6. 且 y + x < 2 * s / x, 那么z应该是多少呢？
7. 假设x可以整除s
8. s = 6, x = 3, 按照第一个方案 2 * 6 / 3 = 4
9. 如果设置一条边为 3, 另外两条边 0.5