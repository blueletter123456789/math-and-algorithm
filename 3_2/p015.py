a, b = map(int, input().split())

while a and b:
    if a >= b:
        a %= b
    else:
        b %= a

if a >= b:
    print(a)
else:
    print(b)
