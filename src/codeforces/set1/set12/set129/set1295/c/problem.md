You are given two strings 𝑠
 and 𝑡
 consisting of lowercase Latin letters. Also you have a string 𝑧
 which is initially empty. You want string 𝑧
 to be equal to string 𝑡
. You can perform the following operation to achieve this: append any subsequence of 𝑠
 at the end of string 𝑧
. A subsequence is a sequence that can be derived from the given sequence by deleting zero or more elements without changing the order of the remaining elements. For example, if 𝑧=𝑎𝑐
, 𝑠=𝑎𝑏𝑐𝑑𝑒
, you may turn 𝑧
 into following strings in one operation:

𝑧=𝑎𝑐𝑎𝑐𝑒
 (if we choose subsequence 𝑎𝑐𝑒
);
𝑧=𝑎𝑐𝑏𝑐𝑑
 (if we choose subsequence 𝑏𝑐𝑑
);
𝑧=𝑎𝑐𝑏𝑐𝑒
 (if we choose subsequence 𝑏𝑐𝑒
).
Note that after this operation string 𝑠
 doesn't change.

Calculate the minimum number of such operations to turn string 𝑧
 into string 𝑡
.

Input
The first line contains the integer 𝑇
 (1≤𝑇≤100
) — the number of test cases.

The first line of each testcase contains one string 𝑠
 (1≤|𝑠|≤105
) consisting of lowercase Latin letters.

The second line of each testcase contains one string 𝑡
 (1≤|𝑡|≤105
) consisting of lowercase Latin letters.

It is guaranteed that the total length of all strings 𝑠
 and 𝑡
 in the input does not exceed 2⋅105
.

Output
For each testcase, print one integer — the minimum number of operations to turn string 𝑧
 into string 𝑡
. If it's impossible print −1
.

### ideas
1. 首先需要确定一个点，是不是贪心的匹配越长子串越好？
2. 假设当前的z的前缀时x，如果能够匹配xa, 是不是匹配掉更好？
3. z = xay, 假设 xa，ay 都可以被匹配到，那么优先匹配掉xa是更安全的做法，这样子y更有可能被匹配到
4. 所有有一个贪心的方案，就是尽量得匹配z的前缀，然后从头来过
5. 为了匹配的更快，需要记录next[i][x]