Roman and Denis are on the trip to the programming competition. Since the trip was long, they soon got bored, and hence decided to came up with something. Roman invented a pizza's recipe, while Denis invented a string multiplication. According to Denis, the result of multiplication (product) of strings 𝑠
 of length 𝑚
 and 𝑡
 is a string 𝑡+𝑠1+𝑡+𝑠2+…+𝑡+𝑠𝑚+𝑡
, where 𝑠𝑖
 denotes the 𝑖
-th symbol of the string 𝑠
, and "+" denotes string concatenation. For example, the product of strings "abc" and "de" is a string "deadebdecde", while the product of the strings "ab" and "z" is a string "zazbz". Note, that unlike the numbers multiplication, the product of strings 𝑠
 and 𝑡
 is not necessarily equal to product of 𝑡
 and 𝑠
.

Roman was jealous of Denis, since he invented such a cool operation, and hence decided to invent something string-related too. Since Roman is beauty-lover, he decided to define the beauty of the string as the length of the longest substring, consisting of only one letter. For example, the beauty of the string "xayyaaabca" is equal to 3
, since there is a substring "aaa", while the beauty of the string "qwerqwer" is equal to 1
, since all neighboring symbols in it are different.

In order to entertain Roman, Denis wrote down 𝑛
 strings 𝑝1,𝑝2,𝑝3,…,𝑝𝑛
 on the paper and asked him to calculate the beauty of the string (…(((𝑝1⋅𝑝2)⋅𝑝3)⋅…)⋅𝑝𝑛
, where 𝑠⋅𝑡
 denotes a multiplication of strings 𝑠
 and 𝑡
. Roman hasn't fully realized how Denis's multiplication works, so he asked you for a help. Denis knows, that Roman is very impressionable, he guarantees, that the beauty of the resulting string is at most 109
.


### ideas
1. s * t => 把t用s的每个位置分割后，再连接在一起
2. 假设原来t的beauty是tx, 不会变差,
3. 比如 t = aaa, s = bb, 那么 s * t => aaabaaabaaa
4. 重要的是t的前后缀
5. 如果t的最后一个字符是x,且s[?] = x, 那么beauty至少加1，如果前缀也一样
6. 那么再加上前缀的长度
7. 且如果t不是全部一样的字符，那么前后缀是不会变的
8. 但是完蛋了（每个p的前后缀都需要考虑）
9. 感觉最后的结果只和pn有关？
10. 如果pn不是有单子字符组成的，那么首先计算它自己的beauty，然后考虑它的前后缀，是否可以被扩展
11. 只有当pn是单个字符组成的时候，才有难度
12. 就需要计算出，前面的乘法，能够出现多少个连接在一起的，和pn字符相同的字符串
13. 