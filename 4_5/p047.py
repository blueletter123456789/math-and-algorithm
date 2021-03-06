import sys
sys.setrecursionlimit(1000000)

n, m = map(int, input().split())
# adjacency list
al = [list() for _ in range(n+1)]
g1 = set() # グループ1
g2 = set() # グループ2
for _ in range(m):
    a, b = map(int, input().split())
    al[a].append(b)
    al[b].append(a)
vis = [False]*(n+1)
vis[0] = True
def dfs(pos, g_flg):
    for i in al[pos]:
        if g_flg:
            g1.add(i)
        else:
            g2.add(i)
        if not vis[i]:
            vis[i] = True
            dfs(i, not g_flg)
flg = True

for i in range(len(vis)):
    if not vis[i]:
        dfs(i, True)
        if len(g1) != len(g1-g2):
            flg = False
            break
        g1 = set()
        g2 = set()

if flg:
    print('Yes')
else:
    print('No')


###########################################
import sys

sys.setrecursionlimit(210000)

def dfs(pos):
    for i in g[pos]:
        if color[i] == 0:
            color[i] = 3-color[pos]
            dfs(i)

n, m = map(int, input().split())
g = [list() for _ in range(n+1)]
al = [-1]*(m)
bl = [-1]*(m)
for i in range(m):
    al[i], bl[i] = map(int, input().split())
    g[al[i]].append(bl[i])
    g[bl[i]].append(al[i])

color = [0]*(n+1)

for i in range(1, n+1):
    if color[i] == 0:
        color[i] = 1
        dfs(i)
ans = True
for i in range(m):
    if color[al[i]] == color[bl[i]]:
        ans = False
        break
if ans:
    print('Yes')
else:
    print('No')


###########################################
import sys
from collections import deque 
 
def main():
    def judge():
        N, M = map(int, readline().split())
        G = [[] for _ in range(N)]
        for _ in range(M):
            a, b = map(int, readline().split())
            G[a - 1].append(b - 1)  # a <-> b
            G[b - 1].append(a - 1)
 
        colors = [0] * N
        for i in range(N):
            if colors[i] != 0:
                continue
            que = deque((i,))
            colors[i] = 1
            while que:
                u = que.popleft()
                cu = colors[u]
                for v in G[u]:
                    if cu == colors[v]:
                        return False
                    if colors[v] == 0:
                        colors[v] = -cu
                        que.append(v)
        return True
 
    print("Yes" if judge() else "No")
 
 
if __name__ == '__main__':
    readline = sys.stdin.readline
    readlines = sys.stdin.readlines
    main()
