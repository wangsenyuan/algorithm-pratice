This time, the brothers are dealing with a strange piece of wood marked with their names. This plank of wood can be represented as a string of 𝑛
 characters. Each character is either a 'D' or a 'K'. You want to make some number of cuts (possibly 0
) on this string, partitioning it into several contiguous pieces, each with length at least 1
. Both brothers act with dignity, so they want to split the wood as evenly as possible. They want to know the maximum number of pieces you can split the wood into such that the ratios of the number of occurrences of 'D' to the number of occurrences of 'K' in each chunk are the same.

Kaeya, the curious thinker, is interested in the solution for multiple scenarios. He wants to know the answer for every prefix of the given string. Help him to solve this problem!

For a string we define a ratio as 𝑎:𝑏
 where 'D' appears in it 𝑎
 times, and 'K' appears 𝑏
 times. Note that 𝑎
 or 𝑏
 can equal 0
, but not both. Ratios 𝑎:𝑏
 and 𝑐:𝑑
 are considered equal if and only if 𝑎⋅𝑑=𝑏⋅𝑐
.

For example, for the string 'DDD' the ratio will be 3:0
, for 'DKD' — 2:1
, for 'DKK' — 1:2
, and for 'KKKKDD' — 2:4
. Note that the ratios of the latter two strings are equal to each other, but they are not equal to the ratios of the first two strings.

### ideas
1. 假设前面的被分成了很多段，且每段的比例是 a:b, 那么整段的比例也是 a:b
2. 那么对于一个给定的prefix来说，就是不断的分解出a/b的个数
3. got