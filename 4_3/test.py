from decimal import Decimal
a = 2
r = 2

for _ in range(5):
    zahyo_x = a
    zahyo_y = a ** 2
    sessen_a = 2 * zahyo_x
    sessen_b = zahyo_y - sessen_a*zahyo_x

    next_a = Decimal((r-sessen_b)/sessen_a)

    a = next_a
    print(a)
