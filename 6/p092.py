# import math
# n = int(input())
# max_num = int(math.sqrt(n))+1
# ans = 10**15
# for i in range(1, max_num):
#     tgt = 0
#     if n % i == 0:
#         tgt = n / i
#         ans = min((2*tgt+2*i), ans)
# print(int(ans))

n = int(input())
max_num = int(n**0.5)+1
ans = 10**15
for i in range(1, max_num):
    tgt = 0
    if n % i == 0:
        tgt = n // i
        ans = min((tgt+i)*2, ans)
print(int(ans))
