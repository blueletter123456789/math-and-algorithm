def new_list():
    return [[0]*2 for _ in range(2)]

def multiple(a, b, m):
    ans = new_list()
    for i in range(2):
        for j in range(2):
            for k in range(2):
                ans[i][j] += a[i][k]*b[k][j]
            ans[i][j] = ans[i][j] % m
    return ans

def bin_pow(r, e, m):
    res = r[:]
    flg = False
    for i in range(65):
        if e & 1 << i:
            if not flg:
                res = r[:]
                flg = True
            else:
                res = multiple(res, r, m)
        r = multiple(r, r, m)
    return res

n = int(input())
A = new_list()
A[0][0], A[0][1], A[1][0] = 1, 1, 1
M = 10 ** 9
ans = bin_pow(A, n-1, M)
ans = (ans[1][0]+ans[1][1]) % M
print(ans)
