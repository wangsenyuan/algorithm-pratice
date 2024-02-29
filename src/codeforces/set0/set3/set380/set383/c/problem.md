Iahub likes trees very much. Recently he discovered an interesting tree named propagating tree. The tree consists of n
nodes numbered from 1 to n, each node i having an initial value ai. The root of the tree is node 1.

This tree has a special property: when a value val is added to a value of node i, the value -val is added to values of
all the children of node i. Note that when you add value -val to a child of node i, you also add -(-val) to all children
of the child of node i and so on. Look an example explanation to understand better how it works.

This tree supports two types of queries:

"1 x val" — val is added to the value of node x;
"2 x" — print the current value of node x.
In order to help Iahub understand the tree better, you must answer m queries of the preceding type.

### thoughts

1. 用eular序，记录节点的in[u], out[u]
2. 构造两颗fenwick，代表奇数和偶数层的节点
3. 然后对于查询u = same-parity-pre-sum - other-parity-pre-sum