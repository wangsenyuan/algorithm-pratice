VK news recommendation system daily selects interesting publications of one of 𝑛
 disjoint categories for each user. Each publication belongs to exactly one category. For each category 𝑖
 batch algorithm selects 𝑎𝑖
 publications.

The latest A/B test suggests that users are reading recommended publications more actively if each category has a different number of publications within daily recommendations. The targeted algorithm can find a single interesting publication of 𝑖
-th category within 𝑡𝑖
 seconds.

What is the minimum total time necessary to add publications to the result of batch algorithm execution, so all categories have a different number of publications? You can't remove publications recommended by the batch algorithm.

### ideas
1. 只能增加额外的题目（x * t[i])
2. 使得所有的种类的题目的数量不同
3. 增加的题目越少，所花的时间应该是越少的，在必须改变时，应该选择花费时间最小的那个
4. 所以按照数量排序，然后如果同一个数量下，多余一个，应该选择最小的花费的那些，进行增加，一次增加1....(x)
5. 