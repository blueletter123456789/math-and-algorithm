def multipule(a, b):
    c = [[0]*3 for _ in range(3)]
    for i in range(3):
        for k in range(3):
            for j in range(3):
                c[i][j] += a[i][k]*b[k][j]
    return c

def pow_matrix(a, n):
    p = a
    q = list()
    flg = False
    for i in range(60):
        if n & (1 << i):
            if not flg:
                q = p
                flg = True
            else:
                q = multipule(q, p)
        p = multipule(p, p)
    return q

q = int(input())
for _ in range(q):
    a, b, c, t = input().split()
    a = float(a)
    b = float(b)
    c = float(c)
    A = [[1-a, b, 0], 
        [0, 1-b, c], 
        [a, 0, 1-c]]
    ans = pow_matrix(A, int(t))
    print("{:.15f} {:.15f} {:.15f}".format(sum(ans[0]), sum(ans[1]), sum(ans[2])))
