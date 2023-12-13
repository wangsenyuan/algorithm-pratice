Recall that string 𝑎
is lexicographically smaller than string 𝑏
if 𝑎
is a prefix of 𝑏
(and 𝑎≠𝑏
), or there exists an index 𝑖
(1≤𝑖≤min(|𝑎|,|𝑏|)
) such that 𝑎𝑖<𝑏𝑖
, and for any index 𝑗
(1≤𝑗<𝑖
) 𝑎𝑗=𝑏𝑗
.

Consider a sequence of strings 𝑠1,𝑠2,…,𝑠𝑛
, each consisting of lowercase Latin letters. String 𝑠1
is given explicitly, and all other strings are generated according to the following rule: to obtain the string 𝑠𝑖
, a character is removed from string 𝑠𝑖−1
in such a way that string 𝑠𝑖
is lexicographically minimal.

For example, if 𝑠1=dacb
, then string 𝑠2=acb
, string 𝑠3=ab
, string 𝑠4=a
.

After that, we obtain the string 𝑆=𝑠1+𝑠2+⋯+𝑠𝑛
(𝑆
is the concatenation of all strings 𝑠1,𝑠2,…,𝑠𝑛
).

You need to output the character in position 𝑝𝑜𝑠
of the string 𝑆
(i. e. the character 𝑆𝑝𝑜𝑠
).

### thoughts

1. 根据pos可以确定是si
2. n + n - 1 + ... + n - i <= pos => i等于多少
3. 第一次删除的，必然是第一个 s[j] > s[j+1]的位置
4. ...
5. 使用一个stack