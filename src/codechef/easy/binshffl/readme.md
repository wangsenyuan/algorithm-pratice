* 1\. problem statement:
Chef has two integers A and B. He can perform the following operation on A an arbitrary number of times (including zero):

  * 1\. write A as a binary number with an arbitrary number of leading zeroes (possibly without any)
  * 2\. shuffle the binary digits of A in an arbitrary way, obtaining a number s
  * 3\. replace A by s+1

Chef is wondering about the minimum number of operations he has to perform on A in order to obtain B. Compute this number or determine that it is impossible.

* 2\. solution
  * 1\. 假设两个数字（二进制表示）有相同数目的1，那么在不考虑第三步的情况下，通过第二步的shuffle，可以使A和B相同；
  * 2\. 如果执行了一次操作，那么第三步都会被执行，也就是最后的结果都会+1
  * 3\. 考虑A和B-1, 当使得A和B-1相同的时候，再执行一步+1， 就可以使得A和B相同；
  * 4\. 如何使得X和Y，（在不执行最后一步+1的情况下）相同呢？因为上一步是shuffle，只要X和Y的1的位数相同，那么通过shuffle就可以使它们相同；每一步最多添加一个1；
    * 1\. 如果one_count(X) == one_count(Y), 那么不需要添加任何1，shuffle，就可以使X，Y相同；
    * 2\. 如果one_count(X) < one_count(Y), 那么需要添加one_count(Y) - one_count(X)个1，最后再shuffle即可；
    * 3\. 如果one_count(X) > one_count(Y), 那么只需要一步，就可以将X和Y的1的个数相同。 假设X为100010110, Y为0010100, 即X比Y多2个位数1. 那么可以在以下的一个操作内，将X的1位减少2:
      * 1\. 添加更多的leading zeors， （不需要）
      * 2\. shuffle后得到 1000000111
      * 3\. +1 后得到 1000001000
