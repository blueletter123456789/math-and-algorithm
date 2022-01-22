a = list(map(int, input().split()))
b = list(map(int, input().split()))
c = list(map(int, input().split()))

ba = [a[0]-b[0], a[1]-b[1]]
bc = [c[0]-b[0], c[1]-b[1]]
ca = [a[0]-c[0], a[1]-c[1]]
cb = [b[0]-c[0], b[1]-c[1]]
babc = ba[0]*bc[0] + ba[1]*bc[1]
cacb = ca[0]*cb[0] + ca[1]*cb[1]
if babc <= 0:
    print((ba[0]**2 + ba[1]**2)**0.5)
elif cacb <= 0:
    print((ca[0]**2 + ca[1]**2)**0.5)
else:
    bottom = (bc[0]**2 + bc[1]**2)**0.5
    area = abs(ba[0]*bc[1] - ba[1]*bc[0])
    print(area/bottom)
