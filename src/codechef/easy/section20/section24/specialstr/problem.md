Consider the string of infinite length obtained by repeating abcdefghijklmnopqrstuvwxyz infinitely many times. That is, it looks like abc...xyzabc...xyzabc...

Let string B be the prefix of this string of length M. For example,

if M=6, B= abcdef
if M=29, B= abcdefghijklmnopqrstuvwxyzabc
Vedansh calls a string X special if it contains B as a substring. Formally, X is special if for some 1≤L≤R≤|X|, XL…R=B.

Vedansh has string S of length N (consisting of lowercase Latin letters only). He wants to convert S into a special string. To do so he can perform the following operation:

Pick an i such that 1≤i≤N and delete Si. The remaining parts of S are concatenated.
Since Vedansh is lazy, he wants to do this in the minimum number of operations or determine if S cannot be converted to a special string. Help Vedansh in doing so.

