You are given a string 𝑠
, consisting of lowercase Latin letters. While there is at least one character in the string 𝑠
 that is repeated at least twice, you perform the following operation:

you choose the index 𝑖
 (1≤𝑖≤|𝑠|
) such that the character at position 𝑖
 occurs at least two times in the string 𝑠
, and delete the character at position 𝑖
, that is, replace 𝑠
 with 𝑠1𝑠2…𝑠𝑖−1𝑠𝑖+1𝑠𝑖+2…𝑠𝑛
.
For example, if 𝑠=
"codeforces", then you can apply the following sequence of operations:

𝑖=6⇒𝑠=
"codefrces";
𝑖=1⇒𝑠=
"odefrces";
𝑖=7⇒𝑠=
"odefrcs";
Given a given string 𝑠
, find the lexicographically maximum string that can be obtained after applying a certain sequence of operations after which all characters in the string become unique.

A string 𝑎
 of length 𝑛
 is lexicographically less than a string 𝑏
 of length 𝑚
, if:

there is an index 𝑖
 (1≤𝑖≤min(𝑛,𝑚)
) such that the first 𝑖−1
 characters of the strings 𝑎
 and 𝑏
 are the same, and the 𝑖
-th character of the string 𝑎
 is less than 𝑖
-th character of string 𝑏
;
or the first min(𝑛,𝑚)
 characters in the strings 𝑎
 and 𝑏
 are the same and 𝑛<𝑚
.
For example, the string 𝑎=
"aezakmi" is lexicographically less than the string 𝑏=
"aezus".

### ideas
1. 显然首字母是出现的最大的那个字符
2. 然后是这个字符后面的，第二大的那个字符。
3. 如果是这样，岂不是太简单了？
4. 不对， 只有一次的是不能删的，能删的是那些重复了多次的
5. 并且每个字符最后都会剩下一个，所以整个字符的长度 <= 26
6. 但是这个状态怎么表示呢？怎么开始操作呢？
7. 从左到右，如果现在的字符可以删除（repeat了多次）怎么判断，要不要删除呢？
8. 一个规则是，如果它后面的比它大，且不能删除的话，那么删除它是更优的
9. 但是如果它后面的比它小，且不能删除的话，就不能删除它；
10. 但是还有其他的情况，特别是后续的字符也能删除的情况下。
11. 比如后面的比较大，但是通过删除它，已经其他的一些，字符，还是能够获得更好的结果
12. 如果处理到i时， 要判断是否删除s[i]，前面的都已经处理好了，但后面字符的长度也知道了（假设）
13. 比如判断能否在s[i]后面找到足够的字符串（长度为l，且它的头部 > s[i])那么就删除它，否则就保留
14. 进了一步；
15. 要判断s[i]后面是否有足够的字符串，且这个字符串的长度 > s[i], 定义一个状态suf[i][k] = x
16. 表示在后缀s[i:]在长度至少为k时的最大头部是什么？
17. suf[i][k] = max(s[i] if s[i...] 有至少k个distinct 字符, s[i-1][k])
18. 这里没有考虑，前面保留某个字符，意味着后面这个字符就不能用的情况