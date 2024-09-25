Demiurges Shambambukli and Mazukta love to watch the games of ordinary people. Today, they noticed two men who play the following game.

There is a rooted tree on n nodes, m of which are leaves (a leaf is a nodes that does not have any children), edges of the tree are directed from parent to children. In the leaves of the tree integers from 1 to m are placed in such a way that each number appears exactly in one leaf.

Initially, the root of the tree contains a piece. Two players move this piece in turns, during a move a player moves the piece from its current nodes to one of its children; if the player can not make a move, the game ends immediately. The result of the game is the number placed in the leaf where a piece has completed its movement. The player who makes the first move tries to maximize the result of the game and the second player, on the contrary, tries to minimize the result. We can assume that both players move optimally well.

Demiurges are omnipotent, so before the game they can arbitrarily rearrange the numbers placed in the leaves. Shambambukli wants to rearrange numbers so that the result of the game when both players play optimally well is as large as possible, and Mazukta wants the result to be as small as possible. What will be the outcome of the game, if the numbers are rearranged by Shambambukli, and what will it be if the numbers are rearranged by Mazukta? Of course, the Demiurges choose the best possible option of arranging numbers.

### ideas
1. f(u) 表示当前player在u节点时,能够得到的最大值，g(u)表示能够得到最小值
2. f(u) = max(g(v)) (当前player只能到达一个子节点，且轮到了对手，所以只能用最小值)
3. g(u) = min(f(v)) (...)
4. 在给定了叶子节点的排列后，可以这样计算。
5. 但是现在是要找出一种排列，以找到最大的数（和最小的数）
6. 一个观察就是，为了找到最大的数（假设root有两颗子树，那么最大的哪一撮应该在同一边）（这样可以保证最小值最大）
7. 假设到了下一次操作（在范围（l...r)内产生最后的结果)，second player，最好的策略也是把这个范围分成连续的段，且进入那个最小的段中
8. 每次似乎都应该选择那个有最小范围的段？似乎和高度也有关系，如果高度是偶数（也就是最后一次由first player选择）那么对于第一个player有利
9. 假设f(u) = 进入这个子树能够获得的最大值（这个和叶子节点的数量有关系，和结构也有关系）
10. 且和其他子树没关系（那么first player就应该选择最大的那个子树，且考虑整体得到一个结果）这个搞定
11. 现在考虑如何安排使得得到的最大值最小？
12. 为了使得最后的结果尽量的小，使用连续的区间，似乎就不大可行了？
13. 因为使用连续的区间，对于第一个player来说，他还是会选择那个让他的结果最大的那个区间
14. 假设有(w个子树)，那么让最小值，分别在每个区间是更好的
15. 应该按照g(u)排序（也就是能够取得的最小值的值，然后倒过来分配就好了）got