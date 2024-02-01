After Zhongkao examination, Daniel and his friends are going to have a party. Everyone will come with some candies.

There will be 𝑛
people at the party. Initially, the 𝑖
-th person has 𝑎𝑖
candies. During the party, they will swap their candies. To do this, they will line up in an arbitrary order and
everyone will do the following exactly once:

Choose an integer 𝑝
(1≤𝑝≤𝑛
) and a non-negative integer 𝑥
, then give his 2𝑥
candies to the 𝑝
-th person. Note that one cannot give more candies than currently he has (he might receive candies from someone else
before) and he cannot give candies to himself.
Daniel likes fairness, so he will be happy if and only if everyone receives candies from exactly one person. Meanwhile,
his friend Tom likes average, so he will be happy if and only if all the people have the same number of candies after
all swaps.

Determine whether there exists a way to swap candies, so that both Daniel and Tom will be happy after the swaps.

### thoughts

1. 两个条件， 所有人都必须给别人 2 ** x的部分,也必须接受别人给的 2 ** y的部分
2. 所有人最后都必须相等
3. 所有人最后的状态是确认的 avg = sum() / n (如果不能整除 => false)
4. 在能整除的情况下， 最后的 avg = x[i] - give[i] + receive[i]
5. receive[i] - give[i] = avg - x[i]
6. avg = x[i] 的部分，相当于对receive和give没有限制，收到多少就可以给出多少，先放一边
7. 我们考虑diff[i] = avg - x[i]
8. 如果它是一个只有一位的情况, 比如2，那么要能接受，且送出去，那么只能接受4, 送出2
8. 如果它是一个只有两位的情况, 比如6， 那么只能接受8送出2, 如果5,似乎就不行了
9. 也就是说它只能是一个连续的一位，接受1 << (hi + 1), 送出 1 << lo, 其他情况不可以
10. 然后分组后，判断给出和接受的情况即可