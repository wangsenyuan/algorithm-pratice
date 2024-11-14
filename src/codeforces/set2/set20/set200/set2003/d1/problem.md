One day, Turtle was playing with 𝑛
 sequences. Let the length of the 𝑖
-th sequence be 𝑙𝑖
. Then the 𝑖
-th sequence was 𝑎𝑖,1,𝑎𝑖,2,…,𝑎𝑖,𝑙𝑖
.

Piggy gave Turtle a problem to solve when Turtle was playing. The statement of the problem was:

There was a non-negative integer 𝑥
 at first. Turtle would perform an arbitrary number (possibly zero) of operations on the integer.
In each operation, Turtle could choose an integer 𝑖
 such that 1≤𝑖≤𝑛
, and set 𝑥
 to mex†(𝑥,𝑎𝑖,1,𝑎𝑖,2,…,𝑎𝑖,𝑙𝑖)
.
Turtle was asked to find the answer, which was the maximum value of 𝑥
 after performing an arbitrary number of operations.
Turtle solved the above problem without difficulty. He defined 𝑓(𝑘)
 as the answer to the above problem when the initial value of 𝑥
 was 𝑘
.

Then Piggy gave Turtle a non-negative integer 𝑚
 and asked Turtle to find the value of ∑𝑖=0𝑚𝑓(𝑖)
 (i.e., the value of 𝑓(0)+𝑓(1)+…+𝑓(𝑚)
). Unfortunately, he couldn't solve this problem. Please help him!

†mex(𝑐1,𝑐2,…,𝑐𝑘)
 is defined as the smallest non-negative integer 𝑥
 which does not occur in the sequence 𝑐
. For example, mex(2,2,0,3)
 is 1
, mex(1,2)
 is 0
.



### ideas
1. 这个题目看着就很难，咋只有1500 呢
2. 对于a[i], 第i个序列, 它的mex可以提前算出来，且可以算出第二个mex(也就是填上第一个mex后，所能得到的)
3. 然后可以把这些序列组成一个树,如果mex2(a[1]) = 3, 那么如果mex1(a[i]) = 3, 那么可以把它们接起来
4. 从而可以计算出，加入某个x后，能够得到的最大mex， 这个就是f(x) 
5. 然后当我们提供x的时候，就可以把这个值放入即可
6. 其他的是x的值是?
7. 假设一个序列的mex = 4, 如果选择一个x = 4， 那么可以得到它第二个mex（这个没有问题）
8. 但是如果将x> 4, 那么，第一次操作后就可以得到4，然后就可以得到同样第二个mex
9. 也就是说满足 f[i] >= f[i-1]的这种情况