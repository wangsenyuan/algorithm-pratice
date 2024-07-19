Polycarp loves to listen to music, so he never leaves the player, even on the way home from the university. Polycarp overcomes the distance from the university to the house in exactly 𝑇
 minutes.

In the player, Polycarp stores 𝑛
 songs, each of which is characterized by two parameters: 𝑡𝑖
 and 𝑔𝑖
, where 𝑡𝑖
 is the length of the song in minutes (1≤𝑡𝑖≤50
), 𝑔𝑖
 is its genre (1≤𝑔𝑖≤3
).

Polycarp wants to create such a playlist so that he can listen to music all the time on the way from the university to his home, and at the time of his arrival home, the playlist is over. Polycarp never interrupts songs and always listens to them from beginning to end. Thus, if he started listening to the 𝑖
-th song, he would spend exactly 𝑡𝑖
 minutes on its listening. Polycarp also does not like when two songs of the same genre play in a row (i.e. successively/adjacently) or when the songs in his playlist are repeated.

Help Polycarpus count the number of different sequences of songs (their order matters), the total duration is exactly 𝑇
, such that there are no two consecutive songs of the same genre in them and all the songs in the playlist are different.

### ideas
1. 这个list里面的歌曲不需要是顺序的；所以不能简单的用dp[x][g]来表示（这样假定了歌曲是按顺序处理）
2. 假设cnt[0], cnt[1], cnt[2]是三类歌曲的数目
3. 似乎可以用熔斥定理？
4. 不考虑限制的话是 n!中
5. 其中歌曲0连续出现两次的是 (n - 1) * (n - 2)!
6. 还是有点不清楚。
7. 如果 cnt[0] > n / 2 => 0 （假设cnt[0] 最多）
8. 考虑n是偶数的情况，简单点，把它们分成h = n/2个格子。每个格子里面，必须放入两种不同的歌曲（顺序先不考虑）
9. 这个时候，好像可以用dp了。
10. dp[i][j][k][x][s] 表示使用共i个0类歌曲，j个1类，k个2类歌曲时，最后一个是哪一类, 且总和是s时的总计数
11. 如果x!= 0, 那么现在可以放入一个新的0类歌曲，放入它的话，ns = s + t[i], 且计数 * （i+1）（因为有(i+1)个0类歌曲的位置）
12. 50 * 50 * 50 * 3 * 2500 = 1e10 TLE
13. 15 * 15 * 15 * 3 * 225 = 还是TLE
14. 我们可以考虑这样一个情况就是，这三个如何搭配，就不会出现连续的两个
15. 假设共有m个格子 (m <= n / 2) 其中x个分给多数的歌曲，剩下 (m - x)个不包括多数, 剩下x个位置（和多数匹配的位置）
16. 