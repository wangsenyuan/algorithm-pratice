Today, Cat and Fox found an array 𝑎
 consisting of 𝑛
 non-negative integers.

Define the loneliness of 𝑎
 as the smallest positive integer 𝑘
 (1≤𝑘≤𝑛
) such that for any two positive integers 𝑖
 and 𝑗
 (1≤𝑖,𝑗≤𝑛−𝑘+1
), the following holds:
𝑎𝑖|𝑎𝑖+1|…|𝑎𝑖+𝑘−1=𝑎𝑗|𝑎𝑗+1|…|𝑎𝑗+𝑘−1,
where 𝑥|𝑦
 denotes the bitwise OR of 𝑥
 and 𝑦
. In other words, for every 𝑘
 consecutive elements, their bitwise OR should be the same. Note that the loneliness of 𝑎
 is well-defined, because for 𝑘=𝑛
 the condition is satisfied.

Cat and Fox want to know how lonely the array 𝑎
 is. Help them calculate the loneliness of the found array.

### ideas
1. 如果 k = x, 成立， k = x + 1 成立吗？
2. 似乎是成立的
3. 考虑第一个和第二个序列， 
4.     a[1] | ... | a[x] = a[2] | ... | a[x+1]
5.     a[2] | ... | a[x+1] = a[3] | ... | a[x+2] 
6.  上下两边相与，左右还是相等的
7.  所以可以二分