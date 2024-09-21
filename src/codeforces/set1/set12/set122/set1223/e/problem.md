You are given a weighted tree consisting of 𝑛
 vertices. Recall that a tree is a connected graph without cycles. Vertices 𝑢𝑖
 and 𝑣𝑖
 are connected by an edge with weight 𝑤𝑖
.

Let's define the 𝑘
-coloring of the tree as an assignment of exactly 𝑘
 colors to each vertex, so that each color is used no more than two times. You can assume that you have infinitely many colors available. We say that an edge is saturated in the given 𝑘
-coloring if its endpoints share at least one color (i.e. there exists a color that is assigned to both endpoints).

Let's also define the value of a 𝑘
-coloring as the sum of weights of saturated edges.

Please calculate the maximum possible value of a 𝑘
-coloring of the given tree.

You have to answer 𝑞
 independent queries.

 ### ideas.
 1. 每个节点，都要配置k个颜色, 比如1...k， 2...k+1, 类似这样
 2. 但是，同一个颜色，只能最多出现一次（那么为了产生最多的符合条件的边，可以假定都出现了2次）
 3. 然后，那条，两个端点，有（至少1个）公共颜色的，就可以被计入
 4. 那么显然，越大的weight的，越应该被计入
 5. 所以，按照weight降序，如果这条边的两个端点，还有容量，就给它们分配一个共同的颜色
 6. 如果没有，就继续。最后把剩余的分配上去就可以了。
 7. 如果这样子，是不是太简单了？
 8. 和点的度数有关系吗？好像也没啥关系，还是要挑出这个点，最大的那些边，先处理