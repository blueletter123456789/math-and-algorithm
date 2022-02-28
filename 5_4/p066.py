n, k = map(int, input().split())
ans = 0
for a in range(1, n+1):
    for b in range(max(1, a-(k-1)), min(n+1, a+k)):
        for c in range(max(1, a-(k-1)), min(n+1, a+k)):
            if abs(b-c) < k:
                ans += 1
print(n**3 - ans)
