You are given an array 𝑎
 of length 𝑛
, a positive integer 𝑚
, and a string of commands of length 𝑛
. Each command is either the character 'L' or the character 'R'.

Process all 𝑛
 commands in the order they are written in the string 𝑠
. Processing a command is done as follows:

First, output the remainder of the product of all elements of the array 𝑎
 when divided by 𝑚
.
Then, if the command is 'L', remove the leftmost element from the array 𝑎
, if the command is 'R', remove the rightmost element from the array 𝑎
.
Note that after each move, the length of the array 𝑎
 decreases by 1
, and after processing all commands, it will be empty.

Write a program that will process all commands in the order they are written in the string 𝑠
 (from left to right).