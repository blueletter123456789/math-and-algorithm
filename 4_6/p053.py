def bin_pow(b, e, m):
    res = 1
    for i in range(65):
        if e & 1 << i:
            res *= b
            res %= m
        b *= b
        b %= m
    return res

n = int(input())
r = 4
M = 1000000007
child = bin_pow(r, n+1, M) - 1
parent = bin_pow(3, M-2, M)
print((child*parent)%M)
