Polycarp has 𝑛
 different binary words. A word called binary if it contains only characters '0' and '1'. For example, these words are binary: "0001", "11", "0" and "0011100".

Polycarp wants to offer his set of 𝑛
 binary words to play a game "words". In this game, players name words and each next word (starting from the second) must start with the last character of the previous word. The first word can be any. For example, these sequence of words can be named during the game: "0101", "1", "10", "00", "00001".

Word reversal is the operation of reversing the order of the characters. For example, the word "0111" after the reversal becomes "1110", the word "11010" after the reversal becomes "01011".

Probably, Polycarp has such a set of words that there is no way to put them in the order correspondent to the game rules. In this situation, he wants to reverse some words from his set so that:

the final set of 𝑛
 words still contains different words (i.e. all words are unique);
there is a way to put all words of the final set of words in the order so that the final sequence of 𝑛
 words is consistent with the game rules.
Polycarp wants to reverse minimal number of words. Please, help him.

## ideas
1. 这里有两个限定因素，一个是从第二个开始，要能头尾连接起来
2. 第二个是，要保证最后的words都是unique的
3. 似乎第二个更难处理
4. 并不需要按顺序连接
5. 所以，猜一下的话，应该是贪心处理
6. 0， 1， 01， 10
7. 0  01 10 01 ... 111
8. 也就是说，可以把所有0的部分放在最前面，把所有1的部分放在最后面
9. 0/1的部分都不需要翻转，增加翻转次数，且可能产生重复的字符
10. 01 和 10 
11. 如果前面有0，后面有1，那么01/10的必须相差1
12. 0...011001...1
13. 如果前面没有0， 10011001...1 可以相等（当然也可以相差1）
14. 如果后面没有1，
15. 考虑既有0，又有1的情况 （cnt(01) + cnt(10)) % 2 = 1 (else no solution)
16. 如果10多于01， 那么最好1在前面（只需要把部分10改成01）
17. 