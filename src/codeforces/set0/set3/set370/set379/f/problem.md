https://codeforces.com/problemset/problem/379/F

你有一棵树，一开始有 4 个节点，编号为 1,2,3,4，其中 2,3,4 都和 1 相连。

输入 q(1≤q≤5e5) 表示有 q 次操作。
每次操作，输入 v(1≤v≤当前树的大小)，保证 v 是叶子。
在叶子 v 的下面添加两个新的节点与 v 相连，编号分别为 n+1 和 n+2，其中 n 是当前树的大小。

每次操作后，输出树的直径长度。

输入
5
2
3
4
8
5
输出
3
4
4
5
6

【灵茶の试炼】题目&题解
https://docs.qq.com/sheet/DWGFoRGVZRmxNaXFz


### ideas
1. 先根据输入构造出树，对树进行heave-light分解；维护节点在分段中的位置（初始值）
2. 在添加节点的过程中，对路径上的关键节点进行加1，并计算和（分段上）后面的最大值
3. 但是位置的值只能被添加一次（第一个节点被加入时）