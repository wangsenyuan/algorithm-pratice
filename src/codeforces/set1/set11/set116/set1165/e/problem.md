You are given two arrays 𝑎
 and 𝑏
, both of length 𝑛
.

Let's define a function 𝑓(𝑙,𝑟)=∑𝑙≤𝑖≤𝑟𝑎𝑖⋅𝑏𝑖
.

Your task is to reorder the elements (choose an arbitrary order of elements) of the array 𝑏
 to minimize the value of ∑1≤𝑙≤𝑟≤𝑛𝑓(𝑙,𝑟)
. Since the answer can be very large, you have to print it modulo 998244353
. Note that you should minimize the answer but not its remainder.