Cirno and Daiyousei are playing a game with a tree∗
 of 𝑛
 nodes, rooted at node 1
. The value of the 𝑖
-th node is 𝑤𝑖
. They take turns to play the game; Cirno goes first.

In each turn, assuming the opponent chooses 𝑗
 in the last turn, the player can choose any remaining node 𝑖
 satisfying 𝑤𝑖>𝑤𝑗
 and delete the subtree†
 of node 𝑖
. In particular, Cirno can choose any node and delete its subtree in the first turn.

The first player who can not operate wins, and they all hope to win. Find one of the possible nodes Cirno may choose so that she wins if both of them play optimally.

### ideas
1. 如果直接选择节点1,肯定是不行的， 这样马上就输掉了（所以一个节点时，必输）
2. 如果考虑1只有一个子树u（且限定只能在子树u中选择），如果可以在里面进行奇数次操作，那么就可以逼迫对手选择节点1
3. 从而获胜
4. 如果有x个子树，如果有两个可以进行奇数次选择，那么就抵消了
5. 在子树u中，如果只有一个节点w[?] < w[1], 那么就是一次
6. 否则也还是输掉了
7. 假设按照w[?]升序排节点
8. 如果正好是一条直线，且从底部到root
9. 那么C肯定可以获胜，它可以把倒数第二个选中，然后留给对手root
10. 如果倒数第二个包含了所有的比它的节点，那么此时C必须选择倒数第三的
11. 如果能找到这样一个节点，那么它就获胜了，否则肯定就失败了