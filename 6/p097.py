# TLE source code
# def prime_calc(n):
#     if n == 0:
#         return 0
#     seen = set([1])
#     limit = int(n ** 0.5)
#     for i in range(2, limit+1):
#         for j in range(i*2, n+1, i):
#             if j not in seen:
#                 seen.add(j)
#     return n - len(seen)

# l, r = map(int, input().split())
# print(prime_calc(r) - prime_calc(l-1))

l, r = map(int, input().split())
prime = [True]*(r-l+1)
if l == 1:
    prime[0] = False
limit = int(r ** 0.5)
for i in range(2, limit+1):
    min_val = int(((l + i - 1) // i ) * i)
    for j in range(min_val, r+1, i):
        if i == j:
            continue
        prime[j-l] = False
ans = 0
for i in range(r-l+1):
    if prime[i]:
        ans += 1
print(ans)

