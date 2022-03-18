M = 1000000007

def bin_pow(a, r):
    res = 1
    p = a
    for i in range(30):
        if r & 1 << i:
            res *= p
            res %= M
        p *= p
        p %= M
    return res

def div(a, b):
    return (a * bin_pow(b, M-2)) % M

def ncr(n, r):
    return div(bc_list[n], (bc_list[r]*bc_list[n-r])%M)

n = int(input())
l = list(map(int, input().split()))
bc_list = [0]*(n+1)
bc_list[0] = 1
for i in range(1, n+1):
    bc_list[i] = (i * bc_list[i-1]) % M
ans = 0
for i in range(n):
    ans += l[i] * ncr(n-1, i)
    ans %= M
print(ans)
