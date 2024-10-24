You are given a Young diagram.

Given diagram is a histogram with n
 columns of lengths a1,a2,…,an
 (a1≥a2≥…≥an≥1
).

Young diagram for a=[3,2,2,2,1]
.
Your goal is to find the largest number of non-overlapping dominos that you can draw inside of this histogram, a domino is a 1×2
 or 2×1
 rectangle.

 ### ideas
 1. 要用到dp，但是怎么进行状态迁移呢？
 2. 感觉从前面的状态迁移，最优的覆盖方式下，始终能够填满，要么在右下角留下一个空位
 3. 要么把空位留在中间（无法被利用到）
 4. 要么把空位留在右下角
 5. 要么把空位留在右下角的上面（2和3是排斥的，不应该出现两个位置都为空的情况，否则的话，加一块新的更优）
 6. dp[i][0]表示没有可利用空位的状态，dp[i][1]表示留了右下角的空位，dp[i][2]表示留了右下角上方的空位（如果存在的话）
 7. i => i+1的转移  
 8.  (如果a[i]是奇数) dp[i][1] = max(dp[i-1][0/1/2] )+ a[i] / 2
 9.     dp[i][2] = dp[i-1][1] + 1 + (a[i] - 2) / 2 (这种情况下无法出现) (虽然有两个位置)
 10.    dp[i][0] = max of dp[i][1], dp[i-1][1] + 1 + a[i] / 2 (和前面留空的右下角匹配)
 11. 如果a[i]是偶数 dp[i][1] = dp[i-2] + 1 + (a[i] - 2) / 2
 12.  dp[i][2] = dp[i-1][1] + 1 + (a[i] - 2) / 2
 13.  dp[i][0] = max(dp[i][1], dp[i][2], dp[i-1][0/1/2] + a[i] / 2) 
 14.  