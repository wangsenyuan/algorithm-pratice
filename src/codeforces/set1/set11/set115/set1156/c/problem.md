You are given a set of points 𝑥1
, 𝑥2
, ..., 𝑥𝑛
 on the number line.

Two points 𝑖
 and 𝑗
 can be matched with each other if the following conditions hold:

neither 𝑖
 nor 𝑗
 is matched with any other point;
|𝑥𝑖−𝑥𝑗|≥𝑧
.
What is the maximum number of pairs of points you can match with each other?


### ideas
1. sort it first
2. 那么第一个肯定要匹配掉， 如果它不匹配掉，假设在最优答案中，使用了x[2]是第一个，他们某个x[i]匹配了
3. 那么把x[2]换成x[1]， 不会更差，且因为x[2]被空出来了，一定程度上还更优了
4. 确实不对的，考虑这样一种情况z = 3, x = [1, 4, 4, 4, 7, 7]