There are 𝑛
 persons who initially don't know each other. On each morning, two of them, who were not friends before, become friends.

We want to plan a trip for every evening of 𝑚
 days. On each trip, you have to select a group of people that will go on the trip. For every person, one of the following should hold:

Either this person does not go on the trip,
Or at least 𝑘
 of his friends also go on the trip.
Note that the friendship is not transitive. That is, if 𝑎
 and 𝑏
 are friends and 𝑏
 and 𝑐
 are friends, it does not necessarily imply that 𝑎
 and 𝑐
 are friends.

For each day, find the maximum number of people that can go on the trip on that day.

### ideas
1. 假设知道目前所有人的状态（他有多少个朋友），
2. 如何计算有多少个人去trip呢？
3. 参与的人的去参加的朋友的个数 >= k才行
4. 在第i-1天参与的人，在第i天仍然时参与的
5. 那么在第i天，新增的一对(a, b)，首先会影响到(a, b)
6. 如果如果如果a, b 已经参与了，那么这里没有影响
7. 如果a, b都没有参与，那么a至少要k-1个人，b也要至少k-1个人？
8. 或者a已经参与了，b是这次参与
9. 假设x=(a, b) 满足参与的条件， 那么因为x的参与，又会去更新和x有关联的那些人
10. 前面那个也不成立的，因为a，b加入，可以会让一票人都满足条件
11. 首先新加入的x，比如是k-1个friends在了，否则没法继续
12. 然后因为x的加入，y是x之前的一个邻居，它也必须是k-1个在群内，因为x加入，所有它可以了
13. 但是z，如果z既是x的朋友，又是y的朋友，那么它需要的就是k-2个在群里面
14. 用一个heap？
15. 新加入的人员，如果他的准备阶段的朋友的数量+已经参加的朋友的数量 >= k， 那么他就可以加入
16. 准备阶段的，不一定马上也能加入呐
17. 那么对于x来说，他能加入的时间，就等于他的前k个朋友的最晚加入时间
18. 先把所有的都加入到一个图里面，只保留那些deg >= k 的节点（这些是最后能够参加的人）其他的事件可以直接忽略掉
19. 