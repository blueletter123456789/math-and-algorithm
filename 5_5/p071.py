N = int(input())
l = [list(map(int, input().split())) for _ in range(N)]

ans = 0
for i in range(N):
    for j in range(i+1, N):
        if l[i][1]/l[i][0] == l[j][1]/l[j][0]:
            continue
        x = (l[i][2]*l[j][1]-l[j][2]*l[i][1])/(l[i][0]*l[j][1]-l[j][0]*l[i][1])
        y = (l[i][2]*l[j][0]-l[j][2]*l[i][0])/(l[i][1]*l[j][0]-l[j][1]*l[i][0])
        flg = True
        for k in range(N):
            if l[k][0]*x + l[k][1]*y > l[k][2]:
                flg = False
        if flg:
            ans = max(ans, x+y)
print(ans)
