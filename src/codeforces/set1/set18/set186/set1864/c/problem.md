You are given an integer 𝑥
. Your task is to reduce 𝑥
to 1
.

To do that, you can do the following operation:

select a divisor 𝑑
of 𝑥
, then change 𝑥
to 𝑥−𝑑
, i.e. reduce 𝑥
by 𝑑
. (We say that 𝑑
is a divisor of 𝑥
if 𝑑
is an positive integer and there exists an integer 𝑞
such that 𝑥=𝑑⋅𝑞
.)
There is an additional constraint: you cannot select the same value of 𝑑
more than twice.

For example, for 𝑥=5
, the following scheme is invalid because 1
is selected more than twice: 5−→−−14−→−−13−→−−12−→−−11
. The following scheme is however a valid one: 5−→−−14−→−−22−→−−11
.

Output any scheme which reduces 𝑥
to 1
with at most 1000
operations. It can be proved that such a scheme always exists.

### thoughts

1. 不考虑次数限制的情况下, 要怎么处理呢？
2. 因为所有的都是成立的，所以，先使用最大的divisor来处理， 这样得到的肯定不能被同一个divisor继续处理
3. 假设可以，那么表示x % (div * div) = 0,然后到某个数无法处理时，最大的divisor是1（不包括自己）
4. 这个数是质数，因为不会出现两个连续的质数，（除了2，3) 所以，不会出现超过两次1
5. 搞错了～
6. 是同一个数，不能出现超过2次，而不是连续两次
7. 如果是减少到0，将非常容易
8. got it, 先减小到最高位去，然后，从最高位依次降下来即可