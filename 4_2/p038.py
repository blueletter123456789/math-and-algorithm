n, q = map(int, input().split())
l = list(map(int, input().split()))
t = [0]*(n+1)
for i in range(1, n+1):
    t[i] = t[i-1]+l[i-1]
for j in range(q):
    l, r = map(int, input().split())
    print(t[r]-t[l-1])
