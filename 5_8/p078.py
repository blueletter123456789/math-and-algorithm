from collections import deque

n, m = map(int, input().split())
al = [list() for _ in range(n)]
for i in range(m):
    a, b = map(int, input().split())
    al[a-1].append(b-1)
    al[b-1].append(a-1)
ans = [-1]*n
que = deque([0])
ans[0] = 0
while len(que):
    p = que.popleft()
    for i in al[p]:
        if ans[i] >= 0:
            continue
        ans[i] = min(120, ans[p]+1)
        que.append(i)
for i in ans:
    print(120 if i == -1 else i)
