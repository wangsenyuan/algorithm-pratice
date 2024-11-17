You are given two arrays 𝑎
and 𝑏
, both of length 𝑛
. All elements of both arrays are from 0
to 𝑛−1
.

You can reorder elements of the array 𝑏
(if you want, you may leave the order of elements as it is). After that, let array 𝑐
be the array of length 𝑛
, the 𝑖
-th element of this array is 𝑐𝑖=(𝑎𝑖+𝑏𝑖)%𝑛
, where 𝑥%𝑦
is 𝑥
modulo 𝑦
.

Your task is to reorder elements of the array 𝑏
to obtain the lexicographically minimum possible array 𝑐
.

Array 𝑥
of length 𝑛
is lexicographically less than array 𝑦
of length 𝑛
, if there exists such 𝑖
(1≤𝑖≤𝑛
), that 𝑥𝑖<𝑦𝑖
, and for any 𝑗
(1≤𝑗<𝑖
) 𝑥𝑗=𝑦𝑗
.

### ideas

1. c[i] = (a[i] + b[i]) % n
2. 显然希望c[0] = 0 (如果可以的话)
3. 那么b[0] = n - a[i] 