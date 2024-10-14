Tsumugi brought 𝑛
 delicious sweets to the Light Music Club. They are numbered from 1
 to 𝑛
, where the 𝑖
-th sweet has a sugar concentration described by an integer 𝑎𝑖
.

Yui loves sweets, but she can eat at most 𝑚
 sweets each day for health reasons.

Days are 1
-indexed (numbered 1,2,3,…
). Eating the sweet 𝑖
 at the 𝑑
-th day will cause a sugar penalty of (𝑑⋅𝑎𝑖)
, as sweets become more sugary with time. A sweet can be eaten at most once.

The total sugar penalty will be the sum of the individual penalties of each sweet eaten.

Suppose that Yui chooses exactly 𝑘
 sweets, and eats them in any order she wants. What is the minimum total sugar penalty she can get?

Since Yui is an undecided girl, she wants you to answer this question for every value of 𝑘
 between 1
 and 𝑛
.

### ideas
1. 显然，a[i]越大的越应该被吃掉
2.  haha，不对～
3.  是前k个最小中的，倒序排列