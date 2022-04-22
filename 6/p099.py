import sys

sys.setrecursionlimit(100000)

def dfs(current):
    child[current] = 1
    res = 0
    for v in al[current]:
        if child[v] != -1:
            continue
        res += dfs(v)
        child[current] += child[v]
        res += child[v] * (n - child[v])
    return res

n = int(input())
al = [list() for _ in range(n)]
for i in range(n-1):
    u, v = map(int, input().split())
    al[u-1].append(v-1)
    al[v-1].append(u-1)

ans = 0
child = [-1]*n
ans = dfs(0)
print(ans)
