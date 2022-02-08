from dis import dis
import queue

h, w = map(int, input().split())
sx, sy = map(int, input().split())
gx, gy = map(int, input().split())
start = sx * w + sy
goal = gx * w + gy

c = [['#']*(w+1) for _ in range(h+1)]
for i in range(1, h+1):
    l = list(input())
    for j in range(1, w+1):
        c[i][j] = l[j-1]

lim = (h+1)*(w+1)
g = [list() for _ in range(lim)]
for i in range(1, h+1):
    for j in range(1, w):
        idx1 = i * w + j
        idx2 = i * w + (j + 1)
        if c[i][j] == '.' and c[i][j+1] == '.':
            g[idx1].append(idx2)
            g[idx2].append(idx1)

for i in range(1, h):
    for j in range(1, w+1):
        idx1 = i * w + j
        idx2 = (i + 1) * w + j
        if c[i][j] == '.' and c[i+1][j] == '.':
            g[idx1].append(idx2)
            g[idx2].append(idx1)

dist = [-1] * lim
que = queue.Queue()
que.put(start)
dist[start] = 0
while not que.empty():
    x = que.get()
    for nex in g[x]:
        if dist[nex] == -1:
            dist[nex] = dist[x] + 1
            que.put(nex)
print(dist[goal])
