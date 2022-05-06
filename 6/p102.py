# 全探索 #################################
# def calc(i):
#     return (3 - i) % 3

# num_set = {
#     'B': 0,
#     'W': 1,
#     'R': 2
# }

# n = int(input())
# in_lst = [num_set[i] for i in list(input())]

# for i in range(n):
#     for j in range(n-i-1):
#         in_lst[j] = calc(in_lst[j]+in_lst[j+1])

# for k, v in num_set.items():
#     if in_lst[0] == v:
#         print(k)

# WA answer #############################
# M = 1000000007

# def modpow(a, b, m):
#     p = a
#     ans = 1
#     for i in range(30):
#         if b & 1 << i:
#             ans *= p
#             ans %= m
#         p *= p
#         p %= m
#     return ans

# def div(a, b, m):
#     return (a * modpow(b, m-2, m)) % M

# def ncr(n, r):
#     return div(cache[n], cache[r]*cache[n-r] % M, M)

# cache = [1]
# for i in range(1, 10**5):
#     cache.append(i * cache[i-1] % M)

# num_set = {
#     'B': 0,
#     'W': 1,
#     'R': 2
# }

# n = int(input())
# in_lst = [num_set[i] for i in list(input())]

# ans = 0
# for i in range(1, n+1):
#     ans += in_lst[i-1] * ncr(n-1, i-1)
# ans = (3*(n*(n-1))//2 - ans) % 3
# for k, v in num_set.items():
#     if ans == v:
#         print(k)

def ncr_lucas(x, y):
    if x < 3 and y < 3:
        if x < y: return 0
        if x == 2 and y == 1: return 2
        return 1
    return (ncr_lucas(x//3, y//3) * ncr_lucas(x%3, y%3)) % 3

n = int(input())
c = input()
ans = 0

num_set = {
    'B': 0,
    'W': 1,
    'R': 2
}

for i in range(n):
    ans += num_set[c[i]] * ncr_lucas(n-1, i)
    ans %= 3

if n % 2 == 0:
    ans = (3 - ans) % 3

print('BWR'[ans])
