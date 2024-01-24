Recently, Kolya found out that a new movie theatre is going to be opened in his city soon, which will show a new movie
every day for 𝑛
days. So, on the day with the number 1≤𝑖≤𝑛
, the movie theatre will show the premiere of the 𝑖
-th movie. Also, Kolya found out the schedule of the movies and assigned the entertainment value to each movie, denoted
by 𝑎𝑖
.

However, the longer Kolya stays without visiting a movie theatre, the larger the decrease in entertainment value of the
next movie. That decrease is equivalent to 𝑑⋅𝑐𝑛𝑡
, where 𝑑
is a predetermined value and 𝑐𝑛𝑡
is the number of days since the last visit to the movie theatre. It is also known that Kolya managed to visit another
movie theatre a day before the new one opened — the day with the number 0
. So if we visit the movie theatre the first time on the day with the number 𝑖
, then 𝑐𝑛𝑡
— the number of days since the last visit to the movie theatre will be equal to 𝑖
.

For example, if 𝑑=2
and 𝑎=[3,2,5,4,6]
, then by visiting movies with indices 1
and 3
, 𝑐𝑛𝑡
value for the day 1
will be equal to 1−0=1
and 𝑐𝑛𝑡
value for the day 3
will be 3−1=2
, so the total entertainment value of the movies will be 𝑎1−𝑑⋅1+𝑎3−𝑑⋅2=3−2⋅1+5−2⋅2=2
.

Unfortunately, Kolya only has time to visit at most 𝑚
movies. Help him create a plan to visit the cinema in such a way that the total entertainment value of all the movies he
visits is maximized.

### thoughts

1. 假设最后一次visit剧院的日期是i那么cnt = i - m + m = i
2. 其中m是访问的次数
3. 居然是固定的，那就是迭代i从m开始到n，找出最大的 sum(max m) - i * d
