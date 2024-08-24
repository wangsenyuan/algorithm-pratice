Berland year consists of 𝑚
 months with 𝑑
 days each. Months are numbered from 1
 to 𝑚
. Berland week consists of 𝑤
 days. The first day of the year is also the first day of the week. Note that the last week of the year might be shorter than 𝑤
 days.

A pair (𝑥,𝑦)
 such that 𝑥<𝑦
 is ambiguous if day 𝑥
 of month 𝑦
 is the same day of the week as day 𝑦
 of month 𝑥
.

Count the number of ambiguous pairs.



### ideas
1. 如果pair(x, y) is ambiguous 是说day x of month y in the week, = day y of month in the week
2. (x * d + y - y * d - x) % w = 0
  
### solution
Let the month, the days in them and the days of the week be numbered 0
-based.

Translate the 𝑥
-th day of the 𝑦
-th month to the index of that day in a year; that would be 𝑦𝑑+𝑥
. Thus, the corresponding day of the week is (𝑦𝑑+𝑥) 𝑚𝑜𝑑 𝑤
.

So we can rewrite the condition for a pair as 𝑥𝑑+𝑦≡𝑦𝑑+𝑥 (𝑚𝑜𝑑 𝑤)
. That's also (𝑥𝑑+𝑦)−(𝑦𝑑+𝑥)≡0 (𝑚𝑜𝑑 𝑤)
. Continue with (𝑥−𝑦)(𝑑−1)≡0 (𝑚𝑜𝑑 𝑤)
. So (𝑥−𝑦)(𝑑−1)
 should be divisible by 𝑤
. 𝑑−1
 is fixed and some prime divisors of 𝑤
 might have appeared in it already. If we remove them from 𝑤
 then (𝑥−𝑦)
 should just be divisible by the resulting number. So we can divide 𝑤
 by 𝑔𝑐𝑑(𝑤,𝑑−1)
 and obtain that 𝑤′
.

Now we should just count the number of pairs (𝑥,𝑦)
 such that 𝑥−𝑦
 is divisible by 𝑤′
. We know that the difference 𝑥−𝑦
 should be from 1
 to 𝑚𝑖𝑛(𝑑,𝑚)
. So we can fix the difference and add the number of pairs for that difference. That would be 𝑚𝑛−𝑘
 for a difference 𝑘
. Finally, the answer is ∑𝑖=1𝑚𝑖𝑛(𝑑,𝑚)/𝑤′𝑚𝑛−𝑖⋅𝑤′
. Use the formula for the sum of arithmetic progression to solve that in 𝑂(1)
.

Overall complexity: 𝑂(log(𝑤+𝑑))
 per testcase.