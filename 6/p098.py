"""
def cross(ax, ay, bx, by, cx, cy):
    return (bx-ax)*(by-cy) - (by-ay)*(bx-cx)

def convex_hull(points, n, a, b):
    from collections import deque

    points.sort(key=lambda x: (x[0],x[1]))
    if len(points) <= 1:
        return points
    lower = deque()
    for x, y in points:
        while len(lower) >= 2 and cross(*lower[-1], *lower[-2], x, y) <= 0:
            lower.pop()
        lower.append((x, y))
    upper = deque()
    for x, y in reversed(points):
        while len(upper) >= 2 and cross(*upper[-1], *upper[-2], x, y) <= 0:
            upper.pop()
        upper.append((x, y))
    lower.pop()
    upper.pop()
    return lower + upper


if __name__ == '__main__':
    n = int(input())
    lst = [tuple(map(int, input().split())) for _ in range(n)]
    ab = tuple(map(int, input().split()))
    lst.append(ab)

    res = convex_hull(lst, n, *ab)
    if ab not in set(res):
        print('INSIDE')
    else:
        print('OUTSIDE')
"""

# sample source
n = int(input())
X = list()
Y = list()
for _ in range(n):
    x, y = map(int, input().split())
    X.append(x)
    Y.append(y)
a, b = map(int, input().split())

cnt = 0
for i in range(n):
    xa, ya = X[i] - a, Y[i] - b
    xb, yb = X[(i+1)%n] - a, Y[(i+1)%n] - b
    if ya > yb:
        xa, xb = xb, xa
        ya, yb = yb, ya
    if ya <= 0 and yb > 0 and xa*yb - ya*xb < 0:
        cnt += 1
print('INSIDE' if cnt % 2 else 'OUTSIDE')
