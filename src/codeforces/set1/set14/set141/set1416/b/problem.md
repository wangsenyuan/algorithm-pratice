You are given an array 𝑎
 consisting of 𝑛
 positive integers, numbered from 1
 to 𝑛
. You can perform the following operation no more than 3𝑛
 times:

choose three integers 𝑖
, 𝑗
 and 𝑥
 (1≤𝑖,𝑗≤𝑛
; 0≤𝑥≤109
);
assign 𝑎𝑖:=𝑎𝑖−𝑥⋅𝑖
, 𝑎𝑗:=𝑎𝑗+𝑥⋅𝑖
.
After each operation, all elements of the array should be non-negative.

Can you find a sequence of no more than 3𝑛
 operations after which all elements of the array are equal?


 ### ideas
 1. 最后的结果 a[i] = sum / n
 2. 如果 sum % n != 0 => false
 3. else 始终存在吗？
 4. 是否可以用1作为过渡？
 5. 比如如果 a[1] > avg, 那么就可以把 a[1] - avg 替换到小的那部分上面去
 6. 想到一个方案，就是把所有大于 avg 的部分， a[j] - j * x 换到 a[1]上去
 7. a[j] - j * x > 0 就可以了
 8. 然后，在把i= 1即可。 这样子，只需要2 * n的过程即可