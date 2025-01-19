A frog is initially at position 0
 on the number line. The frog has two positive integers 𝑎
 and 𝑏
. From a position 𝑘
, it can either jump to position 𝑘+𝑎
 or 𝑘−𝑏
.

Let 𝑓(𝑥)
 be the number of distinct integers the frog can reach if it never jumps on an integer outside the interval [0,𝑥]
. The frog doesn't need to visit all these integers in one trip, that is, an integer is counted if the frog can somehow reach it if it starts from 0
.

Given an integer 𝑚
, find ∑𝑚𝑖=0𝑓(𝑖)
. That is, find the sum of all 𝑓(𝑖)
 for 𝑖
 from 0
 to 𝑚
.



### ideas
1. m * a - n * b = x