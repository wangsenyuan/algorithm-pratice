Alice, the president of club FCB, wants to build a team for the new volleyball tournament. The team should consist of 𝑝
 players playing in 𝑝
 different positions. She also recognizes the importance of audience support, so she wants to select 𝑘
 people as part of the audience.

There are 𝑛
 people in Byteland. Alice needs to select exactly 𝑝
 players, one for each position, and exactly 𝑘
 members of the audience from this pool of 𝑛
 people. Her ultimate goal is to maximize the total strength of the club.

The 𝑖
-th of the 𝑛
 persons has an integer 𝑎𝑖
 associated with him — the strength he adds to the club if he is selected as a member of the audience.

For each person 𝑖
 and for each position 𝑗
, Alice knows 𝑠𝑖,𝑗
  — the strength added by the 𝑖
-th person to the club if he is selected to play in the 𝑗
-th position.

Each person can be selected at most once as a player or a member of the audience. You have to choose exactly one player for each position.

Since Alice is busy, she needs you to help her find the maximum possible strength of the club that can be achieved by an optimal choice of players and the audience.

### ideas
1. dp[i][j][state] = 到i为止选择了j个观众和state表示的position的人
2. 但是这个是 n * n * 128 的复杂性
3. 不大行
4. 假设选择了k+p个人，那么在这k+p个人里面，可以按照 dp[state]来处理
5. 这k+p个人，是不是肯定是a[i]最大的k+p个人呢？
6. 不对，但是如果选择了p个人，那么剩下的k个人，肯定是最大的a[i]的k个人
7. 按照a[i]升序排列
8. dp[i][state] 如果 n - i >= k， 那么k个观众肯定从后面选出来
9. 如果 n - i < k， 那么后面的肯定要全部选择作为观众
10. 且再选择 k - (n - i)个人作为观众（这些人就是没有被选为player的j个靠右的人）
11. dp[i][state] = dp[i-1][state] (当前的人不作为选手) 
12.              = dp[i-1][state^pi] (当前的人作为选手) + s[i][pi] - a[i] (他不能作为观众) + a[?] (要选择一个人作为观众才行)
13. 但是那个要被代替的人是谁呢？应该是state中pi 对应的那个人
14. 这个state不对，它一开始就是全的（前p个人，排列后对应的位置，获得最大值）
15. 然后就可以算出最近的位置是什么
16. 如果 i >= n - k, 那么它的贡献 = -a[i] + s[i][?] (它被分配给某个位置的时候)
17. 否则就是 a[i]
18. 换个角度，就是 dp[i][state] 表示到i为止，满足state的最大值
19. dp[i][state] = dp[i-1][state] or i参与为止j