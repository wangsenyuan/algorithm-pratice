There are n students studying in the 6th grade, in group "B" of a berland secondary school. Every one of them has exactly one friend whom he calls when he has some news. Let us denote the friend of the person number i by g(i). Note that the friendships are not mutual, i.e. g(g(i)) is not necessarily equal to i.

On day i the person numbered as ai learns the news with the rating of bi (bi ≥ 1). He phones the friend immediately and tells it. While he is doing it, the news becomes old and its rating falls a little and becomes equal to bi - 1. The friend does the same thing — he also calls his friend and also tells the news. The friend of the friend gets the news already rated as bi - 2. It all continues until the rating of the news reaches zero as nobody wants to tell the news with zero rating.

More formally, everybody acts like this: if a person x learns the news with a non-zero rating y, he calls his friend g(i) and his friend learns the news with the rating of y - 1 and, if it is possible, continues the process.

Let us note that during a day one and the same person may call his friend and tell him one and the same news with different ratings. Thus, the news with the rating of bi will lead to as much as bi calls.

Your task is to count the values of resi — how many students learned their first news on day i.

The values of bi are known initially, whereas ai is determined from the following formula:


where mod stands for the operation of taking the excess from the cleavage, res0 is considered equal to zero and vi — some given integers.

### ideas
1. 因为每个人只有一个friend, 且不能是自己，所以最终会进入一个环
2. 所以要计算，每个人要进入的环的距离，以及入口
3. 如果进入环了，需要知道是否跑了一圈
4. 这个题目好难理解啊。它问的是第i天，有多少个人是第一次听到news
5. 要找到哪些人听到了这个消息是比较容易的。但是哪些人第一次听到呢？
6. 要记录这条路径上，目前最后一个听到消息的人， 如果是一个圈，也是一样的
7. 