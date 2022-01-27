n = int(input())
l = list(map(int, input().split()))
m = int(input())
cs = [0]*(n+1)
for i in range(1,n):
    cs[i+1] = cs[i] + l[i-1]
old = int(input())
ans = 0
for _ in range(m-1):
    next = int(input())
    ans += abs(cs[next]-cs[old])
    old = next
print(ans)
