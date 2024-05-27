You are given a tree, consisting of 𝑛
 vertices, numbered from 1
 to 𝑛
. Every vertex is colored in some color, denoted by an integer from 1
 to 𝑛
.

A simple path of the tree is called beautiful if:

it consists of at least 2
 vertices;
the first and the last vertices of the path have the same color;
no other vertex on the path has the same color as the first vertex.
Count the number of the beautiful simple paths of the tree. Note that paths are considered undirected (i. e. the path from 𝑥
 to 𝑦
 is the same as the path from 𝑦
 to 𝑥
).

### ideas
1. 如果固定一个端点为u，那么它的贡献，就是它四周和它颜色相同的最近的那些点的个数
2. 假设以0位root，计算出了每个颜色的个数
3. 现在移动到0的某个child u， 怎么更新这个计数？
4. 这里分u内部，和u外部的情况，u外部的计数，除了0的颜色，其他的都是不变的
5. 也就是假设cnt[0][x 0的颜色] = 1
6. 但是有个问题是，如何从0中吧u的贡献给取消掉？
7. u的颜色，对于0来说，它的贡献是1，但是对于u来说，它的贡献应该是它内部的相同颜色的个数
8. 大概知道了