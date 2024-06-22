Oleg received a permutation 𝑎
 of length 𝑛
 as a birthday present.

Oleg's friend Nechipor asks Oleg 𝑞
 questions, each question is characterized by two numbers 𝑙
 and 𝑟
, in response to the question Oleg must say the number of sets of indices (𝑡1,𝑡2,…,𝑡𝑘)
 of any length 𝑘≥1
 such that:

𝑙≤𝑡𝑖≤𝑟
 for each 𝑖
 from 1
 to 𝑘
.
𝑡𝑖<𝑡𝑖+1
 for each 𝑖
 from 1
 to 𝑘−1
.
𝑎𝑡𝑖+1
 is divisible by 𝑎𝑡𝑖
 for each 𝑖
 from 1
 to 𝑘−1
.
Help Oleg and answer all of Nechipor's questions.


### ideas
1. 描述倒是很清楚。但是没有任何想法～～～
2. 一个点，就是a是一个permuation，所以各不相同
3. 在一个区间内，l...r， 如果存在一个长度为k的序列，依次能够整除，
4. 这个序列，基本可以被(x, m) x是第一个元素，m是这个序列的长度
5. 对于任何一个这样的序列，它的贡献都是 2 ** m - 1, 因为这个序列中的任何一个元素，在/或不在都不影响这个结果
6. 学习一下solution