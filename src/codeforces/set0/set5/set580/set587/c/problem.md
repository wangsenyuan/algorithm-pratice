Recently Duff has been a soldier in the army. Malek is her commander.

Their country, Andarz Gu has n cities (numbered from 1 to n) and n - 1 bidirectional roads. Each road connects two different cities. There exist a unique path between any two cities.

There are also m people living in Andarz Gu (numbered from 1 to m). Each person has and ID number. ID number of i - th person is i and he/she lives in city number ci. Note that there may be more than one person in a city, also there may be no people living in the city.


Malek loves to order. That's why he asks Duff to answer to q queries. In each query, he gives her numbers v, u and a.

To answer a query:

Assume there are x people living in the cities lying on the path from city v to city u. Assume these people's IDs are p1, p2, ..., px in increasing order.

If k = min(x, a), then Duff should tell Malek numbers k, p1, p2, ..., pk in this order. In the other words, Malek wants to know a minimums on that path (or less, if there are less than a people).

Duff is very busy at the moment, so she asked you to help her and answer the queries.

### ideas
1. 假设每个city最多只有一个人，那么，就很简单了。
2. 找到p = lca(u, v), 然后分别在路径p...u, p...v上找出x个，然后再合并即可
3. got了。类似的