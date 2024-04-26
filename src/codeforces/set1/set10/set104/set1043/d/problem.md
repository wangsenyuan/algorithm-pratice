Acingel is a small town. There was only one doctor here — Miss Ada. She was very friendly and nobody has ever said
something bad about her, so who could've expected that Ada will be found dead in her house? Mr Gawry, world-famous
detective, is appointed to find the criminal. He asked 𝑚
neighbours of Ada about clients who have visited her in that unlucky day. Let's number the clients from 1
to 𝑛
. Each neighbour's testimony is a permutation of these numbers, which describes the order in which clients have been
seen by the asked neighbour.

However, some facts are very suspicious – how it is that, according to some of given permutations, some client has been
seen in the morning, while in others he has been seen in the evening? "In the morning some of neighbours must have been
sleeping!" — thinks Gawry — "and in the evening there's been too dark to see somebody's face...". Now he wants to delete
some prefix and some suffix (both prefix and suffix can be empty) in each permutation, so that they'll be non-empty and
equal to each other after that — some of the potential criminals may disappear, but the testimony won't stand in
contradiction to each other.

In how many ways he can do it? Two ways are called different if the remaining common part is different.

### ideas

1. 先理解一下这个题目，假设剩余了k个人， 那么这个k个人在m组中出现的顺序应该是一样的
2. 假设k是最长的那个，那么 1...k-1也满足条件
3. 所以以第一组为标尺，对于l应该尽可能的延展r，直到不可以，然后从下一个r开始
4. 