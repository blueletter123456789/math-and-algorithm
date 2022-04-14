n = int(input())
t = list(map(int, input().split()))
sum_num = sum(t)
base = sum_num//2

dp = [[False]*(base+1) for _ in range(n+1)]
dp[0][0] = True
for i in range(1, n+1):
    for j in range(base+1):
        if dp[i-1][j]:
            dp[i][j] = True
        if t[i-1] <= j and dp[i-1][j-t[i-1]]:
            dp[i][j] = True

ans = 0
for i in range(base+1):
    if dp[n][i]:
        ans = i
print(sum_num - ans)
