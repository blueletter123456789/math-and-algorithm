n = int(input())
a = [0]*(n+1)
a[1], a[2] = 1, 1
for i in range(3, n+1):
    a[i] = (a[i-1]+a[i-2])%1000000007
print(a[n])
