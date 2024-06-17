Find the minimum height of a rooted tree†
 with 𝑎+𝑏+𝑐
 vertices that satisfies the following conditions:

𝑎
 vertices have exactly 2
 children,
𝑏
 vertices have exactly 1
 child, and
𝑐
 vertices have exactly 0
 children.
If no such tree exists, you should report it.


### ideas
1. 显然c是叶子节点的个数
2. b是只有一个子child的内部节点
3. 它们只能排成一条线
4. (b + c - 1) / c + 1 是它们的最低高度
5. 先假设能整除，此时可以认为仍然有c的叶子节点
6. 然后就是a个有两个节点的内部节点。这样的节点可以组成一个完全二叉树, 
7. a <= c, 