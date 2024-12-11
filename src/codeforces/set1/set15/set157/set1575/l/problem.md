Mr. Chanek gives you a sequence 𝑎
 indexed from 1
 to 𝑛
. Define 𝑓(𝑎)
 as the number of indices where 𝑎𝑖=𝑖
.

You can pick an element from the current sequence and remove it, then concatenate the remaining elements together. For example, if you remove the 3
-rd element from the sequence [4,2,3,1]
, the resulting sequence will be [4,2,1]
.

You want to remove some elements from 𝑎
 in order to maximize 𝑓(𝑎)
, using zero or more operations. Find the largest possible 𝑓(𝑎)
.

### ideas
1. 如果删除i，那么i后面所有的数的位置前移1
2. 使用b[i] = i - a[i]
3. 只需要考虑那些 >= 0 的位置
4. divide and conquer but how?
5. dfs(l, r) => freq 
6. 表示，不管用什么方式，得到的在这个区间内部， b[i] = x 的最大值
7. 其中有一些是删除了区间里面的某些位置后得到的 (不对， 因为如果这个区间里面删除了， 它的父区间会被影响到)
8. dfs(l, r) => dp 
9. 表示dp[0] 表示在这个区间内部不删除任何元素时的最大值
10. dp[1]表示，删除了某个位置后得到的最大值
11. dp[i]表示删除了i个位置后得到的最大值
12. 怎么merge
13. dp, fp分别表示左右的结果，结果 = gp
14. dp[i] 表示左边删除了i, gp[i+j] = max(dp[i] + fp[i+j])
15. 这里有个关键的就是， 左边删除的，其实对右边没有影响？不对上面那个式子是不对的
16. fp[x]表示的是在区间总共删除了x个，有可能在位置mid删除， mid+1处删除，也可能是在比较靠后的位置删除
17. 那么此时在他们的前面删除, 那么那些本来符合条件的，就变成不符合条件了
18. 那这些不符合条件的是不是正好等于 fp[i]呢？还是 sum(fp[...i])?
19. 比如  0 1 2 2, 1
20. fp[0] = 1, fp[1] = 1, fp[2] = 2
21. 如果前面删除掉1, 且计算 dp[3] = 1 (因为删除掉1个，-1，0, 1, 1)
22. 且再删除掉两个2（那么就只能是1, 1)
23. gp[j] = dp[i] + fp[j-i]似乎不成立
24. 前面删掉1个后， 0不满足条件了，1满足条件了，但是比如 fp[1]， 有可能删掉的就是第一个1, 且0可能是被算入的
25. 只需要关心 >= 0 的那些数字, dfs(l, r)处理后，返回两个结果
26. 1个是在这个区间的最优解,2是只保留>=0 的那些数字的位置（和顺序的情况）的数组
27. best = max(l-best, r-best)
28. 两边数组merge在一起后，得到一个新的best？
29. 假设 0,0, 1, 1,2, 2,  0, 1, 3 假设是这样的
30. 假设要计算删除w个元素的最优解
31. gp[w][a] = dp[w][a] + fp[0][a+w]
32.          = dp[w-1][a] + fp[1][a + w - 1]
33.          ...
34.          = dp[w-i][a] + fp[i][a + w - i]
35. 左边进行w-i次操作，右边进行i次操作后，最后的数组中产生最多gp[w][a]个数 = a
36. 直接爆炸
37. b[i] = i - a[i] 
38. 按照 b[i] 升序进行处理
39. 假设有一些b[i] = 0, 这些0一定会被保留吗？
40. 比如 0, 0, 0, 1, 1, 1, 2， 2, 2
41. 那么删除掉一个0， 可以得到3个1，删除掉一个1，可以得到3个2
42. 这样是更优的选择
43. 从后往前处理，也是不行的（因为会影响到后面已经ok的）那么从前往后处理
44. dp[i][j]表示在前i个数中删除了j个后，得到的最大值
45. dp[i][j] = dp[i-1][j] + 1 如果 b[i] == j else 0 (不删除a[i])
46.          = dp[i-1][j-1] (删除a[i]时)
47.          = dp[i-1][j] 不删除a[i]
48.  好像有戏了
49.  看起来, dp[i]只依赖于dp[i-1], 要么等于j,要么等于 j-1
50.  从i-1到i转移的时候， 先记录下 dp[j] 的值 v（j = b[i])
51.  然后所有的下标右移（怎么处理）然后更新 j = b[i], v + 1
52.  dp[j] = max(dp[j], dp[j-1]) 这个要怎么处理呢？
53.  2 1 4 2 5 3 7
54.  1 2 3 4 5 6 7
55.  -1, 1, -1, 2, 0, 3, 0
56.  fp[x]表示最后一个数是x时的最大值
57.  b[i] = v, fp[v] = fp[v-1] + 1(不删除 b[i]的情况下)
58.  删除b[i]的情况下，如果删除， 后面所有的v都需要-1
59.  