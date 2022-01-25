import math

a, b, h, m = map(int, input().split())
ha = h * 30 + 0.5 * m
mb = m * 6
d = 0
if ha > mb:
    d = ha-mb
else:
    d = mb-ha
ab = a * b * math.cos(math.radians(d))
if ab < 0:
    d = 360-d
c = math.sqrt(a**2+b**2-(2*a*b*math.cos(math.radians(d))))
print(c)


# ans2
a, b, h, m = map(int, input().split())
ha = h * 30 + 0.5 * m
mb = m * 6
ax = a*math.cos(math.radians(ha))
ay = a*math.sin(math.radians(ha))
bx = b*math.cos(math.radians(mb))
by = b*math.sin(math.radians(mb))
d = math.sqrt((ax-bx)**2 + (ay-by)**2)
print(d)
