n = int(input())
l = list(map(int, input().split()))
# print(sum(l))
ans = 0
for i in range(n):
    ans += l[i]
print(ans)
