import sys

sys.setrecursionlimit(1000000)

def dfs(pos):
    vis[pos] = True
    for i in al[pos]:
        if vis[i] == False:
            dfs(i)

n, m = map(int, input().split())
al = [list() for _ in range(n+1)]

for i in range(m):
    a, b = map(int, input().split())
    al[a].append(b)
    al[b].append(a)

vis = [False]*(n+1)
dfs(1)
ans = True
for i in range(1, n+1):
    if not vis[i]:
        ans = False
        break
if ans:
    print('The graph is connected.')
else:
    print('The graph is not connected.')
