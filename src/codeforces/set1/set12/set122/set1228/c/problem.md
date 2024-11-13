Let's introduce some definitions that will be needed later.

Let 𝑝𝑟𝑖𝑚𝑒(𝑥)
 be the set of prime divisors of 𝑥
. For example, 𝑝𝑟𝑖𝑚𝑒(140)={2,5,7}
, 𝑝𝑟𝑖𝑚𝑒(169)={13}
.

Let 𝑔(𝑥,𝑝)
 be the maximum possible integer 𝑝𝑘
 where 𝑘
 is an integer such that 𝑥
 is divisible by 𝑝𝑘
. For example:

𝑔(45,3)=9
 (45
 is divisible by 32=9
 but not divisible by 33=27
),
𝑔(63,7)=7
 (63
 is divisible by 71=7
 but not divisible by 72=49
).
Let 𝑓(𝑥,𝑦)
 be the product of 𝑔(𝑦,𝑝)
 for all 𝑝
 in 𝑝𝑟𝑖𝑚𝑒(𝑥)
. For example:

𝑓(30,70)=𝑔(70,2)⋅𝑔(70,3)⋅𝑔(70,5)=21⋅30⋅51=10
,
𝑓(525,63)=𝑔(63,3)⋅𝑔(63,5)⋅𝑔(63,7)=32⋅50⋅71=63
.
You have integers 𝑥
 and 𝑛
. Calculate 𝑓(𝑥,1)⋅𝑓(𝑥,2)⋅…⋅𝑓(𝑥,𝑛)mod(109+7)
.

### ideas
1. f(x, y) = f(gcd(x, y), y) = g(y, p1) * g(y, p2) ...
2. x的prime factors可以计算出来，不超过30个
3. 然后对于每个p，需要计算它们的贡献
4. 比如factor 2， 那么在2,4,6,8,10,12,14,17... 都有贡献， 
5. 2, 6, 10, 14, 
6. 需要倒过来算