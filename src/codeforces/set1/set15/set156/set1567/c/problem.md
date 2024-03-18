Alice has just learned addition. However, she hasn't learned the concept of "carrying" fully — instead of carrying to
the next column, she carries to the column two columns to the left.

For example, the regular way to evaluate the sum 2039+2976
would be as shown:

However, Alice evaluates it as shown:

In particular, this is what she does:

add 9
and 6
to make 15
, and carry the 1
to the column two columns to the left, i. e. to the column "0
9
";
add 3
and 7
to make 10
and carry the 1
to the column two columns to the left, i. e. to the column "2
2
";
add 1
, 0
, and 9
to make 10
and carry the 1
to the column two columns to the left, i. e. to the column above the plus sign;
add 1
, 2
and 2
to make 5
;
add 1
to make 1
.
Thus, she ends up with the incorrect result of 15005
.
Alice comes up to Bob and says that she has added two numbers to get a result of 𝑛
. However, Bob knows that Alice adds in her own way. Help Bob find the number of ordered pairs of positive integers such
that when Alice adds them, she will get a result of 𝑛
. Note that pairs (𝑎,𝑏)
and (𝑏,𝑎)
are considered different if 𝑎≠𝑏
.

### ideas

1. a + b => n
2. 考虑最低位x, 如果在这位产生了进位，也就是在2（+1）,或者没有进位 (+0)
3. dp[i][0/1] 表示在第i位没有（0）或者有（1）进位时的计数
4. 那这里要从高到低吗？
5. dfs(i, carry)
6. 但是问题是当前位是被i-2位影响到的，所以要把这个状态传递下去， carry需要用0,1,2,3 来表示
7. 