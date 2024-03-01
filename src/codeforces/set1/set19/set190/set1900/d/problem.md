Let 𝑎
, 𝑏
, and 𝑐
be integers. We define function 𝑓(𝑎,𝑏,𝑐)
as follows:

Order the numbers 𝑎
, 𝑏
, 𝑐
in such a way that 𝑎≤𝑏≤𝑐
. Then return gcd(𝑎,𝑏)
, where gcd(𝑎,𝑏)
denotes the greatest common divisor (GCD) of integers 𝑎
and 𝑏
.

So basically, we take the gcd
of the 2
smaller values and ignore the biggest one.

You are given an array 𝑎
of 𝑛
elements. Compute the sum of 𝑓(𝑎𝑖,𝑎𝑗,𝑎𝑘)
for each 𝑖
, 𝑗
, 𝑘
, such that 1≤𝑖<𝑗<𝑘≤𝑛
.

More formally, compute
∑𝑖=1𝑛∑𝑗=𝑖+1𝑛∑𝑘=𝑗+1𝑛𝑓(𝑎𝑖,𝑎𝑗,𝑎𝑘).
Input