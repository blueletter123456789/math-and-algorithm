n = int(input())
l = list(map(int, input().split()))

a = l[0]
b = 0
for i in range(1, n):
    b = l[i]
    while a and b:
        if a >= b:
            a %= b
        else:
            b %= a
    if a < b:
        a = b

if a >= b:
    print(a)
else:
    print(b)
