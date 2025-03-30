In anticipation of a duel with his old friend Fernan, Edmond is preparing an energy drink called "Mishkin Energizer". The drink consists of a string 𝑠
 of length 𝑛
, made up only of the characters L, I, and T, which correspond to the content of three different substances in the drink.

We call the drink balanced if it contains an equal number of all substances. To boost his aura and ensure victory in the duel, Edmond must make the initial string balanced by applying the following operation:

Choose an index 𝑖
 such that 𝑠𝑖≠𝑠𝑖+1
 (where 𝑖+1
 must not exceed the current size of the string).
Insert a character 𝑥
, either L, I, or T, between them such that 𝑥≠𝑠𝑖
 and 𝑥≠𝑠𝑖+1
.
Help Edmond make the drink balanced and win the duel by performing no more than 2𝐧
 operations. If there are multiple solutions, any one of them can be output. If it is impossible, you must report this.

 ### ideas
 1. LIT  (s[i] != s[i+1])
 2. 当前的为n，那么需要的结果 = 当前最多的字符 * 3 = 3 * m
 3. 如果 3 * m - n > 2 * n
 4. 3 * m > 3 * n
 5. 只有一个字符的时候，没有答案
 6. 否则肯定是有答案的
 7. S[i] != S[i+1] LT => LIT => 如果是L不够， LILT,