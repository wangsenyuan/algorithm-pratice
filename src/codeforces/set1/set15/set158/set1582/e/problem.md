Pchelyonok decided to give Mila a gift. Pchelenok has already bought an array 𝑎
 of length 𝑛
, but gifting an array is too common. Instead of that, he decided to gift Mila the segments of that array!

Pchelyonok wants his gift to be beautiful, so he decided to choose 𝑘
 non-overlapping segments of the array [𝑙1,𝑟1]
, [𝑙2,𝑟2]
, …
 [𝑙𝑘,𝑟𝑘]
 such that:

the length of the first segment [𝑙1,𝑟1]
 is 𝑘
, the length of the second segment [𝑙2,𝑟2]
 is 𝑘−1
, …
, the length of the 𝑘
-th segment [𝑙𝑘,𝑟𝑘]
 is 1
for each 𝑖<𝑗
, the 𝑖
-th segment occurs in the array earlier than the 𝑗
-th (i.e. 𝑟𝑖<𝑙𝑗
)
the sums in these segments are strictly increasing (i.e. let 𝑠𝑢𝑚(𝑙…𝑟)=∑𝑖=𝑙𝑟𝑎𝑖
 — the sum of numbers in the segment [𝑙,𝑟]
 of the array, then 𝑠𝑢𝑚(𝑙1…𝑟1)<𝑠𝑢𝑚(𝑙2…𝑟2)<…<𝑠𝑢𝑚(𝑙𝑘…𝑟𝑘)
).
Pchelenok also wants his gift to be as beautiful as possible, so he asks you to find the maximal value of 𝑘
 such that he can give Mila a gift!


 ### ideas
 1. 对于给定的k, 这些segments的总长度 = (k + 1) * k / 2 <= n =》 k 的上界
 2. k >= 1
 3. 假设对于给定的k来说，s[1] 越小越好，r[1]越小越好
 4. s[1]越小，后面的sum的压力越小，r[1]越小，后面越有足够的空间得到k-1
 5. dp[i][j] 表示后缀i...n， 分割出j段后，最大的sum
 6. dp[i][j] = dp[i+1][j] 或者 s = 如果 sum(i...i+j-1)  < dp[i+j][j-1]
 7.  