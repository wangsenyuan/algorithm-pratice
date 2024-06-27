You are given a string 𝑠
 consisting of lowercase Latin characters. Count the number of nonempty strings 𝑡≠
 "𝚊
" such that it is possible to partition†
 𝑠
 into some substrings satisfying the following conditions:

each substring either equals 𝑡
 or "𝚊
", and
at least one substring equals 𝑡
.
†
 A partition of a string 𝑠
 is an ordered sequence of some 𝑘
 strings 𝑡1,𝑡2,…,𝑡𝑘
 (called substrings) such that 𝑡1+𝑡2+…+𝑡𝑘=𝑠
, where +
 represents the concatenation operation.


 ### ideas
 1. 有点迷糊
 2. 对于这样一个string t
 3. s = （'a')(t)('a')(t)('a') ?
 4. 如果s不是以a开头，那么t是不是一定是s的一个前缀，且这个前缀的后面必须是'a'或者是另外一个t
 5. 如果s以a开头，似乎有变的麻烦起来
 6. 但是有个点就是，如果s全部是a，这个号处理 （n - 1), t != 'a'
 7. 如果s包含非'a'， 那么t肯定是某个包含了所有非‘a'的最短字符串（且可以增长）
 8. 那么从s中把非a的摘出来，它们组成新的x，那么t肯定是x的重复子串
 9. 但是回到s会还有问题没有解决，比如 s = abcaabca,  x = bcbc
 10. bc显然是一个答案，但是 abc,也是一个答案, bca也是一个答案， abca也是
 11. 假设通过某种方式，找到了一个t怎么去验证它？
 12. 使用kmp算法，假设在为止i找到了一个匹配，那么上个匹配j到i-len(t)之间的，只能是'a'
 13. 然后开始从头匹配
 14. 这个匹配是 O(n)的，但是即使是按照x的子串这样的方式去找t，这个时间也是 O(sqrt(n))的
 15. 等等，其实 如果 bc是一个可能的答案的话，那么如果每个bc后面有多少个a，b前面有多少个a，是不是也能算出来的？
 16. 假设每个bc前/后的a的个数确定了
 17. 个数可以这样统计 （0, c后面最小的数） + （1，min(last, 中间最小的数-1)） + 。。。 min(first， 中间最小的数), v)
 18. 假设前后(u, v)个a，那么有 u <= min(最开始前面的a的个数, 中间的a)， v <= min(最后面a的个数， 中间的a)
 19. 且 u + v <= (中间a的个数)
 20. 这个可以算 x + y = w, 且 x<= f, y <= s 的面积计算出来
 21. 那么还剩下一个问题，就是如何验证bc可以满足条件
 22. 用上面kmp的方式
 23. 整体是 o(sqrt(n)) * o(n)
 24. 居然搞对了。😂

### solution
Special case: the string consists of only "𝚊
", then answer is 𝑛−1
.

Otherwise, 𝑡
 has to contain a character that is not "𝚊
". Let's consider one approach of counting. Let's force all 𝑡
 to start with the first non-a character, and see how many work.

To see if one of these strings 𝑡
 will work, we can start with 𝑖=0
. As long as 𝑖
 does not point to the end of the string, we can find the next non-a character in 𝑠
, and then see if the next |𝑡|
 characters in 𝑠
 matches 𝑡
. The last check mentioned can be done through hashing or using the Z-function. If it doesn't, this value of 𝑡
 doesn't work. Otherwise, we update 𝑖
 to the new current position and continue checking.

Now, if we find a string 𝑡
 that does work, we need to count its contribution to the answer. We can just add 1
, but obviously not all working 𝑡
 will start with a non-a character. Instead, we can find the minimum unused "𝚊
"s before all of the substrings equal to 𝑡
 (call this 𝑚
), and the current 𝑡
 will be able to extend out up to 𝑚
 "𝚊
"s at the beginning of the string. Thus, the contribution is 𝑚+1
 to the answer.

How fast is this? Because we are checking at most one string 𝑡
 of each length, and the check itself can be made to take 𝑂(𝑛|𝑡|)
, we have a total time complexity of 𝑂(𝑛log𝑛)
 due to harmonic sums.