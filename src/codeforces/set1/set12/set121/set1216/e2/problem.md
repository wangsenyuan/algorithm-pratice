You are given an infinite sequence of form "112123123412345…
" which consist of blocks of all consecutive positive integers written one after another. The first block consists of all numbers from 1
 to 1
, the second one — from 1
 to 2
, the third one — from 1
 to 3
, …
, the 𝑖
-th block consists of all numbers from 1
 to 𝑖
.

So the first 56
 elements of the sequence are "11212312341234512345612345671234567812345678912345678910". Elements of the sequence are numbered from one. For example, the 1
-st element of the sequence is 1
, the 3
-rd element of the sequence is 2
, the 20
-th element of the sequence is 5
, the 38
-th element is 2
, the 56
-th element of the sequence is 0
.

Your task is to answer 𝑞
 independent queries. In the 𝑖
-th query you are given one integer 𝑘𝑖
. Calculate the digit at the position 𝑘𝑖
 of the sequence.

 ### ideas
 1. 只能二分了
 2. 先检查 <= num 的共有多少个digits
 3. 1有num个，2有num-1个。。。。9有num-9个， 10有 2 *（num - 10）个
 4. f(num) = f(num - 1) + digitsof(num)