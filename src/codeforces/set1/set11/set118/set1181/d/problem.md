Misha was interested in water delivery from childhood. That's why his mother sent him to the annual Innovative Olympiad in Irrigation (IOI). Pupils from all Berland compete there demonstrating their skills in watering. It is extremely expensive to host such an olympiad, so after the first 𝑛
 olympiads the organizers introduced the following rule of the host city selection.

The host cities of the olympiads are selected in the following way. There are 𝑚
 cities in Berland wishing to host the olympiad, they are numbered from 1
 to 𝑚
. The host city of each next olympiad is determined as the city that hosted the olympiad the smallest number of times before. If there are several such cities, the city with the smallest index is selected among them.

Misha's mother is interested where the olympiad will be held in some specific years. The only information she knows is the above selection rule and the host cities of the first 𝑛
 olympiads. Help her and if you succeed, she will ask Misha to avoid flooding your house.

### ideas
1. 到目前为止的频率可以计算出来
2. 然后可以计算出下一次所有人都一致的时间吗？
3. 然后再之后的就是依次(1, 2, ... m)循环
4. 假设当前的最大频率是k, 那么再 m * k 次后，开始循环
5. 在这之前，要怎么计算呢？
6. 首先肯定是最少且（同样数量时，最左边的）的城市举行
7. 但是这样子的问题是枚举起来太慢了
8. 假设最后是在城市i举行第k次，要满足什么样的条件呢？
9. 所有它左边的城市，举行过的次数大于它，且所有右边的城市 >= 它的举办次数
10. 假设在城市i举行前第k次时的频率是x,
11. 那么有(先要把右边填平到x) 这个可以算出来
12. 左边也要填平 (x + 1)
13. （x + 1) * cl - sum(...) + x * cr - sum(..) = k - 1
14. 假设按照最少最左的规则排好序
15. 那么到i能够第一次能够排到的次数，是可以计算出来的, 就是上面那个式子
16. 根据k分布在哪个区间，再二分查找就可以了
17. 