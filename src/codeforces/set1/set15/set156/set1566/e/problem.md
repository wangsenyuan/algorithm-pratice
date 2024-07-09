A tree is a connected graph without cycles. A rooted tree has a special vertex called the root. The parent of a vertex 𝑣
 (different from root) is the previous to 𝑣
 vertex on the shortest path from the root to the vertex 𝑣
. Children of the vertex 𝑣
 are all vertices for which 𝑣
 is the parent.

A vertex is a leaf if it has no children. We call a vertex a bud, if the following three conditions are satisfied:

it is not a root,
it has at least one child, and
all its children are leaves.
You are given a rooted tree with 𝑛
 vertices. The vertex 1
 is the root. In one operation you can choose any bud with all its children (they are leaves) and re-hang them to any other vertex of the tree. By doing that you delete the edge connecting the bud and its parent and add an edge between the bud and the chosen vertex of the tree. The chosen vertex cannot be the bud itself or any of its children. All children of the bud stay connected to the bud.

What is the minimum number of leaves it is possible to get if you can make any number of the above-mentioned operations (possibly zero)?


### ideas
1. 移动一个bud的时候，这个bud u的children的数量不会减少
2. 如果把它移动到某个leaf上，那可以减少一个leaf
3. 那岂不是都是可以减少到2了？
4. 不对，因为能够选择的只有bud这样的节点，也就是倒数第二层的节点
5. 这样的节点，只能移动一次（移动多次似乎没有意义）
6. 假设将u移动到v（v是一个leaf）那么p(v)就不能移动了
7. 那么u可以接到某个非leaf节点吗？似乎没有意义，因为这样子，叶子节点的数目没有减少，且对其他节点（包括它新的parent）
8. 但也不完全是，因为这时，原来u的爷爷节点，就有可能可以移动了
9. 貌似是的parent节点
10. 有点混乱呐。移动的过程中，不变的是什么？变化的又是什么呢？
11. 而且，似乎在一定程度上，选定一颗子树，是可以把这颗子树完全移走的
12. 是不是就模拟这个过程，就可以了？
13. 每次选度数最低的节点，然后找到某个固定的叶子节点连接上去
14. 因为这个固定的叶子节点，它没有用处
15. 但是如果没有这样的节点呢？似乎相互对接就没啥区别了，就把头部的两个连接起来就可以。但这时可能会出现新的固定节点，或者bud