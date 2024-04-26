At a break Vanya came to the class and saw an array of 𝑛
𝑘
-bit integers 𝑎1,𝑎2,…,𝑎𝑛
on the board. An integer 𝑥
is called a 𝑘
-bit integer if 0≤𝑥≤2𝑘−1
.

Of course, Vanya was not able to resist and started changing the numbers written on the board. To ensure that no one
will note anything, Vanya allowed himself to make only one type of changes: choose an index of the array 𝑖
(1≤𝑖≤𝑛
) and replace the number 𝑎𝑖
with the number 𝑎𝑖⎯⎯⎯⎯⎯
. We define 𝑥⎯⎯⎯
for a 𝑘
-bit integer 𝑥
as the 𝑘
-bit integer such that all its 𝑘
bits differ from the corresponding bits of 𝑥
.

Vanya does not like the number 0
. Therefore, he likes such segments [𝑙,𝑟]
(1≤𝑙≤𝑟≤𝑛
) such that 𝑎𝑙⊕𝑎𝑙+1⊕…⊕𝑎𝑟≠0
, where ⊕
denotes the bitwise XOR operation. Determine the maximum number of segments he likes Vanya can get applying zero or more
operations described above.

### ideas

1. 先不考虑reverse的情况， 什么情况下，一段的xor = 0？
2. 当这段中每位的出现次数是偶数的时候
3. 好难呐，完全没有想法～
4. 每个数，都可以保持不变，或者取反。对自己有影响，对别人也有影响
5. 比如更改了a[i]后，那么所有包含a[i]的序列都会被影响，但也不是它单独就可以，其他元素也会跟着变化
6. 我们考虑k = 1 的情况，假设要最大化非0的子序列，应怎么处理？
7. 假设1的个数是x，0的个数是y，那么结果是否= x * y?
8. 如果全部是1， ans = odd[x] * (odd[x] + 1) / 2 + even[x] * (even[x] + 1) / 2 最大
9. 111111

### solution

We need to maximize the number of segments with XOR-sum of its elements not equal to 0
. Let's note that the total number of segments is (𝑛+12)=𝑛⋅(𝑛+1)/2
so we need to minimize the number of segments with XOR-sum of its elements equal to 0
. Let's call such segments bad.

Consider the prefix XOR sums of array 𝑎
. Let's define 𝑝𝑖=𝑎1⊕⋯⊕𝑎𝑖
for any 0≤𝑖≤𝑛
. Note that the segment [𝑙,𝑟]
is bad if 𝑝𝑟𝑙−1=𝑝𝑟𝑟
.

How does an array 𝑝
changes in a single operation with an array 𝑎
? Let's note that 𝑥⎯⎯⎯=𝑥⊕(2𝑘−1)
. If we do operation with the element 𝑖
then in array 𝑝
all elements 𝑝𝑗
(𝑗≥𝑖
) will change to 𝑝𝑗⊕(2𝑘−1)
or 𝑝𝑗⎯⎯⎯⎯⎯
. So after all operations, element of 𝑝
will be either equal to itself before operations or will be equal to itself before operations but with all changed bits.
But 𝑝0=0
still.

On the other hand. Suppose that there are some changes of this kind in the array 𝑝
. Let's prove that we could make some operations in the array 𝑎
so that 𝑝
will be changed and became exactly the same. To prove this, note that 𝑎𝑖=𝑝𝑖−1⊕𝑝𝑖
, so all values of 𝑎𝑖
have either not changed or have changed to numbers with all changed bits. So let's just assume that we are changing the
array 𝑝
(all its elements except the element with index 0
), not the array 𝑎
.

Let's note that two numbers 𝑥
and 𝑦
can become equal after all operations with them if they are equal or 𝑥=𝑦⎯⎯⎯
. So we can replace 𝑝𝑖
with 𝑚𝑖𝑛(𝑝𝑖,𝑝𝑖⎯⎯⎯⎯⎯)
for any 𝑖
in the array 𝑝
. The answer will be the same if we do this. But it will be impossible to get equal numbers from different numbers in
the new array.

Let's calculate for each number 𝑐
the number of its occurrences in the array 𝑝
and denote it for 𝑘
. Some of this 𝑘
numbers we can leave equal to 𝑐
, and some of them we change to 𝑐⎯⎯⎯
. Let's change 0≤𝑎≤𝑘
numbers. Then we want to minimize the number of pairs of equal numbers, which is (𝑘−𝑎+12)+(𝑎+12)
. To do this, we need to change and not change the approximately equal number of numbers, that is, the minimum is
reached at 𝑎=⌊𝑘/2⌋
. It can be proved simply. So we should divide the numbers this way for any 𝑐
and sum the answers.

For each number 𝑐
, the number of its occurrences in the array can be calculated using the sorting algorithm or the std::map structure
into the C++ language.

Time complexity: 𝑂(𝑛⋅𝑙𝑜𝑔(𝑛))
.