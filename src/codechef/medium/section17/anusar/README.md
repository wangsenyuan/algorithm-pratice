#### statement

Given a string S and several frequencies Fi. For each Fi output the number of substrings of S (the characters of
substring should be contiguous) that occur at least Fi times in S. Note, that we consider two substrings distinct if
they have distinct length, or they have distinct starting indices.

##### constraints

1. 1 ≤ |S| ≤ 200000
2. 1 ≤ Q ≤ 200000
3. 1 ≤ Fi ≤ 200000

#### 测试数据

```
Input:
1
aaeddf
4
1
2
3
4
Output:
21
4
0
0
```

#### 思考

1. F[i] 越大,子串的数目会越小
2. 假设有一个子串 s[i..j], 重复了x次, 则所有的F[i] <= x, 需要计入
3. 如果 s[i...j] 重复x次， 则 s[i..j-1]也至少重复了x次, 所以对于 S[i...j] (重复x次)， 计入的次数是 j - i + 1
4. 所以，对于任何一个i, 找到j, s[i...j] 不重复 (x = 1), 然后就可以停止了（j到n, x = 1）
5. (i...j) 的过程中，计算x
6. 然后直接到j的位置重新开始
7. 然后就是如何知道 s[i...j]是否重复，已经重复几次
8. 还是有些问题


#### solutions
```c++
#include <cstdio>
#include <cstring>
#include <algorithm>
#include <iostream>
#include <vector>
#include <map>
using namespace std;

#define SA_N 220000
#define SA_LN 20
int lg, len, p[SA_LN][SA_N], sa[SA_LN];
int la[SA_N], lb[SA_N], lc[SA_N];
int bucket[SA_N], tla[SA_N], tlb[SA_N], tlc[SA_N];
void constructSA(int *a, int leng) {
	len = leng;
	lg = 1;
	for(int i=0; i<len; i++) p[0][i] = a[i];
	sort(a, a+len);
	map <int, int> M;
	for(int i=0; i<len; i++) if(i==0 || (a[i] != a[i-1])) M[a[i]] = i;
	for(int i=0; i<len; i++) p[0][i] = M[p[0][i]];
	for(int skip=1; skip/2 < len; skip<<=1, lg++) {
		for(int i=0; i<len; i++) {
			la[i] = p[lg-1][i];
			lb[i] = i+skip < len ? p[lg-1][i+skip] : -1;
			lc[i] = i;
		}
		for(int i=0; i<len+10; i++) bucket[i] = 0;
		for(int i=0; i<len; i++) bucket[lb[i]+1]++;
		for(int i=1; i<len+10; i++) bucket[i] += bucket[i-1];
		for(int i=len-1; i>=0; i--) {
			int wer = (bucket[lb[i]+1]--)-1;
			tla[wer] = la[i]; tlb[wer] = lb[i]; tlc[wer] = lc[i];
		}
		for(int i=0; i<len; i++) la[i] = tla[i], lb[i] = tlb[i], lc[i] = tlc[i];
		for(int i=0; i<len+10; i++) bucket[i] = 0;
		for(int i=0; i<len; i++) bucket[la[i]]++;
		for(int i=1; i<len+10; i++) bucket[i] += bucket[i-1];
		for(int i=len-1; i>=0; i--) {
			int wer = (bucket[la[i]]--)-1;
			tla[wer] = la[i]; tlb[wer] = lb[i]; tlc[wer] = lc[i];
		}
		for(int i=0; i<len; i++) la[i] = tla[i], lb[i] = tlb[i], lc[i] = tlc[i];
		for(int i=0; i<len; i++) {
			p[lg][lc[i]] = (i>0 && la[i-1]==la[i] && lb[i-1]==lb[i])? p[lg][lc[i-1]] : i;
		}
	}
	lg--;
	for(int i=0; i<len; i++) sa[p[lg][i]] = i;
	// for(int i=0; i<len; i++) printf("%d ",p[lg][i]); printf("\n");
	//p[lg][] has the rank of every index.
}
void constructSA(char *t, int leng) {
	int *a = new int[leng];
	for(int i=0; i<leng; i++) a[i] = int(t[i]);
	constructSA(a, leng);
	delete a;
}
int lcp(int i, int j) {
	if(i==j) return len - i;
	int ans = 0;
	for(int k = lg; k>=0 && i<len && j<len; k--)
		if(p[k][i] == p[k][j])
			ans += (1<<k), i += (1<<k), j += (1<<k);
	return ans;
}

#define N SA_N
int stk[N], LCP[N], L[N], R[N];
vector <int> li[N];
char s[N];
long long d[N];
int main() {
	int t;
	scanf("%d", &t);
	while(t--) {
		scanf("%s", s);
		constructSA(s, strlen(s));
		// cerr << string(s) << endl;
		for(int i=0; i<=len+2; i++) li[i].clear(), d[i] = 0;
		for(int i=0; i<len-1; i++) li[ (LCP[i] = lcp(sa[i], sa[i+1])) ].push_back(i);
		int top = 0;
		for(int i=0; i<len-1; i++) {
			while(top && LCP[stk[top-1]] >= LCP[i]) top--;
			if(top == 0) L[i] = -1;
			else L[i] = stk[top-1];
			stk[top++] = i;
		}
		top = 0;
		for(int i=len-2; i>=0; i--) {
			while(top && LCP[stk[top-1]] >= LCP[i]) top--;
			if(top == 0) R[i] = len-1;
			else R[i] = stk[top-1];
			stk[top++] = i;
		}
		for(int i=len; i; i--) {
			int j = 0;
			while(j < li[i].size()) {
				int mini = i;
				int k = li[i][j];
				if(L[k] != -1 && mini > i - LCP[L[k]]) mini = i - LCP[L[k]];
				if(R[k] != len-1 && mini > i - LCP[R[k]]) mini = i - LCP[R[k]];
				int range = R[k] - L[k] - 1 + 1;
				d[range] += (long long)mini * range;
				while(j < li[i].size() && li[i][j] < R[k]) j++;
			}
		}
		for(int i=len-1; i; i--) d[i] += d[i+1];
		d[1] = (long long)len * (len + 1);
		d[1] >>= 1;
		int q;
		scanf("%d", &q);
		while(q--) {
			int x;
			scanf("%d", &x);
			printf("%lld\n", x>len?0:d[x]);
		}
	}
	return 0;
}
```

