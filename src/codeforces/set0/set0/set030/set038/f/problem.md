Once Petya and Vasya invented a new game and called it "Smart Boy". They located a certain set of words — the dictionary — for the game. It is admissible for the dictionary to contain similar words.

The rules of the game are as follows: first the first player chooses any letter (a word as long as 1) from any word from the dictionary and writes it down on a piece of paper. The second player adds some other letter to this one's initial or final position, thus making a word as long as 2, then it's the first player's turn again, he adds a letter in the beginning or in the end thus making a word as long as 3 and so on. But the player mustn't break one condition: the newly created word must be a substring of a word from a dictionary. The player who can't add a letter to the current word without breaking the condition loses.

Also if by the end of a turn a certain string s is written on paper, then the player, whose turn it just has been, gets a number of points according to the formula:


where

 is a sequence number of symbol c in Latin alphabet, numbered starting from 1. For example, , and .
 is the number of words from the dictionary where the line s occurs as a substring at least once.
Your task is to learn who will win the game and what the final score will be. Every player plays optimally and most of all tries to win, then — to maximize the number of his points, then — to minimize the number of the points of the opponent.


### ideas
1. 构造一个图？
2. a, b, c 三个串，要找到 b,c的最长公共子串， 然后构建 a -> d -> (b, c)
3. 怎么构造这个图呢～
4. 假设当前得到的是 s1, sum(s1) * max(s1) + len(s1)
5. 然后添加一个x后, s2, sum(s2) * max(s2) + len(s2)
6. 增加的部分, 如果x不是s2的最大值
7. delta = 1 + max(s1) * x
8. 如果x是s2的最大值, y是s1的最大值
9. detail = 1 + (x - y) * sum(s1) +
10. 
