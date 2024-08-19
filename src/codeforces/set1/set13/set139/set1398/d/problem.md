You are given three multisets of pairs of colored sticks:

𝑅
 pairs of red sticks, the first pair has length 𝑟1
, the second pair has length 𝑟2
, …
, the 𝑅
-th pair has length 𝑟𝑅
;
𝐺
 pairs of green sticks, the first pair has length 𝑔1
, the second pair has length 𝑔2
, …
, the 𝐺
-th pair has length 𝑔𝐺
;
𝐵
 pairs of blue sticks, the first pair has length 𝑏1
, the second pair has length 𝑏2
, …
, the 𝐵
-th pair has length 𝑏𝐵
;
You are constructing rectangles from these pairs of sticks with the following process:

take a pair of sticks of one color;
take a pair of sticks of another color different from the first one;
add the area of the resulting rectangle to the total area.
Thus, you get such rectangles that their opposite sides are the same color and their adjacent sides are not the same color.

Each pair of sticks can be used at most once, some pairs can be left unused. You are not allowed to split a pair into independent sticks.

What is the maximum area you can achieve?

### ideas
1. 先要将每种颜色的stick降序排，用长的肯定比用短的好
2. 如果不考虑复杂性dp[i][j][k] 表示当前各剩余i,j,k时的状态
3. dp[i][j][k] => dp[i-1][j-1][k], dp[i-1][j][k-1], dp[i][j-1][k-1] 
4. 但是这个是 n * n * n 的复杂性， boom
5. 没问题， 200 * 200 * 200 = 8 * 1e6
6. 