Let's assume that we have a pair of numbers (a, b). We can get a new pair (a + b, b) or (a, a + b) from the given pair in a single step.

Let the initial pair of numbers be (1,1). Your task is to find number k, that is, the least number of steps needed to transform (1,1) into the pair where at least one number equals n.

### ideas
1. 好像不大对
2. (1, 1) -> (1, 2) -> (1, 3) ... (1, n) (n - 1 步)
3. (1, 1) -> (1, 2) -> (2, 3) ... (?. n)
4. gcd(n, x) -> gcd(x, n - x) -> ... -> gcd(1, 1) = 1
5. 如果x是n的一个因数, 那么肯定是不行的
6. 比如 n = 4, x = 2, gc(4, 2) => gcd(2, 2) => gcd(0, 2)
7. 所以x不能是n个一个除数, gcd(n, x) -> gcd(n - x, x)
8. x必须是一个 gcd(n, x) = 1的数
9. 这样的一个数, gcd(n, x) => gcd(n - x, x) -> gcd(n - 2 * x, x) .... gcd(1, x)
10. 和n互质的数， 它不包括的质数的
11. 假设求出了 f(x) (如果x和y互质，那么 f(y) = f(x) + y / x)
12. 但是怎么找到y呢？迭代肯定不大行
13. 