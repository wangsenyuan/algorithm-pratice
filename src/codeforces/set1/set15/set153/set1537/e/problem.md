This is the easy version of the problem. The only difference is the constraints on 𝑛
 and 𝑘
. You can make hacks only if all versions of the problem are solved.

You have a string 𝑠
, and you can do two types of operations on it:

Delete the last character of the string.
Duplicate the string: 𝑠:=𝑠+𝑠
, where +
 denotes concatenation.
You can use each operation any number of times (possibly none).

Your task is to find the lexicographically smallest string of length exactly 𝑘
 that can be obtained by doing these operations on string 𝑠
.

A string 𝑎
 is lexicographically smaller than a string 𝑏
 if and only if one of the following holds:

𝑎
 is a prefix of 𝑏
, but 𝑎≠𝑏
;
In the first position where 𝑎
 and 𝑏
 differ, the string 𝑎
 has a letter that appears earlier in the alphabet than the corresponding letter in 𝑏
.


### ideas
1. 完全没想法～
2. 我们考虑一个场景，就是操作1，不可能发生在中间
3. 也就是说， 这个操作的顺序应该是 1111222221111
4. 假设在2的中间出现了操作1，有两种情况，一种是删除到很短， 长度<= 复制前的字符串，那么没有意义
5. 如S[0] = 'a' 没有啥问题，删除到1， 然后重复k次
6. 或者 s[0]是s中最小的字符
7. 所以，这里就是最多删除到s[i], s[i]是最小的那个字符，
8. 然后重复,最后把多余的删除掉
9. 但是删除到哪个位置呢？ 如果 s[i+1] > s[0], 应该在i+1处
10. 如果s[i+1] < s[0], 应该要继续
11. 如果s[i+1] == s[0], 可能要继续
12. 也就是说，要找到 s[i....j] >= s[0...(j - i + 1)]的位置  