h, w = map(int, input().split())
R = [0]*h
L = [0]*w
A = [list(map(int, input().split())) for _ in range(h)]
for i in range(h):
    sum_num = 0
    for j in range(w):
        sum_num += A[i][j]
    R[i] = sum_num

for j in range(w):
    sum_num = 0
    for i in range(h):
        sum_num += A[i][j]
    L[j] = sum_num

for i in range(h):
    ans = [0]*w
    for j in range(w):
        ans[j] = str(R[i] + L[j] - A[i][j])
    print(' '.join(ans))
