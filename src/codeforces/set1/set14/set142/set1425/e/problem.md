Mr. Chanek is currently participating in a science fair that is popular in town. He finds an exciting puzzle in the fair and wants to solve it.

There are 𝑁
 atoms numbered from 1
 to 𝑁
. These atoms are especially quirky. Initially, each atom is in normal state. Each atom can be in an excited. Exciting atom 𝑖
 requires 𝐷𝑖
 energy. When atom 𝑖
 is excited, it will give 𝐴𝑖
 energy. You can excite any number of atoms (including zero).

These atoms also form a peculiar one-way bond. For each 𝑖
, (1≤𝑖<𝑁)
, if atom 𝑖
 is excited, atom 𝐸𝑖
 will also be excited at no cost. Initially, 𝐸𝑖
 = 𝑖+1
. Note that atom 𝑁
 cannot form a bond to any atom.

Mr. Chanek must change exactly 𝐾
 bonds. Exactly 𝐾
 times, Mr. Chanek chooses an atom 𝑖
, (1≤𝑖<𝑁)
 and changes 𝐸𝑖
 to a different value other than 𝑖
 and the current 𝐸𝑖
. Note that an atom's bond can remain unchanged or changed more than once. Help Mr. Chanek determine the maximum energy that he can achieve!

note: You must first change exactly 𝐾
 bonds before you can start exciting atoms.


 ### ideas
 1. 如果K % 2 == 0, 那么可以使用初始状态的结果 sum(A) - D[1]
 2. 如果K % 2 = 1, 那么就应该找出 max pref(A[i]) - D[i] (将i指向1) + max(suf)
 3. K = 2时，有必要改变两个吗？
 4. 还是有可能的，就是分成K段
 5. 如果把i改为指向前面，那么肯定是指向越前面越好
 6. 因为改变的这个，比如是那个激活最低的那个（否则肯定不是最优的）
 7. 假设有两段（那么如果能连起来，那么就最好连起来，这样子，可以少激活一次）
 8. 如果K奇数，那么只需要把 E[n-1] => 1即可，然后激活最小的D[i]即可
 9. 还有一种情况，就是把这个最小的留着，从它那里开始激活直到最后，当时把它前面的链接到1，用前面最小的去激活
 10. 这两种情况都可以在o(n)的时间内处理
 11. 当k是偶数的时候，就是在前面选择两个最小的，然后从当前位置开始处理