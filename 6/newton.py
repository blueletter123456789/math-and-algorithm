import math

r = 2
a = 1
for i in range(5):
    z_x = a
    z_y = math.exp(a)

    s_a = z_y
    s_b = z_y - s_a*z_x

    a = (r-s_b) / s_a
    print(f'{i=}, {a=}')
print(math.log(2))
