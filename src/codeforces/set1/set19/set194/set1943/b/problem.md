A string 𝑡
 is said to be 𝑘
-good if there exists at least one substring†
 of length 𝑘
 which is not a palindrome‡
. Let 𝑓(𝑡)
 denote the sum of all values of 𝑘
 such that the string 𝑡
 is 𝑘
-good.

You are given a string 𝑠
 of length 𝑛
. You will have to answer 𝑞
 of the following queries:

Given 𝑙
 and 𝑟
 (𝑙<𝑟
), find the value of 𝑓(𝑠𝑙𝑠𝑙+1…𝑠𝑟)
.
†
 A substring of a string 𝑧
 is a contiguous segment of characters from 𝑧
. For example, "𝚍𝚎𝚏𝚘𝚛
", "𝚌𝚘𝚍𝚎
" and "𝚘
" are all substrings of "𝚌𝚘𝚍𝚎𝚏𝚘𝚛𝚌𝚎𝚜
" while "𝚌𝚘𝚍𝚎𝚜
" and "𝚊𝚊𝚊
" are not.

‡
 A palindrome is a string that reads the same backwards as forwards. For example, the strings "𝚣
", "𝚊𝚊
" and "𝚝𝚊𝚌𝚘𝚌𝚊𝚝
" are palindromes while "𝚌𝚘𝚍𝚎𝚏𝚘𝚛𝚌𝚎𝚜
" and "𝚊𝚋
" are not.


### ideas
1. 先理解一下题目， 一个s是k-good，表示，它至少有一个k长度的substring，不是回文
2. f(s) = 所有的k-good的k的和
3. 如果s是k-good的，那么它一定是 k-1good的吗？
4. s是k-good，表示它有(至少)一个k长度的子串不是回文，但是这里能确认k-1的也不是回文吗？
5. 如果有k-1，那是不是一定有k？不一定，比如 32不是回文, 323是回文
6. k和k-1似乎也没有关系，比如 323是回文， 32不是回文，比如 322 不是回文，但是 22是回文（当然32不是回文）
7. 如果s都是由同一个字符组成，那么f(s) = 0
8. 如果s由两个字符组成，那么它肯定是2-good的，因为必然存在一个边界
9. 它是3-good的吗？必须存在aab，或者 abc 这样的结构
10. 肯定没法每个k都从左到有检查一遍，甚至没有办法每个k都检查一遍
11. 肯定有啥东西被遗漏了
12. 除非s由全a组成，否则，偶数长度的k-good，都存在，2，已经确认了
13. 比如对于4, abba 那么bba？,肯定不是4长度的回文
14. 对于6， a1234a,因为要求是6-good，
15. 1 = 4, 2 = 3, 移动到1开始后，2 = a => 3也是a，再移动2开始后，会要求4 = a => 1 也等于4
16. 如果k=3呢？ 
17. 如果它是ababa... 这样结构的，它就不是3-good的，
18. 同时它也不是5-good的。。。。
19. 如果不是这个结构, 那么它肯定是3-good的，那它肯定也是5-good的， 7-good的。。。
20. 这是因为因为它不是3-good，所以它肯定存在aab，或者abd，或者其他的结构，只要某个长度包括这个子串，那么它就肯定不是回文
21. 所以，总结起来就是
22.   如果 s = (a) 全有a组成，f(s) = 0, 否则，它的是偶数-good的（也就是存在偶数长度的子串，不是回文） 
23.   如果 s = (ababa) 交替子串，那么它的奇数长度的子串(3,5...)都是回文，否则它所有奇数长度-good的（包括3）
24.   s本身必须单独判断
25.   这个可以用hash来判断（正向计算一次hash，反向计算一次hash，比较是否相同）
26.   