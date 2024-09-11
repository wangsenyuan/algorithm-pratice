Gildong recently learned how to find the longest increasing subsequence (LIS) in 𝑂(𝑛log𝑛)
 time for a sequence of length 𝑛
. He wants to test himself if he can implement it correctly, but he couldn't find any online judges that would do it (even though there are actually many of them). So instead he's going to make a quiz for you about making permutations of 𝑛
 distinct integers between 1
 and 𝑛
, inclusive, to test his code with your output.

The quiz is as follows.

Gildong provides a string of length 𝑛−1
, consisting of characters '<' and '>' only. The 𝑖
-th (1-indexed) character is the comparison result between the 𝑖
-th element and the 𝑖+1
-st element of the sequence. If the 𝑖
-th character of the string is '<', then the 𝑖
-th element of the sequence is less than the 𝑖+1
-st element. If the 𝑖
-th character of the string is '>', then the 𝑖
-th element of the sequence is greater than the 𝑖+1
-st element.

He wants you to find two possible sequences (not necessarily distinct) consisting of 𝑛
 distinct integers between 1
 and 𝑛
, inclusive, each satisfying the comparison results, where the length of the LIS of the first sequence is minimum possible, and the length of the LIS of the second sequence is maximum possible.

### ideas
1. 什么情况下可以取到最小值呢？
2. 考虑n的位置， 如果存在一个最左边的i, s[i] = '<', and s[i+1] = '>' 那么这里就可以放置n
3. 然后再放置n-1。。。 
4. 如果不存在这样一个位置，（所有的都是 ‘《’），只能放在目前最后面的位置
5. 比如 n = 4, <>>>, 4123
6. 这样子得到的最lis = 最长的"<"的序列的长度+1
7. 怎么得到最大值呢？
8. 这里，最好是能够把所有的 <都能连起来
9. 就是把所有<号所在的位置，都放置目前的最小值，然后>号所在的位置，放置最大值即可
10. 最小的lis的结论是成立的，就是最长的递增序列的长度+1
11. 但是怎么构造出来呢？
12. 假设到目前为止已经构造了一个合理的序列，现在出现了>> 递减序列（那么它们肯定要用最小的数字，出现在前面会造成新的Lis)
13. 先不管咋加入的。
14. 然后出现了<<， 理论上这里，应该增加 1,2,3 这样的序列
15. 最后一个肯定是1，所以1，已经存在了，但是2，3在前方已经出现过了
16. 那么只要将前面所有的数字+1即可
17. 