You are playing a strange game with Li Chen. You have a tree with 𝑛
nodes drawn on a piece of paper. All nodes are unlabeled and distinguishable. Each of you independently labeled the
vertices from 1
to 𝑛
. Neither of you know the other's labelling of the tree.

You and Li Chen each chose a subtree (i.e., a connected subgraph) in that tree. Your subtree consists of the vertices
labeled 𝑥1,𝑥2,…,𝑥𝑘1
in your labeling, Li Chen's subtree consists of the vertices labeled 𝑦1,𝑦2,…,𝑦𝑘2
in his labeling. The values of 𝑥1,𝑥2,…,𝑥𝑘1
and 𝑦1,𝑦2,…,𝑦𝑘2
are known to both of you.

The picture shows two labelings of a possible tree: yours on the left and Li Chen's on the right. The selected trees are
highlighted. There are two common nodes.
You want to determine whether your subtrees have at least one common vertex. Luckily, your friend Andrew knows both
labelings of the tree. You can ask Andrew at most 5
questions, each of which is in one of the following two forms:

A x: Andrew will look at vertex 𝑥
in your labeling and tell you the number of this vertex in Li Chen's labeling.
B y: Andrew will look at vertex 𝑦
in Li Chen's labeling and tell you the number of this vertex in your labeling.
Determine whether the two subtrees have at least one common vertex after asking some questions. If there is at least one
common vertex, determine one of your labels for any of the common vertices.

### ideas

1. 完全看不懂这个题目
2. 这里有用的信息是
    + 选中的那些节点是一个连通的子图
    + 自己的那颗tree是清楚的
    + 5 = 10 / 2 and 2 ** 10 > 1000
3. 知道对方完整的🌲是不可能的，但是假设r1是自己的root，ask (A, r1) => r2 是对方的root
4. 知道这个信息后，有什么帮助吗？
5. 如果r2不符合条件，那么对于r1的所有的size + 1 <= size(B)的子树，都可以被确定，它们不会包含重叠的点
6. 如果有重叠的点，也必然出现在剩下的那些forest中
    + 包含至少 sz(B) 个节点，
    + 且其中存在选中的节点
7. 如果选择的r1，能够将mark的节点二分，那么每次范围就会缩小一半
8. 考虑一条先 1 - 2 - 3 - 4 - 5 - 6
9. 选中的节点是 [2, 3, 4]
10. 选中的节点，组成一棵树，这个信息要怎么利用？
11. 反过来 ask B会怎样？如果正好问到了一个重叠的点 => good
12. 如果没有， 它的点，比如确定是这边的x，那么就可以确定的一点是包含x的一个子集，就是b选中的点
13. 假设离x最近的点是y，如果x到y的距离 > sz(B) => -1
14. else ask A y 是不是就可以了？ done
15. 甚至，我都不需要查这个距离
