Bob came to a cash & carry store, put n items into his trolley, and went to the checkout counter to pay. Each item is described by its price ci and time ti in seconds that a checkout assistant spends on this item. While the checkout assistant is occupied with some item, Bob can steal some other items from his trolley. To steal one item Bob needs exactly 1 second. What is the minimum amount of money that Bob will have to pay to the checkout assistant? Remember, please, that it is Bob, who determines the order of items for the checkout assistant.

Input
The first input line contains number n (1 ≤ n ≤ 2000). In each of the following n lines each item is described by a pair of numbers ti, ci (0 ≤ ti ≤ 2000, 1 ≤ ci ≤ 109). If ti is 0, Bob won't be able to steal anything, while the checkout assistant is occupied with item i.

Output
Output one number — answer to the problem: what is the minimum amount of money that Bob will have to pay.


### ideas
1. 假设bob总共花费了x时间来通过checkout，那么就是说，要从其中找出两部分item
2. 分为A/B，其中在a中的时间之和 = x, 在b的个数 <= x, 且希望b中的金额之和最大
3. dp[i][j] 表示前i个物品，偷窃其中j个时，所能节省的最大金额
4. dp[i][j] = max(dp[i-1][j], dp[i-1][j-1] + c[i] if pref of sum of t >= j)