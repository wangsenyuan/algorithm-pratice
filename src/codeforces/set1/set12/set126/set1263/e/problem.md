The development of a text editor is a hard problem. You need to implement an extra module for brackets coloring in text.

Your editor consists of a line with infinite length and cursor, which points to the current character. Please note that it points to only one of the characters (and not between a pair of characters). Thus, it points to an index character. The user can move the cursor left or right one position. If the cursor is already at the first (leftmost) position, then it does not move left.

Initially, the cursor is in the first (leftmost) character.

Also, the user can write a letter or brackets (either (, or )) to the position that the cursor is currently pointing at. A new character always overwrites the old value at that position.

Your editor must check, whether the current line is the correct text. Text is correct if the brackets in them form the correct bracket sequence.

Formally, correct text (CT) must satisfy the following rules:

any line without brackets is CT (the line can contain whitespaces);
If the first character of the string — is (, the last — is ), and all the rest form a CT, then the whole line is a CT;
two consecutively written CT is also CT.
Examples of correct texts: hello(codeforces), round, ((i)(write))edi(tor)s, ( me). Examples of incorrect texts: hello)oops(, round), ((me).

The user uses special commands to work with your editor. Each command has its symbol, which must be written to execute this command.

The correspondence of commands and characters is as follows:

L — move the cursor one character to the left (remains in place if it already points to the first character);
R — move the cursor one character to the right;
any lowercase Latin letter or bracket (( or )) — write the entered character to the position where the cursor is now.
For a complete understanding, take a look at the first example and its illustrations in the note below.

You are given a string containing the characters that the user entered. For the brackets coloring module's work, after each command you need to:

check if the current text in the editor is a correct text;
if it is, print the least number of colors that required, to color all brackets.
If two pairs of brackets are nested (the first in the second or vice versa), then these pairs of brackets should be painted in different colors. If two pairs of brackets are not nested, then they can be painted in different or the same colors. For example, for the bracket sequence ()(())()() the least number of colors is 2
, and for the bracket sequence (()(()())())(()) — is 3
.

Write a program that prints the minimal number of colors after processing each command.


### ideas
1. 如何判断现在的text是CT？
2. 假设左括号为+1，右括号为-1， 那么右括号形成的前缀和，满足条件，最后的sum = 0， 且没有<0的情况， CT
3. 假设把一个普通字符变成左括号（相当于从该处，后面所有的数字+1)
4. segment tree with lazy update