Petya learned a new programming language CALPAS. A program in this language always takes one non-negative integer and
returns one non-negative integer as well.

In the language, there are only three commands: apply a bitwise operation AND, OR or XOR with a given constant to the
current integer. A program can contain an arbitrary sequence of these operations with arbitrary constants from 0 to

1023. When the program is run, all operations are applied (in the given order) to the argument and in the end the result
      integer is returned.

Petya wrote a program in this language, but it turned out to be too long. Write a program in CALPAS that does the same
thing as the Petya's program, and consists of no more than 5 lines. Your program should return the same integer as
Petya's program for all arguments from 0 to 1023.

### ideas

1. 很有意思的一个题目
2. 其中有些lines应该是多于的（甚至大部分是多余的）
3. 每一位单独处理
4. 如果它依次处理，最后的结果始终是0，那么就是 | 1, & 0, 如果始终是1，就 | 1
5. 如果和原来的一致， 那么就 & 1, 如果和原来的相反就 xor 1
6. 但是问题是如何保证只有5条指令呢？
7. 似乎三条指令就够了，分别是 |, &, ^, 然后根据上面的情况，设置对应的位即可
8. 