Phoenix has a string 𝑠
 consisting of lowercase Latin letters. He wants to distribute all the letters of his string into 𝑘
 non-empty strings 𝑎1,𝑎2,…,𝑎𝑘
 such that every letter of 𝑠
 goes to exactly one of the strings 𝑎𝑖
. The strings 𝑎𝑖
 do not need to be substrings of 𝑠
. Phoenix can distribute letters of 𝑠
 and rearrange the letters within each string 𝑎𝑖
 however he wants.

For example, if 𝑠=
 baba and 𝑘=2
, Phoenix may distribute the letters of his string in many ways, such as:

ba and ba
a and abb
ab and ab
aa and bb
But these ways are invalid:

baa and ba
b and ba
baba and empty string (𝑎𝑖
 should be non-empty)
Phoenix wants to distribute the letters of his string 𝑠
 into 𝑘
 strings 𝑎1,𝑎2,…,𝑎𝑘
 to minimize the lexicographically maximum string among them, i. e. minimize 𝑚𝑎𝑥(𝑎1,𝑎2,…,𝑎𝑘)
. Help him find the optimal distribution and print the minimal possible value of 𝑚𝑎𝑥(𝑎1,𝑎2,…,𝑎𝑘)
.

String 𝑥
 is lexicographically less than string 𝑦
 if either 𝑥
 is a prefix of 𝑦
 and 𝑥≠𝑦
, or there exists an index 𝑖
 (1≤𝑖≤𝑚𝑖𝑛(|𝑥|,|𝑦|))
 such that 𝑥𝑖
 < 𝑦𝑖
 and for every 𝑗
 (1≤𝑗<𝑖)
 𝑥𝑗=𝑦𝑗
. Here |𝑥|
 denotes the length of the string 𝑥
.

### ideas
1. 显然只和词频有关系
2. 如果cnt[a] >= k, 那么它可以尽量的分布均匀；
3. 如果不行，那么就需要使用下一个，比如b
4. 还有点难呐
5. 貌似除了第一位，后面的所有位，都需要放在同一个字符串上面，这样的好处是，小的字符尽量的出现在了后面
6. 但是abab, 2这个 
7. (abb, a) 和 (ab, ab) 
8. 显然是ab更小。这里的区别是什么呢？区别就在b的后面没有更多的字符
9. 如果b的后面有c，那么显然abbc是更好的选择 （否则就会出现 abc这样子的结果）