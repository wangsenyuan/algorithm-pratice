You are given a tree with 𝑛
 vertices.

You need to construct an array 𝑎1,𝑎2,…,𝑎𝑛
 of length 𝑛
, consisting of unique integers from 1
 to 2⋅𝑛
, and such that for each edge 𝑢𝑖↔𝑣𝑖
 of the tree, the value |𝑎𝑢𝑖−𝑎𝑣𝑖|
 is not a prime number.

Find any array that satisfies these conditions, or report that there is no such array.

### ideas
1. 完全没想法～
2. 如果两个节点相差为1，是最好的，但是似乎没法保证这一点
3. 假设当前a[u] = x, 那么就使用x + 1(如果可以用), x + 4, x + 6, x + 8, x + 9, ...
4. 如果都不能用就返回false？
5. 