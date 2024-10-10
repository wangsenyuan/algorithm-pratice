This problem is split into two tasks. In this task, you are required to find the minimum possible answer. In the task Village (Maximum) you are required to find the maximum possible answer. Each task is worth 50
 points.

There are 𝑁
 houses in a certain village. A single villager lives in each of the houses. The houses are connected by roads. Each road connects two houses and is exactly 1
 kilometer long. From each house it is possible to reach any other using one or several consecutive roads. In total there are 𝑁−1
 roads in the village.

One day all villagers decided to move to different houses — that is, after moving each house should again have a single villager living in it, but no villager should be living in the same house as before. We would like to know the smallest possible total length in kilometers of the shortest paths between the old and the new houses for all villagers.


Example village with seven houses

For example, if there are seven houses connected by roads as shown on the figure, the smallest total length is 8
 km (this can be achieved by moving 1→6
, 2→4
, 3→1
, 4→2
, 5→7
, 6→3
, 7→5
).

Write a program that finds the smallest total length of the shortest paths in kilometers and an example assignment of the new houses to the villagers.

Input
The first line contains an integer 𝑁
 (1<𝑁≤105
). Houses are numbered by consecutive integers 1,2,…,𝑁
.

Then 𝑁−1
 lines follow that describe the roads. Each line contains two integers 𝑎
 and 𝑏
 (1≤𝑎,𝑏≤𝑁
, 𝑎≠𝑏
) denoting that there is a road connecting houses 𝑎
 and 𝑏
.

Output
In the first line output the smallest total length of the shortest paths in kilometers.

In the second line describe one valid assignment of the new houses with the smallest total length: 𝑁
 space-separated distinct integers 𝑣1,𝑣2,…,𝑣𝑁
. For each 𝑖
, 𝑣𝑖
 is the house number where the villager from the house 𝑖
 should move (𝑣𝑖≠𝑖
). If there are several valid assignments, output any of those.


### ideas
1. 感觉无从下手啊
2. 在一个子树u中,只要size(u) > 1, 那么在内部肯定是可以实现move的
3. dp(u) = 在子树u中完成转换
4. 但是存在一种情况， u需要进入某个子树v， 那么此时v中就会出来一个节点
5. 但是出来的这个v, 怎么移动呢？最好是移动到u，但是存在3个的那种情况
6. dp[u] = sum of dp[c][0] + dp[v][1] + 2
7. size(u) % 2 = 0, 那么只需要把u和其中的某个sz[v] 是奇数的部分进行交换即可
8. 那是不是一定存在奇数个数的子树呢？肯定是存在的（否则， 都是偶数的话，它们加起来也是偶数， 再加1就变奇数了）
9. dp[u][1]表示这棵树内部所有的节点都move后，但是u保持位置的，最优解
10. dp[u][0]表示这棵树内部所有节点都move后，u也移动到某个位置的最优解
11. size[u] 是偶数时, dp[u][0] = sum of dp[c][0] + dp[v][1] + 2 (u/v互换位置)
12.                  dp[u][1] = sum of dp[c][0]
13.         是奇数时， dp[u][1] = sum of dp[c][0] 
14.                 dp[u][0] = ？ dp[v][1], dp[w][1] + 4 (u移动到v， v 移动到w（2）， w移动到u)
15.                    dp[v][1] - dp[v][0], dp[w][1] - dp[w][0] 最小的两个数
16. 好像也不用区分奇偶性，