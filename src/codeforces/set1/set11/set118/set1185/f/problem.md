A company of 𝑛
 friends wants to order exactly two pizzas. It is known that in total there are 9
 pizza ingredients in nature, which are denoted by integers from 1
 to 9
.

Each of the 𝑛
 friends has one or more favorite ingredients: the 𝑖
-th of friends has the number of favorite ingredients equal to 𝑓𝑖
 (1≤𝑓𝑖≤9
) and your favorite ingredients form the sequence 𝑏𝑖1,𝑏𝑖2,…,𝑏𝑖𝑓𝑖
 (1≤𝑏𝑖𝑡≤9
).

The website of CodePizza restaurant has exactly 𝑚
 (𝑚≥2
) pizzas. Each pizza is characterized by a set of 𝑟𝑗
 ingredients 𝑎𝑗1,𝑎𝑗2,…,𝑎𝑗𝑟𝑗
 (1≤𝑟𝑗≤9
, 1≤𝑎𝑗𝑡≤9
) , which are included in it, and its price is 𝑐𝑗
.

Help your friends choose exactly two pizzas in such a way as to please the maximum number of people in the company. It is known that a person is pleased with the choice if each of his/her favorite ingredients is in at least one ordered pizza. If there are several ways to choose two pizzas so as to please the maximum number of friends, then choose the one that minimizes the total price of two pizzas.

### ideas
1. 理解一下题目，购买两个pizza，对于每个人来说，如果他喜欢的调料，在其中一个pizza中都存在，那么他就是高兴的
2. 第一个目标是最大化高兴的人的数量
3. 在购买一个pizza的情况下，能够高兴的人的数量是很好算的
4. 如果是两个（因为只有9种调料，所以最多有 1 << 9 = 512种组合)
5. 最多只需要考虑9个pizza. 
6. 考虑这是一颗树，如果某个内部节点存在，那么就可以直接使用这个节点，当这个内部节点不存在时，才需要进入子节点
7. 最多只有9个节点（单个位为true的情况）
8. 最大值很好算
9. 然后就是算最小cost
10. dp[state] = 这个状态的最小cost
11. 如果state1 & state2 = 0, and state1 | state2 能够得到最大的满意人数，那么 dp[state1] + dp[state2] 可以算入
12. state1 & state2 = 0 这个条件是必须的吗？似乎也不是
13. fp[state] 表示这个状态的最大化的人员数量
14. fp[state1] + fp[state2] - fp[satte1 & state2] = 最大化的人？
15. 似乎是的
16. fp[state] 怎么算
17. fp[state] = max(fp[sub]) + 当前state下的人数
18. 这个题目对中文不友好。其中的 **each***表示的是，单个ingredient在两个里面有出现即可
19. fp[state] 表示得到这个state能够使用户满意的最大人数
20. dp[state] = 最多两个pizza时能够得到这个状态的最小cost
21. dp[state] = dp[s1] + dp[s2] 如果 s1 | s2 = state