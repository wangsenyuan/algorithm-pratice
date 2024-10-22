You are given two strings s
 and t
 both of length n
 and both consisting of lowercase Latin letters.

In one move, you can choose any length len
 from 1
 to n
 and perform the following operation:

Choose any contiguous substring of the string s
 of length len
 and reverse it;
at the same time choose any contiguous substring of the string t
 of length len
 and reverse it as well.
Note that during one move you reverse exactly one substring of the string s
 and exactly one substring of the string t
.

Also note that borders of substrings you reverse in s
 and in t
 can be different, the only restriction is that you reverse the substrings of equal length. For example, if len=3
 and n=5
, you can reverse s[1…3]
 and t[3…5]
, s[2…4]
 and t[2…4]
, but not s[1…3]
 and t[1…2]
.

Your task is to say if it is possible to make strings s
 and t
 equal after some (possibly, empty) sequence of moves.

You have to answer q
 independent test cases.

 ### ideas
 1. 完全没想法
 2. 同一个位置有可能被翻转很多次，且被移动到不知道那里去？
 3. 能不能固定len = 2, 那这样子，就相当于将相邻的位置进行交换？
 4. 那这样子，只要s和t排序后，相同就可以了？
 5. 不对。 还必须s/t的inversion count要一致
 6. 假设len = 3, 那么可以拆成2次len = 2的，不改变结果，更长的也是如此