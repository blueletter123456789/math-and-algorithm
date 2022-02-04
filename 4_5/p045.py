import queue

sp = 1
n, m = map(int, input().split())
al = [list() for _ in range(n+1)]
for _ in range(m):
    a, b = map(int, input().split())
    al[a].append(b)
    al[b].append(a)
cl = [-1]*(n+1)

cl[sp] = 0
rst = 0
que = queue.Queue()
que.put(sp)
while not que.empty():
    tgt = que.get()
    nxt = al[tgt]
    cnt = 0
    for i in nxt:
        if tgt > i:
            cnt += 1
        if cl[i] == -1:
            cl[i] = 0
            que.put(i)
    cl[tgt] = cnt
    if cnt == 1:
        rst += 1
print(rst)


###########################################
n, m = map(int, input().split())
cnt = [0] * (n+1)
for _ in range(m):
    a, b = map(int, input().split())
    if a < b:
        cnt[b] += 1
    else:
        cnt[a] += 1
print(cnt.count(1))
