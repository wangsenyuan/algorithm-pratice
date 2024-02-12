Pak Chanek is the chief of a village named Khuntien. On one night filled with lights, Pak Chanek has a sudden and
important announcement that needs to be notified to all of the 𝑛
residents in Khuntien.

First, Pak Chanek shares the announcement directly to one or more residents with a cost of 𝑝
for each person. After that, the residents can share the announcement to other residents using a magical helmet-shaped
device. However, there is a cost for using the helmet-shaped device. For each 𝑖
, if the 𝑖
-th resident has got the announcement at least once (either directly from Pak Chanek or from another resident), he/she
can share the announcement to at most 𝑎𝑖
other residents with a cost of 𝑏𝑖
for each share.

If Pak Chanek can also control how the residents share the announcement to other residents, what is the minimum cost for
Pak Chanek to notify all 𝑛
residents of Khuntien about the announcement?

### thoughts

1. 假设现在知道消息的人的集合为set，剩余w个人
2. 选谁来通知呢？是cost最低的那个人吗？还是 a[i] * b[i] 最小的那个人呢？
3. 或者让pak使用p来通知