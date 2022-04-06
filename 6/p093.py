M = 10**18

def LCM(a, b):
    if b == 0:
        return a
    return LCM(b, a%b)

a, b = map(int, input().split())
res = LCM(a, b)
ans = (a*b)//res
if ans <= M:
    print(ans)
else:
    print('Large')
