Vika came to her favorite cosmetics store "Golden Pear". She noticed that the prices of 𝑛
items have changed since her last visit.

She decided to analyze how much the prices have changed and calculated the difference between the old and new prices for
each of the 𝑛
items.

Vika enjoyed calculating the price differences and decided to continue this process.

Let the old prices be represented as an array of non-negative integers 𝑎
, and the new prices as an array of non-negative integers 𝑏
. Both arrays have the same length 𝑛
.

In one operation, Vika constructs a new array 𝑐
according to the following principle: 𝑐𝑖=|𝑎𝑖−𝑏𝑖|
. Then, array 𝑐
renamed into array 𝑏
, and array 𝑏
renamed into array 𝑎
at the same time, after which Vika repeats the operation with them.

For example, if 𝑎=[1,2,3,4,5,6,7]
; 𝑏=[7,6,5,4,3,2,1]
, then 𝑐=[6,4,2,0,2,4,6]
. Then, 𝑎=[7,6,5,4,3,2,1]
; 𝑏=[6,4,2,0,2,4,6]
.

Vika decided to call a pair of arrays 𝑎
, 𝑏
dull if after some number of such operations all elements of array 𝑎
become zeros.

Output "YES" if the original pair of arrays is dull, and "NO" otherwise.

### thoughts

1. 对于每个i，单独判断
2. 对于数字a,b, 如果都是正数 a % b = 0 or b % a = 0
3. 比如 (12, 4) => (4, 8) => (8, 4) => (4, 4) => (4, 0) => (0, 4)
4. 如果都是负数呢？两次操作后，就变成了正数
5. 一个正数，一个负数，也会很快变成正数
6. 但是考虑（12， 11）=> (11, 1) => (1, 10) 这样也是可以的
7. 所有的数，都是可以变成0的，但变成0的次数不一样