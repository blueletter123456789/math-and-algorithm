from dis import dis
import queue


import queue

start_point = 1
n, m = map(int, input().split())
al = [list() for _ in range(n+1)]
for i in range(m):
    a, b = map(int, input().split())
    al[a].append(b)
    al[b].append(a)
dist = [-1]*(n+1)
dist[start_point] = 0
bfs = queue.Queue()
bfs.put(start_point)
while not bfs.empty():
    pos = bfs.get()
    nex = al[pos]
    for i in nex:
        if dist[i] == -1:
            dist[i] = dist[pos]+1
            bfs.put(i)

for i in range(1, n+1):
    print(dist[i])
