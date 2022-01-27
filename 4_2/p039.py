n, q = map(int, input().split())
t = [0]*(n+1)
for i in range(q):
    l, r, x = map(int, input().split())
    t[l-1] += x
    t[r] -= x
for j in range(1, n):
    if t[j] > 0:
        print('<', end='')
    elif t[j] < 0:
        print('>', end='')
    else:
        print('=', end='')
