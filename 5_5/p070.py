N, K = map(int, input().split())
X = [0]*N
Y = [0]*N
for i in range(N):
    X[i], Y[i] = map(int, input().split())

ans = 10 ** 20
for i in range(N):
    for j in range(N):
        for k in range(N):
            for l in range(N):
                xl, xr, yb, yt = X[i], X[j], Y[k], Y[l]
                if xl > xr or yb > yt:
                    continue
                cnt = 0
                for m in range(N):
                    if xl <= X[m] <= xr and yb <= Y[m] <= yt:
                        cnt += 1
                if cnt >= K:
                    s = abs(xr-xl)*abs(yt-yb)
                    ans = min(ans, s)
print(ans)
