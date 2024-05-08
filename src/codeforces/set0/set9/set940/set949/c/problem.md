BigData Inc. is a corporation that has n data centers indexed from 1 to n that are located all over the world. These
data centers provide storage for client data (you can figure out that client data is really big!).

Main feature of services offered by BigData Inc. is the access availability guarantee even under the circumstances of
any data center having an outage. Such a guarantee is ensured by using the two-way replication. Two-way replication is
such an approach for data storage that any piece of data is represented by two identical copies that are stored in two
different data centers.

For each of m company clients, let us denote indices of two different data centers storing this client data as ci, 1 and
ci, 2.

In order to keep data centers operational and safe, the software running on data center computers is being updated
regularly. Release cycle of BigData Inc. is one day meaning that the new version of software is being deployed to the
data center computers each day.

Data center software update is a non-trivial long process, that is why there is a special hour-long time frame that is
dedicated for data center maintenance. During the maintenance period, data center computers are installing software
updates, and thus they may be unavailable. Consider the day to be exactly h hours long. For each data center there is an
integer uj (0 ≤ uj ≤ h - 1) defining the index of an hour of day, such that during this hour data center j is
unavailable due to maintenance.

Summing up everything above, the condition uci, 1 ≠ uci, 2 should hold for each client, or otherwise his data may be
unaccessible while data centers that store it are under maintenance.

Due to occasional timezone change in different cities all over the world, the maintenance time in some of the data
centers may change by one hour sometimes. Company should be prepared for such situation, that is why they decided to
conduct an experiment, choosing some non-empty subset of data centers, and shifting the maintenance time for them by an
hour later (i.e. if uj = h - 1, then the new maintenance hour would become 0, otherwise it would become uj + 1).
Nonetheless, such an experiment should not break the accessibility guarantees, meaning that data of any client should be
still available during any hour of a day after the data center maintenance times are changed.

Such an experiment would provide useful insights, but changing update time is quite an expensive procedure, that is why
the company asked you to find out the minimum number of data centers that have to be included in an experiment in order
to keep the data accessibility guarantees.

### ideas

1. 这个题目有点理解不了。主要是不理解 答案不应该是0吗？
2. client 的要求，就是在两个节点间创建了一个连接，如果改变其中一个，如果它造成和它连接的节点冲突了
3. 那么对应的节点也应该改变
4. 但是这里涉及到一个点，就是 a -> b, b不一定改变a
5. a -> b -.... -> a 了，组成了一个环，那么就必须同时变更这一圈
6. 如果存在某条line， 那么答案是1， 最后一个节点
7. 如果大家都到了一个圈，那么就是最小的圈的大小
8. got
9. 