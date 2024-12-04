You have a connected undirected graph made of 𝑛
 nodes and 𝑚
 edges. The 𝑖
-th node has a value 𝑣𝑖
 and a target value 𝑡𝑖
.

In an operation, you can choose an edge (𝑖,𝑗)
 and add 𝑘
 to both 𝑣𝑖
 and 𝑣𝑗
, where 𝑘
 can be any integer. In particular, 𝑘
 can be negative.

Your task to determine if it is possible that by doing some finite number of operations (possibly zero), you can achieve for every node 𝑖
, 𝑣𝑖=𝑡𝑖
.



### ideas
1. 在这个过程中不变的是什么？
2. 假设只有两个节点，由一条线连接
3. 那么显然有 t1 - v1 = t2 - v2 (否则肯定不行)
4. 假设在t1 - v1, 那么必然有一条线（或者多条线的变化之和 = t1 - v1)
5. 如果一个节点只有一条线的时候，那么只有一个选择，就是ti - vi
6. 那么这样子，相当于将父节点直接变化 ti - vi
7. 最后剩下的节点，都在圈上
8. 考虑最简单的圈，(t1, v1) (t2, v2) (t3, v3)
9. (1， -〉 4)，(2 -> 5) ( 0, 0)
10. 考虑一个圈 a -> b -> c -> ... -> a
11. a+1, b+1/-1, c-1/+1, d+1/-1, e-1/+1
12. 任何两个中间有偶数个节点的节点对，都可以同时变化
13. (a, b, c)个一个三角形
14. 是不是保留一个生成树，然后从叶子节点开始处理看最后是否成立就可以了？
15. 