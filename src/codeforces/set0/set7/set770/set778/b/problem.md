Bob recently read about bitwise operations used in computers: AND, OR and XOR. He have studied their properties and
invented a new game.

Initially, Bob chooses integer m, bit depth of the game, which means that all numbers in the game will consist of m
bits. Then he asks Peter to choose some m-bit number. After that, Bob computes the values of n variables. Each variable
is assigned either a constant m-bit number or result of bitwise operation. Operands of the operation may be either
variables defined before, or the number, chosen by Peter. After that, Peter's score equals to the sum of all variable
values.

Bob wants to know, what number Peter needs to choose to get the minimum possible score, and what number he needs to
choose to get the maximum possible score. In both cases, if there are several ways to get the same score, find the
minimum number, which he can choose.

Input
The first line contains two integers n and m, the number of variables and bit depth, respectively (1 ≤ n ≤ 5000; 1 ≤ m ≤
1000).

The following n lines contain descriptions of the variables. Each line describes exactly one variable. Description has
the following format: name of a new variable, space, sign ":=", space, followed by one of:

Binary number of exactly m bits.
The first operand, space, bitwise operation ("AND", "OR" or "XOR"), space, the second operand. Each operand is either
the name of variable defined before or symbol '?', indicating the number chosen by Peter.
Variable names are strings consisting of lowercase Latin letters with length at most 10. All variable names are
different.

Output
In the first line output the minimum number that should be chosen by Peter, to make the sum of all variable values
minimum possible, in the second line output the minimum number that should be chosen by Peter, to make the sum of all
variable values maximum possible. Both numbers should be printed as m-bit binary numbers.

### ideas

1. bit by bit from the least to most
2. in order to get min value, need to get least 1 in this bit
3. max from most 1