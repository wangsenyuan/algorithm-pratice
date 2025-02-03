https://codeforces.com/problemset/problem/1996/C

输入 T(≤1e3) 表示 T 组数据。所有数据的 n 之和 ≤2e5，q 之和 ≤2e5。
每组数据输入 n(1≤n≤2e5)，q(1≤q≤2e5)，长为 n 的字符串 s，长为 n 的字符串 t，都只包含小写英文字母。
然后输入 q 个询问，每个询问输入两个数 L 和 R，表示下标从 L 到 R 的子串 (1≤L≤R≤n) s[L..R] 和 t[L..R]。

对于每个询问，你可以修改 s[L..R] 中的若干字母，使得 s[L..R] 和 t[L..R] 排序后相等。
输出最小修改次数。

注意：每个询问互相独立。

进阶：值域更大的情况 https://atcoder.jp/contests/abc367/tasks/abc367_f

输入
3
5 3
abcde
edcba
1 5
1 4
3 3
4 2
zzde
azbe
1 3
1 4
6 3
uwuwuw
wuwuwu
2 4
1 3
1 6
输出
0
1
0
2
2
1
1
0

【灵茶の试炼】题目&题解
https://docs.qq.com/sheet/DWGFoRGVZRmxNaXFz