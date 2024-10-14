You play a strategic video game (yeah, we ran out of good problem legends). In this game you control a large army, and your goal is to conquer 𝑛
 castles of your opponent.

Let's describe the game process in detail. Initially you control an army of 𝑘
 warriors. Your enemy controls 𝑛
 castles; to conquer the 𝑖
-th castle, you need at least 𝑎𝑖
 warriors (you are so good at this game that you don't lose any warriors while taking over a castle, so your army stays the same after the fight). After you take control over a castle, you recruit new warriors into your army — formally, after you capture the 𝑖
-th castle, 𝑏𝑖
 warriors join your army. Furthermore, after capturing a castle (or later) you can defend it: if you leave at least one warrior in a castle, this castle is considered defended. Each castle has an importance parameter 𝑐𝑖
, and your total score is the sum of importance values over all defended castles. There are two ways to defend a castle:

if you are currently in the castle 𝑖
, you may leave one warrior to defend castle 𝑖
;
there are 𝑚
 one-way portals connecting the castles. Each portal is characterised by two numbers of castles 𝑢
 and 𝑣
 (for each portal holds 𝑢>𝑣
). A portal can be used as follows: if you are currently in the castle 𝑢
, you may send one warrior to defend castle 𝑣
.
Obviously, when you order your warrior to defend some castle, he leaves your army.

You capture the castles in fixed order: you have to capture the first one, then the second one, and so on. After you capture the castle 𝑖
 (but only before capturing castle 𝑖+1
) you may recruit new warriors from castle 𝑖
, leave a warrior to defend castle 𝑖
, and use any number of portals leading from castle 𝑖
 to other castles having smaller numbers. As soon as you capture the next castle, these actions for castle 𝑖
 won't be available to you.

If, during some moment in the game, you don't have enough warriors to capture the next castle, you lose. Your goal is to maximize the sum of importance values over all defended castles (note that you may hire new warriors in the last castle, defend it and use portals leading from it even after you capture it — your score will be calculated afterwards).

Can you determine an optimal strategy of capturing and defending the castles?

### ideas
1. 守卫堡垒有两种方式，一种是马上守卫（留下一个人），另外一种是延后守卫
2. 如果某个堡垒可以被延后守卫，似乎延后是个更好的选择（有更多的人手去占领下一个）
3. k + sum(b[i]) <= 5000
4. dp[i][x]表示到达i处时，剩余x个士兵时的最大收益
5. 不大对，应该是尽量的留下一个士兵去守卫，当人手不够的时候，抽调最低价值的士兵，这个士兵是可以算在最后的队伍里面
6. dp[i][x] 表示到达i时，守护最优价值的x个目标时，所能剩下的最多的士兵的数量？
7. dp[i][x] 如果
8. 如果 dp[i-1][j] >= a[i], dp[i][j] = dp[i-1][j] - a[i] + b[i] (单纯的占领但是不守卫)
9. 然后，再选择 dp[i][j] 个，能够到达的位置，去守卫
10. 还是不对。因为如果只是守卫自己，是很简单的状态 dp[i][j] - 1(如果当前是在第j大)
11. 但是如何去守卫之前的那些castle呢？除非这些都是额外的？
12. 理解错了两个点，一个是，必须要征服所有的城堡，否则就lose了 => -1
13. 还有一个就是，每个城堡至少要a[i]个人，才能征服，但是却不会失去这些人
14. 所以，人数只会增加（除去守卫castle)
15. dp[i][x]表示到达i时，能够守卫的castle的最多数量以及价值
16. dp[i][x] = dp[i-1][x-1] + 1, 如果可以从i这里去守卫第x-1个城堡
17. 主要是这些portals怎么处理？如果直接看i-1的状态，这些portals有可能被守卫了
18. 如果一个castle可以被延迟，应该一致延迟到最后
19. 也就是说，每个v，只有一个u需要考虑（它自己，或者最后那个可以影响到它的）
20. 然后把这些价值，加到u上去就好了
21. 但是这里，还是 n * n * n 的复杂性～
22. 对于i，最多安排 sum - a[i+1]个人去defend
23. 然后，反悔掉？