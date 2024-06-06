Harry Potter is hiking in the Alps surrounding Lake Geneva. In this area there are 𝑚
 cabins, numbered 1 to 𝑚
. Each cabin is connected, with one or more trails, to a central meeting point next to the lake. Each trail is either short or long. Cabin 𝑖
 is connected with 𝑠𝑖
 short trails and 𝑙𝑖
 long trails to the lake.

Each day, Harry walks a trail from the cabin where he currently is to Lake Geneva, and then from there he walks a trail to any of the 𝑚
 cabins (including the one he started in). However, as he has to finish the hike in a day, at least one of the two trails has to be short.

How many possible combinations of trails can Harry take if he starts in cabin 1 and walks for 𝑛
 days?

Give the answer modulo 109+7
.


### ideas
1. 如果 n = 1? 如果他选择short, 那么有 s1, * sum(li) + l1 * sum(si)
2. 所以，必须知道他在第i天的时候处在什么位置
3. dp[i][j] 表示第i天的时候处在j, dp[i+1][x] = dp[i][j] * (s[j] * l[x] + l[j] * s[x])