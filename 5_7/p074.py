n = int(input())
l = list(map(int, input().split()))
ans = 0
for i in range(n):
    ans += l[i] * (2 * i + 1 - n)
print(ans)
