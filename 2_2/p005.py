n = int(input())
l = list(map(int, input().split()))
# s = sum(l)
s = 0
for i in range(n):
    s += l[i]
print(s%100)
