def cross(x1, y1, x2, y2):
    return x1*y2 - x2*y1

ax, ay = map(int, input().split())
bx, by = map(int, input().split())
cx, cy = map(int, input().split())
dx, dy = map(int, input().split())

k1 = cross(bx-ax, by-ay, cx-ax, cy-ay)
k2 = cross(bx-ax, by-ay, dx-ax, dy-ay)
k3 = cross(dx-cx, dy-cy, ax-cx, ay-cy)
k4 = cross(dx-cx, dy-cy, bx-cx, by-cy)

if not k1 and not k2 and not k3 and not k4:
    p1 = (ax, ay)
    p2 = (bx, by)
    p3 = (cx, cy)
    p4 = (dx, dy)
    if p1 > p2:
        p1, p2 = p2, p1
    if p3 > p4:
        p3, p4 = p4, p3
    if max(p1, p3) <= min(p2, p4):
        print('Yes')
    else:
        print('No')
else:
    res1 = res2 = False
    if k1 >= 0 and k2 <= 0:
        res1 = True
    elif k1 <= 0 and k2 >= 0:
        res1 = True
    if k3 >= 0 and k4 <= 0:
        res2 = True
    elif k3 <= 0 and k4 >= 0:
        res2 = True
    
    if res1 and res2:
        print('Yes')
    else:
        print('No')
