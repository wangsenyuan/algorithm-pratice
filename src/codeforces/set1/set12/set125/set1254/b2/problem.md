Christmas is coming, and our protagonist, Bob, is preparing a spectacular present for his long-time best friend Alice. This year, he decides to prepare 𝑛
 boxes of chocolate, numbered from 1
 to 𝑛
. Initially, the 𝑖
-th box contains 𝑎𝑖
 chocolate pieces.

Since Bob is a typical nice guy, he will not send Alice 𝑛
 empty boxes. In other words, at least one of 𝑎1,𝑎2,…,𝑎𝑛
 is positive. Since Alice dislikes coprime sets, she will be happy only if there exists some integer 𝑘>1
 such that the number of pieces in each box is divisible by 𝑘
. Note that Alice won't mind if there exists some empty boxes.

Charlie, Alice's boyfriend, also is Bob's second best friend, so he decides to help Bob by rearranging the chocolate pieces. In one second, Charlie can pick up a piece in box 𝑖
 and put it into either box 𝑖−1
 or box 𝑖+1
 (if such boxes exist). Of course, he wants to help his friend as quickly as possible. Therefore, he asks you to calculate the minimum number of seconds he would need to make Alice happy.

### ideas
1. 这里a[i] 大于1， 所以会存在，把其中的部分，迁移出来，或者加入部分，产生k的倍数的情况
2. 