a, b, c = map(int, input().split())
ans = 0
if c > 1:
    p = 1
    flg = False
    for _ in range(b+1):
        if a < p:
            flg = True
            break
        p *= c
    if flg:
        print('Yes')
    else:
        print('No')
else:
    print('No')
