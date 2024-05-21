There is a fun game where you need to feed cats that come and go. The level of the game consists of 𝑛
 steps. There are 𝑚
 cats; the cat 𝑖
 is present in steps from 𝑙𝑖
 to 𝑟𝑖
, inclusive. In each step, you can feed all the cats that are currently present or do nothing.

If you feed the same cat more than once, it will overeat, and you will immediately lose the game. Your goal is to feed as many cats as possible without causing any cat to overeat.

Find the maximum number of cats you can feed.

Formally, you need to select several integer points from the segment from 1
 to 𝑛
 in such a way that among given segments, none covers two or more of the selected points, and as many segments as possible cover one of the selected points.


### ideas
1. 对猫按照r[i]升序排
2. 如果选择喂第i个猫，那么在l[i]后面，不能喂过猫，否则会让猫i吃多了
3. 假设在时刻r[i]喂猫，dp[r[i]] = max(dp[0...l[i]-1]) + 1
4. 假设在dp[i]处（之前）在不违法过量的情况下，共喂了dp[i]只猫
5. dp[i] = dp[i-1] 不喂猫
6. 或者 dp[j] + x j < 目前仍然在场的猫的最小的l,x是目前在场猫的数量