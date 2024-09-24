You are given n numbers a1, a2, ..., an. You can perform at most k operations. For each operation you can multiply one of the numbers by x. We want to make  as large as possible, where  denotes the bitwise OR.

Find the maximum possible value of  after performing at most k operations optimally.

### ideas
1. k <= 10, x <= 8
2. 是不是brute force就可以了？
3. 假设a按照倒序牌，目前的数字 * x 后，如果能贡献一个新的最高位，就应该使用这个操作
4. 不对,貌似就应该在最大的数上，一直乘上去。。
5. 