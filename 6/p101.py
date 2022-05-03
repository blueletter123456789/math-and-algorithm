M = 1000000007

def modpow(a, b, m):
    p = a
    ans = 1
    for i in range(30):
        if b & 1 << i:
            ans *= p
            ans %= m
        p *= p
        p %= m
    return ans

def div(a, b, m):
    return (a * modpow(b, m-2, m)) % M

def ncr(n, r):
    return div(cache[n], cache[r]*cache[n-r] % M, M)

cache = [1]
for i in range(1, 10**5+10):
    cache.append(i * cache[i-1] % M)

n = int(input())
for i in range(1, n+1):
    ans = 0
    for j in range(1, ((n+i-1)//i)+1):
        ans += ncr(n - (i-1) * (j-1), j)
        ans %= M
    print(ans)
