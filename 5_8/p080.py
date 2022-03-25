# one case was failed
from collections import defaultdict, deque
 
n, m = map(int, input().split())
al = [defaultdict(int) for _ in range(n)]
for i in range(m):
    a, b, c = map(int, input().split())
    al[a-1][b-1] = c
    al[b-1][a-1] = c
dist = [-1]*n
seen = [False]*n
dist[0], seen[0] = 0, True
que = deque([0])
while len(que):
    p = que.popleft()
    for k, v in al[p].items():
        cost = dist[p]+v
        if dist[k] == -1:
            dist[k] = cost
        else:
            dist[k] = min(dist[k], cost)
        if seen[k]:
            continue
        seen[k] = True
        que.append(k)
print(dist[n-1])

print('###################################')

import heapq

N, M = map(int, input().split())
A, B, C = [ None ] * M, [ None ] * M, [ None ] * M
for i in range(M):
	A[i], B[i], C[i] = map(int, input().split())

G = [ list() for i in range(N + 1) ]
for i in range(M):
	G[A[i]].append((B[i], C[i]))
	G[B[i]].append((A[i], C[i]))

max_num = 10 ** 19
dist = [ max_num ] * (N + 1)
used = [ False ] * (N + 1)
Q = list()
dist[1] = 0
heapq.heappush(Q, (0, 1))

while len(Q) >= 1:
	pos = heapq.heappop(Q)[1]
	if used[pos] == True:
		continue
	used[pos] = True
	for i in G[pos]:
		to = i[0]
		cost = dist[pos] + i[1]
		if dist[to] > cost:
			dist[to] = cost
			heapq.heappush(Q, (dist[to], to))

if dist[N] != max_num:
	print(dist[N])
else:
	print(-1)
