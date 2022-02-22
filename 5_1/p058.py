n, x, y = map(int, input().split())
xy = abs(x) + abs(y)
if xy <= n and n % 2 == xy % 2:
    print('Yes')
else:
    print('No')
