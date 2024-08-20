You are given a sequence 𝑎
 consisting of 𝑛
 positive integers.

Let's define a three blocks palindrome as the sequence, consisting of at most two distinct elements (let these elements are 𝑎
 and 𝑏
, 𝑎
 can be equal 𝑏
) and is as follows: [𝑎,𝑎,…,𝑎𝑥,𝑏,𝑏,…,𝑏𝑦,𝑎,𝑎,…,𝑎𝑥]
. There 𝑥,𝑦
 are integers greater than or equal to 0
. For example, sequences []
, [2]
, [1,1]
, [1,2,1]
, [1,2,2,1]
 and [1,1,2,1,1]
 are three block palindromes but [1,2,3,2,1]
, [1,2,1,2,1]
 and [1,2]
 are not.

Your task is to choose the maximum by length subsequence of 𝑎
 that is a three blocks palindrome.

You have to answer 𝑡
 independent test cases.

Recall that the sequence 𝑡
 is a a subsequence of the sequence 𝑠
 if 𝑡
 can be derived from 𝑠
 by removing zero or more elements without changing the order of the remaining elements. For example, if 𝑠=[1,2,1,3,1,2,1]
, then possible subsequences are: [1,1,1,1]
, [3]
 and [1,2,1,3,1,2,1]
, but not [3,2,3]
 and [1,1,1,1,2]
.


### ideas
1. 这里就是选择a,b (26 * 26)然后找到最长的subsequence符合 aba 的
2. 寻找aba，可以以a为基准，两边x个a的时候，有多少个b在中间