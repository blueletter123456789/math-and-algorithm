M = 1000000007
n = int(input())
l = list(map(int, input().split()))
ans = 0
base = 1
for i in range(n):
    ans += l[i]%M * base
    ans %= M
    base *= 2
    base %= M
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
