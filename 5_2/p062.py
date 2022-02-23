n, k = map(int, input().split())
l = list(map(int, input().split()))
dl = list([1])
es = set([1])
cl = list()
tm = dict()
for i, v in enumerate(l):
    tm[i+1] = v
cur = tm[1]
for _ in range(1, n):
    if cur == 1 or cur in es:
        idx = dl.index(cur)
        cl = dl[:idx]
        dl = dl[idx:]
        break
    dl.append(cur)
    es.add(cur)
    cur = tm[cur]

if len(cl) > k:
    print(cl[k])
else:
    k -= len(cl)
    k %= len(dl)
    print(dl[k])

print('#######################################')

n, k = map(int, input().split())
l = list(map(int, input().split()))
First = [-1]*(n+1)
Second = [-1]*(n+1)
cnt = 0
cur = 1
while True:
    if First[cur] == -1:
        First[cur] = cnt
    elif Second[cur] == -1:
        Second[cur] = cnt
    if cnt == k:
        print(cur)
        break
    elif (Second[cur] != -1 and 
        (k-First[cur]) % (First[cur]-Second[cur]) == 0):
        print(cur)
        break
    cur = l[cur-1]
    cnt += 1
