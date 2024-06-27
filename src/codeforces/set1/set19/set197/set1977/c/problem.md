Nikita is a student passionate about number theory and algorithms. He faces an interesting problem related to an array of numbers.

Suppose Nikita has an array of integers 𝑎
 of length 𝑛
. He will call a subsequence†
 of the array special if its least common multiple (LCM) is not contained in 𝑎
. The LCM of an empty subsequence is equal to 0
.

Nikita wonders: what is the length of the longest special subsequence of 𝑎
? Help him answer this question!

†
 A sequence 𝑏
 is a subsequence of 𝑎
 if 𝑏
 can be obtained from 𝑎
 by the deletion of several (possibly, zero or all) elements, without changing the order of the remaining elements. For example, [5,2,3]
 is a subsequence of [1,5,7,8,2,4,3]
.

### ideas
1. lcm和顺序没有关系，这里可以先排序
2. 如果一个set的lcm在a中存在，假设这个lcm是a[i]，那么这个set就是a[i]的因数集合,
3. 如果在这个set中添加任何一个非a[i]因数的数，那么a[i]就不再是它们的lcm
4. 但此时不能保证新的set的最大因数不在后面存在
5. 如果在a中distinct的数字，超过30个，那么整个arr肯定是满足条件的
6. 所以，这个数组的长度不会超过30
7. 然后将那些小的数，算到它们的大的倍数里面
8. 但是可能会存在（2， 3， 6）这样的情况，那么6应该算成是3
9. 然后检查最大的那些数的lcm