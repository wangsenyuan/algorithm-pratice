Furik loves writing all sorts of problems, especially such that he can't solve himself. You've got one of his problems, the one Furik gave to Rubik. And Rubik asks you to solve it.

There is integer n and array a, consisting of ten integers, indexed by numbers from 0 to 9. Your task is to count the number of positive integers with the following properties:

the number's length does not exceed n;
the number doesn't have leading zeroes;
digit i (0 ≤ i ≤ 9) occurs in the number at least a[i] times.
Input
The first line contains integer n (1 ≤ n ≤ 100). The next line contains 10 integers a[0], a[1], ..., a[9] (0 ≤ a[i] ≤ 100) — elements of array a. The numbers are separated by spaces.


### ideas
1. 长度为n, 但是i至少出现多少次