x, y = map(int, input().split())
M = 1000000007
c = x + y
child = 1
parent = 1
for i in range(1, c+1):
    child *= (i % M)
    child %= M
for i in range(1, x+1):
    parent *= (i % M)
    parent %= M
for i in range(1, y+1):
    parent *= (i % M)
    parent %= M
ans = 1
p = parent
for i in range(30):
    if M-2 & 1 << i:
        ans *= p
        ans %= M
    p *= p
    p %= M
ans = (child * ans) % M
print(ans)
