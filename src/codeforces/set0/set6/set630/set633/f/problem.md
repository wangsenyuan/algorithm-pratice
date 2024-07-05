Alice and Bob have a tree (undirected acyclic connected graph). There are ai chocolates waiting to be picked up in the i-th vertex of the tree. First, they choose two different vertices as their starting positions (Alice chooses first) and take all the chocolates contained in them.

Then, they alternate their moves, selecting one vertex at a time and collecting all chocolates from this node. To make things more interesting, they decided that one can select a vertex only if he/she selected a vertex adjacent to that one at his/her previous turn and this vertex has not been already chosen by any of them during other move.

If at any moment one of them is not able to select the node that satisfy all the rules, he/she will skip his turns and let the other person pick chocolates as long as he/she can. This goes on until both of them cannot pick chocolates any further.

Due to their greed for chocolates, they want to collect as many chocolates as possible. However, as they are friends they only care about the total number of chocolates they obtain together. What is the maximum total number of chocolates they may pick?

### solution

如果只有一条路径，做法类似求 树的直径 ，把路径分解成左右两条链，dfs 返回从叶子到当前节点的最长链。

方法一：两次 DFS（换根 DP）

提示：枚举边，删除这条边，把这棵树一分为二，每部分只需要计算一条最大路径和。

分类讨论：
如果两条路径分别在两棵不相交子树中，那么做法和求直径是一样的，设当前节点为 v，维护前面遍历的 v 的儿子子树的最大路径和，加上当前儿子子树的最大路径和。
dfs 除了返回最长链，还要返回子树最大路径和。

如果两条路径分别在两棵相交子树中，考虑换根 DP，一条路径在子树 w 中，另一条路径从 v 开始的以下三条链中选两条：
v - w1 子树中的最长链（其中 w1 != w）
v - w2 子树中的最长链（其中 w2 != w）
v - 从父节点过来的最长链（在换根的过程中计算）

代码一

方法二：一次 DFS

如右图所示，下文把：
从叶子往上到 v 的路径称作「链」。
由两条链拼成的路径叫做「路」。
由链 + 路组成的复合结构叫做「链路」。
注意链路中的链和路不一定在一起，可能间隔很远。

分类讨论，遍历 v 的子树列表时，假设当前遍历到子树 w，得到其「w 链」「w 路」「w 链路」，同时维护遍历过的子树的「最大链」「含 v 最大路」「不含 v 最大路」「最大链路」。其中「最大链」「最大链路」已经把 a[v] 算进去了。

如右图，答案是以下三种情况的最大值：
max(含 v 最大路, 不含 v 最大路) + w 路
最大链 + w 链路
最大链路 + w 链

其中「最大链路」是以下三种情况的最大值：
w 链路 + a[v]
不含 v 最大路 + w 链 + a[v]
最大链 + w 路
注意链路这一复合结构，一定是链在上，路在下，不然无法往上加点。
不能是【含 v 最大路 + w 链】。
这也是为什么我们要区分这两种最大路。

含 v 最大路的计算方式同直径，取最大链 + w 链的最大值。
而最大链又是 w 链 + a[v] 的最大值。注意变量更新的顺序。
不含 v 最大路是 w 路的最大值。

代码二