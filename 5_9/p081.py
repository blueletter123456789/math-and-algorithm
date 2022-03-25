n = int(input())
l = [10000, 5000, 1000]
ans = 0
for i in l:
    cnt = n // i
    if cnt:
        ans += cnt
        n %= i
print(ans)
