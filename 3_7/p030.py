n, w = map(int, input().split())
l = [list(map(int, input().split())) for _ in range(n)]
dp = [[0]*(w+1) for _ in range(n+1)]
for i in range(1, n):
    for j in range(w):
        wi = l[i][0]
        if dp[i-1][wi] - l[i][0] <= w:
            dp[i][wi] = dp[i-1][wi] + l[i][1]
print(dp, l)

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
