Anya loves watching horror movies. In the best traditions of horror, she will be visited by m ghosts today. To meet them, Anya has prepared many candles, each of which burns for exactly t seconds. In one second, she manages to light one candle. More formally, Anya can spend one second to light one candle, after which this candle burns for exactly t seconds, and then goes out and is no longer usable.

For each of the m ghosts, Anya knows the moment in time at which it will arrive: the i -th visit will occur w i seconds after midnight, and all w i are different. Each visit lasts exactly one second.

What is the minimum number of candles Anya needs to use so that at least r candles are burning during each visit ? Anya can start lighting a candle at any time that is an integer number of seconds after midnight, including before midnight .

### ideas
1. 显然必须保证在w[1]前有r个candle在燃烧
2. 那么共需要r秒来点， 那么必须从 w[1] - r 开始点一根，知道w[1]
3. 如果 t < r, 那么 -1
4. 然后后面始终在下一个 max(w[i+1] - r, w[i])的时候点一根就可以了