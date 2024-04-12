Imur Ishakov decided to organize a club for people who love to play the famous game «The hat». The club was visited by n
students, where n is even. Imur arranged them all in a circle and held a draw to break the students in pairs, but
something went wrong. The participants are numbered so that participant i and participant i + 1 (1 ≤ i ≤ n - 1) are
adjacent, as well as participant n and participant 1. Each student was given a piece of paper with a number in such a
way, that for every two adjacent students, these numbers differ exactly by one. The plan was to form students with the
same numbers in a pair, but it turned out that not all numbers appeared exactly twice.

As you know, the most convenient is to explain the words to the partner when he is sitting exactly across you. Students
with numbers i and sit across each other. Imur is wondering if there are two people sitting across each other with the
same numbers given. Help him to find such pair of people if it exists.

You can ask questions of form «which number was received by student i?», and the goal is to determine whether the
desired pair exists in no more than 60 questions.

### ideas

1. 60这个上限，说明是要用二分的方式
2. 但是完全没有头绪～
3. 但是这里有个关键的信息是，相邻两个数的相差为1，也就是说，如果a[1] = 0, 那么 a[2] = 1/-1
4. let size = 2 * n
4. 如果a[n+1] = 0, 那么就找到了一个pair
5. diff = a[n+1] - a[1], 假设 a[n+1] > a[1]
6. h 是 1...n的中间
7. a[n+h] - a[h] > diff

### solution

Let a(i) be a number given to the i-th student. Let's introduce function . As n is even, is integer, and b(i) is defined
correctly. Notice two facts: first, , and second, . The problem is to find an i such that b(i)= 0.

Second note leads us to observation that all b(i) have the same oddity. So, let's find b(0), and if it is odd, then
there is no answer, and we need to print - 1 — all b(i) have the same oddity (odd), and zero, as a number of different
oddity, won't appear in b(i).

Otherwise, suppose we know b(0), and it is equal to x (without loss of generality x>0). Notice that . As we remember
from second observation that all b(i) have the same oddity, and neighboring numbers differ by no more than 2, we can use
discrete continuity.

Lemma: if you have two indices i and j with values b(i) and b(j), then on segment between i and j there are all values
with the same oddity from min(b(i), b(j)) to max(b(i), b(j)). Indeed, as neighboring b differ by no more than 2, we
couldn't skip any number with the same oddity.

Now we can use a binary search. At the start l = 0, , and values b(l) and b(r) are even numbers with the different
signs. By lemma, on segment between l and r there is a zero that we want to find. Let's take . If b(m)= 0, we found the
answer. Otherwise, based on sign of b(m) we can replace one of l, r in binary search to m, and reduce the problem to the
same with smaller size. There will be no more than iterations, and as calculation of b(i) requires two queries, we
solved the problem in queries.