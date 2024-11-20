First-rate specialists graduate from Berland State Institute of Peace and Friendship. You are one of the most talented students in this university. The education is not easy because you need to have fundamental knowledge in different areas, which sometimes are not related to each other.

For example, you should know linguistics very well. You learn a structure of Reberland language as foreign language. In this language words are constructed according to the following rules. First you need to choose the "root" of the word — some string which has more than 4 letters. Then several strings with the length 2 or 3 symbols are appended to this word. The only restriction — it is not allowed to append the same string twice in a row. All these strings are considered to be suffixes of the word (this time we use word "suffix" to describe a morpheme but not the few last characters of the string as you may used to).

Here is one exercise that you have found in your task list. You are given the word s. Find all distinct strings with the length 2 or 3, which can be suffixes of this word according to the word constructing rules in Reberland language.

Two strings are considered distinct if they have different length or there is a position in which corresponding characters do not match.

Let's look at the example: the word abacabaca is given. This word can be obtained in the following ways: , where the root of the word is overlined, and suffixes are marked by "corners". Thus, the set of possible suffixes for this word is {aca, ba, ca}.

### ideas
1. 如果 xy 可以作为一个suffix，那么前面可以不管，但是后面在2，3这种组合中，不能出现重复的（不是和xy重复）
2. dp[i] 表示后缀[i...]是否可以work，
3. dp[i] = dp[i+2] 并且后面没有xy 或者 dp[i+3] and 后面没有 xyz
4. 不对的， 比如 dp[i+2]成立，且后面有 xy, 但是如果把x,y分在前后两组里面也是OKde
5. 理解错误题目了， 是不能连续添加，不是不能添加两个相同的。。。
6. 😢