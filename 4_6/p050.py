a, b = map(int, input().split())

M = 1000000007
ans = 1
p = a
for i in range(35):
    if b & 1 << i:
        ans *= p
        ans %= M
    p *= p
    p %= M
print(ans%M)
