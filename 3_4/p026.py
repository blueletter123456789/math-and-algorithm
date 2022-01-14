n = int(input())
ans = 1
for i in range(1, n):
    ans += 1/(1-(i/n))
print(ans)
