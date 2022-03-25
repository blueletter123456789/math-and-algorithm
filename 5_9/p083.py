n = int(input())
a = list(map(int, input().split()))
b = list(map(int, input().split()))
a.sort()
b.sort()
ans = 0
for i, j in zip(a, b):
    ans += abs(i-j)
print(ans)
