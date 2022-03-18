n = int(input())
l = list(map(int, input().split()))
l.sort()
ans = 0
for i in range(n):
    ans += l[i] * (i * 2 + 1 - n)
print(ans)
