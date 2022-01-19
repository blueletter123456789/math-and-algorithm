n, s = map(int, input().split())
l = list(map(int, input().split()))
dp = [[0]*(s+1) for _ in range(n+1)]
for i in range(1, n+1):
    for j in range(s+1):
        if j < l[i-1]:
            dp[i][j] = dp[i-1][j]
        else:
            dp[i][j] = max(dp[i-1][j], dp[i-1][j-(l[i-1])] + l[i-1])
ans = 'No'
for i in dp[n]:
    if i == s:
        ans = 'Yes'
print(ans)
