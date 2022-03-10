def mul(n, r, M):
    res = 1
    p = 2
    for i in range(n):
        if r & 1 << i:
            res *= p
            res %= M
        p *= p
        p %= M
    return res

n = int(input())
l = list(map(int, input().split()))
l.sort()
M = 1000000007
ans = 0
r = [0]*n
for i in range(n):
    r[i] = 2 ** i
    # r[i] = mul(n, i, M)

for i in range(n):
    # ans += l[i]%M * mul(n, i, M)
    ans += l[i]%M * r[i]
    ans %= M
print(ans)

# 全探索では
# res = 0
# for i in range(1 << n):
#     t = list()
#     for j in range(n):
#         if i & 1 << j:
#             t.append(l[j])
#     print(t)
#     if t:
#         res += max(t)
# print(res)
