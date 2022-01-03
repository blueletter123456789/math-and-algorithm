n, s = map(int, input().split())
cnt = 0
for i in range(1, n+1):
    for j in range(1, n+1):
        if s >= i+j:
            cnt += 1
print(cnt)
