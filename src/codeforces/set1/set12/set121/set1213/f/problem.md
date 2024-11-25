Authors have come up with the string 𝑠
 consisting of 𝑛
 lowercase Latin letters.

You are given two permutations of its indices (not necessary equal) 𝑝
 and 𝑞
 (both of length 𝑛
). Recall that the permutation is the array of length 𝑛
 which contains each integer from 1
 to 𝑛
 exactly once.

For all 𝑖
 from 1
 to 𝑛−1
 the following properties hold: 𝑠[𝑝𝑖]≤𝑠[𝑝𝑖+1]
 and 𝑠[𝑞𝑖]≤𝑠[𝑞𝑖+1]
. It means that if you will write down all characters of 𝑠
 in order of permutation indices, the resulting string will be sorted in the non-decreasing order.

Your task is to restore any such string 𝑠
 of length 𝑛
 consisting of at least 𝑘
 distinct lowercase Latin letters which suits the given permutations.

If there are multiple answers, you can print any of them.

### ideas
1. 感觉是个图的问题
2. s[p[i]] <= s[p[i+1]]
3. 那么在 p[i] -> p[i+1]中间加一条边。 同样的对q也进行处理
4. 如果出现了环，那么就智能用同一个字符
5. 所以，如果最后环的个数，不到k => false
6. 否则就肯定可以