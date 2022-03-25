a, b, c = map(int, input().split())
tgt = c-a-b
if tgt >= 0 and 4 * a * b < tgt**2:
    print('Yes')
else:
    print('No')

# not correct answer
from decimal import Decimal
import math
if Decimal(math.sqrt(a)) + Decimal(math.sqrt(b)) < Decimal(math.sqrt(c)):
    print('Yes')
else:
    print('No')
