import math

x1, y1, r1 = map(int, input().split())
x2, y2, r2 = map(int, input().split())
d12 = math.sqrt((x1-x2)**2 + (y1-y2)**2)
if r1+r2 < d12:
    print(5)
elif r1+r2 == d12:
    print(4)
elif d12+r1 == r2 or d12+r2 == r1:
    print(2)
elif d12+r1 < r2 or d12+r2 < r1:
    print(1)
else:
    print(3)
