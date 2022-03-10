import math

a, b = map(int, input().split())
ans = 1
for i in range(1, b+1):
    tb = b // i
    ta = math.ceil(a / i)
    if tb - ta >= 1:
        ans = i
print(ans)
