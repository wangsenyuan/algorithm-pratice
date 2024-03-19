Chouti was tired of the tedious homework, so he opened up an old programming problem he created years ago.

You are given a connected undirected graph with 𝑛
vertices and 𝑚
weighted edges. There are 𝑘
special vertices: 𝑥1,𝑥2,…,𝑥𝑘
.

Let's define the cost of the path as the maximum weight of the edges in it. And the distance between two vertexes as the
minimum cost of the paths connecting them.

For each special vertex, find another special vertex which is farthest from it (in terms of the previous paragraph, i.e.
the corresponding distance is maximum possible) and output the distance between them.

The original constraints are really small so he thought the problem was boring. Now, he raises the constraints and hopes
you can solve it for him.

### ideas

1. 如果一个个去跑， k * (m + k) 复杂性，显然是不行的
2. 考虑任意两点间的距离 = min of cost of path between them
3. but cost of path = max weight of the edge in it
4. 有点想法了，把edge按照weight升序排，然后不断的合并
5. 假设当前的edge的weight = x，合并后，把a和b连接到了一起，那么它们之间的distance = x
6. 在这个过程中，合并到一起的特殊节点的最远距离更新为x；也不用更新，就记录这个集合的答案=x即可