n, r = map(int, input().split())
a = b = 1
for i in range(n-r+1, n+1):
    a *= i
for j in range(1, r+1):
    b *= j
print(a//b)
