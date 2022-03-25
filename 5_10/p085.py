n, x, y = map(int, input().split())
flg = False
for i in range(1, n+1):
    for j in range(i, n+1):
        for k in range(j, n+1):
            s = x - (i + j + k)
            if s <= 0 or s > n:
                continue
            if i * j * k * s == y:
                flg = True
                break
if flg:
    print('Yes')
else:
    print('No')
