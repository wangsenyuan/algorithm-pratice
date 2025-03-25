Right now you are to solve a very, very simple problem — to crack the safe. Four positive integers stand one by one on a circle protecting the safe. You know that to unlock this striking safe you have to make all four numbers equal to one. Operations are as follows: you may choose two adjacent numbers and increase both by one; you may choose two adjacent even numbers and divide both by two. Nothing else. Crack the safe!



### ideas
1. 如果相邻的两个数都是偶数，最好的选择是除以2
2. 直到两个数都变成0，或者一个是偶数，一个是奇数
3. 如果两个都是奇数(大于1)，先+1，再除以2
4. 如果一个奇数，一个偶数
5. 所以可以看到，在所有的数变成1以前，始终是可以操作的