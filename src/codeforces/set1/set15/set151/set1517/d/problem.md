You are wandering in the explorer space of the 2050 Conference.

The explorer space can be viewed as an undirected weighted grid graph with size 𝑛×𝑚
. The set of vertices is {(𝑖,𝑗)|1≤𝑖≤𝑛,1≤𝑗≤𝑚}
. Two vertices (𝑖1,𝑗1)
 and (𝑖2,𝑗2)
 are connected by an edge if and only if |𝑖1−𝑖2|+|𝑗1−𝑗2|=1
.

At each step, you can walk to any vertex connected by an edge with your current vertex. On each edge, there are some number of exhibits. Since you already know all the exhibits, whenever you go through an edge containing 𝑥
 exhibits, your boredness increases by 𝑥
.

For each starting vertex (𝑖,𝑗)
, please answer the following question: What is the minimum possible boredness if you walk from (𝑖,𝑗)
 and go back to it after exactly 𝑘
 steps?

You can use any edge for multiple times but the boredness on those edges are also counted for multiple times. At each step, you cannot stay on your current vertex. You also cannot change direction while going through an edge. Before going back to your starting vertex (𝑖,𝑗)
 after 𝑘
 steps, you can visit (𝑖,𝑗)
 (or not) freely.

Input
The first line contains three integers 𝑛
, 𝑚
 and 𝑘
 (2≤𝑛,𝑚≤500,1≤𝑘≤20
).

The 𝑗
-th number (1≤𝑗≤𝑚−1
) in the 𝑖
-th line of the following 𝑛
 lines is the number of exibits on the edge between vertex (𝑖,𝑗)
 and vertex (𝑖,𝑗+1)
.

The 𝑗
-th number (1≤𝑗≤𝑚
) in the 𝑖
-th line of the following 𝑛−1
 lines is the number of exibits on the edge between vertex (𝑖,𝑗)
 and vertex (𝑖+1,𝑗)
.

The number of exhibits on each edge is an integer between 1
 and 106
.

Output
Output 𝑛
 lines with 𝑚
 numbers each. The 𝑗
-th number in the 𝑖
-th line, 𝑎𝑛𝑠𝑤𝑒𝑟𝑖𝑗
, should be the minimum possible boredness if you walk from (𝑖,𝑗)
 and go back to it after exactly 𝑘
 steps.

If you cannot go back to vertex (𝑖,𝑗)
 after exactly 𝑘
 steps, 𝑎𝑛𝑠𝑤𝑒𝑟𝑖𝑗
 should be −1
.

### ideas
1. k must be even
2. 而且肯定是到达某个地方，原路返回
3. 假设不是，意味着返回的那条路更便宜，所以就可以选择那条路径出发
4. 所以可以， dp[i][j][k] 表示从(i, j) 出发最远走k步的最小值