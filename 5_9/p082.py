n = int(input())
l = [list(map(int, input().split())) for _ in range(n)]
l.sort(key=lambda x: x[1])
end = l[0][1]
ans = 1
for i in range(1, n):
    if end <= l[i][0]:
        ans += 1
        end = l[i][1]
print(ans)
