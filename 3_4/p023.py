n = int(input())
b = [int(i) for i in input().split()]
r = [int(i) for i in input().split()]
ans = sum(b)/n + sum(r)/n
print(ans)

# n = int(input())
# b = list(map(int, input().split()))
# r = list(map(int, input().split()))
# ansb = ansr = 0
# for i in range(n):
#     ansb += b[i]/n
#     ansr += r[i]/n
# print(ansb+ansr)

# n = int(input())
# b = list(map(int, input().split()))
# r = list(map(int, input().split()))
# ans = 0
# for x in b:
#     for y in r:
#         ans += (x + y)/(n**2)
# print(ans)
