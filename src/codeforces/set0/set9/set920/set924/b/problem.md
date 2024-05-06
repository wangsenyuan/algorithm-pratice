An atom of element X can exist in n distinct states with energies E1<E2<...<En. Arkady wants to build a laser on this
element, using a three-level scheme. Here is a simplified description of the scheme.

Three distinct states i, j and k are selected, where i<j<k. After that the following process happens:

initially the atom is in the state i,
we spend Ek - Ei energy to put the atom in the state k,
the atom emits a photon with useful energy Ek - Ej and changes its state to the state j,
the atom spontaneously changes its state to the state i, losing energy Ej - Ei,
the process repeats from step 1.
Let's define the energy conversion efficiency as , i. e. the ration between the useful energy of the photon and spent
energy.

Due to some limitations, Arkady can only choose such three states that Ek - Ei ≤ U.

Help Arkady to find such the maximum possible energy conversion efficiency within the above constraints.

### ideas

1. x = (E[k] - E[j]) / (E[k] - E[i])
2. 首先有个观察，在给定i, k的情况下，就是j其实越小越好，这样子 E[k] - E[j] 最大
3. 所以j = i + 1
4. 然后就是在E[k] - E[i] <= U 的前提下，找到使得x最大的k
5. (E[k] - a) / (E[k] - b) = x
6. E[k] - a = x * E[k] - x * b
7. (1 - x) * E[k] = a - x * b
8. E[k] = (a - x * b) / (1 - x) <= U + b
9. 数学不好，分析不下去了
10. 上式中，在a, b, U确定的情况下，x的范围也是确定的（或者是可解的）
11. 但是如何确定k呢？
12. 另 y = (x - a) / (x - b) (a > b)
13. x = b 时，上式接近无穷大，x = a时上式= 0
14. 那时是不是它是单调递降的？
15. k越大，分母越大，分子也越大
16. k居然是越大越好

### solution

First of all, you can note that for fixed i and k setting j = i + 1 is always the best choice. Indeed, if X>Y, then for
positive B.

Then, let's fix i. Then j = i + 1, and what is the optimal k? We can define the energy loss as . As we need to minimize
the loss, it's obvious that we should maximize Ek, so we should choose as large k as possible satisfying Ek - Ei ≤ U.
This is a classic problem that can be solved with two pointers approach, this leads to O(n) solution, or with binary
search approach, this leads to solution. Both are acceptable.