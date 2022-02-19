def new_list(n):
    return [[0]*n for _ in range(n)]


def div_matrix(a, b, m):
    ans = new_list(len(a))
    for i in range(len(a)):
        for j in range(len(a)):
            for k in range(len(a)):
                ans[i][j] += a[i][k] * b[k][j]
                ans[i][j] %= m
    return ans
    

def bin_pow(r, e, m):
    res = r[:]
    flg = True
    for i in range(65):
        if e & 1 << i:
            if flg:
                res = r[:]
                flg = False
            else:
                res = div_matrix(res, r, m)
        r = div_matrix(r, r, m)
    return res


k, n = map(int, input().split())
e = 2**k
r = new_list(e)
M = 1000000007
if k == 2:
    r[0][3] = 1
    r[1][2] = 1
    r[2][1] = 1
    r[3][0], r[3][3] = 1, 1
elif k == 3:
    r[0][7] = 1
    r[1][6] = 1
    r[2][5] = 1
    r[3][4], r[3][7] = 1, 1
    r[4][3] = 1
    r[5][2] = 1
    r[6][1], r[6][7] = 1, 1
    r[7][0], r[7][3], r[7][6] = 1, 1, 1
else:
    r[0][15] = 1
    r[1][14] = 1
    r[2][13] = 1
    r[3][12], r[3][15] = 1, 1
    r[4][11] = 1
    r[5][10] = 1
    r[6][9], r[6][15] = 1, 1
    r[7][8], r[7][11], r[7][14] = 1, 1, 1
    r[8][7] = 1
    r[9][6] = 1
    r[10][5] = 1
    r[11][4], r[11][7] = 1, 1
    r[12][3], r[12][15] = 1, 1
    r[13][2], r[13][14] = 1, 1
    r[14][1], r[14][7], r[14][13] = 1, 1, 1
    r[15][0], r[15][3], r[15][6], r[15][12], r[15][15] = 1, 1, 1, 1, 1
ans = bin_pow(r, n, M)
print(ans[e-1][e-1])
