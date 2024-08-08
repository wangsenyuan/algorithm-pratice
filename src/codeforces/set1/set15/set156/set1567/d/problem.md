On the board, Bob wrote 𝑛
 positive integers in base 10
 with sum 𝑠
 (i. e. in decimal numeral system). Alice sees the board, but accidentally interprets the numbers on the board as base-11
 integers and adds them up (in base 11
).

What numbers should Bob write on the board, so Alice's sum is as large as possible?

Input
The input consists of multiple test cases. The first line contains an integer 𝑡
 (1≤𝑡≤100
) — the number of test cases. The description of the test cases follows.

The only line of each test case contains two integers 𝑠
 and 𝑛
 (1≤𝑠≤109
; 1≤𝑛≤min(100,𝑠)
) — the sum and amount of numbers on the board, respectively. Numbers 𝑠
 and 𝑛
 are given in decimal notation (base 10
).

Output
For each test case, output 𝑛
 positive integers — the numbers Bob should write on the board, so Alice's sum is as large as possible. If there are multiple answers, print any of them.