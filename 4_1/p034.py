n = int(input())
l = [list(map(int, input().split())) for _ in range(n)]
ans = 2*(10**6)
for i in range(n):
    for j in range(i+1, n):
        ax = l[i][0]
        ay = l[i][1]
        bx = l[j][0]
        by = l[j][1]
        tmp = ((ax-bx)**2 + (ay-by)**2)**0.5
        if ans > tmp:
            ans = tmp
print(ans)


# answer
import sys
 
int1 = lambda x: int(x)-1
pDB = lambda *x: print(*x, end="\n", file=sys.stderr)
p2D = lambda x: print(*x, sep="\n", end="\n\n", file=sys.stderr)
def II(): return int(sys.stdin.readline())
def LI(): return list(map(int, sys.stdin.readline().split()))
def LLI(rows_number): return [LI() for _ in range(rows_number)]
def LI1(): return list(map(int1, sys.stdin.readline().split()))
def LLI1(rows_number): return [LI1() for _ in range(rows_number)]
def SI(): return sys.stdin.readline().rstrip()
inf = 18446744073709551615
md = 10**9+7
 
def min_dist(xy):
    n = len(xy)
    if n < 2: return inf
 
    d = min(min_dist(xy[:n//2]), min_dist(xy[n//2:]))
 
    mid = xy[n//2][0]
    nxy = []
    for x, y in xy:
        if x < mid-d: continue
        if x > mid+d: break
        nxy.append((x, y))
 
    nxy.sort(key=lambda x: x[1])
    n = len(nxy)
    for i in range(n):
        x, y = nxy[i]
        for j in range(i+1, n):
            s, t = nxy[j]
            if t-y > d: break
            d = min(d, ((x-s)**2+(y-t)**2)**0.5)
    return d
 
n = II()
xy = LLI(n)
xy.sort(key=lambda x: x[0])
print(min_dist(xy))