You are given a rebus of form ? + ? - ? + ? = n, consisting of only question marks, separated by arithmetic operation '+' and '-', equality and positive integer n. The goal is to replace each question mark with some positive integer from 1 to n, such that equality holds.

Input
The only line of the input contains a rebus. It's guaranteed that it contains no more than 100 question marks, integer n is positive and doesn't exceed 1 000 000, all letters and integers are separated by spaces, arithmetic operations are located only between question marks.

Output
The first line of the output should contain "Possible" (without quotes) if rebus has a solution and "Impossible" (without quotes) otherwise.

If the answer exists, the second line should contain any valid rebus with question marks replaced by integers from 1 to n. Follow the format given in the samples.


### ideas
1. 所有的数要 <= n
2. 只有+/1
3. 所以可以分成两类，那些正数/负数
4. 正数和负数的和 = n
5. 然后在正数和负数中分配