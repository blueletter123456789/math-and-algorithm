n, w = map(int, input().split())
dp = [[0]*(w+1) for _ in range(n+1)]
for i in range(1, n+1):
    a, b = map(int, input().split())
    for j in range(w+1):
        if j < a:
            dp[i][j] = dp[i-1][j]
        else:
            dp[i][j] = max(dp[i-1][j], dp[i-1][j-a]+b)
print(max(dp[n]))


# bit全探索

# l = [list(map(int, input().split())) for _ in range(n)]
# ans_v = 0
# for i in range(1 << n):
#     x = y = 0
#     for j in range(n):
#         if i & (1 << j):
#             x += l[j][0]
#             y += l[j][1]
#     if x <= w and ans_v < y:
#         ans_v = y
# print(ans_v)
