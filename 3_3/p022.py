# n = int(input())
# l = list(map(int, input().split()))
# d = {}
# cache = {}
# cnt = 0
# for i in l:
#     key = 100000 - i
#     val = d.get(key, 0)
#     t = cache.get(key, 0)
#     if val != 0 and t == 0:
#         cnt += 1
#     else:
#         d[i] = key
#     cache[key] = cache.get(key, 0) + 1
# if cache.get(50000, 0) > 1:
#     cnt += 1
# print(cnt)

MAX_NUM = 100000
n = int(input())
l = list(map(int, input().split()))
d = [0 for _ in range(MAX_NUM)]
for i in range(n):
    d[l[i]] += 1
cnt = 0
for i in range(1, MAX_NUM//2):
    cnt += d[i] * d[MAX_NUM-i]
cnt += d[MAX_NUM//2] * (d[MAX_NUM//2]-1)//2
print(cnt)
