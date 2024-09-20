You are given string 𝑠
 of length 𝑛
 consisting of 0-s and 1-s. You build an infinite string 𝑡
 as a concatenation of an infinite number of strings 𝑠
, or 𝑡=𝑠𝑠𝑠𝑠…
 For example, if 𝑠=
 10010, then 𝑡=
 100101001010010...

Calculate the number of prefixes of 𝑡
 with balance equal to 𝑥
. The balance of some string 𝑞
 is equal to 𝑐𝑛𝑡0,𝑞−𝑐𝑛𝑡1,𝑞
, where 𝑐𝑛𝑡0,𝑞
 is the number of occurrences of 0 in 𝑞
, and 𝑐𝑛𝑡1,𝑞
 is the number of occurrences of 1 in 𝑞
. The number of such prefixes can be infinite; if it is so, you must say that.

A prefix is a string consisting of several first letters of a given string, without any reorders. An empty prefix is also a valid prefix. For example, the string "abcd" has 5 prefixes: empty string, "a", "ab", "abc" and "abcd".

### ideas
1. 如果在s中0的1的数目是相同的，且x <= 最大的前缀差，那么就有无数个前缀=x
2. 如果数目相同，且x >最大的前缀差， 那么答案是0
3. 如果s中0的数目不相同，假如是w，那么每重复一次，这个差就会增加w
4. 假设在l的时候, l * w + ? = x
5. 也要分好几种情况