/**
* Super nice dp problem! Let's try to walkthrough this, as the tutorial really isn't good.
* Directly couting for permutations must be difficult, so we first observe that the chain
* of gcd looks sth like 12,12,6,2,2,1,1, where every num divides the previous one. Namely,
* g_i := gcd(a_1, a_2, ..., a_i) and we know g_i | g_i-1.
*
* So, regardless of len we use max_sum[t] to denote the maximum gcd prefixed sum that can
* be achieved with the last gcd being exactly t. Clearly we have to greedily include all
* multiples of t for max_sum[t], but how shall we order them?
* The key is, we wanna have the min num of "t"s and go to a t * k (k > 1) as soon as
* possible. So in our example, max_sum[1] would be 2 ones and the rest would be
* max_sum[2]. Again we may not know the num of elements included in max_sum[2], but we
* just know that we have to minimize ones and the rest is just stored in max_sum[2].
*
* This is a bit inaccurately though, since the gcd may not be exactly t but a multiple of
* it. Anyway, let's do it formally:
*      let max_sum[t] to be the maximum prefix gcd, in a certain order, for all
*      elements that are divisible by t (we don't restrict it to having gcd t).
*      Suppose that the prev num greater t is t', so all numbers divisible by t'
*      shall be placed at the front, i.e. it's optimal to make t' multiples
*      included in max_sum[t'], so naturally we have
*            max_sum[t] = max{ max_sum[t'] + #{ x : t|x but t'∤x } * t }   (t | t')
*                       = max{ max_sum[t'] + (#{x : t|x } - #{x : t'|x }) * t }
*      In the end, the answer is in max_sum[1] by definition.
* You may be curious about why the t in  max_sum[t] need not be the gcd. For example in
* {18, 12, 6, 7, 6} max_sum[6] = gcd(18)+gcd(18,12)+gcd(18,12,6)+gcd(18,12,6,6) = 36
* and max_sum[2] has the same value (36), though of course gcd(18,12,6,6) is 6. But anyway
* as we find max_sum[1], by which we include the 7, referecing to max_sum[6] yields a
* result larger than max_sum[2]. So everything fits!
*
* This already yields a O(A_MAX * log(A_MAX)) solution without sieve, but for D2 we have
* to go a bit further. Notice (again) that for a t | t', t' has at least 1 more prime
* divisor than t, so there exists prime p s.t. pt | t', which means max_sum[t'] is already
* "covered" by max_sum[pt]. Therefore, it suffices to iterate through all t' = tp. How
* brilliant!
* But wait, there's some impl details in the code comments. I think I've learnt loads from
* it.
*
* Time Complexity: (unable to determine)   (at least o(A_MAX * log(log(A_MAX))))
* Implementation 1
  */

#include <bits/stdc++.h>

typedef int64_t     int_t;
typedef std::vector<int_t>  vec;

const int_t A_MAX = 2 * 1e7;
const int_t EXP_PRIME = 1270607;


int main() {
std::ios_base::sync_with_stdio(false);
std::cin.tie(NULL);

    int_t n;
    std::cin >> n;
    vec count(A_MAX + 1, 0);    // count[t] = num of values divisible by t
    for (int_t k = 0; k < n; k++) {
        int_t value;
        std::cin >> value;
        count[value]++;
    }


    vec primes;
    primes.reserve(EXP_PRIME);
    std::vector<bool> sieve(A_MAX + 1, false);
    for (int_t t = 2; t <= A_MAX; t++) {
        if (!sieve[t]) {
            primes.emplace_back(t);
            for (int_t k = 2; k * t <= A_MAX; k++)
                sieve[k * t] = true;
        }
    }
    std::cerr << "prime sieve finished with " << primes.size() << " primes" << std::endl;

    // Extremely IMPORTANT here!
    // The loop for iterating through primes is the outer one. This ensures no double
    // counting, eg 12 counted in count[2] twice as 2*3|12 and 2*2|12. Also we iterate
    // from large t to small ones.
    //      We can't write the outer loop as the t one and the inner loop as
    //      prime one.
    for (int_t p : primes) {
        for (int_t t = A_MAX / p; t >= 1; t--)
            count[t] += count[t * p];
    }
    std::cerr << "divisor counting finished" << std::endl;

    vec m_sum(A_MAX + 1);   // hope it won't be MLE
    for (int_t t = A_MAX; t >= 1; t--) {
        // makes sure that we've included those gcd(a_1) as t_new > A_MAX would
        // just break the whole thing~ A little lower bound to include the n
        // given values.
        m_sum[t] = t * count[t];    // a lower bound
        // num of primes is difficult to estimate so for time complexity i wrote unable...
        for (int_t p : primes) {
            int_t t_new = t * p;
            if (t_new > A_MAX)
                break;
            m_sum[t] = std::max(m_sum[t], m_sum[t_new] + t * (count[t] - count[t_new]));
        }
    }
    std::cerr << "max gcd sum computed" << std::endl;
    std::cout << m_sum[1] << '\n';
}