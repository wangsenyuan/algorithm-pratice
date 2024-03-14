You are given two arrays: an array 𝑎
consisting of 𝑛
zeros and an array 𝑏
consisting of 𝑛
integers.

You can apply the following operation to the array 𝑎
an arbitrary number of times: choose some subsegment of 𝑎
of length 𝑘
and add the arithmetic progression 1,2,…,𝑘
to this subsegment — i. e. add 1
to the first element of the subsegment, 2
to the second element, and so on. The chosen subsegment should be inside the borders of the array 𝑎
(i.e., if the left border of the chosen subsegment is 𝑙
, then the condition 1≤𝑙≤𝑙+𝑘−1≤𝑛
should be satisfied). Note that the progression added is always 1,2,…,𝑘
but not the 𝑘,𝑘−1,…,1
or anything else (i.e., the leftmost element of the subsegment always increases by 1
, the second element always increases by 2
and so on).

Your task is to find the minimum possible number of operations required to satisfy the condition 𝑎𝑖≥𝑏𝑖
for each 𝑖
from 1
to 𝑛
. Note that the condition 𝑎𝑖≥𝑏𝑖
should be satisfied for all elements at once.

### ideas

1. 每次都加了 (k +1 ) * k / 2
2. 对于b[0] 来说，它只能添加b[0]次，然后贪心的处理下一个
3. 这里需要知道a[i]被增加了多少
4. a[i]到a[i+k]被分别增加了1。。。。k次
5. 肯定不能每个都处理一次
6. 需要倒过来处理，这样可以保证，在i处是贪心的得到结果
7. 所以需要知道在它的前面的哪些位置进行了多少次操作
8. 假设在j处进行了x次操作，对i的影响是
9. dist = i - j < k
10. x * (dist + 1) = x * (i - j + 1) = x * i - x * j + x
11. 假设知道这个区间[i - k, i]区间那的sum 和 count