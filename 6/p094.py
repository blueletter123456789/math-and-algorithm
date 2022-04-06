n = int(input())
B = list(map(int, input().split()))
tgt = B[0]
ans = 0
for i in B:
    ans += min(i, tgt)
    tgt = i
ans += tgt
print(ans)
