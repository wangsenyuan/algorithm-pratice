Bob is playing a game of Spaceship Solitaire. The goal of this game is to build a spaceship. In order to do this, he first needs to accumulate enough resources for the construction. There are 𝑛
 types of resources, numbered 1
 through 𝑛
. Bob needs at least 𝑎𝑖
 pieces of the 𝑖
-th resource to build the spaceship. The number 𝑎𝑖
 is called the goal for resource 𝑖
.

Each resource takes 1
 turn to produce and in each turn only one resource can be produced. However, there are certain milestones that speed up production. Every milestone is a triple (𝑠𝑗,𝑡𝑗,𝑢𝑗)
, meaning that as soon as Bob has 𝑡𝑗
 units of the resource 𝑠𝑗
, he receives one unit of the resource 𝑢𝑗
 for free, without him needing to spend a turn. It is possible that getting this free resource allows Bob to claim reward for another milestone. This way, he can obtain a large number of resources in a single turn.

The game is constructed in such a way that there are never two milestones that have the same 𝑠𝑗
 and 𝑡𝑗
, that is, the award for reaching 𝑡𝑗
 units of resource 𝑠𝑗
 is at most one additional resource.

A bonus is never awarded for 0
 of any resource, neither for reaching the goal 𝑎𝑖
 nor for going past the goal — formally, for every milestone 0<𝑡𝑗<𝑎𝑠𝑗
.

A bonus for reaching certain amount of a resource can be the resource itself, that is, 𝑠𝑗=𝑢𝑗
.

Initially there are no milestones. You are to process 𝑞
 updates, each of which adds, removes or modifies a milestone. After every update, output the minimum number of turns needed to finish the game, that is, to accumulate at least 𝑎𝑖
 of 𝑖
-th resource for each 𝑖∈[1,𝑛]
.

Input
The first line contains a single integer 𝑛
 (1≤𝑛≤2⋅105
) — the number of types of resources.

The second line contains 𝑛
 space separated integers 𝑎1,𝑎2,…,𝑎𝑛
 (1≤𝑎𝑖≤109
), the 𝑖
-th of which is the goal for the 𝑖
-th resource.

The third line contains a single integer 𝑞
 (1≤𝑞≤105
) — the number of updates to the game milestones.

Then 𝑞
 lines follow, the 𝑗
-th of which contains three space separated integers 𝑠𝑗
, 𝑡𝑗
, 𝑢𝑗
 (1≤𝑠𝑗≤𝑛
, 1≤𝑡𝑗<𝑎𝑠𝑗
, 0≤𝑢𝑗≤𝑛
). For each triple, perform the following actions:

First, if there is already a milestone for obtaining 𝑡𝑗
 units of resource 𝑠𝑗
, it is removed.
If 𝑢𝑗=0
, no new milestone is added.
If 𝑢𝑗≠0
, add the following milestone: "For reaching 𝑡𝑗
 units of resource 𝑠𝑗
, gain one free piece of 𝑢𝑗
."
Output the minimum number of turns needed to win the game.


### ideas
1. 同一个resource， s可以有多个组合（s, t)可以有多对
2. 如果是 (s, t, 0) => 删除 (s, t)
3. 在确定的组合时，如何处理呢？
4. 如果没有milestone =》 sum(ai)
5. 另外，能够用到milestone的时候，就应该使用。除非已经满足条件了
6. 完全没有想法
7. (s, t, u) 造成u的个数增加1
8. (u, tx, v) 造成v的个数增加1（那么为了减少turns， 那么最好让u的状态等在tx - 1 的状态
9. 假设这样子组成了一条线（先不考虑环的情况）
10. 那么 ts + tu - 1 + tv - 1 ++ ... => ts + tu + tv ...
11. 换句话说，这条线省掉的turns的数量 = 线的数量？
12. 也就是组成一个越长的线越好
13. 考虑换 (s, 1, s), (s, 2, s), (s, 3, s)
14. 这个线也是可以算入的
15. 至少有点思路了
16. 那么就需要计算有多少个这样的图，节省的turns = 图中线条的数量
17. 但似乎也不完全是这样
18. 