```
#include <bits/stdc++.h>
using namespace std;

typedef long long LL;

// suffix array O(log^2(N)) algorithm.
struct entry {
    int nr[2];
    int p;
};

bool cmp (entry a, entry b) {
    if (a.nr[0] == b.nr[0]) return a.nr[1] < b.nr[1];
    else return a.nr[0] < b.nr[0];
}

const int MAXN = 200005;
const int MAXLOG = 20;

char s[MAXN];
entry L[MAXN];
int P[MAXLOG][MAXN];

int stp, cnt;
int N;

int findLCP (int x, int y) {
    int ret = 0;
    if (x == y) return N - x;

    for (int k = stp - 1; k >= 0 && x < N && y < N; k--)
        if (P[k][x] == P[k][y]) {
            x += (1 << k);
            y += (1 << k);
            ret += (1 << k);
        }

    return ret;
}

void suffixArray() {
      for (int i = 0; i < N; i++)
            P[0][i] = (int) (s[i] - 'a');

      for (stp = 1, cnt = 1; cnt >> 1 < N; stp ++, cnt *= 2) {
            // compute L
            for (int i = 0; i < N; i++) {
                L[i].nr[0] = P[stp - 1][i];
                L[i].nr[1] = i + cnt < N ? P[stp - 1][i + cnt] : -1;
                L[i].p = i;
            }

            sort (L, L + N, cmp);

            for (int i = 0; i < N; i++) {
                if (i > 0 && L[i].nr[0] == L[i - 1].nr[0] && L[i].nr[1] == L[i - 1].nr[1])
                        P[stp][L[i].p] = P[stp][L[i - 1].p];
                else P[stp][L[i].p] = i;
            }
      }
}

struct maxSegmentTree {
      vector <int> data;
      int n;

      maxSegmentTree(int _n) {
            n = _n;
            data.resize(4 * n);
            // initialize with -1.
            build(1, 1, n);
      }

      void build(int k, int lo, int hi) {
            if (lo == hi) data[k] = - 1;
            else {
                  int mid = (lo + hi) / 2;
                  build(2 * k, lo, mid);
                  build(2 * k + 1, mid + 1, hi);
                  data[k] = max(data[2 * k], data[2 * k + 1]);
            }
      }

      void add(int pos, int val) {
            // increase pos to make in the range [1, n]
            update(1, 1, n, pos + 1, val);
      }

      void update(int k, int lo, int hi, int pos, int val) {
            if (lo == hi && lo == pos) {
                  data[k] = max(data[k], val);
            } else {
                  int mid = (lo + hi) / 2;
                  if (pos <= mid) update(2 * k, lo, mid, pos, val);
                  else if (pos > mid) update(2 * k + 1, mid + 1, hi, pos, val);
                  data[k] = max(data[2 * k], data[2 * k + 1]);
            }
      }

      int ask(int pos) {
            if (pos < 0)
                  return -1;
            // increase pos .
            return query(1, 1, n, 1, pos + 1);
      }

      int query(int k, int lo, int hi, int left, int right) {
            if (lo == left && hi == right) {
                  return data[k];
            } else {
                  int mid = (lo + hi) / 2;
                  if (right <= mid) return query(2 * k, lo, mid, left, right);
                  else if (left > mid) return query(2 * k + 1, mid + 1, hi, left, right);
                  else {
                        int ans1 = query(2 * k, lo, mid, left, mid);
                        int ans2 = query(2 * k + 1, mid + 1, hi, mid + 1, right);
                        return max(ans1, ans2);
                  }
            }
      }
};

struct minSegmentTree {
      vector <int> data;
      int n;

      minSegmentTree(int _n) {
            n = _n;
            data.resize(4 * n);
            // initialize with n-1.
            build(1, 1, n);
      }

      void build(int k, int lo, int hi) {
            if (lo == hi) data[k] = n - 1;
            else {
                  int mid = (lo + hi) / 2;
                  build(2 * k, lo, mid);
                  build(2 * k + 1, mid + 1, hi);
                  data[k] = min(data[2 * k], data[2 * k + 1]);
            }
      }

      void add(int pos, int val) {
            // increase pos to make in the range [1, n]
            update(1, 1, n, pos + 1, val);
      }

      void update(int k, int lo, int hi, int pos, int val) {
            if (lo == hi && lo == pos) {
                  data[k] = min (data[k], val);
            } else {
                  int mid = (lo + hi) / 2;
                  if (pos <= mid) update(2 * k, lo, mid, pos, val);
                  else if (pos > mid) update(2 * k + 1, mid + 1, hi, pos, val);
                  data[k] = min(data[2 * k], data[2 * k + 1]);
            }
      }

      int ask(int pos) {
            if (pos < 0)
                  return n - 1;
            // increase pos .
            return query(1, 1, n, 1, pos + 1);
      }

      int query(int k, int lo, int hi, int left, int right) {
            if (lo == left && hi == right) {
                  return data[k];
            } else {
                  int mid = (lo + hi) / 2;
                  if (right <= mid) return query(2 * k, lo, mid, left, right);
                  else if (left > mid) return query(2 * k + 1, mid + 1, hi, left, right);
                  else {
                        int ans1 = query(2 * k, lo, mid, left, mid);
                        int ans2 = query(2 * k + 1, mid + 1, hi, mid + 1, right);
                        return min(ans1, ans2);
                  }
            }
      }
};

int main() {
      int T;
      scanf ("%d", &T);
      while (T--) {
            scanf ("%s", s);
            N = strlen(s);

            suffixArray();

            vector <int> a;
            for (int i = 0; i + 1 < N; i++)
                  a.push_back(findLCP(L[i].p, L[i + 1].p));


            vector <int> mn;
            maxSegmentTree maxSeg(N);
            for (int i = 0; i < a.size(); i++) {
                  int val = maxSeg.ask(a[i] - 1);
                  mn.push_back(val);
                  maxSeg.add(a[i], i);
            }

            vector <int> mx;
            minSegmentTree minSeg(N);
            for (int i = a.size() - 1; i >= 0; i--) {
                  int val = minSeg.ask(a[i] - 1);
                  mx.push_back(val);
                  minSeg.add(a[i], i);
            }
            reverse(mx.begin(), mx.end());

            vector<vector<int> > indices(N);
            for (int i = 0; i < a.size(); i++) {
                  int id = a[i];
                  indices[id].push_back(i);
            }

            vector<LL> D(N + 1);
            // D[i] denotes number substrings which repeats i times exactly.
            // update the D array by the two pointer method as explained in editorial.
            for (int i = 1; i < N; i++) {
                  int right = 0;
                  for (int j = 0; j < indices[i].size(); j++) {
                        int id = indices[i][j];
                        if (id >= right) {
                              int lo = mn[id], hi = mx[id];
                              int t = hi - lo;
                              int mn = i;
                              if (0 <= hi && hi < a.size()) {
                                    assert (i >= a[hi]);
                                    mn = min(mn, i - a[hi]);
                              }
                              if (lo >= 0 && lo < a.size()) {
                                    assert(i >= a[lo]);
                                    mn = min(mn, i - a[lo]);
                              }
                              assert(mn >= 0);
                              D[t] += (LL)t * (LL)mn;
                              right = hi;
                        }
                  }
            }


            LL tot = accumulate(D.begin() + 2, D.end(), 0LL);
            D[1] = (LL) N * ((LL) N + 1) / 2 - tot;

            LL totalSum = accumulate(D.begin(), D.end(), 0LL);
            vector<LL> sum(N + 1);
            for (int i = 1; i <= N; i++) {
                  sum[i] = sum[i - 1] + D[i];
            }

            // process the queries.
            int queries;
            scanf ("%d", &queries);
            while (queries --) {
                  int freq;
                  scanf ("%d", &freq);
                  LL ans = 0;
                  if (freq > N) {
                        ans = 0;
                  } else {
                        ans = totalSum - sum[freq - 1];
                  }
                  assert(ans >= 0);
                  printf ("%lld\n", ans);
            }
      }

      return 0;
}
```