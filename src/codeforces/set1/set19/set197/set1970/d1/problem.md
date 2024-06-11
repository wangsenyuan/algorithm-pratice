Professor Vector is preparing to teach her Arithmancy class. She needs to prepare 𝑛
 distinct magic words for the class. Each magic word is a string consisting of characters X and O. A spell is a string created by concatenating two magic words together. The power of a spell is equal to the number of its different non-empty substrings. For example, the power of the spell XOXO is equal to 7, because it has 7 different substrings: X, O, XO, OX, XOX, OXO and XOXO.

Each student will create their own spell by concatenating two magic words. Since the students are not very good at magic yet, they will choose each of the two words independently and uniformly at random from the 𝑛
 words provided by Professor Vector. It is therefore also possible that the two words a student chooses are the same. Each student will then compute the power of their spell, and tell it to Professor Vector. In order to check their work, and of course to impress the students, Professor Vector needs to find out which two magic words and in which order were concatenated by each student.

Your program needs to perform the role of Professor Vector: first, create 𝑛
 distinct magic words, and then handle multiple requests where it is given the spell power and needs to determine the indices of the two magic words, in the correct order, that were used to create the corresponding spell.

Interaction
This is an interactive problem.

First, your program should read a single integer 𝑛
 (1≤𝑛≤3
), the number of magic words to prepare. Then, it should print 𝑛
 magic words it has created, one per line. The magic words must be distinct, each magic word must have at least 1 and at most 30⋅𝑛
 characters, and each character must be either X or O. We will denote the 𝑖
-th magic word you printed as 𝑤𝑖
 (1≤𝑖≤𝑛
).

Then, your program should read a single integer 𝑞
 (1≤𝑞≤1000
), the number of students in the class. Then, it should repeat the following process 𝑞
 times, one per student.

For the 𝑗
-th student, it should first read a single integer 𝑝𝑗
, the power of their spell. It is guaranteed that this number is computed by choosing two indices 𝑢𝑗
 and 𝑣𝑗
 independently and uniformly at random between 1 and 𝑛
 inclusive, concatenating 𝑤𝑢𝑗
 and 𝑤𝑣𝑗
, and finding the number of different non-empty substrings of the resulting string. Then, your program must print the numbers 𝑢𝑗
 and 𝑣𝑗
, in this order (1≤𝑢𝑗,𝑣𝑗≤𝑛
).

Note that it is not enough to find any two magic words that concatenate into a spell with the given power. You must find the exact words used by the student in the exact order.

Remember to flush the output stream after printing all magic words and after printing 𝑢𝑗
 and 𝑣𝑗
 for each student.


 ### ideas
 1. x, o, xx
 2.  if n = 1, just x?  xx => 2
 3.  if n = 2, x and o => xo => x, o, xo = 3, ox => o, x, ox = 3 不行
 4.     xx, o => xxxx = 4, xxo (x, xx, xo, xxo) 4, 也不行
 5.     xo, o => xoxo = 7, xoo (x, o, oo, xoo, xo) 5, oo (o oo) 2, oxo (o, ox, xo, oxo, x) 5 （wrong）
 6. 这个题目完全没想法啊
 7.  n = 2， XOXO X, 这个例子可以用来解决 n = 2, 但是规律是什么？
 8.     XOXOX (x, xo, xox, oxo, xoxo, oxox, xoxox)
 9.     xoxoxoxo (它包含上面的，且多了几个, 所以可以区分出来)
 10.    xxoxo  (x, xx, xo, ox, xox, oxo, xxox, xoxo, xxoxo)