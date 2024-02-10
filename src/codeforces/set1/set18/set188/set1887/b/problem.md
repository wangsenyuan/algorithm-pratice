Berland is a country with ancient history, where roads were built and destroyed for centuries. It is known that there
always were 𝑛
cities in Berland. You also have records of 𝑡
key moments in the history of the country, numbered from 1
to 𝑡
. Each record contains a list of bidirectional roads between some pairs of cities, which could be used for travel in
Berland at a specific moment in time.

You have discovered a time machine that transports you between key moments. Unfortunately, you cannot choose what point
in time to end up at, but you know the order consisting of 𝑘
moments in time 𝑎𝑖
, in which the machine will transport you. Since there is little time between the travels, when you find yourself in the
next key moment in time (including after the last time travel), you can travel on at most one existing road at that
moment, coming out from the city you were in before time travel.

Currently, you are in city 1
, and the time machine has already transported you to moment 𝑎1
. You want to reach city 𝑛
as quickly as possible. Determine the minimum number of time travels, including the first one, that you need to make in
order to reach city 𝑛
.

### thoughts

1. 好复杂的一个题目啊
2. 假设在时刻t，you可以到达的city的集合为set（包括stay在那里的）
3. 那么在下一个时刻，可以扩展这个集合，但是问题是这个复杂性太高
4. 对于每个点，反过来计算它最快什么时候可以激活
5. 对于点u，假设和它相邻的点v在时刻t激活，那么u激活的时间，必然是下次和v在同一个时间内，存在路径的时刻
6. 那么对于每条边(u, v)记录一个列表，它们相邻的时刻，然后更新即可