Dasha logged into the system and began to solve problems. One of them is as follows:

Given two sequences a and b of length n each you need to write a sequence c of length n, the i-th element of which is calculated as follows: ci = bi - ai.

About sequences a and b we know that their elements are in the range from l to r. More formally, elements satisfy the following conditions: l ≤ ai ≤ r and l ≤ bi ≤ r. About sequence c we know that all its elements are distinct.


Dasha wrote a solution to that problem quickly, but checking her work on the standard test was not so easy. Due to an error in the test system only the sequence a and the compressed sequence of the sequence c were known from that test.

Let's give the definition to a compressed sequence. A compressed sequence of sequence c of length n is a sequence p of length n, so that pi equals to the number of integers which are less than or equal to ci in the sequence c. For example, for the sequence c = [250, 200, 300, 100, 50] the compressed sequence will be p = [4, 3, 5, 2, 1]. Pay attention that in c all integers are distinct. Consequently, the compressed sequence contains all integers from 1 to n inclusively.

Help Dasha to find any sequence b for which the calculated compressed sequence of sequence c is correct.

### ideas
1. 每个C有一个取值范围，[a[i] + l, a[i] + r], 且当前的C要大于前一个C（按照P排列后）
2. 假设第一个取值= a[pos[1]] + l, 它肯定在范围内 
3. 第二个如果可以取值 a[pos[2]] + l, 当然很好（它对后面的数没有那么大的压力）
4. 否则的话，就应该取值a[pos[2]] + l + 1, 2, 3... a[pos[2]] + r
5. 其实就是看max(a[pos[2]] + l, c[1] + 1)