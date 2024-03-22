You have an array of positive integers a[1], a[2], ..., a[n] and a set of bad prime numbers b1, b2, ..., bm. The prime
numbers that do not occur in the set b are considered good. The beauty of array a is the sum , where function f(s) is
determined as follows:

f(1)= 0;
Let's assume that p is the minimum prime divisor of s. If p is a good prime, then , otherwise .
You are allowed to perform an arbitrary (probably zero) number of operations to improve array a. The operation of
improvement is the following sequence of actions:

Choose some number r (1 ≤ r ≤ n) and calculate the value g = GCD(a[1], a[2], ..., a[r]).
Apply the assignments: , , ..., .
What is the maximum beauty of the array you can get?

### ideas

1. f(p) 如果p是个good prime number f(p) = 2 else f(p) = -1
2. 考虑一个数num， 进行素数化p1 ** c1 * p2 ** c2...
3. 那么结果必然是 good(p1) * c1 + good(p2) * c2 ...
4. 那么一次操作后，应该尽量的删除掉bad prime number 这个cost是可以计算出来的
5. 这里还有一个问题，就是从左往右，还是从右往左呢？
6. 从左往右的问题是，有可能无法处理后面的bad prime number，所以感觉还是从右往左更好