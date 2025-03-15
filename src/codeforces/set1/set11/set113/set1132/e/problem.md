You have a set of items, each having some integer weight not greater than 8
. You denote that a subset of items is good if total weight of items in the subset does not exceed 𝑊
.

You want to calculate the maximum possible weight of a good subset of items. Note that you have to consider the empty set and the original set when calculating the answer.

Input
The first line contains one integer 𝑊
 (0≤𝑊≤1018
) — the maximum total weight of a good subset.

The second line denotes the set of items you have. It contains 8
 integers 𝑐𝑛𝑡1
, 𝑐𝑛𝑡2
, ..., 𝑐𝑛𝑡8
 (0≤𝑐𝑛𝑡𝑖≤1016
), where 𝑐𝑛𝑡𝑖
 is the number of items having weight 𝑖
 in the set.

Output
Print one integer — the maximum possible weight of a good subset of items.


Let's consider the optimal answer. Suppose we take 𝑐𝑖
 items of weight 𝑖
.

Let 𝐿
 be the least common multiple of all weights (that is 840
). Then we may represent 𝑐𝑖
 as 𝑐𝑖=𝐿𝑖𝑝𝑖+𝑞𝑖
, where 0≤𝑞<𝐿𝑖
. Let's do the following trick: we will take 𝑞𝑖
 items of weight 𝑖
, and all the remaining items of this weight can be merged into some items of weight 𝐿
.

Then we can write a brute force solution that picks less than 𝐿𝑖
 items of each weight, transforms the remaining ones into items of weight 𝐿
 as much as possible, and when we fix the whole subset, adds maximum possible number of items of weight 𝐿
 to the answer. This works in something like ∏𝑖=18𝐿𝑖=𝐿88!
 operations, which is too much.

How can we speed it up? Rewrite it using dynamic programming! When we have fixed the number of items we take from 𝑥
 first sets, the only two things that matter now are the current total weight of taken items and the number of items of weight 𝐿
 we can use; and it's obvious that the more items of weight 𝐿
 we can use, the better. So let's write the following dynamic programming solution: 𝑑𝑝[𝑥][𝑦]
 — maximum number of items of weight 𝐿
 we can have, if we processed first 𝑥
 types of items, and current total weight is 𝑦
. Note that the second dimension should have size 8𝐿
.