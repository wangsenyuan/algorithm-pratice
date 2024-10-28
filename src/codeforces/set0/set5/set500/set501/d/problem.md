Let's define the sum of two permutations p and q of numbers 0, 1, ..., (n - 1) as permutation , where Perm(x) is the x-th lexicographically permutation of numbers 0, 1, ..., (n - 1) (counting from zero), and Ord(p) is the number of permutation p in the lexicographical order.

For example, Perm(0) = (0, 1, ..., n - 2, n - 1), Perm(n! - 1) = (n - 1, n - 2, ..., 1, 0)

Misha has two permutations, p and q. Your task is to find their sum.

Permutation a = (a0, a1, ..., an - 1) is called to be lexicographically smaller than permutation b = (b0, b1, ..., bn - 1), if for some k following conditions hold: a0 = b0, a1 = b1, ..., ak - 1 = bk - 1, ak < bk.

### ideas
1. 关键是先要求出order，然后要能从order恢复
2. 求出order，可以从左往右，如果当前是目前最小的数字（比如1）那么就到下一位（order没有变）
3. 否则 order += diff * (n-1)!, 然后移动到下一位
4. 不过没法直接计算出(n-1)！，这个太大了，肯定会溢出
5. 但是可以记录下来(n-1)的系数（diff）
6. 但是要咋恢复呢？
7. 考虑一个例子[0, 1, 2] 
8. [0, 1, 2], [0, 2, 1], [1, 0, 2], [1, 2, 0], [2, 0, 1], [2, 1, 0] (3!= 6)
9. (0, 1, 2, 3, 4, 5)  
10. (5 + 5) % 6 = 4
11. (3 + 4) % 6 = 1
12. 所以不能计算完再恢复。搞不定
13. 换个角度，可以考虑p的互补的位置在哪里？也就是 order(p) + order(p^) = 0
14. [0, 1, 2] vs [2, 1, 0]
15. [1, 0, 2] vs [1, 2, 0]
16.  (n - 1 - i) % (n - 1) ?
17. [0, 1, 2, 3], [0, 1, 3, 2]                  [3, 2, 0, 1], [3, 2, 1, 0]
18. 好像是的
19. 如果 q <= p^, 那么就可以用order的系数直接加起来（不会溢出）
20. 否则的话，可以计算和 q 和 p^ 的系数之差？
21. q > p^ (系数之差要怎么算呢？)
22. 是不是不用这么麻烦。所有的位上，都取余(n-1)即可？