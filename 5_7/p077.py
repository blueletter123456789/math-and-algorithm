n = int(input())
X = [0] * n
Y = [0] * n
for i in range(n):
    X[i], Y[i] = map(int, input().split())
X.sort()
Y.sort()
ans = 0
for i in range(n):
    ans += X[i]*(i * 2 + 1 - n) + Y[i]*(i * 2 + 1 - n)
print(ans)
