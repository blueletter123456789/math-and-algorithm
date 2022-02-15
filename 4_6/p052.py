x, y = map(int, input().split())
M = 1000000007
dp = [[0]*(x+1) for _ in range(y+1)]
dp[1][2], dp[2][1] = 1, 1
for i in range(1, x+1):
    for j in range(1, y+1):
        if j >= 2:
            if dp[i][j] < dp[i-1][j-2]:
                dp[i][j] = (dp[i-1][j-2] + 1) % M
        if i >= 2:
            if dp[i][j] < dp[i-2][j-1]:
                dp[i][j] = (dp[i-2][j-1] + 1) % M
print(dp[x][y])
