n = int(input())
h = list(map(int, input().split()))
d = [0]*n
d[1] = abs(h[0] - h[1])
for i in range(2, n):
    step1 = d[i-2] + abs(h[i-2] - h[i])
    step2 = d[i-1] + abs(h[i-1] - h[i])
    d[i] = min(step1, step2)
print(d[n-1])
