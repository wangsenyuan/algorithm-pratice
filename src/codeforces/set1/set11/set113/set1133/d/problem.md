You are given two arrays 𝑎
 and 𝑏
, each contains 𝑛
 integers.

You want to create a new array 𝑐
 as follows: choose some real (i.e. not necessarily integer) number 𝑑
, and then for every 𝑖∈[1,𝑛]
 let 𝑐𝑖:=𝑑⋅𝑎𝑖+𝑏𝑖
.

Your goal is to maximize the number of zeroes in array 𝑐
. What is the largest possible answer, if you choose 𝑑
 optimally?

 ### ideas
 1. c = d * a + b = 0
 2. d * a = -b
 3. d = - (b / a)
 4. 如果a = 0， 那么要求b = 0
 5. 如果b = 0, 要么要求 a * d = 0
 6. 如果a != 0, b != 0, 那么要求 d * a = -b