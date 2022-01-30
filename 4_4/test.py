import math
from decimal import Decimal

def calc1(x):
    return x**4 / 4 + 3*x**3/3 + 3*x**2/2 + x

a = 5
b = 3
print(calc1(5)-calc1(3))


def calc2(x):
    return math.pow(2, x**2)

n = 10**7
ans = 0
for i in range(0, n):
    x = Decimal(2*i+1)/Decimal(2*n)
    ans += calc2(x)
print(Decimal(ans/n))
