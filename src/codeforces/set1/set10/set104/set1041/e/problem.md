Monocarp has drawn a tree (an undirected connected acyclic graph) and then has given each vertex an index. All indices
are distinct numbers from 1
to 𝑛
. For every edge 𝑒
of this tree, Monocarp has written two numbers: the maximum indices of the vertices of the two components formed if the
edge 𝑒
(and only this edge) is erased from the tree.

Monocarp has given you a list of 𝑛−1
pairs of numbers. He wants you to provide an example of a tree that will produce the said list if this tree exists. If
such tree does not exist, say so.

### ideas

1. 考虑包含1的那条边
2. (1, ?) ？肯定是n， 且它只能连接到1，（当然也可能1在内部）
3. 所以找到包含最小数字的那个pair (x, ?) 另外一个必须是n
4. x是一个叶子节点，然后下一个最小的(y, ?) 这个有可能是个叶子节点， 也有可能不是
5. 它可能是连接(x - y),
6. 可以确定的n应该出现在所有的pair中
7. 那其他的出现了一次的，都作为一个叶子节点（但是有部分会出现多次）
8. 出现多次的，那个可以，必须有对应比u小的个数存在
9. 比如 如果3出现了2次，那么就必须要有1\2没有被使用
10. 