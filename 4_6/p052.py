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

#########################################################
x, y = map(int, input().split())

M = 1000000007
flg = True
if (2*y-x)%3 != 0 or (2*x-y)%3 != 0:
    flg = False
elif (2*y-x) < 0 or (2*x-y) < 0:
    flg = False
if flg:
    child = parent = 1
    a = (2*y-x)//3
    b = (2*x-y)//3
    for i in range(1,a+b+1):
        child *= i % M
        child %= M
    for i in range(1,a+1):
        parent *= i % M
        parent %= M
    for i in range(1, b+1):
        parent *= i % M
        parent %= M
    ans = 1
    p = parent
    for i in range(30):
        if M-2 & 1 << i:
            ans *= p
            ans %= M
        p *= p
        p %= M
    ans = (child * ans) % M
    print(ans)
else:
    print(0)
