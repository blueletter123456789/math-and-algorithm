n = int(input())
a = [int(i) for i in input().split()]
b = [int(i) for i in input().split()]
ansa = ansb = 0
for i in range(n):
    ansa += 1*a[i]/3
    ansb += 2*b[i]/3
print(ansa+ansb)
