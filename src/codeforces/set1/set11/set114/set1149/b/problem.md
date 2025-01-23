During the archaeological research in the Middle East you found the traces of three ancient religions: First religion, Second religion and Third religion. You compiled the information on the evolution of each of these beliefs, and you now wonder if the followers of each religion could coexist in peace.

The Word of Universe is a long word containing the lowercase English characters only. At each moment of time, each of the religion beliefs could be described by a word consisting of lowercase English characters.

The three religions can coexist in peace if their descriptions form disjoint subsequences of the Word of Universe. More formally, one can paint some of the characters of the Word of Universe in three colors: 1
, 2
, 3
, so that each character is painted in at most one color, and the description of the 𝑖
-th religion can be constructed from the Word of Universe by removing all characters that aren't painted in color 𝑖
.

The religions however evolve. In the beginning, each religion description is empty. Every once in a while, either a character is appended to the end of the description of a single religion, or the last character is dropped from the description. After each change, determine if the religions could coexist in peace.



### ideas
1. 考虑增加一个字符x
2. 如果 religion i 增加了字符x，假设在当前的状态下， word能够匹配到的状态是 dp[i][j][k] (也就是1匹配到了i, 2匹配到了j，3匹配到了k)
3. 然后1增加x，然后如何得到新的 i,j,k的状态呢？
4. dp[i][j][k]表示匹配到了最短的word的位置， 前i/j/k个字符当前字符的长度
5. 如果增加一个新的字符x， 如果在dp[i][j][k]后面有一个x, 那么更新到这个位置（如果没有呢？）
6. 如果没有的话，就是不能共存，且不能更新位置，直到删除到这里为止