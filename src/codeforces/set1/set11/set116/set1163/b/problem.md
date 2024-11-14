Shiro's just moved to the new house. She wants to invite all friends of her to the house so they can play monopoly. However, her house is too small, so she can only invite one friend at a time.

For each of the 𝑛
 days since the day Shiro moved to the new house, there will be exactly one cat coming to the Shiro's house. The cat coming in the 𝑖
-th day has a ribbon with color 𝑢𝑖
. Shiro wants to know the largest number 𝑥
, such that if we consider the streak of the first 𝑥
 days, it is possible to remove exactly one day from this streak so that every ribbon color that has appeared among the remaining 𝑥−1
 will have the same number of occurrences.

For example, consider the following sequence of 𝑢𝑖
: [2,2,1,1,5,4,4,5]
. Then 𝑥=7
 makes a streak, since if we remove the leftmost 𝑢𝑖=5
, each ribbon color will appear exactly twice in the prefix of 𝑥−1
 days. Note that 𝑥=8
 doesn't form a streak, since you must remove exactly one day.

Since Shiro is just a cat, she is not very good at counting and needs your help finding the longest streak.


### ideas
1. 假设找到了这样一个x, 且删除了一个后，剩余x-1
2. 如果删除的是最后一个， 那么最后一个的freq[u[x]] - 1 要能整除 x-1
3. 如果删除的不是最后一个，那么最后一个的freq[u[x]] 要能整除 x- 1
4. 这是一个检查，且我们可以找到目前所有有效的出现过的color, 计算他们的最小freq和最大freq