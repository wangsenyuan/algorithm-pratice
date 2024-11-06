One important contest will take place on the most famous programming platform (Topforces) very soon!

The authors have a pool of 𝑛
 problems and should choose at most three of them into this contest. The prettiness of the 𝑖
-th problem is 𝑎𝑖
. The authors have to compose the most pretty contest (in other words, the cumulative prettinesses of chosen problems should be maximum possible).

But there is one important thing in the contest preparation: because of some superstitions of authors, the prettinesses of problems cannot divide each other. In other words, if the prettinesses of chosen problems are 𝑥,𝑦,𝑧
, then 𝑥
 should be divisible by neither 𝑦
, nor 𝑧
, 𝑦
 should be divisible by neither 𝑥
, nor 𝑧
 and 𝑧
 should be divisible by neither 𝑥
, nor 𝑦
. If the prettinesses of chosen problems are 𝑥
 and 𝑦
 then neither 𝑥
 should be divisible by 𝑦
 nor 𝑦
 should be divisible by 𝑥
. Any contest composed from one problem is considered good.

Your task is to find out the maximum possible total prettiness of the contest composed of at most three problems from the given pool.

You have to answer 𝑞
 independent queries.

### ideas
1. a[i] <= 1e5 
2. 如果是单个数，那么就是最大值
3. 如果有两个数，如果x是较小的那个数，就是找出，不是x的倍数的，最大的数（似乎找出第2大的数）？
4. a + b < c + d 且假设b是最大的数，
5. a < d < b, 没有选择d是，因为b % d = 0
6. 且c < d 所以，c + d < d + d < b (因为d能整除b)
7. 所以，结论是，如果是两个数，那么其中一个肯定是最大的数
8. 现在考虑3个数的情况
9. 假设 a, b, c是一个最优解。
10. 例子2可以看出，不一定要选最大值
11. 但是