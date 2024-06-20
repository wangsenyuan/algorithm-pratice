You are given a string 𝑠
, consisting of lowercase Latin letters and/or question marks.

A tandem repeat is a string of an even length such that its first half is equal to its second half.

A string 𝑎
 is a substring of a string 𝑏
 if 𝑎
 can be obtained from 𝑏
 by the deletion of several (possibly, zero or all) characters from the beginning and several (possibly, zero or all) characters from the end.

Your goal is to replace each question mark with some lowercase Latin letter in such a way that the length of the longest substring that is a tandem repeat is maximum possible.

### ideas
1. 假设最长的substring的长度是 2 * k
2. 显然k成立的时候，k-1也是成立的
3. 二分查找检查的时候，如果能确定第二段的起点时，就可以很快的检查出来
4. 弄错了， substring指的是中间的部分
5. 那现在就可以二分了
6. 在给定开始的位置，检查，是否存在一个2 * k的substring，满足条件
7. 现在不能二分了，因为k成立的时候， k-1不一定成立
8. 假设[l...r]是l开始的最长的tandem repeat substring
9. 那么当l => l+1的时候，会发生什么？
10. 下一个位置，仍然可以从l...r中间的某个位置开始
11. 从中间开始