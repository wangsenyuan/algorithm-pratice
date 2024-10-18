You are given a directed graph with 𝑛
 vertices and 𝑚
 directed edges without self-loops or multiple edges.

Let's denote the 𝑘
-coloring of a digraph as following: you color each edge in one of 𝑘
 colors. The 𝑘
-coloring is good if and only if there no cycle formed by edges of same color.

Find a good 𝑘
-coloring of given digraph with minimum possible 𝑘
.

### ideas
1. 找出最大的圈？
2. 最大的ssc？
3. 题目理解错了。
4. 要求没有一个同色的圈，不是要求，同一个圈里面的边的颜色都不一样～
5. 只要有一个圈，答案至少为2
6. 那么2是否是最小的答案呢？
7. 如果两个圈不共享边，那么它们的着色方案是没有关系的
8. 如果有共享，那么k=2是够的，
9. 如果有3个圈共享，类似奥运五环，似乎还是2？
10. 考虑一个完全图，也就是所有的边都存在，那么答案似乎是3？
11. 