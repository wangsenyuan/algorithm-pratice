Alice likes snow a lot! Unfortunately, this year's winter is already over, and she can't expect to have any more of it.
Bob has thus bought her a gift — a large snow maker. He plans to make some amount of snow every day. On day i he will
make a pile of snow of volume Vi and put it in her garden.

Each day, every pile will shrink a little due to melting. More precisely, when the temperature on a given day is Ti,
each pile will reduce its volume by Ti. If this would reduce the volume of a pile to or below zero, it disappears
forever. All snow piles are independent of each other.

Note that the pile made on day i already loses part of its volume on the same day. In an extreme case, this may mean
that there are no piles left at the end of a particular day.

You are given the initial pile sizes and the temperature on each day. Determine the total volume of snow melted on each
day.

### thoughts

1. 假设到目前为止，还没有融化的雪人为set
2. 那么在day i, set += V[i] and all reduce by T[i], 然后求和
3. 所有小于等于T[i]的要退出，剩下的部分求和 - cnt * T[i]即可
4. 还有问题，就是不同不同雪人融化的量是不同的
5. 假设第一个雪人volume很大，到目前一直没有融化，但第二个雪人很小，很块就融化了
6. 所以，需要知道每个雪人融化的时机，没有融化的那些的sum里面要减去cnt * T[i]