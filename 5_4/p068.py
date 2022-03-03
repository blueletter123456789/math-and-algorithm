def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a%b)

def mul(a, b):
    return a * b // gcd(a, b)

n, k = map(int, input().split())
l = list(map(int, input().split()))
ans = 0

for i in range(1, 1 << k):
    cnt, a = 0, 1
    for j in range(k):
        if i & 1 << j:
            cnt += 1
            a = mul(a, l[j])
    if cnt % 2 == 0:
        ans -= n // a
    else:
        ans += n // a

print(ans)
