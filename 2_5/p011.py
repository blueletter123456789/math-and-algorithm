n = int(input())
l = ['2']
for i in range(2, n+1):
    for j in range(2, i):
        if i % j == 0:
            break
        elif i-j == 1:
            l.append(str(i))
print(' '.join(l))
